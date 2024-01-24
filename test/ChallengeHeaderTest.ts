import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { setupCanonicalStateChain } from "./lib/chain";
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

    it("cannot invalidate a valid header", async function () {
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

    it("invalid header with incorrect L2 height", async function () {
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

    it("invalidate header even after a newer block", async function () {
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

      const validHeader: Header = {
        epoch: BigInt(5),
        l2Height: invalidHeader.l2Height + BigInt(1),
        prevHash: hashHeader(invalidHeader),
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
        .send(invalidHeader);

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
  });
});
