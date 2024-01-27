import { ethers } from "hardhat";
import { expect } from "chai";
import { Contract } from "ethers";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { makeNextBlock, setupCanonicalStateChain } from "./lib/chain";
import { Header, hashHeader } from "./lib/header";
import { ChainOracle } from "../typechain-types";

describe("ChainOracle", function () {
  let chainOracle: ChainOracle;
  let owner: HardhatEthersSigner;
  let publisher: HardhatEthersSigner;

  beforeEach(async function () {
    [owner, publisher] = await ethers.getSigners();
    const _chain = await setupCanonicalStateChain(owner, publisher.address);
    const canonicalStateChain = _chain.canonicalStateChain as any;

    const _MockDaOracle = await ethers.getContractFactory("MockDAOracle");
    const mockDaOracle = (await _MockDaOracle.deploy()) as any;

    const RLPReaderFactory = await ethers.getContractFactory("RLPReader");
    const rlpReader = await RLPReaderFactory.deploy();

    const ChainOracleFactory = await ethers.getContractFactory("ChainOracle");
    const chainOracleDeployed = await ChainOracleFactory.deploy(
      await canonicalStateChain.getAddress(),
      await mockDaOracle.getAddress(),
      await rlpReader.getAddress()
    );

    chainOracle = new ethers.Contract(
      await chainOracleDeployed.getAddress(),
      chainOracleDeployed.interface,
      owner
    ) as any;
  });

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

  it("should be able to decode header", async function () {
    const TestRLPHeader =
      "0xf9021aa0ce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acca0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48fa00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48403b5351d83e4e1c08084659bf868a00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aa";

    const header = await chainOracle.decodeRLPHeader(TestRLPHeader);
    expect(header).to.not.be.undefined;
    expect(header[0]).to.be.equal(
      // parentHash
      "0xce095cb5cd4725f71278ce79cb4589e5a87147fcc148fdf587292a540ee15acc"
    );
    expect(header[1]).to.be.equal(
      // uncleHash
      "0x0000000000000000000000000000000000000000000000000000000000000000"
    );
    expect(header[2]).to.be.equal(
      // coinbase
      "0xdFaD157B8D4e58c26Bf9b947f8e75b5AdbC7822B"
    );
    expect(header[3]).to.be.equal(
      // stateRoot
      "0x3903de7f5290e9ef5974c2789c47778c69bff45299b10c2c2046774a6baec48f"
    );
    expect(header[4]).to.be.equal(
      // transactionsRoot
      "0x0000000000000000000000000000000000000000000000000000000000000000"
    );
    expect(header[7]).to.be.equal(
      // difficulty
      500n
    );
    expect(header[8]).to.be.equal(
      // timestamp
      62207261n
    );
  });
});
