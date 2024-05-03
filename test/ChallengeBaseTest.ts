import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract, toBigInt } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { setupCanonicalStateChain } from "./lib/chain";
import {
  CanonicalStateChain,
  Challenge,
  Challenge__factory,
} from "../typechain-types";

import { pushRandomHeader } from "./lib/chain";
import { proxyDeployAndInitialize } from "../scripts/lib/deploy";

type Header = CanonicalStateChain.HeaderStruct;

// Challenge Status
const STATUS_NONE = 0;
const STATUS_INITIATED = 1;
const STATUS_CHALLENGER_WON = 2;
const STATUS_DEFENDER_WON = 3;

// TODO test challenge fees are sent
describe("ChallengeDataAvailability", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;
  let otherAccount: HardhatEthersSigner;
  let challengeOwner: HardhatEthersSigner;
  let challengeFee = ethers.parseEther("1");

  let genesisHeader: Header;
  let genesisHash: string;
  let canonicalStateChain: CanonicalStateChain;
  let mockDaOracle: Contract;
  let challenge: Challenge;

  beforeEach(async function () {
    [owner, publisher, otherAccount, challengeOwner] =
      await ethers.getSigners();

    // 1. Setup canonical state chain
    const _chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = _chain.canonicalStateChain;
    genesisHash = _chain.genesisHash;
    genesisHeader = _chain.genesisHeader;

    // 2. setup mock DAOracle
    const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
    mockDaOracle = (await _MockDaOracle.deploy()) as any;

    const deployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("Challenge"),
      [
        await canonicalStateChain.getAddress(),
        await mockDaOracle.getAddress(),
        ethers.ZeroAddress, // chainOracle not needed for this test
      ],
    );

    challenge = Challenge__factory.connect(deployment.address, owner);

    // set challenge contract in canonical state chain
    await canonicalStateChain.setChallengeContract(deployment.address);
    // set Challenge fee
    await challenge.setChallengeFee(challengeFee);
    // set publisher as defender
    await challenge.setDefender(publisher.address);
    // set isDAChallengeEnabled to true
    await challenge.getFunction("toggleDAChallenge").send(true);
  });

  describe("setChallengeWindow", function () {
    it("should set the challenge window", async function () {
      const window = 12 * 60 * 60; // 12 hours
      await challenge.setChallengeWindow(window);
      expect(await challenge.challengeWindow()).to.be.equal(window);
    });
    it("should revert as window < minimum threshold (12 hours)", async function () {
      const window = 12 * 60 * 60 - 1; // 12 hours - 1 second
      await expect(challenge.setChallengeWindow(window)).to.be.revertedWith(
        "challenge window must be between 12 hours and 3 weeks",
      );
    });
    it("should revert as window > maximum threshold (3 weeks)", async function () {
      const window = 3 * 7 * 24 * 60 * 60 + 1; // 3 weeks + 1 second
      await expect(challenge.setChallengeWindow(window)).to.be.revertedWith(
        "challenge window must be between 12 hours and 3 weeks",
      );
    });
    it("should revert if not called by owner", async function () {
      await expect(
        challenge.connect(otherAccount).setChallengeWindow(10),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });
  });

  describe("setChallengePeriod", function () {
    it("should set the challenge period", async function () {
      const period = 12 * 60 * 60; // 12 hours
      await challenge.setChallengePeriod(period);
      expect(await challenge.challengePeriod()).to.be.equal(period);
    });
    it("should revert as period < minimum threshold (12 hours)", async function () {
      const period = 12 * 60 * 60 - 1; // 12 hours - 1 second
      await expect(challenge.setChallengePeriod(period)).to.be.revertedWith(
        "challenge period must be between 12 hours and 3 weeks",
      );
    });
    it("should revert as period > maximum threshold (3 weeks)", async function () {
      const period = 3 * 7 * 24 * 60 * 60 + 1; // 3 weeks + 1 second
      await expect(challenge.setChallengePeriod(period)).to.be.revertedWith(
        "challenge period must be between 12 hours and 3 weeks",
      );
    });
    it("should revert if not called by owner", async function () {
      await expect(
        challenge.connect(otherAccount).setChallengePeriod(10),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });
  });

  describe("setChallengeFee", function () {
    it("should set the challenge fee", async function () {
      const fee = ethers.parseEther("1");
      await challenge.setChallengeFee(fee);
      expect(await challenge.challengeFee()).to.be.equal(fee);
    });
    it("should revert if not called by owner", async function () {
      await expect(
        challenge.connect(otherAccount).setChallengeFee(ethers.parseEther("1")),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });
  });

  describe("setChallengeReward", function () {
    it("should set the challenge reward", async function () {
      const reward = ethers.parseEther("1");
      await challenge.setChallengeReward(reward);
      expect(await challenge.challengeReward()).to.be.equal(reward);
    });
    it("should revert if not called by owner", async function () {
      await expect(
        challenge
          .connect(otherAccount)
          .setChallengeReward(ethers.parseEther("1")),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });
  });

  describe("setDefender", function () {
    it("should set the defender", async function () {
      await challenge.setDefender(otherAccount.address);
      expect(await challenge.defender()).to.be.equal(otherAccount.address);
    });
    it("should revert if not called by owner", async function () {
      await expect(
        challenge.connect(otherAccount).setDefender(otherAccount.address),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });
  });
});
