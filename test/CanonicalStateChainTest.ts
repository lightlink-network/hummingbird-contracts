import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { pushRandomHeader, setupCanonicalStateChain } from "./lib/chain";
import { CanonicalStateChain } from "../typechain-types";

describe("CanonicalStateChain", function () {
  let canonicalStateChain: CanonicalStateChain;
  let owner: HardhatEthersSigner,
    publisher: HardhatEthersSigner,
    otherAccount: HardhatEthersSigner,
    challengeContract: HardhatEthersSigner,
    _chain: any;

  beforeEach(async function () {
    [owner, publisher, otherAccount, challengeContract] =
      await ethers.getSigners();

    _chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = _chain.canonicalStateChain;

    await canonicalStateChain.setChallengeContract(challengeContract.address);
  });

  describe("Deployment", function () {
    it("Should set the right publisher", async function () {
      expect(await canonicalStateChain.publisher()).to.equal(publisher.address);
    });

    it("Should have the correct genesis hash", async function () {
      expect(await canonicalStateChain.chain(0)).to.equal(_chain.genesisHash);
    });

    it("maxPointers var should be 7 by default", async function () {
      expect(await canonicalStateChain.maxPointers()).to.eq(7);
    });

    it("Should not be allowed to initialize twice", async function () {
      let genesisHeader: CanonicalStateChain.HeaderStruct = {
        epoch: BigInt(0),
        l2Height: BigInt(1),
        prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [],
      };

      await expect(
        canonicalStateChain
          .connect(otherAccount)
          .getFunction("initialize")
          .send(owner, genesisHeader),
      ).to.be.revertedWithCustomError(
        canonicalStateChain,
        "InvalidInitialization",
      );
    });
  });

  describe("pushBlock", function () {
    it("Should revert if called by an account other than the publisher", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 1,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        shareRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1, shareStart: 1, shareLen: 1 }],
      };

      await expect(
        canonicalStateChain
          .connect(otherAccount)
          .getFunction("pushBlock")
          .send(header),
      ).to.be.revertedWith("only publisher can add blocks");
    });

    it("Should add a block when called by the publisher", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 1,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        shareRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1, shareStart: 1, shareLen: 1 }],
      };

      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header),
      )
        .to.emit(canonicalStateChain, "BlockAdded")
        .withArgs(1);
    });

    it("Should revert as epoch is too low", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 0,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1, shareStart: 1, shareLen: 1 }],
      };

      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header),
      ).to.be.revertedWith("epoch must be greater than previous epoch");
    });

    it("Should revert as the prevHash is not correct", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 1,
        l2Height: 1,
        prevHash:
          "0x55eb99d77b0e1ed261c0a8d11f026f811b8af01455a2b45189bcc87b93dfbbb7",
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [{ height: 1, shareStart: 1, shareLen: 1 }],
      };

      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header),
      ).to.be.revertedWith("prevHash must be the previous block hash");
    });

    it("Should revert as < 1 celestia pointers", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 1,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [],
      };

      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header),
      ).to.be.revertedWith("block must have atleast one celestia pointer");
    });

    it("Should revert as > 7 (maxPointers) celestia pointers", async function () {
      const header: CanonicalStateChain.HeaderStruct = {
        epoch: 1,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaPointers: [
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
          { height: 1, shareStart: 1, shareLen: 1 },
        ],
      };

      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("pushBlock")
          .send(header),
      ).to.be.revertedWith("block has too many celestia pointers");
    });
  });

  describe("rollback", function () {
    it("Should revert if called by an account other than the challenge contract", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);
      await pushRandomHeader(publisher, canonicalStateChain);

      await expect(canonicalStateChain.rollback(1)).to.be.revertedWith(
        "only challenge contract can rollback chain",
      );
    });

    it("Should emit a RolledBack event when called by the challenge address", async function () {
      await pushRandomHeader(publisher, canonicalStateChain);
      await pushRandomHeader(publisher, canonicalStateChain);

      await expect(
        canonicalStateChain
          .connect(challengeContract)
          .getFunction("rollback")
          .send(1),
      ).to.emit(canonicalStateChain, "RolledBack");
    });

    it("Should not be able to rollback if no blocks added", async function () {
      await expect(
        canonicalStateChain.connect(challengeContract).rollback(0),
      ).to.be.revertedWith("block number must be less than chain head");
    });

    it("should rollback to the correct block", async function () {
      const [goodBlockHash] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );
      await pushRandomHeader(publisher, canonicalStateChain);

      await canonicalStateChain.connect(challengeContract).rollback(1);
      expect(await canonicalStateChain.chainHead()).to.equal(
        1,
        "chain head should be 1",
      );

      const headHash = await canonicalStateChain.chain(
        await canonicalStateChain.chainHead(),
      );
      expect(headHash).to.equal(
        goodBlockHash,
        "chain head should be the good block hash",
      );
    });
  });

  describe("setPublisher", function () {
    it("Should revert if called by an account other than the owner", async function () {
      await expect(
        canonicalStateChain
          .connect(otherAccount)
          .getFunction("setPublisher")
          .send(otherAccount.address),
      ).to.be.reverted;
    });

    it("Should set a new publisher when called by the owner", async function () {
      await expect(canonicalStateChain.setPublisher(otherAccount.address))
        .to.emit(canonicalStateChain, "PublisherChanged")
        .withArgs(otherAccount.address);
    });
  });

  describe("setChallengeContract", function () {
    it("Should revert if called by an account other than the owner", async function () {
      await expect(
        canonicalStateChain
          .connect(otherAccount)
          .getFunction("setChallengeContract")
          .send(otherAccount.address),
      ).to.be.reverted;
    });

    it("Should set a new challenge contract when called by the owner", async function () {
      await expect(
        canonicalStateChain.setChallengeContract(challengeContract.address),
      )
        .to.emit(canonicalStateChain, "ChallengeChanged")
        .withArgs(challengeContract.address);
    });

    it("challenge contract should be able to rollback", async function () {
      // add a header
      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send({
          epoch: 1,
          l2Height: 1,
          prevHash: _chain.genesisHash,
          stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
          shareRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
          celestiaPointers: [{ height: 1, shareStart: 1, shareLen: 1 }],
        });

      await expect(
        canonicalStateChain.connect(publisher).getFunction("rollback").send(0),
      ).to.be.revertedWith("only challenge contract can rollback chain");

      await canonicalStateChain.setChallengeContract(challengeContract.address);

      await expect(
        canonicalStateChain
          .connect(challengeContract)
          .getFunction("rollback")
          .send(0),
      ).to.emit(canonicalStateChain, "RolledBack");
    });
  });
  describe("setMaxPointers", function () {
    it("setMaxPointers should update maxPointers var", async function () {
      expect(
        await canonicalStateChain
          .connect(owner)
          .getFunction("setMaxPointers")
          .send(1),
      );
      expect(await canonicalStateChain.maxPointers()).to.equal(1);
    });

    it("setMaxPointers should revert for non owner", async function () {
      await expect(
        canonicalStateChain
          .connect(publisher)
          .getFunction("setMaxPointers")
          .send(1),
      ).to.be.revertedWithCustomError(
        canonicalStateChain,
        "OwnableUnauthorizedAccount",
      );
    });
  });
});
