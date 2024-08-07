import { expect } from "chai";
import { ethers } from "hardhat";
import { JsonRpcProvider, Log, EventLog } from "ethers";
import { startNetworks } from "../../scripts/hardhat/lib/startNetworks";
import { ChildProcess } from "child_process";
import { setupCanonicalStateChain, makeNextBlock } from "../lib/chain";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
  BridgeProofHelper,
  Challenge,
  L2CrossDomainMessenger,
  L1CrossDomainMessenger,
  L1Block,
  PingPong,
  L2CrossDomainMessenger__factory,
  L2ToL1MessagePasser__factory,
  L1Block__factory,
} from "../../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  proxyDeployAndInitialize,
  uupsProxyDeployAndInitialize,
} from "../../scripts/hardhat/lib/deploy";
import {
  initiateWithdraw,
  getWithdrawalProofs,
  sendMessageL2ToL1,
} from "../lib/bridge";
import { assert } from "console";
import { Proxy } from "../../typechain-types/contracts/universal";
import { Proxy__factory } from "../../typechain-types/factories/contracts/universal";

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
    const networks = await startNetworks({
      genesisFile: "test/tests/genesis.json",
    });
    l1Network = networks.l1Network;
    l2Network = networks.l2Network;

    console.log("Anvil networks started");

    // Set up L1 provider and signer
    l1Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8545");
    l1Deployer = (await l1Provider.getSigner(0)) as any;

    // Set up L2 provider and signer
    l2Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8546");
    l2Deployer = (await l2Provider.getSigner(0)) as any;

    console.log("L1 ChainID", (await l1Provider.getNetwork()).chainId);
    console.log("L2 ChainID", (await l2Provider.getNetwork()).chainId);

    // Deploy L1 contracts

    // - CanonicalStateChain
    const _chain = await setupCanonicalStateChain(
      l1Deployer,
      l1Deployer.address,
    );
    canonicalStateChain = _chain.canonicalStateChain;
    console.log("CanonicalStateChain deployed");

    // - Challenge
    const challengeDeployment = await uupsProxyDeployAndInitialize(
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
        l1Deployer.address,
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

    // Link L2 contract predeploys
    l2ToL1MessagePasser = L2ToL1MessagePasser__factory.connect(
      "0x4200000000000000000000000000000000000016",
      l2Deployer,
    );

    l1Block = L1Block__factory.connect(
      "0x4200000000000000000000000000000000000015",
      l2Deployer,
    );

    l2CrossDomainMessenger = L2CrossDomainMessenger__factory.connect(
      "0x4200000000000000000000000000000000000007",
      l2Deployer,
    );

    // - Deploy cross domain messenger on L1
    console.log("Deploying cross domain messengers");

    const L1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("L1CrossDomainMessenger"),
      [await lightLinkPortal.getAddress()],
    );
    l1CrossDomainMessenger =
      L1CrossDomainMessengerDeployment.contract as L1CrossDomainMessenger;

    // Impersonate l2 Depositor account
    console.log("Impersonating L2 depositor account");
    await l2Provider.send("hardhat_impersonateAccount", [
      "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
    ]);

    l2Depositor = (await l2Provider.getSigner(
      "0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001",
    )) as any;

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
      const { withdrawalProof, outputProof, outputRoot } =
        await getWithdrawalProofs(
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
    it("Correct owner", async function () {
      expect(await lightLinkPortal.connect(l1Deployer).owner()).to.equal(
        l1Deployer.address,
      );
    });

    it("Can Pause and Upause", async function () {
      await lightLinkPortal.connect(l1Deployer).pause();
      expect(await lightLinkPortal.connect(l1Deployer).paused()).to.be.true;
      await lightLinkPortal.connect(l1Deployer).unpause();
      expect(await lightLinkPortal.connect(l1Deployer).paused()).to.be.false;
    });

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
      const { withdrawalProof, outputProof, outputRoot } =
        await getWithdrawalProofs(
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
      const { withdrawalProof, outputProof, outputRoot } =
        await getWithdrawalProofs(
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

      // get finalization seconds from challenge
      const finalizationSeconds = await challenge
        .connect(l1Deployer)
        .finalizationSeconds();

      // move time forward
      await l1Provider.send("evm_increaseTime", [
        Number(finalizationSeconds) * 2,
      ]);
      await l1Provider.send("evm_mine", []);

      // fund the contract
      await lightLinkPortal
        .connect(l1Deployer)
        .donateETH({ value: ethers.parseEther("3") });

      // Finalize withdrawal
      const finalizeTx = await lightLinkPortal
        .connect(l1Deployer)
        .finalizeWithdrawalTransaction(withdrawal.withdrawalTx);
      expect(finalizeTx, "Failed to finalize withdrawal").to.emit(
        lightLinkPortal,
        "WithdrawalFinalized",
      );

      expect(await l1Provider.getBalance(recipient)).to.equal(
        ethers.parseEther("1"),
      );
    });

    it("Deposit", async function () {
      const recipient = randomAddress();
      const value = ethers.parseEther("1");

      const depositTx = await lightLinkPortal
        .connect(l1Deployer)
        .depositTransaction(recipient, value, 21000, false, "0x", {
          value: value,
        });

      expect(depositTx, "Failed to deposit").to.emit(
        lightLinkPortal,
        "TransactionDeposited",
      );
    });
  }); // describe("LightLinkPortal")

  describe("L1CrossDomainMessenger", function () {
    it("Pong", async function () {
      // deploy pingpong contract to l2
      const PingPongFactory = await ethers.getContractFactory("PingPong");
      const pingPong = (await PingPongFactory.connect(
        l2Deployer,
      ).deploy()) as PingPong;

      // encode call: `ping("Hello L2!")`
      const callData = pingPong.interface.encodeFunctionData("ping", [
        "Hello L2!",
      ]);

      // estimate gas
      const gasEstimate = await l2Deployer.estimateGas({
        to: await pingPong.getAddress(),
        data: callData,
      });

      const msgTx = await l1CrossDomainMessenger
        .connect(l1Deployer)
        .sendMessage(await pingPong.getAddress(), callData, gasEstimate);

      expect(msgTx, "Failed to send message").to.emit(
        lightLinkPortal,
        "TransactionDeposited",
      );
    });
  }); // describe("L1CrossDomainMessenger")

  describe("L2CrossDomainMessenger", function () {
    it("Pong", async function () {
      // deploy pingpong contract to l1
      const PingPongFactory = await ethers.getContractFactory("PingPong");
      const pingPong = (await PingPongFactory.connect(
        l1Deployer,
      ).deploy()) as PingPong;

      // encode call: `ping("Hello L1!")`
      const callData = pingPong.interface.encodeFunctionData("ping", [
        "Hello L1!",
      ]);

      // send message
      const withdrawal = await sendMessageL2ToL1(
        l2CrossDomainMessenger,
        l2ToL1MessagePasser,
        l2Deployer,
        l1Provider,
        await pingPong.getAddress(),
        callData,
      );

      // Generate withdrawal proofs
      const { withdrawalProof, outputProof, outputRoot } =
        await getWithdrawalProofs(
          l2Provider,
          withdrawal.sendMessageTx.blockNumber ?? "latest",
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

      // get finalization seconds from challenge
      const finalizationSeconds = await challenge
        .connect(l1Deployer)
        .finalizationSeconds();

      // move time forward
      await l1Provider.send("evm_increaseTime", [
        Number(finalizationSeconds) * 2,
      ]);
      await l1Provider.send("evm_mine", []);

      // Finalize withdrawal
      const finalizeTx = await lightLinkPortal
        .connect(l1Deployer)
        .finalizeWithdrawalTransaction(withdrawal.withdrawalTx);
      expect(finalizeTx, "Failed to finalize withdrawal").to.emit(
        lightLinkPortal,
        "WithdrawalFinalized",
      );
      expect(finalizeTx, "Failed to call ping").to.emit(pingPong, "Pong");
    });
  }); // describe("L2CrossDomainMessenger")

  describe("Proxy", async function () {
    it("Upgrade LightLinkPortal", async function () {
      // step 1: deploy new implementation
      const newLightLinkPortalFactory =
        await ethers.getContractFactory("LightLinkPortal");
      const newLightLinkPortal = (await newLightLinkPortalFactory
        .connect(l1Deployer)
        .deploy()) as LightLinkPortal;

      // step 2: upgrade proxy contract
      const proxy = Proxy__factory.connect(
        await lightLinkPortal.getAddress(),
        l1Deployer,
      );
      const upgradeTx = await proxy
        .connect(l1Deployer)
        .upgradeTo(await newLightLinkPortal.getAddress());

      const upgradeTxReceipt = await upgradeTx.wait();
      expect(upgradeTxReceipt, "Failed to upgrade proxy").to.emit(
        proxy,
        "Upgraded",
      );
    });
  });
});

const randomAddress = () => ethers.Wallet.createRandom().address;
