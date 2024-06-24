import { expect } from "chai";
import { ethers } from "hardhat";
import { JsonRpcProvider, Log, EventLog } from "ethers";
import { startNetworks } from "../../scripts/lib/startNetworks";
import { ChildProcess } from "child_process";
import {
  setupCanonicalStateChain,
  pushRandomHeader,
  makeNextBlock,
} from "../lib/chain";
import { BigNumber } from "@ethersproject/bignumber";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
  BridgeProofHelper,
} from "../../typechain-types";
import { Types } from "../../typechain-types/contracts/L1/test/BridgeProofHelper";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { proxyDeployAndInitialize } from "../../scripts/lib/deploy";
import {
  MessagePassedEvent,
  L2ToL1MessagePasserInterface,
} from "../../typechain-types/contracts/L2/L2ToL1MessagePasser";
import exp from "constants";

import { makeStateTrieProof, hashMessageHash, initiateWithdraw, getWithdrawalProofs } from "../lib/bridge";

describe("Cross-chain interaction", function () {
  // Networks
  let l1Network: ChildProcess;
  let l2Network: ChildProcess;

  // Signers
  let l1Deployer: HardhatEthersSigner;
  let l2Deployer: HardhatEthersSigner;

  // Providers
  let l1Provider: JsonRpcProvider;
  let l2Provider: JsonRpcProvider;

  // Contracts
  let canonicalStateChain: CanonicalStateChain;
  let lightLinkPortal: LightLinkPortal;
  let l2ToL1MessagePasser: L2ToL1MessagePasser;
  let BridgeProofHelper: BridgeProofHelper;

  before(async function () {
    // Start Anvil network instances
    const networks = await startNetworks();
    l1Network = networks.l1Network;
    l2Network = networks.l2Network;

    console.log("Anvil networks started");

    // Set up L1 provider and signer
    l1Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8545");
    l1Deployer = (await l1Provider.getSigner(0)) as any;

    // Set up L2 provider and signer
    l2Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8546");
    l2Deployer = (await l2Provider.getSigner(0)) as any;

    // Deploy L1 contracts

    // CanonicalStateChain
    const _chain = await setupCanonicalStateChain(
      l1Deployer,
      l1Deployer.address,
    );
    canonicalStateChain = _chain.canonicalStateChain;

    console.log("CanonicalStateChain deployed");

    // LightLinkPortal
    const lightLinkPortalDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("LightLinkPortal"),
      [
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ],
    );
    lightLinkPortal = lightLinkPortalDeployment.contract as LightLinkPortal;

    // BridgeProofHelper
    const bridgeProofHelperFactory = await ethers.getContractFactory(
      "contracts/L1/test/BridgeProofHelper.sol:BridgeProofHelper",
      l1Deployer,
    );
    BridgeProofHelper = (await bridgeProofHelperFactory.deploy()) as any;
    await BridgeProofHelper.waitForDeployment();

    // Deploy L2 contracts
    const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
      "contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser",
      l2Deployer,
    );
    l2ToL1MessagePasser = (await L2ToL1MessagePasserFactory.deploy()) as any;
    await l2ToL1MessagePasser.waitForDeployment();
  });

  after(async function () {
    // Kill Anvil network instances
    l1Network.kill();
    l2Network.kill();
  });

  describe("BridgeProofHelper", function () {
    it("Should verify a proof", async function () {
      // Initiate withdrawal from L2
      const withdrawal = await initiateWithdraw(
        l2ToL1MessagePasser,
        l2Deployer,
        ethers.ZeroAddress,
        0,
        "0x",
      );

      // Verify initiateWithdrawal event
      expect(withdrawal.withdrawalTx).to.not.be.undefined;
      expect(withdrawal.withdrawalTx).to.not.be.undefined;
      expect(
        withdrawal.withdrawalTx.data,
        "MessagePassed event: incorrect data",
      ).to.equal("0x");
      expect(
        withdrawal.withdrawalTx.gasLimit,
        "MessagePassed event: incorrect gas limit",
      ).to.equal(0);
      expect(
        withdrawal.withdrawalTx.nonce,
        "MessagePassed event: incorrect nonce",
      ).to.not.be.undefined;
      expect(
        withdrawal.withdrawalTx.sender,
        "MessagePassed event: incorrect sender",
      ).to.equal(l2Deployer.address);
      expect(
        withdrawal.withdrawalTx.target,
        "MessagePassed event: incorrect target",
      ).to.equal(ethers.ZeroAddress);

      // Generate proof
      const { withdrawalProof, outputProof, outputRoot } = await getWithdrawalProofs(
        l2Provider,
        withdrawal.initiateTx.blockNumber ?? "latest",
        l2ToL1MessagePasser,
        withdrawal.messageSlot,
      );

      // Verify proof
      const verified = await BridgeProofHelper.connect(l1Deployer).checkProof(
        outputRoot,
        outputProof,
        withdrawal.withdrawalHash,
        withdrawalProof.storageProof,
      );
      expect(verified, "Proof verification failed").to.be.true;
    }); // it("Should verify a proof")
  }); // describe("BridgeProofHelper")

  describe("LightLinkPortal", function () {
    it("Prove withdrawal", async function () {

      // Initiate withdrawal from L2
      const withdrawal = await initiateWithdraw(
        l2ToL1MessagePasser,
        l2Deployer,
        ethers.ZeroAddress,
        0,
        "0x",
      );

      // Generate proofs
      const { withdrawalProof, outputProof, outputRoot } = await getWithdrawalProofs(
        l2Provider,
        withdrawal.initiateTx.blockNumber ?? "latest",
        l2ToL1MessagePasser,
        withdrawal.messageSlot,
      );

      // Push a new header to L1
      const [nextHeader] = await makeNextBlock(l1Deployer, canonicalStateChain);
      nextHeader.outputRoot = outputRoot;
      const pushTx = await canonicalStateChain
        .connect(l1Deployer)
        .pushBlock(nextHeader);
      expect(pushTx, "Failed to push block").to.emit(
        canonicalStateChain,
        "BlockAdded",
      );

      // Prove withdrawal
      const proveTx = await lightLinkPortal
        .connect(l1Deployer)
        .proveWithdrawalTransaction(
          withdrawal.withdrawalTx,
          await canonicalStateChain.chainHead(),
          outputProof,
          withdrawalProof.storageProof,
        );
      expect(proveTx, "Failed to prove withdrawal").to.emit(
        lightLinkPortal,
        "WithdrawalProven",
      );
    });
  }); // describe("LightLinkPortal")
});
