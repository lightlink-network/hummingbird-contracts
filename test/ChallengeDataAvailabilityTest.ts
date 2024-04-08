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
        ethers.ZeroAddress,
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

  describe("deployment", function () {
    it("should set the correct owner", async function () {
      expect(await challenge.owner()).to.equal(owner.address);
    });

    it("should set the correct canonical state chain", async function () {
      expect(await challenge.chain()).to.equal(
        await canonicalStateChain.getAddress(),
      );
    });

    it("should set the correct DAOracle", async function () {
      expect(await challenge.daOracle()).to.equal(
        await mockDaOracle.getAddress(),
      );
    });
  });

  describe("challengeDataRootInclusion", function () {
    it("should not allow challenging the genesis header", async function () {
      await expect(
        challenge.connect(challengeOwner).challengeDataRootInclusion(0, 0),
      ).to.be.revertedWith("cannot challenge genesis block");
    });

    it("should not allow challenging a header that is not canonical", async function () {
      await expect(
        challenge.connect(challengeOwner).challengeDataRootInclusion(100, 0),
      ).to.be.revertedWith("block not in the chain yet");
    });

    it("should not allow challenge if challenge fee is not paid", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);

      await expect(
        challenge.connect(challengeOwner).challengeDataRootInclusion(1, 0),
      ).to.be.revertedWith("challenge fee not paid");
    });

    it("should allow challenge (if challenge fee is paid)", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      await expect(
        challenge
          .connect(challengeOwner)
          .challengeDataRootInclusion(1, 0, { value: challengeFee }),
      ).to.not.be.reverted;

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      const challengeInfo = await challenge.daChallenges(challengeKey);
      expect(
        challengeInfo.challenger,
        "expect: daChallenges(hash).challenger = challengeOwner.address",
      ).to.equal(challengeOwner.address);
      expect(
        challengeInfo.blockIndex,
        "expect: daChallenges(hash).blockIndex = 1",
      ).to.equal(1);
      expect(
        challengeInfo.status,
        "expect: daChallenges(hash).status = 1",
      ).to.equal(STATUS_INITIATED);
    });

    it("should not allow two challenges on the same block", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      await expect(
        challenge
          .connect(challengeOwner)
          .challengeDataRootInclusion(1, 0, { value: challengeFee }),
      ).to.be.revertedWith("challenge already exists");
    });
  });

  describe("defendDataRootInclusion", function () {
    it("should not allow defending a non-existent challenge", async function () {
      const proof = { ...EXAMPLE_PROOF };

      const challengeKey = ethers.keccak256(
        ethers.toUtf8Bytes("None existent challenge"),
      );

      await expect(
        challenge
          .connect(publisher)
          .defendDataRootInclusion(challengeKey, proof),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not allow defending a challenge with incorrect proof", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      // set mock daoracle to return false result
      await mockDaOracle.getFunction("setResult").send(false);

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = toBigInt(header.celestiaPointers[0].height);

      await expect(
        challenge
          .connect(publisher)
          .defendDataRootInclusion(challengeKey, proof),
      ).to.be.revertedWith("invalid proof");
    });

    it("should be able to defend a challenge", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = toBigInt(header.celestiaPointers[0].height);

      const prebalance = await challengeOwner.provider.getBalance(
        publisher.address,
      );

      await challenge
        .connect(publisher)
        .defendDataRootInclusion(challengeKey, proof);

      const c = await challenge.daChallenges(challengeKey);
      expect(c.status, "expect: daChallenges(hash).status = 3").to.equal(
        STATUS_DEFENDER_WON,
      );

      const postbalance = await challengeOwner.provider.getBalance(
        publisher.address,
      );

      expect(postbalance).to.be.greaterThan(prebalance);
    });

    it("should not allow defending a challenge twice", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = toBigInt(header.celestiaPointers[0].height);

      await challenge
        .connect(publisher)
        .defendDataRootInclusion(challengeKey, proof);

      await expect(
        challenge
          .connect(publisher)
          .defendDataRootInclusion(challengeKey, proof),
      ).to.be.revertedWith("challenge is not in the correct state");
    });
  });

  describe("settleDataRootInclusion", function () {
    it("should not allow settling a non-existent challenge", async function () {
      const challengeKey = ethers.keccak256(
        ethers.toUtf8Bytes("None existent challenge"),
      );

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challengeKey),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not allow settling a challenge that is already defended", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = toBigInt(header.celestiaPointers[0].height);

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      await challenge
        .connect(publisher)
        .defendDataRootInclusion(challengeKey, proof);

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challengeKey),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not settle challenge if challenge period is not over", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challengeKey),
      ).to.be.revertedWith("challenge has not expired");
    });

    it("should settle challenge if challenge period is over", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = await challenge.dataRootInclusionChallengeKey(
        hash,
        0,
      );

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const prebalance = await challengeOwner.provider.getBalance(
        challengeOwner.address,
      );

      // increase time by 49 hours
      await ethers.provider.send("evm_increaseTime", [49 * 60 * 60]);
      await ethers.provider.send("evm_mine");

      await challenge
        .connect(challengeOwner)
        .settleDataRootInclusion(challengeKey);

      const c = await challenge.daChallenges(challengeKey);
      expect(c.status, "expect: daChallenges(hash).status = 2").to.equal(
        STATUS_CHALLENGER_WON,
      );

      const postbalance = await challengeOwner.provider.getBalance(
        challengeOwner.address,
      );
      expect(postbalance).to.be.greaterThan(prebalance);
    });
  });

  describe("toggleDAChallenge", function () {
    it("toggleDAChallenge should be failed without owner", async function () {
      await expect(
        challenge
          .connect(otherAccount)
          .getFunction("toggleDAChallenge")
          .send(false),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });

    it("toggleDAChallenge should be correct", async function () {
      await challenge.getFunction("toggleDAChallenge").send(true);
      expect(await challenge.isDAChallengeEnabled()).to.equal(true);

      await challenge.getFunction("toggleDAChallenge").send(false);
      expect(await challenge.isDAChallengeEnabled()).to.equal(false);
    });

    it("challenge DA should revert when disabled", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);
      await challenge.getFunction("toggleDAChallenge").send(false);
      await expect(
        challenge
          .connect(challengeOwner)
          .challengeDataRootInclusion(1, 0, { value: challengeFee }),
      ).to.be.revertedWith("DA challenges are disabled");
    });

    it("challenge DA should not revert when enabled", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);
      await challenge.getFunction("toggleDAChallenge").send(true);
      await expect(
        challenge
          .connect(challengeOwner)
          .challengeDataRootInclusion(1, 0, { value: challengeFee }),
      ).to.not.be.reverted;
    });
  });
});

const EXAMPLE_PROOF: ChallengeDAProof = {
  rootNonce: BigInt(500),
  dataRootTuple: {
    height: BigInt(1),
    dataRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
  },
  proof: {
    sideNodes: [],
    key: BigInt(0),
    numLeaves: BigInt(0),
  },
};
