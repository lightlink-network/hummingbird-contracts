import { expect } from "chai";
import { ethers } from "hardhat";
import { JsonRpcProvider, Log, EventLog } from "ethers";
import { startNetworks } from "../scripts/lib/startNetworks";
import { ChildProcess } from "child_process";
import { setupCanonicalStateChain, pushRandomHeader } from "./lib/chain";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
  BridgeProofHelper,
} from "../typechain-types";
import { Types } from "../typechain-types/contracts/L1/test/BridgeProofHelper";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { proxyDeployAndInitialize } from "../scripts/lib/deploy";
import {
  MessagePassedEvent,
  L2ToL1MessagePasserInterface,
} from "../typechain-types/contracts/L2/L2ToL1MessagePasser";
import exp from "constants";

import { makeStateTrieProof, hashMessageHash } from "../test/lib/bridge";

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

  it("Should interact with L1 and L2", async function () {
    // Initiate withdrawal from L2
    await l2ToL1MessagePasser
      .connect(l2Deployer)
      .initiateWithdrawal(ethers.ZeroAddress, 0, "0x");

    // Push a new header to L1
    const [headerHash, header] = await pushRandomHeader(
      l1Deployer,
      canonicalStateChain,
    );

    // log headers
    console.log("Header hash:", headerHash);
    console.log("Header:", header);
  });

  describe("BridgeProofHelper", function () {
    it("Should verify a proof", async function () {
      // Initiate withdrawal from L2
      const initiateTx = await l2ToL1MessagePasser
        .connect(l2Deployer)
        .initiateWithdrawal(ethers.ZeroAddress, 0, "0x");
      const initiateReceipt = await initiateTx.wait();
      const msgPassed = parseMessagePassedEvent(
        l2ToL1MessagePasser.interface,
        initiateReceipt!.logs[0],
      );

      // Verify initiateWithdrawal event
      expect(msgPassed).to.not.be.undefined;
      expect(msgPassed.withdrawalTx).to.not.be.undefined;
      expect(
        msgPassed.withdrawalTx.data,
        "MessagePassed event: incorrect data",
      ).to.equal("0x");
      expect(
        msgPassed.withdrawalTx.gasLimit,
        "MessagePassed event: incorrect gas limit",
      ).to.equal(0);
      expect(
        msgPassed.withdrawalTx.nonce,
        "MessagePassed event: incorrect nonce",
      ).to.not.be.undefined;
      expect(
        msgPassed.withdrawalTx.sender,
        "MessagePassed event: incorrect sender",
      ).to.equal(l2Deployer.address);
      expect(
        msgPassed.withdrawalTx.target,
        "MessagePassed event: incorrect target",
      ).to.equal(ethers.ZeroAddress);

      // Get withdrawal tx hash
      const withdrawalHash = await BridgeProofHelper.connect(
        l1Deployer,
      ).hashWithdrawalTx(msgPassed.withdrawalTx);

      expect(withdrawalHash).to.not.be.undefined;
      expect(withdrawalHash).to.not.be.empty;

      // Calculate message slot
      const messageSlot = hashMessageHash(withdrawalHash);

      // Generate proof
      let withdrawalProof = await makeStateTrieProof(
        l1Provider,
        1, // block number
        await l2ToL1MessagePasser.getAddress(),
        messageSlot,
      );

      // log withdrawal proof
      console.log("Withdrawal proof:", withdrawalProof);
    });
  });

  const parseMessagePassedEvent = (
    iface: L2ToL1MessagePasserInterface,
    log: Log | EventLog,
  ): {
    evt: MessagePassedEvent.Event;
    withdrawalTx: Types.WithdrawalTransactionStruct;
  } => {
    const event = iface.parseLog({
      topics: [...log.topics],
      data: log.data,
    })!;

    return {
      evt: event as unknown as MessagePassedEvent.Event,
      withdrawalTx: {
        data: event.args.data,
        gasLimit: event.args.gasLimit,
        nonce: event.args.nonce,
        sender: event.args.sender,
        target: event.args.target,
        value: event.args.value,
      },
    };
  };
});
