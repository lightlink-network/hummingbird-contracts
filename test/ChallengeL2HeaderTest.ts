import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  CanonicalStateChain,
  ChainOracle,
  ChainOracle__factory,
  Challenge,
  Challenge__factory,
} from "../typechain-types";
import { makeNextBlock, setupCanonicalStateChain } from "./lib/chain";
import { challengeL2HeaderMockData as MOCK_DATA } from "./mock/mock_challengeL2Header";
import { proxyDeployAndInitialize } from "../scripts/lib/deploy";
import { provideHeader } from "./lib/oracle";
import { time } from "@nomicfoundation/hardhat-network-helpers";

type Header = CanonicalStateChain.HeaderStruct;

const CURR_HEADER = 1;
const PREV_HEADER = 0;

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
    // const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
    // const chainOracleFactory = await ethers.getContractFactory("ChainOracle");
    // const chainOracleImplementation = await chainOracleFactory.deploy();
    const chainOracleDeployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("ChainOracle"),
      [
        await chain.getAddress(),
        await mockDaOracle.getAddress(),
        await rlpReader.getAddress(),
      ],
    );

    chainOracle = ChainOracle__factory.connect(
      chainOracleDeployment.address,
      owner,
    );

    // 3. setup challenge
    const challengeDeployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("Challenge"),
      [
        await chain.getAddress(),
        await mockDaOracle.getAddress(),
        await chainOracle.getAddress(),
      ],
    );

    challenge = Challenge__factory.connect(challengeDeployment.address, owner);
    await challenge.connect(owner).setChallengeFee(challengeFee);
    await chain.connect(owner).setChallengeContract(challengeDeployment.address);

    // 4. push next block
    const nextBlock = { ...MOCK_DATA[0].rollupHeader };
    nextBlock.prevHash = genesisHash;
    await chain.connect(publisher).pushBlock(nextBlock);

    // set isL2HeaderChallengeEnabled to true
    await challenge.getFunction("toggleL2HeaderChallenge").send(true);
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
          .challengeL2Header(
            1,
            MOCK_DATA[0].headers[CURR_HEADER].header.number,
          ),
      ).to.be.revertedWith("challenge fee not paid");
    });

    it("should able to challenge if fee is paid", async function () {
      const l2Header = MOCK_DATA[0].headers[1].header;
      const prevNumber = BigInt(l2Header.number) - 1n;

      await expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.emit(challenge, "L2HeaderChallengeUpdate");

      const rBlockHash = await chain.chain(1);
      const hash = await challenge.l2HeaderChallengeHash(
        rBlockHash,
        l2Header.number,
      );
      const challengeData = await challenge.l2HeaderChallenges(hash);

      expect(challengeData.header.rblock).to.be.equal(
        rBlockHash,
        "invalid rblock hash on challenge data",
      );
      expect(challengeData.header.number).to.be.equal(
        BigInt(l2Header.number.toString()),
        "invalid number on challenge data",
      );

      expect(challengeData.prevHeader.rblock).to.be.equal(
        rBlockHash,
        "invalid prev rblock hash on challenge data",
      );
      expect(challengeData.prevHeader.number).to.be.equal(
        prevNumber,
        "invalid prev number on challenge data",
      );
    });

    it("should not able to challenge if rblockNum not canonical (rblock not in chain yet)", async function () {
      const l2Header = MOCK_DATA[0].headers[1].header;

      await expect(
        challenge.connect(owner).challengeL2Header(2, l2Header.number, {
          value: challengeFee,
        }),
      ).to.be.revertedWith("block not in the chain yet");
    });
    it("should not able to challenge if not in challenge window", async function () {
      const l2Header = MOCK_DATA[0].headers[1].header;

      await challenge
        .connect(owner)
        .connect(owner)
        .setChallengeWindow(12 * 60 * 60);

      // advance time by 12 hrs & 1 second and mine a new block
      await time.increase(12 * 60 * 60 + 1);

      await expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.be.revertedWith("block is too old to challenge");
    });

    it("should not able to challenge if challenge already exists", async function () {
      const l2Header = MOCK_DATA[0].headers[1].header;

      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      await expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.be.revertedWith("challenge already exists");
    });

    it("should not able to challenge if l2 header not within rblock bundle range", async function () {
      await expect(
        challenge.connect(owner).challengeL2Header(1, 1, {
          value: challengeFee,
        }),
      ).to.be.revertedWith("L2 header must be within the rblock bundle range");
    });

    it("should not able to challenge the first l2 header in the first bundle", async function () {
      const firstRblock = await chain.getHeaderByNum(0);
      await expect(
        challenge
          .connect(owner)
          .challengeL2Header(1, firstRblock[1] + BigInt(1), {
            value: challengeFee,
          }),
      ).to.be.revertedWith(
        "Cannot challenge the first L2 header in the first rblock",
      );
    });
  });

  describe("defendL2Header", function () {
    it("should not be able to defend non-existing challenge", async function () {
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      const l2HeaderHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;
      const l2PrevHeaderHash =
        MOCK_DATA[0].headers[CURR_HEADER].header.parentHash;

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

    it("should not be able to defend if data of previous l2 header not pre-submitted to chainOracle", async function () {
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      const l2HeaderHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;
      const l2PrevHeaderHash =
        MOCK_DATA[0].headers[CURR_HEADER].header.parentHash;

      const rblockHash = await chain.chain(1);
      const shareProof = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const shareRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;

      // load shares for the current header
      await expect(
        chainOracle
          .connect(owner)
          .provideShares(rblockHash, 0, shareProof),
      ).to.not.be.reverted;

      // load header for the current header
      await expect(
        chainOracle
          .connect(owner)
          .provideHeader(
            await chainOracle.ShareKey(rblockHash, shareProof.data),
            shareRanges,
          ),
      ).to.not.be.reverted;

      // pre compute the challenge hash
      const challengeHash = await challenge.l2HeaderChallengeHash(
        await chain.chain(1),
        l2Header.number,
      );

      // challenge the current header
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      // defend the current header, but the previous header is not loaded
      await expect(
        challenge
          .connect(owner)
          .defendL2Header(challengeHash, l2HeaderHash, l2PrevHeaderHash),
      ).to.be.revertedWith(
        "previous l2 header not loaded for the given rblock",
      );
    });

    it("should not be able to defend if data not pre-submitted to chainOracle", async function () {
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      const l2HeaderHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;
      const l2PrevHeaderHash =
        MOCK_DATA[0].headers[CURR_HEADER].header.parentHash;

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

    it("should have same hash", async function () {
      const header = MOCK_DATA[0].headers[CURR_HEADER].header;
      const headerHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;

      expect(await chainOracle.hashHeader(header)).to.be.equal(headerHash);
    });

    it("should be able to extract data from shares", async function () {
      MOCK_DATA[0].headers[CURR_HEADER].shareProofs.data;
      expect(
        chainOracle
          .connect(owner)
          .extractData(
            MOCK_DATA[0].headers[CURR_HEADER].shareProofs.data,
            MOCK_DATA[0].headers[CURR_HEADER].shareRanges,
          ),
      ).to.not.be.reverted;
    });

    it("should be able to load header shares to chainOracle", async function () {
      const rblockHash = await chain.chain(1);
      const header = MOCK_DATA[0].headers[CURR_HEADER].header;
      const headerHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;
      const shareProof = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const shareRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;

      await expect(
        chainOracle
          .connect(owner)
          .provideShares(rblockHash, 0, shareProof),
      ).to.not.be.reverted;

      const shareKey = await chainOracle.ShareKey(rblockHash, shareProof.data);
      const shares0 = await chainOracle.shares(shareKey, 0);
      expect(shares0).to.be.equal(ethers.hexlify(shareProof.data[0]));
      const shares1 = await chainOracle.shares(shareKey, 1);
      expect(shares1).to.be.equal(ethers.hexlify(shareProof.data[1]));

      await expect(
        chainOracle.connect(owner).provideHeader(shareKey, shareRanges),
      ).to.not.be.reverted;

      // now load the header
      const loadedHeader = await chainOracle.connect(owner).headers(headerHash);
      expect(loadedHeader.number).to.be.equal(BigInt(header.number.toString()));
      expect(loadedHeader.stateRoot).to.be.equal(header.stateRoot);
    });

    it("should be able to defend", async function () {
      const prevHeaderShares = MOCK_DATA[0].headers[PREV_HEADER].shareProofs;
      const prevHeaderRanges = MOCK_DATA[0].headers[PREV_HEADER].shareRanges;
      const headerShares = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;
      const rblockHash = await chain.chain(1);

      // load prev header
      await expect(
        chainOracle
          .connect(owner)
          .provideShares(rblockHash, 0, prevHeaderShares),
      ).to.not.be.reverted;
      const prevShareKey = await chainOracle.ShareKey(
        rblockHash,
        prevHeaderShares.data,
      );
      await expect(
        chainOracle
          .connect(owner)
          .provideHeader(prevShareKey, prevHeaderRanges),
      ).to.not.be.reverted;

      // load header
      await expect(
        chainOracle
          .connect(owner)
          .provideShares(rblockHash, 0, headerShares),
      ).to.not.be.reverted;
      const shareKey = await chainOracle.ShareKey(
        rblockHash,
        headerShares.data,
      );
      await expect(
        chainOracle.connect(owner).provideHeader(shareKey, headerRanges),
      ).to.not.be.reverted;

      // challenge
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      const challengeHash = await challenge.l2HeaderChallengeHash(
        rblockHash,
        l2Header.number,
      );

      // defend
      const l2HeaderHash = MOCK_DATA[0].headers[CURR_HEADER].headerHash;
      const l2PrevHeaderHash =
        MOCK_DATA[0].headers[CURR_HEADER].header.parentHash;
      await expect(challenge
        .connect(publisher)
        .defendL2Header(challengeHash, l2HeaderHash, l2PrevHeaderHash)
      ).to.not.be.reverted;

      const challengeData = await challenge.l2HeaderChallenges(challengeHash);
      expect(challengeData[5], "incorrect challenge state").to.be.equal(3n);

      //check the chain was not rolled back
      const head = await chain.chainHead();
      expect(head, "chain rolled back").to.be.equal(1n);

      // check the challenge cannot be settled
      await expect(
        challenge.connect(owner).settleL2HeaderChallenge(challengeHash),
      ).to.be.revertedWith("challenge is not in the correct state");

      // check the challenge cannot be defended again
      await expect(
        challenge
          .connect(publisher)
          .defendL2Header(challengeHash, l2HeaderHash, l2PrevHeaderHash),
      ).to.be.revertedWith("challenge is not in the correct state");
    });
  });

  describe("settleL2HeaderChallenge", function () {
    it("happy path", async function () {
      const prevHeaderShares = MOCK_DATA[0].headers[PREV_HEADER].shareProofs;
      const prevHeaderRanges = MOCK_DATA[0].headers[PREV_HEADER].shareRanges;
      const headerShares = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;
      const rblockHash = await chain.chain(1);

      // load prev header to chainOracle
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        prevHeaderShares,
        prevHeaderRanges,
      );

      // load current header to chainOracle
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        headerShares,
        headerRanges,
      );

      // set challenge period
      await challenge.connect(owner).setChallengePeriod(12 * 60 * 60);

      // challenge the current header
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      const challengeHash = await challenge.l2HeaderChallengeHash(
        rblockHash,
        l2Header.number,
      );;

      // advance time by 12 hrs & 1 second and mine a new block
      await time.increase(12 * 60 * 60 + 1);

      // settle now that the challenge period has ended
      await expect(challenge.connect(owner).settleL2HeaderChallenge(challengeHash), "challenge failed").to.not.be.reverted;


      const challengeData = await challenge.l2HeaderChallenges(challengeHash);
      expect(challengeData[5], "incorrect challenge hash").to.be.equal(2);

      //check the chain rolled back
      const head = await chain.chainHead();
      expect(head, "chain did not rollback").to.be.equal(0n);

      // claim the challenge
      await expect(challenge.connect(owner).claimL2HeaderChallengeReward(challengeHash)).to.not.be.reverted;

      // check the reward can only be claimed once
      await expect(challenge.connect(owner).claimL2HeaderChallengeReward(challengeHash)).to.be.revertedWith("challenge has already been claimed");
    });

    it("should revert if challenge period has not ended", async function () {
      const prevHeaderShares = MOCK_DATA[0].headers[PREV_HEADER].shareProofs;
      const prevHeaderRanges = MOCK_DATA[0].headers[PREV_HEADER].shareRanges;
      const headerShares = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;
      const rblockHash = await chain.chain(1);

      // load prev header
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        prevHeaderShares,
        prevHeaderRanges,
      );

      // load current header
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        headerShares,
        headerRanges,
      );

      // challenge
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });

      const challengeHash = await challenge.l2HeaderChallengeHash(
        rblockHash,
        l2Header.number,
      );

      // settle
      await expect(
        challenge.connect(owner).settleL2HeaderChallenge(challengeHash),
      ).to.be.revertedWith("challenge period has not ended");
    });

    it("should revert if challenge is not in the correct state", async function () {
      const prevHeaderShares = MOCK_DATA[0].headers[PREV_HEADER].shareProofs;
      const prevHeaderRanges = MOCK_DATA[0].headers[PREV_HEADER].shareRanges;
      const headerShares = MOCK_DATA[0].headers[CURR_HEADER].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[CURR_HEADER].shareRanges;
      const rblockHash = await chain.chain(1);

      
      // load prev header to chainOracle
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        prevHeaderShares,
        prevHeaderRanges,
      );
      
      // load current header to chainOracle
      await provideHeader(
        chainOracle,
        rblockHash,
        0,
        headerShares,
        headerRanges,
      );
      
      // reduce challenge period
      await challenge.connect(owner).setChallengePeriod(12 * 60 * 60);

      // challenge the current header
      const l2Header = MOCK_DATA[0].headers[CURR_HEADER].header;
      await challenge.connect(owner).challengeL2Header(1, l2Header.number, {
        value: challengeFee,
      });
      
      const challengeHash = await challenge.l2HeaderChallengeHash(
        rblockHash,
        l2Header.number,
      );

      // advance time by 12 hrs & 1 second and mine a new block
      await time.increase(12 * 60 * 60 + 1);

      // settle once
      await expect(challenge.connect(owner).settleL2HeaderChallenge(challengeHash)).to.not.be.reverted;

      // settle
      await expect(
        challenge.connect(owner).settleL2HeaderChallenge(challengeHash),
      ).to.be.revertedWith("challenge is not in the correct state");

      const challengeData = await challenge.l2HeaderChallenges(challengeHash);
      expect(challengeData[5]).to.be.equal(2);
    });
  });

  describe("toggleL2HeaderChallenge", function () {
    it("toggleL2HeaderChallenge should be failed without owner", async function () {
      await expect(
        challenge
          .connect(publisher)
          .getFunction("toggleL2HeaderChallenge")
          .send(false),
      ).to.be.revertedWithCustomError(challenge, "OwnableUnauthorizedAccount");
    });

    it("toggleL2HeaderChallenge should be correct", async function () {
      await challenge.getFunction("toggleL2HeaderChallenge").send(true);
      expect(await challenge.isL2HeaderChallengeEnabled()).to.equal(true);

      await challenge.getFunction("toggleL2HeaderChallenge").send(false);
      expect(await challenge.isL2HeaderChallengeEnabled()).to.equal(false);
    });

    it("challenge L2 header should revert when disabled", async function () {
      await challenge.getFunction("toggleL2HeaderChallenge").send(false);
      const l2Header = MOCK_DATA[0].headers[0].header;

      await expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.revertedWith("L2 header challenge is disabled");
    });

    it("challenge L2 header should not revert when enabled", async function () {
      await challenge.getFunction("toggleL2HeaderChallenge").send(true);
      const l2Header = MOCK_DATA[0].headers[0].header;

      await expect(
        challenge.connect(owner).challengeL2Header(1, l2Header.number, {
          value: challengeFee,
        }),
      ).to.emit(challenge, "L2HeaderChallengeUpdate");
    });
  });
});
