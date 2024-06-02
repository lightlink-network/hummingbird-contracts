import { ethers } from "hardhat";
import { expect } from "chai";
import { toBigInt } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { makeNextBlock, setupCanonicalStateChain } from "./lib/chain";
import {
  CanonicalStateChain,
  Challenge,
  Challenge__factory,
  ChallengeTest,
} from "../typechain-types";
import { proxyDeployAndInitialize } from "../scripts/lib/deploy";

type Header = CanonicalStateChain.HeaderStruct;

describe("ChallengeHeader", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;
  let otherAccount: HardhatEthersSigner;
  let challengeOwner: HardhatEthersSigner;

  let genesisHeader: Header;
  let genesisHash: string;
  let canonicalStateChain: CanonicalStateChain;
  let challenge: Challenge;
  let challengeTest: ChallengeTest;

  beforeEach(async function () {
    [owner, publisher, otherAccount, challengeOwner] =
      await ethers.getSigners();

    // Setup canonical state chain
    const _chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = _chain.canonicalStateChain as any;
    genesisHash = _chain.genesisHash;
    genesisHeader = _chain.genesisHeader;

    const deployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("Challenge"),
      [
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress, // chain Oracle not needed for this test
      ],
    );

    challenge = Challenge__factory.connect(deployment.address, owner);

    _chain.canonicalStateChain
      .getFunction("setChallengeContract")
      .send(await challenge.getAddress());

    // Setup challenge test contract allowing to test internal functions
    const challengeTestFactory =
      await ethers.getContractFactory("ChallengeTest");

    challengeTest = await challengeTestFactory.deploy();
    await challengeTest.waitForDeployment();

    await challengeTest.initialize(
      await canonicalStateChain.getAddress(),
      ethers.ZeroAddress,
      ethers.ZeroAddress, // chain Oracle not needed for this test
    );

    // set isHeaderChallengeEnabled to true
    await challenge.getFunction("toggleHeaderChallenge").send(true);
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
  });

  describe("invalidateHeader", function () {
    it("should not allow challenging the genesis header", async function () {
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(0),
      ).to.be.revertedWith("cannot challenge genesis block");
    });

    it("should not allow challenging a header that is not canonical", async function () {
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(100),
      ).to.be.revertedWith("block not in the chain yet");
    });

    it("should not invalidate a valid header", async function () {
      const validHeader: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1),
      ).to.be.revertedWith("header is valid");
    });

    it("should invalidate header with incorrect L2 height", async function () {
      const invalidHeader: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) - BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(invalidHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1),
      ).to.emit(challenge, "InvalidHeader");
    });

    it("should invalidate header even after a newer block", async function () {
      const invalidHeader: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) - BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(invalidHeader);

      const [validHeader] = await makeNextBlock(publisher, canonicalStateChain);

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(2),
      ).to.be.revertedWith("header is valid");

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1),
      ).to.emit(challenge, "InvalidHeader");

      const headIndex = await canonicalStateChain
        .connect(publisher)
        .getFunction("chainHead")
        .call([]);

      expect(headIndex).to.equal(0);
    });

    it("should invalidate 10th block, add new then invalidate 11th", async function () {
      // 1. push 9 blocks.
      for (let i = 0; i < 9; i++) {
        const [header] = await makeNextBlock(publisher, canonicalStateChain);

        await canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header);
      }

      // 2. push invalid header as 10th block
      const [invalidHeader] = await makeNextBlock(
        publisher,
        canonicalStateChain,
      );
      invalidHeader.l2Height = BigInt(1);

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(invalidHeader);

      // 3. should not be able to invalidate 9th block
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(9),
      ).to.be.revertedWith("header is valid");

      // 10. but should be able to invalidate 10th block
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(10),
      ).to.emit(challenge, "InvalidHeader");

      // 11. check the chain actually rolled back to 9th block
      const headIndex = await canonicalStateChain
        .connect(publisher)
        .getFunction("chainHead")
        .call([]);

      expect(headIndex).to.equal(9);

      // 12. add a new block back in as 10th block
      const [validBlock] = await makeNextBlock(publisher, canonicalStateChain);

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validBlock);

      // 13. should not be able to invalidate 10th block
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(10),
      ).to.be.revertedWith("header is valid");

      // 14. add the 11th invalid block
      const [invalidBlock] = await makeNextBlock(
        publisher,
        canonicalStateChain,
      );
      invalidBlock.l2Height = BigInt(1);

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(invalidBlock);

      // 15. should invalidate 11th block
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(11),
      ).to.emit(challenge, "InvalidHeader");
    });
  });

  describe("toggleHeaderChallenge", function () {
    it("toggleHeaderChallenge should be failed without owner", async function () {
      await expect(
        challenge
          .connect(otherAccount)
          .getFunction("toggleHeaderChallenge")
          .send(false),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });

    it("toggleHeaderChallenge should be correct", async function () {
      await challenge.getFunction("toggleHeaderChallenge").send(true);
      expect(await challenge.isHeaderChallengeEnabled()).to.equal(true);

      await challenge.getFunction("toggleHeaderChallenge").send(false);
      expect(await challenge.isHeaderChallengeEnabled()).to.equal(false);
    });

    it("challenge header should revert when disabled", async function () {
      await challenge.getFunction("toggleHeaderChallenge").send(false);

      const validHeader: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1),
      ).to.be.revertedWith("header challenge is disabled");
    });

    it("challenge header should not revert when enabled", async function () {
      await challenge.getFunction("toggleHeaderChallenge").send(true);

      const validHeader: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1),
      ).to.be.revertedWith("header is valid");
    });
  });

  describe("setMaxBundleSize", function () {
    it("setMaxBundleSize should be failed without owner", async function () {
      await expect(
        challenge.connect(otherAccount).getFunction("setMaxBundleSize").send(1),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });

    it("setMaxBundleSize should be correct", async function () {
      await challenge.getFunction("setMaxBundleSize").send(1);
      expect(await challenge.maxBundleSize()).to.equal(1);

      await challenge.getFunction("setMaxBundleSize").send(2);
      expect(await challenge.maxBundleSize()).to.equal(2);
    });
  });

  describe("_isHeaderValid", function () {
    it("test should fail with invalid epoch", async function () {
      const header: Header = {
        epoch: BigInt(0), // epoch same as genesis
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      // expect event emit InvalidHeader reason to be "Invalid Epoch"
      await expect(
        await challengeTest.isHeaderValid(
          header,
          await canonicalStateChain.calculateHeaderHash(header),
          1,
        ),
      )
        .to.emit(challengeTest, "InvalidHeader")
        .withArgs(
          header.epoch,
          await canonicalStateChain.calculateHeaderHash(header),
          0, // Invalid Epoch
        );
    });

    it("test should fail with invalid previous hash", async function () {
      const header: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: ethers.ZeroHash, // invalid previous hash
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1n, shareStart: 1n, shareLen: 1n }],
      };

      // expect event emit InvalidHeader reason to be "Invalid Epoch"
      await expect(
        await challengeTest.isHeaderValid(
          header,
          await canonicalStateChain.calculateHeaderHash(header),
          1,
        ),
      )
        .to.emit(challengeTest, "InvalidHeader")
        .withArgs(
          header.epoch,
          await canonicalStateChain.calculateHeaderHash(header),
          2, // Invalid Previous Hash
        );
    });

    it("test should fail if bundle size > max bundle size", async function () {
      await challengeTest.getFunction("setMaxBundleSize").send(0);

      const header: Header = {
        epoch: BigInt(1),
        l2Height: toBigInt(genesisHeader.l2Height) + BigInt(1),
        prevHash: genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [
          { height: 1n, shareStart: 1n, shareLen: 1n },
          { height: 2n, shareStart: 1n, shareLen: 1n },
        ],
      };

      // expect event emit InvalidHeader reason to be "Invalid Epoch"
      await expect(
        await challengeTest.isHeaderValid(
          header,
          await canonicalStateChain.calculateHeaderHash(header),
          1,
        ),
      )
        .to.emit(challengeTest, "InvalidHeader")
        .withArgs(
          header.epoch,
          await canonicalStateChain.calculateHeaderHash(header),
          3, // Invalid Bundle Size
        );
    });
  });
});
