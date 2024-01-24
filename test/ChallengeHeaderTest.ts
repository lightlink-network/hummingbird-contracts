import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { makeNextBlock, setupCanonicalStateChain } from "./lib/chain";
import { Header, hashHeader } from "./lib/header";

describe("ChallengeHeader", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;
  let otherAccount: HardhatEthersSigner;
  let challengeOwner: HardhatEthersSigner;

  let genesisHeader: Header;
  let genesisHash: string;
  let canonicalStateChain: Contract;
  let challenge: Contract;

  beforeEach(async function () {
    [owner, publisher, otherAccount, challengeOwner] =
      await ethers.getSigners();

    // Setup canonical state chain
    const _chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = _chain.canonicalStateChain as any;
    genesisHash = _chain.genesisHash;
    genesisHeader = _chain.genesisHeader;

    const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
    const challengeFactory: any = await ethers.getContractFactory("Challenge");
    const challengeImplementation = await challengeFactory.deploy();

    const proxy = await proxyFactory.deploy(
      await challengeImplementation.getAddress(),
      challengeImplementation.interface.encodeFunctionData("initialize", [
        ethers.ZeroAddress,
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ])
    );

    challenge = challengeFactory.attach(await proxy.getAddress());

    _chain.canonicalStateChain
      .getFunction("setChallengeContract")
      .send(await challenge.getAddress());
  });

  describe("deployment", function () {
    it("should set the correct owner", async function () {
      expect(await challenge.owner()).to.equal(owner.address);
    });

    it("should set the correct canonical state chain", async function () {
      expect(await challenge.chain()).to.equal(
        await canonicalStateChain.getAddress()
      );
    });
  });

  describe("invalidateHeader", function () {
    it("should not allow challenging the genesis header", async function () {
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(0)
      ).to.be.revertedWith("cannot challenge genesis block");
    });

    it("should not allow challenging a header that is not canonical", async function () {
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(100)
      ).to.be.revertedWith("block not in the chain yet");
    });

    it("should not invalidate a valid header", async function () {
      const validHeader: Header = {
        epoch: BigInt(1),
        l2Height: genesisHeader.l2Height + BigInt(1),
        prevHash: genesisHash,
        txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaHeight: BigInt(5),
        celestiaShareStart: BigInt(1),
        celestiaShareLen: BigInt(1),
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(validHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1)
      ).to.be.revertedWith("header is valid");
    });

    it("should invalidate header with incorrect L2 height", async function () {
      const invalidHeader: Header = {
        epoch: BigInt(1),
        l2Height: genesisHeader.l2Height - BigInt(1),
        prevHash: genesisHash,
        txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaHeight: BigInt(1),
        celestiaShareStart: BigInt(1),
        celestiaShareLen: BigInt(1),
      };

      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(invalidHeader);

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1)
      ).to.emit(challenge, "InvalidHeader");
    });

    it("should invalidate header even after a newer block", async function () {
      const invalidHeader: Header = {
        epoch: BigInt(1),
        l2Height: genesisHeader.l2Height - BigInt(1),
        prevHash: genesisHash,
        txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
        celestiaHeight: BigInt(1),
        celestiaShareStart: BigInt(1),
        celestiaShareLen: BigInt(1),
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
          .send(2)
      ).to.be.revertedWith("header is valid");

      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(1)
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
        canonicalStateChain
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
          .send(9)
      ).to.be.revertedWith("header is valid");

      // 10. but should be able to invalidate 10th block
      await expect(
        challenge
          .connect(challengeOwner)
          .getFunction("invalidateHeader")
          .send(10)
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
          .send(10)
      ).to.be.revertedWith("header is valid");

      // 14. add the 11th invalid block
      const [invalidBlock] = await makeNextBlock(
        publisher,
        canonicalStateChain
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
          .send(11)
      ).to.emit(challenge, "InvalidHeader");
    });
  });
});
