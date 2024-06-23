import { expect } from "chai";
import { ethers } from "hardhat";
import { L2ToL1MessagePasser } from "../../../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { parseMessagePassedEvent } from "../../lib/utils";

describe("L2ToL1MessagePasser", function () {
  // Signer
  let deployer: HardhatEthersSigner;

  // Contracts
  let l2ToL1MessagePasser: L2ToL1MessagePasser;

  before(async function () {
    [deployer] = await ethers.getSigners();

    // Deploy L2 contracts
    const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
      "contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser",
      deployer,
    );
    l2ToL1MessagePasser = (await L2ToL1MessagePasserFactory.deploy()) as any;
    await l2ToL1MessagePasser.waitForDeployment();
  });

  describe("Receive", function () {
    it("Should initiate a withdrawal", async function () {
      // Initiate withdrawal via L2ToL1MessagePasser by sending ether to it
      const tx = await deployer.sendTransaction({
        to: l2ToL1MessagePasser.getAddress(),
        value: ethers.parseEther("1"),
      });
      const initiateReceipt = await tx.wait();
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
      ).to.equal(100000);
      expect(
        msgPassed.withdrawalTx.nonce,
        "MessagePassed event: incorrect nonce",
      ).to.not.be.undefined;
      expect(
        msgPassed.withdrawalTx.sender,
        "MessagePassed event: incorrect sender",
      ).to.equal(deployer.address);
      expect(
        msgPassed.withdrawalTx.target,
        "MessagePassed event: incorrect target",
      ).to.equal(deployer.address);
    });
  });

  describe("Burn", function () {
    it("Should burn native ETH held by this contract", async function () {
      // Verify balance is 1 ether from previous test
      const balanceBeforeBurn = await ethers.provider.getBalance(
        l2ToL1MessagePasser.getAddress(),
      );
      expect(balanceBeforeBurn).to.equal(ethers.parseEther("1"));

      // Burn the ether
      await l2ToL1MessagePasser.burn();

      // Check balance after burn
      const balanceAfter = await ethers.provider.getBalance(
        l2ToL1MessagePasser.getAddress(),
      );

      // Verify balance is 0
      expect(balanceAfter).to.equal(0);
    });
  });

  describe("initiateWithdrawal", function () {
    it("Should initiate a withdrawal", async function () {
      // Initiate withdrawal via L2ToL1MessagePasser by calling initiateWithdrawal
      const tx = await l2ToL1MessagePasser.initiateWithdrawal(
        deployer.address,
        ethers.parseEther("1"),
        "0x",
      );

      // Verify initiateWithdrawal event
      const initiateReceipt = await tx.wait();
      const msgPassed = parseMessagePassedEvent(
        l2ToL1MessagePasser.interface,
        initiateReceipt!.logs[0],
      );
      expect(msgPassed).to.not.be.undefined;
      expect(msgPassed.withdrawalTx).to.not.be.undefined;
      expect(
        msgPassed.withdrawalTx.data,
        "MessagePassed event: incorrect data",
      ).to.equal("0x");
      expect(
        msgPassed.withdrawalTx.gasLimit,
        "MessagePassed event: incorrect gas limit",
      ).to.equal(ethers.parseEther("1"));
      expect(
        msgPassed.withdrawalTx.nonce,
        "MessagePassed event: incorrect nonce",
      ).to.not.be.undefined;
      expect(
        msgPassed.withdrawalTx.sender,
        "MessagePassed event: incorrect sender",
      ).to.equal(deployer.address);
      expect(
        msgPassed.withdrawalTx.target,
        "MessagePassed event: incorrect target",
      ).to.equal(deployer.address);
    });
  });

  describe("messageNonce", function () {
    it("Should return the next message nonce", async function () {
      // Get the next message nonce
      const nonce = await l2ToL1MessagePasser.messageNonce();
      expect(nonce).to.not.be.undefined;
    });
  });
});
