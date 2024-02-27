import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  CanonicalStateChain,
  ChainOracle,
  Challenge,
} from "../typechain-types";
import { makeNextBlock, setupCanonicalStateChain } from "./lib/chain";
import { Header } from "./lib/header";
import { MOCK_DATA } from "./mock/mock_challengeL2Header";

describe("ChallengeL2Header", function () {
  let chain: CanonicalStateChain;
  let genesisHash: string;
  let genesisHeader: Header;
  let chainOracle: ChainOracle;
  let challengeFee = ethers.parseEther("1");
  let challenge: Challenge;
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;

  beforeEach(async () => {
    [owner, publisher] = await ethers.getSigners();

    // 1. setup canonical state chain
    const csc = await setupCanonicalStateChain(owner, publisher.address);
    chain = csc.canonicalStateChain;
    genesisHash = csc.genesisHash;
    genesisHeader = csc.genesisHeader;

    // 2. setup chain oracle
    //  - 2.A deploy mock da oracle
    const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
    const mockDaOracle = (await _MockDaOracle.deploy()) as any;
    //  - 2.B deploy rlp reader
    const RLPReaderFactory = await ethers.getContractFactory("RLPReader");
    const rlpReader = await RLPReaderFactory.deploy();
    //  - 2.C deploy chain oracle implementation
    const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
    const chainOracleFactory = await ethers.getContractFactory("ChainOracle");
    const chainOracleImplementation = await chainOracleFactory.deploy();
    //  - 2.D deploy chain oracle proxy
    const chainOracleProxy = await proxyFactory.deploy(
      await chainOracleImplementation.getAddress(),
      chainOracleImplementation.interface.encodeFunctionData("initialize", [
        await chain.getAddress(),
        await mockDaOracle.getAddress(),
        await rlpReader.getAddress(),
      ]),
    );

    chainOracle = chainOracleFactory.attach(
      await chainOracleProxy.getAddress(),
    ) as any;

    // 3. setup challenge
    const challengeFactory: any = await ethers.getContractFactory("Challenge");
    const challengeImplementation = await challengeFactory.deploy();
    const challengeProxy = await proxyFactory.deploy(
      await challengeImplementation.getAddress(),
      challengeImplementation.interface.encodeFunctionData("initialize", [
        ethers.ZeroAddress, // treasury
        await chain.getAddress(),
        await mockDaOracle.getAddress(),
        ethers.ZeroAddress, // mipsChallenge
        await chainOracle.getAddress(),
      ]),
    );

    challenge = challengeFactory.attach(
      await challengeProxy.getAddress(),
    ) as any;
    await challenge.connect(owner).setChallengeFee(challengeFee);

    // 4. push next block
    const nextBlock = { ...MOCK_DATA.rollupHeaders[0] };
    nextBlock.prevHash = genesisHash;
    await chain.connect(publisher).pushBlock(nextBlock);
  });

  it("should deploy", async function () {
    expect(chainOracle).to.not.be.undefined;
    expect(challenge).to.not.be.undefined;
  });

  describe("challengeL2Header", function () {
    it("should not able to challenge if fee is not paid", async function () {
      expect(
        challenge
          .connect(owner)
          .challengeL2Header(1, MOCK_DATA.l2Headers[0].number),
      ).to.be.revertedWith("Challenge: fee not paid");
    });

    it("should able to challenge if fee is paid", async function () {
      const l2Header = MOCK_DATA.l2Headers[0];
      const prevNumber = BigInt(l2Header.number) - 1n;

      expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.emit(challenge, "L2HeaderChallenge");

      const rBlockHash = await chain.chain(1);
      const hash = await challenge.l2HeaderChallengeHash(
        rBlockHash,
        l2Header.number,
      );
      const challengeData = await challenge.l2HeaderChallenges(hash);

      expect(challengeData.header.rblock).to.be.equal(rBlockHash);
      expect(challengeData.header.number.toString()).to.be.equal(
        l2Header.number.toString(),
      );

      expect(challengeData.prevHeader.rblock).to.be.equal(rBlockHash);
      expect(challengeData.prevHeader.number.toString()).to.be.equal(
        prevNumber.toString(),
      );
    });
  });

  describe("defendL2Header", function () {
    it("should not be able to defend non-existing challenge", async function () {
      const l2Header = MOCK_DATA.l2Headers[1];
      const l2HeaderHash = MOCK_DATA.l2HeaderHashes[1];
      const l2PrevHeaderHash = MOCK_DATA.l2HeaderHashes[0];

      const challengeHash = await challenge.l2HeaderChallengeHash(
        await chain.chain(1),
        l2Header.number,
      );

      await expect(
        challenge
          .connect(owner)
          .defendL2Header(challengeHash, l2HeaderHash, l2PrevHeaderHash),
      ).to.be.revertedWith("challenge is not in the correct state");
    });

    it("should not be able to defend if data not pre-submitted to chainOracle", async function () {
      const l2Header = MOCK_DATA.l2Headers[1];
      const l2HeaderHash = MOCK_DATA.l2HeaderHashes[1];
      const l2PrevHeaderHash = MOCK_DATA.l2HeaderHashes[0];

      const challengeHash = await challenge.l2HeaderChallengeHash(
        await chain.chain(1),
        l2Header.number,
      );

      // challenge
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      await expect(
        challenge
          .connect(owner)
          .defendL2Header(challengeHash, l2HeaderHash, l2PrevHeaderHash),
      ).to.be.revertedWith("l2 header not loaded for the given rblock");
    });

    it("should be have same hash", async function () {
      const header = MOCK_DATA.l2Headers[0];
      const headerHash = MOCK_DATA.l2HeaderHashes[0];

      expect(await chainOracle.hashHeader(header)).to.be.equal(headerHash);
    });

    it("should be able to extra data from shares", async function () {
      const data = await chainOracle
        .connect(owner)
        .extractData(MOCK_DATA.shareProofs[0].data, MOCK_DATA.shareRanges[0]);

      console.log(MOCK_DATA.shareRanges[0]);
      console.log(data);
    });

    it("should be able to load header shares to chainOracle", async function () {
      const rblockHash = await chain.chain(1);
      const header = MOCK_DATA.l2Headers[0];
      const headerHash = MOCK_DATA.l2HeaderHashes[0];
      const shareProof = MOCK_DATA.shareProofs[0];
      const shareRanges = MOCK_DATA.shareRanges[0];

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, shareProof),
      ).to.not.be.reverted;

      const shareKey = await chainOracle.ShareKey(rblockHash, shareProof.data);
      const shares0 = await chainOracle.shares(shareKey, 0);
      console.log("shares0", shares0);
      expect(shares0).to.be.equal(ethers.hexlify(shareProof.data[0]));
      const shares1 = await chainOracle.shares(shareKey, 1);
      expect(shares1).to.be.equal(ethers.hexlify(shareProof.data[1]));

      await expect(
        chainOracle.connect(owner).provideHeader(shareKey, shareRanges),
      ).to.not.be.reverted;

      // now load the header
      const loadedHeader = await chainOracle.connect(owner).headers(headerHash);
      console.log(loadedHeader);
      expect(loadedHeader.number.toString()).to.be.equal(
        header.number.toString(),
      );
      expect(loadedHeader.stateRoot).to.be.equal(header.stateRoot);
    });
  });
});
