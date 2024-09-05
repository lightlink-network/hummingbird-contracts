import { ethers } from "hardhat";
import { expect } from "chai";
import {
  CanonicalStateChain,
  ChainOracle,
  Challenge,
} from "../../../../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
// import { MOCK_DATA } from "./mock/mock_challengeL2Tx";
import { setupCanonicalStateChain } from "../../../lib/chain";
import { provideHeader, provideLegacyTx } from "../../../lib/oracle";

// TODO: re-enable this test after challenge l2 Tx is full implemented.
// describe("ChallengeL2Tx", function () {
//   let chain: CanonicalStateChain;
//   let genesisHash: string;
//   let chainOracle: ChainOracle;
//   let challengeFee = ethers.parseEther("1");
//   let challenge: Challenge;
//   let owner: HardhatEthersSigner;
//   let publisher: HardhatEthersSigner;

//   beforeEach(async () => {
//     [owner, publisher] = await ethers.getSigners();

//     // 1. setup canonical state chain
//     const csc = await setupCanonicalStateChain(owner, publisher.address);
//     chain = csc.canonicalStateChain;
//     genesisHash = csc.genesisHash;

//     // 2. setup chain oracle
//     //  - 2.A deploy mock da oracle
//     const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
//     const mockDaOracle = (await _MockDaOracle.deploy()) as any;
//     //  - 2.B deploy rlp reader
//     const RLPReaderFactory = await ethers.getContractFactory("RLPReader");
//     const rlpReader = await RLPReaderFactory.deploy();
//     //  - 2.C deploy chain oracle implementation
//     const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
//     const chainOracleFactory = await ethers.getContractFactory("ChainOracle");
//     const chainOracleImplementation = await chainOracleFactory.deploy();
//     //  - 2.D deploy chain oracle proxy
//     const chainOracleProxy = await proxyFactory.deploy(
//       await chainOracleImplementation.getAddress(),
//       chainOracleImplementation.interface.encodeFunctionData("initialize", [
//         await chain.getAddress(),
//         await mockDaOracle.getAddress(),
//         await rlpReader.getAddress(),
//       ]),
//     );

//     chainOracle = chainOracleFactory.attach(
//       await chainOracleProxy.getAddress(),
//     ) as any;

//     // 3. setup challenge
//     const challengeFactory: any = await ethers.getContractFactory("Challenge");
//     const challengeImplementation = await challengeFactory.deploy();
//     const challengeProxy = await proxyFactory.deploy(
//       await challengeImplementation.getAddress(),
//       challengeImplementation.interface.encodeFunctionData("initialize", [
//         ethers.ZeroAddress, // treasury
//         await chain.getAddress(),
//         await mockDaOracle.getAddress(),
//         ethers.ZeroAddress, // mipsChallenge
//         await chainOracle.getAddress(),
//       ]),
//     );

//     challenge = challengeFactory.attach(
//       await challengeProxy.getAddress(),
//     ) as any;
//     await challenge.connect(owner).setChallengeFee(challengeFee);

//     // 4. push next block
//     const nextBlock = { ...MOCK_DATA.rollupHeader };
//     nextBlock.prevHash = genesisHash;
//     await chain.connect(publisher).pushBlock(nextBlock);
//   });

//   it("should deploy", async function () {
//     expect(chainOracle).to.not.be.undefined;
//     expect(challenge).to.not.be.undefined;
//   });

//   describe("challengeL2Tx", function () {
//     it("should not be able to challenge without fee paid", async () => {
//       const RBLOCK_NUM = 1;
//       await expect(
//         challenge
//           .connect(owner)
//           .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash),
//       ).to.be.revertedWith("challenge fee not paid");
//     });

//     it("should not be able to challenge without preloaded header", async () => {
//       const RBLOCK_NUM = 1;
//       await expect(
//         challenge
//           .connect(owner)
//           .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//             value: challengeFee,
//           }),
//       ).to.be.revertedWith("l2BlockHash does not exist");
//     });

//     it("should be able to challenge with fee paid and preloaded header", async () => {
//       const RBLOCK_NUM = 1;

//       const rblockHash = await chain.chain(RBLOCK_NUM);

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       await expect(
//         challenge
//           .connect(owner)
//           .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//             value: challengeFee,
//           }),
//       ).to.emit(challenge, "L2TxChallengeUpdate");
//     });
//   });

//   describe("defendL2TxRoot", function () {
//     it("should not be able to defend non-existing challenge", async () => {
//       await expect(
//         challenge.connect(owner).defendL2TxRoot(5, MOCK_DATA.merkleLeaves),
//       ).to.be.revertedWith("challenge not initiated");
//     });

//     it("should not be able to defend with wrong merkle leaves", async () => {
//       const RBLOCK_NUM = 1;
//       const rblockHash = await chain.chain(RBLOCK_NUM);
//       const challengeIndex = 0n;

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       // 2. initiate challenge
//       await challenge
//         .connect(owner)
//         .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//           value: challengeFee,
//         });

//       await expect(
//         challenge
//           .connect(owner)
//           .defendL2TxRoot(challengeIndex, [
//             MOCK_DATA.merkleLeaves[0],
//             MOCK_DATA.merkleLeaves[1],
//           ]),
//       ).to.be.revertedWith("invalid tx root");
//     });

//     it("should be able to defend with correct merkle leaves", async () => {
//       const RBLOCK_NUM = 1;
//       const rblockHash = await chain.chain(RBLOCK_NUM);
//       const challengeIndex = 0n;

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       // 2. initiate challenge
//       await challenge
//         .connect(owner)
//         .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//           value: challengeFee,
//         });

//       await expect(
//         challenge
//           .connect(owner)
//           .defendL2TxRoot(challengeIndex, MOCK_DATA.merkleLeaves),
//       ).to.emit(challenge, "L2TxChallengeUpdate");
//     });
//   });

//   describe("challengeL2TxIndex", function () {
//     it("should not be able to challenge tx index for non-existing challenge", async () => {
//       await expect(
//         challenge.connect(owner).challengeL2TxIndex(5, 1),
//       ).to.be.revertedWith("challenge not in correct state");
//     });

//     it("should be able to challenge tx index", async () => {
//       const RBLOCK_NUM = 1;
//       const rblockHash = await chain.chain(RBLOCK_NUM);
//       const challengeIndex = 0n;

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       // 2. initiate challenge
//       await challenge
//         .connect(owner)
//         .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//           value: challengeFee,
//         });

//       // 3. defend tx root
//       await challenge
//         .connect(owner)
//         .defendL2TxRoot(challengeIndex, MOCK_DATA.merkleLeaves);

//       // 4. challenge tx index
//       await expect(
//         challenge.connect(owner).challengeL2TxIndex(challengeIndex, 0),
//       ).to.emit(challenge, "L2TxChallengeUpdate");
//     });
//   });

//   describe("defendL2TxIndex", function () {
//     it("should not be able to defend non-existing challenge", async () => {
//       await expect(
//         challenge.connect(owner).defendL2TxIndex(1n),
//       ).to.be.revertedWith("challenge not in correct state");
//     });

//     it("should not be able to defend without preloaded tx", async () => {
//       const RBLOCK_NUM = 1;
//       const rblockHash = await chain.chain(RBLOCK_NUM);
//       const challengeIndex = 0n;

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       // 2. initiate challenge
//       await challenge
//         .connect(owner)
//         .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//           value: challengeFee,
//         });

//       // 3. defend tx root
//       await challenge
//         .connect(publisher)
//         .defendL2TxRoot(challengeIndex, MOCK_DATA.merkleLeaves);

//       // 4. challenge tx index
//       await challenge.connect(owner).challengeL2TxIndex(challengeIndex, 0);

//       // 5. defend tx
//       await expect(
//         challenge.connect(publisher).defendL2TxIndex(challengeIndex),
//       ).to.be.revertedWith("tx not pre-submitted to chainOracle");
//     });

//     it("should be able to defend with preloaded tx", async () => {
//       const RBLOCK_NUM = 1;
//       const rblockHash = await chain.chain(RBLOCK_NUM);
//       const challengeIndex = 0n;

//       // 1. preload header
//       await provideHeader(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2HeaderProof,
//         MOCK_DATA.l2HeaderRange,
//       );

//       // 2. initiate challenge
//       await challenge
//         .connect(owner)
//         .challengeL2Tx(RBLOCK_NUM, MOCK_DATA.l2HeaderHash, {
//           value: challengeFee,
//         });

//       // 3. defend tx root
//       await challenge
//         .connect(publisher)
//         .defendL2TxRoot(challengeIndex, MOCK_DATA.merkleLeaves);

//       // 4. challenge tx index
//       await challenge.connect(owner).challengeL2TxIndex(challengeIndex, 0);

//       // 5. defend tx
//       await provideLegacyTx(
//         chainOracle,
//         rblockHash,
//         MOCK_DATA.l2TxProof,
//         MOCK_DATA.l2TxRange,
//       );

//       await expect(
//         challenge.connect(publisher).defendL2TxIndex(challengeIndex),
//       ).to.emit(challenge, "L2TxChallengeUpdate");
//     });
//   });
// });
