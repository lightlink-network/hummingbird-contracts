import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { setupCanonicalStateChain } from "./lib/chain";
import { CanonicalStateChain } from "../typechain-types";
import { Header, hashHeader } from "./lib/header";

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
  });

  describe("Deployment", function () {
    it("Should set the right publisher", async function () {
      expect(await canonicalStateChain.publisher()).to.equal(publisher.address);
    });

    it("Should have the correct genesis hash", async function () {
      expect(await canonicalStateChain.chain(0)).to.equal(_chain.genesisHash);
    });
  });

  describe("pushBlock", function () {
    it("Should revert if called by an account other than the publisher", async function () {
      const header: Header = {
        epoch: 1n,
        l2Height: 1n,
        prevHash: _chain.genesisHash,
        celestiaHeight: 1n,
        celestiaShareStart: 1n,
        celestiaShareLen: 1n,
      };

      await expect(
        canonicalStateChain.connect(otherAccount).pushBlock(header),
      ).to.be.revertedWith("only publisher can add blocks");
    });

    it("Should add a block when called by the publisher", async function () {
      const header: Header = {
        epoch: 1n,
        l2Height: 1n,
        prevHash: _chain.genesisHash,
        celestiaHeight: 1n,
        celestiaShareStart: 1n,
        celestiaShareLen: 1n,
      };
      const hash = hashHeader(header);

      await expect(canonicalStateChain.connect(publisher).pushBlock(header))
        .to.emit(canonicalStateChain, "BlockAdded")
        .withArgs(1);

      expect(await canonicalStateChain.chain(1)).to.equal(hash);
    });
  });

  describe("rollback", function () {
    it("Should revert if called by an account other than the challenge contract", async function () {
      await expect(canonicalStateChain.rollback(0)).to.be.revertedWith(
        "only challenge contract can rollback chain",
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
          .setChallengeContract(otherAccount.address),
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
      await canonicalStateChain.connect(publisher).pushBlock({
        epoch: 1,
        l2Height: 1,
        prevHash: _chain.genesisHash,
        celestiaHeight: 1,
        celestiaShareStart: 1,
        celestiaShareLen: 1,
      });

      await expect(
        canonicalStateChain.connect(publisher).rollback(0),
      ).to.be.revertedWith("only challenge contract can rollback chain");

      await canonicalStateChain.setChallengeContract(challengeContract.address);

      await expect(
        canonicalStateChain.connect(challengeContract).rollback(0),
      ).to.emit(canonicalStateChain, "RolledBack");
    });
  });
});
