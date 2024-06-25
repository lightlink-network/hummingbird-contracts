import { ethers } from "hardhat";
import { expect } from "chai";
import { RLPReader } from "../../../typechain-types";

describe("RLPReader", function () {
  let rlpReader: RLPReader;

  beforeEach(async function () {
    const RLPReaderFactory = await ethers.getContractFactory(
      "contracts/L1/RLPReader.sol:RLPReader",
    );
    rlpReader = (await RLPReaderFactory.deploy()) as any;
  });

  describe("toBlockHeader", function () {
    it("happy path", async function () {
      const headerRLP =
        "0xf9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aa";

      const header = await rlpReader.toBlockHeader(headerRLP);
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
        // difficulty
        500n,
      );
      expect(header[7]).to.be.equal(
        // timestamp
        62207261n,
      );
      expect(header[8]).to.be.equal(
        // gasLimit
        15000000n,
      );
      expect(header[9]).to.be.equal(
        // gasUsed
        0n,
      );
      expect(header[10]).to.be.equal(
        // timestamp
        1704720488n,
      );
      expect(header[11]).to.be.equal(
        // nonce
        14547783457063351210n,
      );
    });
  });

  describe("toLegacyTx", function () {
    it("happy path", async function () {
      const legacyTxRLP =
        "0xf8630182271082520894dfae45f5d42d7d893e15c3f55e947905c0bdec038227108026a00cc9626084e648b362f3358a10f74a9dd7928b262bb233ed9761172508592955a071a3904730383cc844e19ed819e90483493f93ed4db02225c899fe1e";

      const tx = await rlpReader.toLegacyTx(legacyTxRLP);
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
      expect(tx[6]).to.be.equal(
        // v
        38,
      );
      expect(tx[7]).to.be.equal(
        // r
        5783569416312473270300884702647224245515596113943727036258923961084301355349n,
      );
      expect(tx[8]).to.be.equal(
        // s
        51400343732176254478073733380813417096077836432544358073357581244353463975936n,
      );
    });
  });

  describe("toDepositTx", function () {
    it("happy path", async function () {
      // remove the first byte 7e
      const depositTxRLP =
        "0xf9043482076235843b9aca0083989680947f17a74942c5b22b340688f099c99a79426447e18718de76816d8000b903c4c450bc1f000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000030000000000000000000000001146405759e2a20682d9840f19edb93c2d1da0bb000000000000000000000000733910eee5eaba8ae90da5526fdd212d9a3f3ead00000000000000000000000021c184892ca64f5e4f8388970e8696e9a2927cdd0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000041c7f444c528db9a2c293882224564441355a1272480ec16f0d0f5bb45f5bc55b33effd9a430a93f5dddbb99ceeebe1ea1ca7639a2f182b17769bcd70db17448551c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004111b29824cbf7bdb2ca6cd0caa67b2084adc60b276865b6b46ed32585d6c18a58289e484c419ec8d010b36c90f4a0e3de375d0c74dd16cfe832635ac68abf12c81b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000418f374b365beb74f555701a0a3b17162245924ad06726e5012b8b5c3bc297e55f5a96672da276effcd2bb1b390ece1abd2d46ad6155d09922a9a9dc97f4ea835f1c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f99cebba649f375f1877a76112b109dc79fa04b80000000000000000000000003a3c3a5115bc299fe4e03664c9bccf43bc252a7d0000000000000000000000000000000000000000000000000018de76816d800080a0920ac370c3a042f97ff76b637a0a1f15a900d645a89617cac9101546d4c35cc8a0587d594286b222024649202920726e51bb75d2cef54ff59a2a6e6a5e9bf725bb";

      const tx = await rlpReader.toDepositTx(depositTxRLP);
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
        // v
        0n,
      );
      expect(tx[8]).to.be.equal(
        // r
        66056693244458152108319849013085333301746989274264430192031517319879043800264n,
      );
      expect(tx[9]).to.be.equal(
        // s
        40025002607391211296728310437866929214603074978161800677031384171519115994555n,
      );
    });
  });
});