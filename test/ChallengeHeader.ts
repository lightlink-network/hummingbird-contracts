import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { setupCanonicalStateChain } from "./lib/chain";

describe("ChallengeHeader", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;
  let otherAccount: HardhatEthersSigner;
  let challengeOwner: HardhatEthersSigner;

  let genesisHash: string;
  let canonicalStateChain: Contract;
  let challenge: Contract;

  beforeEach(async function () {
    [owner, publisher, otherAccount, challengeOwner] =
      await ethers.getSigners();

    // Setup canonical state chain
    const createStateChain = await setupCanonicalStateChain(owner);
    canonicalStateChain = createStateChain.canonicalStateChain as any;
    genesisHash = createStateChain.genesisHash;

    const challengeFactory: any = await ethers.getContractFactory("Challenge");
    challenge = await challengeFactory.deploy(
      ethers.ZeroAddress,
      await canonicalStateChain.getAddress(),
      ethers.ZeroAddress,
      ethers.ZeroAddress
    );
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
  });
});
