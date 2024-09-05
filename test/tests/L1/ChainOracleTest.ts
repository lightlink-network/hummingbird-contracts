import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  pushRandomHeader,
  setupCanonicalStateChain,
  makeNextBlock,
} from "../../lib/chain";
import { chainOracleMockData as MOCK_DATA } from "../../mock/mock_chainOracle";
import { challengeL2TxMockData as MOCK_DATA2 } from "../../mock/mock_challengeL2Tx";
import {
  CanonicalStateChain,
  ChainOracle,
  MockDAOracle,
  RLPReader,
} from "../../../typechain-types";

describe("ChainOracle", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;

  let chainOracle: ChainOracle;
  let canonicalStateChain: CanonicalStateChain;
  let mockDaOracle: MockDAOracle;
  let rlpReader: RLPReader;

  let chain: any;

  beforeEach(async function () {
    [owner, publisher] = await ethers.getSigners();
    chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = chain.canonicalStateChain as any;

    const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
    mockDaOracle = (await _MockDaOracle.deploy()) as any;

    const RLPReaderFactory = await ethers.getContractFactory(
      "contracts/L1/RLPReader.sol:RLPReader",
    );
    rlpReader = (await RLPReaderFactory.deploy()) as any;

    const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
    const chainOracleFactory = await ethers.getContractFactory("ChainOracle");
    const chainOracleImplementation = await chainOracleFactory.deploy();

    const proxy = await proxyFactory.deploy(
      await chainOracleImplementation.getAddress(),
      chainOracleImplementation.interface.encodeFunctionData("initialize", [
        await canonicalStateChain.getAddress(),
        await mockDaOracle.getAddress(),
        await rlpReader.getAddress(),
      ]),
    );

    chainOracle = chainOracleFactory.attach(await proxy.getAddress()) as any;

    const nextBlock = { ...MOCK_DATA[0].rollupHeader };
    nextBlock.prevHash = chain.genesisHash;
    await canonicalStateChain.connect(publisher).pushBlock(nextBlock);
  });

  describe("Deployment", function () {
    it("should not be allowed to initialize twice", async function () {
      await expect(
        chainOracle
          .connect(owner)
          .getFunction("initialize")
          .send(
            await canonicalStateChain.getAddress(),
            await mockDaOracle.getAddress(),
            await rlpReader.getAddress(),
          ),
      ).to.be.revertedWithCustomError(chainOracle, "InvalidInitialization");
    });
  });

  describe("provideShares", function () {
    it("happy path", async function () {
      const rblockHash = await canonicalStateChain.chain(1);
      const shareProof = MOCK_DATA[0].headers[0].shareProofs;
      const pointerProofs = MOCK_DATA[0].headers[0].pointerProofs;

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, shareProof),
      ).to.not.be.reverted;
    });

    it("should not be allowed to provide shares for unknown rblock", async function () {
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideShares")
          .send(
            "0x0000000000000000000000000000000000000000000000000000000000000000",
            1,
            MOCK_DATA[0].headers[0].shareProofs,
          ),
      ).to.be.revertedWith("rblock not found");
    });

    it("should not be allowed to provide shares if rblock pointer height not equal proof height", async function () {
      // push a new random header to the chain who's height is not the same as the mock proof height
      const [hash] = await pushRandomHeader(publisher, canonicalStateChain);
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideShares")
          .send(hash, 0, MOCK_DATA[0].headers[0].shareProofs),
      ).to.be.revertedWith("rblock height mismatch");
    });

    it("should revert if shares cannot be verified", async function () {
      const rblockHash = await canonicalStateChain.chain(1);
      const shareProof = MOCK_DATA[0].headers[0].shareProofs;
      const pointerProofs = MOCK_DATA[0].headers[0].pointerProofs;

      await mockDaOracle.setResult(false);

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, shareProof),
      ).to.be.revertedWith("shares not verified");
    });
  });

  describe("provideHeader", function () {
    it("happy Path", async function () {
      const nextBlock = { ...MOCK_DATA2.rollupHeader };
      const prevRblockHash = await canonicalStateChain.chain(1);
      nextBlock.prevHash = prevRblockHash;
      await canonicalStateChain.connect(publisher).pushBlock(nextBlock);

      const headerShares = MOCK_DATA2.headers[0].shareProofs;
      const headerRanges = MOCK_DATA2.headers[0].shareRanges;
      const rblockHash = await canonicalStateChain.chain(2);

      // load prev header
      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, headerShares),
      ).to.not.be.reverted;
      const shareKey = await chainOracle.ShareKey(
        rblockHash,
        headerShares.data,
      );
      await expect(
        chainOracle.connect(owner).provideHeader(shareKey, headerRanges),
      ).to.not.be.reverted;
    });

    it("should not be allowed to provide header if its shares are not found", async function () {
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideHeader")
          .send(
            "0x0000000000000000000000000000000000000000000000000000000000000000",
            MOCK_DATA[0].headers[0].shareRanges,
          ),
      ).to.be.revertedWith("share not found");
    });

    // todo: fix this test by adding mock l2 header with index 0
    //   it("Should not be allowed to provide header if its index is 0", async function () {
    //     // first we need to provide shares to the chainOracle for the header
    //     const rblockHash = await canonicalStateChain.chain(1);
    //     const shareProof = MOCK_DATA[0].headers[0].shareProofs;
    //     await chainOracle
    //       .connect(owner)
    //       .provideShares(rblockHash, 0, shareProof);

    //     // then we can provide the header
    //     await expect(
    //       chainOracle
    //         .connect(publisher)
    //         .getFunction("provideHeader")
    //         .send(
    //           "0x0000000000000000000000000000000000000000000000000000000000000000",
    //           MOCK_DATA[0].headers[0].shareRanges,
    //         ),
    //     ).to.be.revertedWith("share not found");
    //   });
    // });

    it("should not be allowed to provide header if it already exists", async function () {
      const headerShares = MOCK_DATA[0].headers[0].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[0].shareRanges;
      const pointerProofs = MOCK_DATA[0].headers[0].pointerProofs;
      const rblockHash = await canonicalStateChain.chain(1);

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, headerShares),
      ).to.not.be.reverted;
      const prevShareKey = await chainOracle.ShareKey(
        rblockHash,
        headerShares.data,
      );
      await expect(
        chainOracle.connect(owner).provideHeader(prevShareKey, headerRanges),
      ).to.not.be.reverted;
      await expect(
        chainOracle.connect(owner).provideHeader(prevShareKey, headerRanges),
      ).to.be.revertedWith("header already exists");
    });
  });

  describe("provideLegacyTx", function () {
    it("happy Path", async function () {
      const nextBlock = { ...MOCK_DATA2.rollupHeader };
      const prevRblockHash = await canonicalStateChain.chain(1);
      nextBlock.prevHash = prevRblockHash;
      await canonicalStateChain.connect(publisher).pushBlock(nextBlock);

      const txShares = MOCK_DATA2.transactions[0][0].shareProofs;
      const txRanges = MOCK_DATA2.transactions[0][0].shareRanges;
      const rblockHash = await canonicalStateChain.chain(2);

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, txShares),
      ).to.not.be.reverted;

      const shareKey = await chainOracle.ShareKey(rblockHash, txShares.data);

      // TODO: fix this test
      await expect(
        chainOracle.connect(publisher).provideLegacyTx(shareKey, txRanges),
      ).to.not.be.reverted;
    });

    it("should revert if shares are not found via shareKey", async function () {
      const headerShares = MOCK_DATA[0].headers[0].shareProofs;
      const headerRanges = MOCK_DATA[0].headers[0].shareRanges;
      const rblockHash = await canonicalStateChain.chain(1);

      const shareKey = await chainOracle.ShareKey(
        rblockHash,
        headerShares.data,
      );

      await expect(
        chainOracle.connect(publisher).provideLegacyTx(shareKey, headerRanges),
      ).to.be.revertedWith("share not found");
    });

    it("should revert as transaction exists", async function () {
      const nextBlock = { ...MOCK_DATA2.rollupHeader };
      const prevRblockHash = await canonicalStateChain.chain(1);
      nextBlock.prevHash = prevRblockHash;
      await canonicalStateChain.connect(publisher).pushBlock(nextBlock);

      const txShares = MOCK_DATA2.transactions[0][0].shareProofs;
      const txRanges = MOCK_DATA2.transactions[0][0].shareRanges;
      const rblockHash = await canonicalStateChain.chain(2);

      await expect(
        chainOracle.connect(owner).provideShares(rblockHash, 0, txShares),
      ).to.not.be.reverted;

      const shareKey = await chainOracle.ShareKey(rblockHash, txShares.data);

      // TODO: fix this test
      await expect(
        chainOracle.connect(publisher).provideLegacyTx(shareKey, txRanges),
      ).to.not.be.reverted;

      await expect(
        chainOracle.connect(publisher).provideLegacyTx(shareKey, txRanges),
      ).to.be.revertedWith("transaction already exists");
    });
  });

  describe("extractData", function () {
    it("should be able to extract data from the share", async function () {
      const TestShares = [
        "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351c83e4e1c08084659bf867a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
        "0x00000000000000000000000000000000000000006c696768746c696e6b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa029c3662fab6869f5192737b6ea4676d1e5c6dc8d417e1cadc6aa6bc54ad6f7eba0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      ];

      const TestRanges = [
        { start: 168, end: 512 },
        { start: 30, end: 227 },
      ];

      const ExpectedData =
        "0xf9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aa";

      const res = await chainOracle.extractData(TestShares, TestRanges);
      expect(res).to.equal(ExpectedData);
    });
    it("should revert if range is not valid for the corresponding raw data", async function () {
      const TestShares = [
        "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351c83e4e1c08084659bf867a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
        "0x00000000000000000000000000000000000000006c696768746c696e6b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa029c3662fab6869f5192737b6ea4676d1e5c6dc8d417e1cadc6aa6bc54ad6f7eba0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      ];

      const TestRanges = [
        { start: 168, end: 1027 },
        { start: 30, end: 227 },
      ];

      await expect(
        chainOracle.extractData(TestShares, TestRanges),
      ).to.be.revertedWith("Invalid range");
    });
  });

  describe("decodeRLPHeader", function () {
    it("should be able to decode header", async function () {
      const TestRLPHeader =
        "0xf9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aa";

      const header = await chainOracle.decodeRLPHeader(TestRLPHeader);
      expect(header).to.not.be.undefined;
      expect(header[0]).to.be.equal(
        // parentHash
        "0xce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acc",
      );
      expect(header[1]).to.be.equal(
        // uncleHash
        "0x0000000000000000000000000000000000000000000000000000000000000000",
      );
      expect(header[2]).to.be.equal(
        // coinbase
        "0xdFaD157B8D4e58c26Bf9b947f8e75b5AdbC7822B",
      );
      expect(header[3]).to.be.equal(
        // stateRoot
        "0x3903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48f",
      );
      expect(header[4]).to.be.equal(
        // transactionsRoot
        "0x0000000000000000000000000000000000000000000000000000000000000000",
      );
      expect(header[5]).to.be.equal(
        // receiptsRoot
        "0x0000000000000000000000000000000000000000000000000000000000000000",
      );
      expect(header[6]).to.be.equal(
        // logsBloom
        "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      );
      expect(header[7]).to.be.equal(
        // difficulty
        500n,
      );
      expect(header[8]).to.be.equal(
        // timestamp
        62207261n,
      );
      expect(header[9]).to.be.equal(
        // gasLimit
        15000000n,
      );
      expect(header[10]).to.be.equal(
        // gasUsed
        0n,
      );
      expect(header[11]).to.be.equal(
        // timestamp
        1704720488n,
      );
      expect(header[12]).to.be.equal(
        // extraData
        "0x",
      );
      expect(header[13]).to.be.equal(
        // mixHash
        "0x0000000000000000000000000000000000000000000000000000000000000000",
      );
      expect(header[14]).to.be.equal(
        // nonce
        14547783457063351210n,
      );
    });
  });

  describe("decodeLegacyTx", function () {
    it("should be able to decode legacy tx", async function () {
      const TestRLPTx =
        "0xf8630182271082520894dfae45f5d42d7d893e15c3f55e947905c0bdec038227108026a00cc9626084e648b362f3358a10f74a9dd7928b262bb233ed9761172508592955a071a3904730383cc844e19ed819e90483493f93ed4db02225c899fe1e";

      const tx = await chainOracle.decodeLegacyTx(TestRLPTx);
      expect(tx).to.not.be.undefined;
      expect(tx[0]).to.be.equal(
        // nonce
        1n,
      );
      expect(tx[1]).to.be.equal(
        // gasPrice
        10000n,
      );
      expect(tx[2]).to.be.equal(
        // gasLimit
        21000n,
      );
      expect(tx[3].toLowerCase()).to.be.equal(
        // to
        "0xdfae45f5d42d7d893e15c3f55e947905c0bdec03".toLowerCase(),
      );
      expect(tx[4]).to.be.equal(
        // value
        10000n,
      );
      expect(tx[5]).to.be.equal(
        // data
        "0x",
      );
    });
  });

  describe("decodeDepositTx", function () {
    it("happy path", async function () {
      // remove the first byte 7e
      const depositTxRLP =
        "0xf9043482076235843b9aca0083989680947f17a74942c5b22b340688f099c99a79426447e18718de76816d8000b903c4c450bc1f000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000030000000000000000000000001146405759e2a20682d9840f19edb93c2d1da0bb000000000000000000000000733910eee5eaba8ae90da5526fdd212d9a3f3ead00000000000000000000000021c184892ca64f5e4f8388970e8696e9a2927cdd0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000041c7f444c528db9a2c293882224564441355a1272480ec16f0d0f5bb45f5bc55b33effd9a430a93f5dddbb99ceeebe1ea1ca7639a2f182b17769bcd70db17448551c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004111b29824cbf7bdb2ca6cd0caa67b2084adc60b276865b6b46ed32585d6c18a58289e484c419ec8d010b36c90f4a0e3de375d0c74dd16cfe832635ac68abf12c81b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000418f374b365beb74f555701a0a3b17162245924ad06726e5012b8b5c3bc297e55f5a96672da276effcd2bb1b390ece1abd2d46ad6155d09922a9a9dc97f4ea835f1c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f99cebba649f375f1877a76112b109dc79fa04b80000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d0000000000000000000000000000000000000000000000000018de76816d800080a0920ac370c3a042f97ff76b637a0a1f15a900d645a89617cac9101546d4c35cc8a0587d594286b222024649202920726e51bb75d2cef54ff59a2a6e6a5e9bf725bb";

      const tx = await chainOracle.decodeDepositTx(depositTxRLP);
      expect(tx).to.not.be.undefined;
      expect(tx[0]).to.be.equal(
        // chainId
        1890n,
      );
      expect(tx[1]).to.be.equal(
        // nonce
        53,
      );
      expect(tx[2]).to.be.equal(
        // gasPrice
        1000000000n,
      );
      expect(tx[3]).to.be.equal(
        // gasLimit
        10000000n,
      );
      expect(tx[4].toLowerCase()).to.be.equal(
        // to
        "0x7F17A74942c5b22b340688f099c99A79426447e1".toLowerCase(),
      );
      expect(tx[5]).to.be.equal(
        // value
        7000000000000000n,
      );
      expect(tx[6]).to.be.equal(
        // data
        "0xc450bc1f000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000030000000000000000000000001146405759e2a20682d9840f19edb93c2d1da0bb000000000000000000000000733910eee5eaba8ae90da5526fdd212d9a3f3ead00000000000000000000000021c184892ca64f5e4f8388970e8696e9a2927cdd0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000041c7f444c528db9a2c293882224564441355a1272480ec16f0d0f5bb45f5bc55b33effd9a430a93f5dddbb99ceeebe1ea1ca7639a2f182b17769bcd70db17448551c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004111b29824cbf7bdb2ca6cd0caa67b2084adc60b276865b6b46ed32585d6c18a58289e484c419ec8d010b36c90f4a0e3de375d0c74dd16cfe832635ac68abf12c81b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000418f374b365beb74f555701a0a3b17162245924ad06726e5012b8b5c3bc297e55f5a96672da276effcd2bb1b390ece1abd2d46ad6155d09922a9a9dc97f4ea835f1c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f99cebba649f375f1877a76112b109dc79fa04b80000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d0000000000000000000000000000000000000000000000000018de76816d8000",
      );
      expect(tx[7]).to.be.equal(
        // r
        66056693244458152108319849013085333301746989274264430192031517319879043800264n,
      );
      expect(tx[8]).to.be.equal(
        // s
        40025002607391211296728310437866929214603074978161800677031384171519115994555n,
      );
      expect(tx[9]).to.be.equal(
        // v
        0,
      );
    });
  });

  describe("getTransaction", function () {
    it("Should return an empty tx as not found", async function () {
      const txn = await chainOracle
        .connect(publisher)
        .getTransaction(
          "0x0000000000000000000000000000000000000000000000000000000000000001",
        );

      expect(txn[0]).to.be.equal(0);
    });
  });

  describe("setRLPReader", function () {
    it("happy path", async function () {
      const RLPReaderFactory = await ethers.getContractFactory(
        "contracts/L1/RLPReader.sol:RLPReader",
      );
      const rlpReader = (await RLPReaderFactory.deploy()) as any;

      await expect(
        chainOracle.connect(owner).setRLPReader(await rlpReader.getAddress()),
      ).to.not.be.reverted;

      const _rlpReader = await chainOracle.rlpReader();
      expect(_rlpReader).to.equal(await rlpReader.getAddress());
    });

    it("should revert as non owner", async function () {
      const RLPReaderFactory = await ethers.getContractFactory(
        "contracts/L1/RLPReader.sol:RLPReader",
      );
      const rlpReader = (await RLPReaderFactory.deploy()) as any;

      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("setRLPReader")
          .send(await rlpReader.getAddress()),
      ).to.be.revertedWithCustomError(
        canonicalStateChain,
        "OwnableUnauthorizedAccount",
      );
    });
  });

  describe("setDAOracle", function () {
    it("happy path", async function () {
      const DAOracleMockAddress = "0xF0c6429ebAB2e7DC6e05DaFB61128bE21f13cb1e";

      await expect(chainOracle.connect(owner).setDAOracle(DAOracleMockAddress))
        .to.not.be.reverted;

      const _DAOracle = await chainOracle.daOracle();
      expect(_DAOracle).to.equal(DAOracleMockAddress);
    });

    it("should revert as non owner", async function () {
      const DAOracleMockAddress = "0xF0c6429ebAB2e7DC6e05DaFB61128bE21f13cb1e";

      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("setDAOracle")
          .send(DAOracleMockAddress),
      ).to.be.revertedWithCustomError(
        canonicalStateChain,
        "OwnableUnauthorizedAccount",
      );
    });
  });
});
