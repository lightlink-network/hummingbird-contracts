import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { setupCanonicalStateChain } from "./lib/chain";
import { Header } from "./lib/header";
import { pushRandomHeader } from "./lib/chain";
import { CanonicalStateChain, Challenge } from "../typechain-types";
import { ChallengeDAProof, getChallengeKey } from "./lib/challenge";
import { asBigInt } from "./lib/utils";

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

    // 3. deploy challenge contract
    const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
    const challengeFactory: any = await ethers.getContractFactory("Challenge");
    const challengeImplementation = await challengeFactory.deploy();

    const proxy = await proxyFactory.deploy(
      await challengeImplementation.getAddress(),
      challengeImplementation.interface.encodeFunctionData("initialize", [
        ethers.ZeroAddress,
        await canonicalStateChain.getAddress(),
        await mockDaOracle.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress, // chainOracle not needed for this test
      ]),
    );

    challenge = challengeFactory.attach(await proxy.getAddress());

    // set challenge contract in canonical state chain
    _chain.canonicalStateChain.setChallengeContract(
      await challenge.getAddress(),
    );

    // set Challenge fee
    await challenge.setChallengeFee(challengeFee);
    // set publisher as defender
    await challenge.setDefender(publisher.address);
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

    // it("should not allow challenge if challenge fee is not paid", async function () {
    //   await pushRandomHeader(publisher, canonicalStateChain);

    //   await expect(
    //     challenge
    //       .connect(challengeOwner)
    //       .challengeDataRootInclusion")
    //       .send(1)
    //   ).to.be.revertedWith("challenge fee not paid");
    // });

    it("should allow challenge (if challenge fee is paid)", async function () {
      console.log("PUSHING RANDOM HEADER");
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = getChallengeKey(headerHash, 0n);

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const c = await challenge.daChallenges(challengeKey);
      expect(
        c.challenger,
        "expect: daChallenges(challengeKey).challenger = challengeOwner.address",
      ).to.equal(challengeOwner.address);
      expect(
        c.blockIndex,
        "expect: daChallenges(challengeKey).blockIndex = 1",
      ).to.equal(1);
      expect(
        c.status,
        "expect: daChallenges(challengeKey).status = 1",
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
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = getChallengeKey(headerHash, BigInt(0));

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const challengeStruct = await challenge.daChallenges(challengeKey);
      expect(
        challengeStruct.status,
        "expect: daChallenges(challengeKey).status = 1",
      ).to.equal(STATUS_INITIATED);

      // set mock daoracle to return false result
      await mockDaOracle.setResult(false);

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = asBigInt(header.pointers[0].celestiaHeight);

      await expect(
        challenge
          .connect(publisher)
          .defendDataRootInclusion(challengeKey, proof),
      ).to.be.revertedWith("invalid proof");
    });

    it("should be able to defend a challenge", async function () {
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = getChallengeKey(headerHash, BigInt(0));

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = asBigInt(header.pointers[0].celestiaHeight);

      const prebalance = await challengeOwner.provider.getBalance(
        publisher.address,
      );

      await challenge
        .connect(publisher)
        .defendDataRootInclusion(challengeKey, proof);

      const c = await challenge.daChallenges(challengeKey);
      expect(
        c.status,
        "expect: daChallenges(challengeKey).status = 3",
      ).to.equal(STATUS_DEFENDER_WON);

      const postbalance = await challengeOwner.provider.getBalance(
        publisher.address,
      );

      expect(postbalance).to.be.greaterThan(prebalance);
    });

    it("should not allow defending a challenge twice", async function () {
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );
      const challengeKey = getChallengeKey(headerHash, BigInt(0));

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = asBigInt(header.pointers[0].celestiaHeight);

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
        ethers.toUtf8Bytes("None existent block"),
      );

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challengeKey),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not allow settling a challenge that is already defended", async function () {
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );
      const challegeKey = getChallengeKey(headerHash, BigInt(0));

      const proof = { ...EXAMPLE_PROOF };
      proof.dataRootTuple.height = asBigInt(header.pointers[0].celestiaHeight);

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      await challenge
        .connect(publisher)
        .defendDataRootInclusion(challegeKey, proof);

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challegeKey),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not settle challenge if challenge period is not over", async function () {
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = getChallengeKey(headerHash, BigInt(0));

      await challenge
        .connect(challengeOwner)
        .challengeDataRootInclusion(1, 0, { value: challengeFee });

      await expect(
        challenge.connect(challengeOwner).settleDataRootInclusion(challengeKey),
      ).to.be.revertedWith("challenge has not expired");
    });

    it("should settle challenge if challenge period is over", async function () {
      const [headerHash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );

      const challengeKey = getChallengeKey(headerHash, BigInt(0));

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
      expect(
        c.status,
        "expect: daChallenges(challengeKey).status = 2",
      ).to.equal(STATUS_CHALLENGER_WON);

      const postbalance = await challengeOwner.provider.getBalance(
        challengeOwner.address,
      );
      expect(postbalance).to.be.greaterThan(prebalance);
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
