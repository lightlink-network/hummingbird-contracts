import { expect } from "chai";
import { ethers } from "hardhat";
import { JsonRpcProvider, Log, EventLog } from "ethers";
import { startNetworks } from "../../scripts/lib/startNetworks";
import { ChildProcess } from "child_process";
import {
  setupCanonicalStateChain,
  makeNextBlock,
} from "../lib/chain";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
  BridgeProofHelper,
  Challenge,
  L2CrossDomainMessenger,
  L1CrossDomainMessenger,
  L1Block,
} from "../../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { proxyDeployAndInitialize } from "../../scripts/lib/deploy";
import { initiateWithdraw, getWithdrawalProofs } from "../lib/bridge";
import { assert } from "console";

describe("Cross-chain interaction", function () {
  // Networks
  let l1Network: ChildProcess;
  let l2Network: ChildProcess;

  // Signers
  let l1Deployer: HardhatEthersSigner;
  let l2Deployer: HardhatEthersSigner;
  let l2Depositor: HardhatEthersSigner;

  // Providers
  let l1Provider: JsonRpcProvider;
  let l2Provider: JsonRpcProvider;

  // Contracts
  let canonicalStateChain: CanonicalStateChain;
  let lightLinkPortal: LightLinkPortal;
  let l2ToL1MessagePasser: L2ToL1MessagePasser;
  let BridgeProofHelper: BridgeProofHelper;
  let challenge: Challenge;
  let l2CrossDomainMessenger: L2CrossDomainMessenger;
  let l1CrossDomainMessenger: L1CrossDomainMessenger;
  let l1Block: L1Block;

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

    // - CanonicalStateChain
    const _chain = await setupCanonicalStateChain(
      l1Deployer,
      l1Deployer.address,
    );
    canonicalStateChain = _chain.canonicalStateChain;
    console.log("CanonicalStateChain deployed");

    // - Challenge
    const challengeDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("Challenge"),
      [
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ],
    );
    challenge = challengeDeployment.contract as Challenge;

    // - LightLinkPortal
    const lightLinkPortalDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("LightLinkPortal"),
      [
        await canonicalStateChain.getAddress(),
        await challengeDeployment.address,
        ethers.ZeroAddress, // L1Block address 
      ],
    );
    lightLinkPortal = lightLinkPortalDeployment.contract as LightLinkPortal;

    // - BridgeProofHelper
    const bridgeProofHelperFactory = await ethers.getContractFactory(
      "contracts/L1/test/BridgeProofHelper.sol:BridgeProofHelper",
      l1Deployer,
    );
    BridgeProofHelper = (await bridgeProofHelperFactory.deploy()) as any;
    await BridgeProofHelper.waitForDeployment();

    // Deploy L2 contracts

    // - L2ToL1MessagePasser
    const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
      "contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser",
      l2Deployer,
    );
    l2ToL1MessagePasser = (await L2ToL1MessagePasserFactory.deploy()) as any;
    await l2ToL1MessagePasser.waitForDeployment();

    // - L1Block 
    const L1BlockFactory = await ethers.getContractFactory(
      "contracts/L2/L1Block.sol:L1Block",
      l2Deployer,
    );
    l1Block = (await L1BlockFactory.deploy()) as any;

    // Cross domain messengers
    // - Infer deployment addresses before deploying
    const l2CrossDomainMessengerAddr = ethers.getCreateAddress({
      from: l2Deployer.address,
      nonce: await l2Provider.getTransactionCount(l2Deployer.address) + 1,
      // +1 because implementation will be deployed first
    });
    const l1CrossDomainMessengerAddr = ethers.getCreateAddress({
      from: l1Deployer.address,
      nonce: await l1Provider.getTransactionCount(l1Deployer.address) + 1,
      // +1 because implementation will be deployed first
    });

    // - Deploy cross domain messengers
    console.log("Deploying cross domain messengers");
    const L2CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
      l2Deployer,
      await ethers.getContractFactory("L2CrossDomainMessenger"),
      [l1CrossDomainMessengerAddr, await l2ToL1MessagePasser.getAddress(), await l1Block.getAddress()],
    );
    l2CrossDomainMessenger = L2CrossDomainMessengerDeployment.contract as L2CrossDomainMessenger;

    const L1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("L1CrossDomainMessenger"),
      [await lightLinkPortal.getAddress(), l2CrossDomainMessengerAddr],
    );
    l1CrossDomainMessenger = L1CrossDomainMessengerDeployment.contract as L1CrossDomainMessenger;

    expect(l2CrossDomainMessengerAddr, "address mismatch").to.equal(
      await l2CrossDomainMessenger.getAddress(),
    );

    expect(l1CrossDomainMessengerAddr, "address mismatch").to.equal(
      await l1CrossDomainMessenger.getAddress(),
    );

    // Impersonate l2 Depositor account
    console.log("Impersonating L2 depositor account");
    await l2Provider.send("hardhat_impersonateAccount", [
      "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
    ]);

    l2Depositor = await l2Provider.getSigner(
      "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
    ) as any;

    console.log("L2 depositor account impersonated - funding...");
    await l2Deployer.sendTransaction({
      to: l2Depositor.address,
      value: ethers.parseEther("1"),
    });

    // Setup GasPayingToken in L1Block
    console.log("Setting up GasPayingToken in L1Block");
    await l1Block.connect(l2Depositor).setGasPayingToken(
      "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", // Constants.ETHER
      18,
      ethers.encodeBytes32String("Ether"),
      ethers.encodeBytes32String("ETH"),
    );
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

    it("Finalize Withdrawal", async function () {

      // Initiate withdrawal from L2

      const recipient = randomAddress();
      const withdrawal = await initiateWithdraw(
        l2ToL1MessagePasser,
        l2Deployer,
        recipient,
        21000,
        "0x",
        {
          value: ethers.parseEther("1"),
        },
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
      const proveTx = await
        lightLinkPortal
          .connect(l1Deployer)
          .proveWithdrawalTransaction(
            withdrawal.withdrawalTx,
            await canonicalStateChain.chainHead(),
            outputProof,
            withdrawalProof.storageProof,
          )
      expect(proveTx, "Failed to prove withdrawal").to.emit(
        lightLinkPortal,
        "WithdrawalProven",
      );

      // get finalization seconds from challenge 
      const finalizationSeconds = await challenge.connect(l1Deployer).finalizationSeconds();

      // move time forward
      await l1Provider.send("evm_increaseTime", [Number(finalizationSeconds) * 2]);
      await l1Provider.send("evm_mine", []);

      // fund the contract
      await lightLinkPortal.connect(l1Deployer).donateETH({ value: ethers.parseEther("3") });

      // Finalize withdrawal
      const finalizeTx = await lightLinkPortal
        .connect(l1Deployer)
        .finalizeWithdrawalTransaction(
          withdrawal.withdrawalTx
        );
      expect(finalizeTx, "Failed to finalize withdrawal").to.emit(
        lightLinkPortal,
        "WithdrawalFinalized",
      );

      expect(await l1Provider.getBalance(recipient)).to.equal(ethers.parseEther("1"));
    });

    it("Deposit", async function () {
      const recipient = randomAddress();
      const value = ethers.parseEther("1");

      const depositTx = await lightLinkPortal
        .connect(l1Deployer)
        .depositTransaction(
          recipient,
          value,
          21000,
          false,
          "0x",
          {
            value: value,
          },
        );

      expect(depositTx, "Failed to deposit").to.emit(
        lightLinkPortal,
        "TransactionDeposited",
      );

    });

  }); // describe("LightLinkPortal")


});


const randomAddress = () => ethers.Wallet.createRandom().address;