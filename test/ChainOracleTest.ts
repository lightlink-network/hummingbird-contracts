import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  pushRandomHeader,
  setupCanonicalStateChain,
  makeNextBlock,
} from "./lib/chain";
import { MOCK_DATA } from "./mock/mock_chainOracle";
import {
  CanonicalStateChain,
  ChainOracle,
  MockDAOracle,
  RLPReader,
} from "../typechain-types";

describe("ChainOracle", function () {
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;

  let chainOracle: ChainOracle;
  let canonicalStateChain: CanonicalStateChain;
  let mockDaOracle: MockDAOracle;
  let rlpReader: RLPReader;

  beforeEach(async function () {
    [owner, publisher] = await ethers.getSigners();
    const _chain = await setupCanonicalStateChain(owner, publisher.address);
    canonicalStateChain = _chain.canonicalStateChain as any;

    const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
    mockDaOracle = (await _MockDaOracle.deploy()) as any;

    const RLPReaderFactory = await ethers.getContractFactory("RLPReader");
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
  });

  describe("Deployment", function () {
    it("Should not be allowed to initialize twice", async function () {
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
    it("Should not be allowed to provide shares for unknown rblock", async function () {
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideShares")
          .send(
            "0x0000000000000000000000000000000000000000000000000000000000000000",
            "0x1",
            MOCK_DATA.l2HeaderProof,
          ),
      ).to.be.revertedWith("rblock not found");
    });

    it("Should not be allowed to provide shares if rblock pointer height not equal proof height", async function () {
      const [hash, header] = await pushRandomHeader(
        publisher,
        canonicalStateChain,
      );
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideShares")
          .send(hash, 0, MOCK_DATA.l2HeaderProof),
      ).to.be.revertedWith("rblock height mismatch");
    });

    it("Should revert if shares cannot be verified", async function () {
      // create next rblock and set pointer 0 height to 1286533
      const [header] = await makeNextBlock(publisher, canonicalStateChain);
      header.celestiaPointers[0].height = 1286533n; // 1286533n is the height of the proof

      // get header hash
      const headerHash = await canonicalStateChain.calculateHeaderHash(header);

      // push header
      await canonicalStateChain
        .connect(publisher)
        .getFunction("pushBlock")
        .send(header);

      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideShares")
          .send(headerHash, 0, MOCK_DATA.l2HeaderProof),
      ).to.be.revertedWith("shares not verified");
    });
  });

  describe("provideHeader", function () {
    it("Should not be allowed to provide header if its shares are not found", async function () {
      await expect(
        chainOracle
          .connect(publisher)
          .getFunction("provideHeader")
          .send(
            "0x0000000000000000000000000000000000000000000000000000000000000000",
            MOCK_DATA.l2HeaderRange,
          ),
      ).to.be.revertedWith("share not found");
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
});
