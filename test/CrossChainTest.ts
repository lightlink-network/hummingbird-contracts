import { expect } from "chai";
import { ethers } from "hardhat";
import { JsonRpcProvider } from "ethers";
import { startNetworks } from "../scripts/lib/startNetworks";
import { ChildProcess } from "child_process";
import { setupCanonicalStateChain, pushRandomHeader } from "./lib/chain";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
  BridgeProofHelper,
} from "../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { proxyDeployAndInitialize } from "../scripts/lib/deploy";

describe("Cross-chain interaction", function () {
  // Networks
  let l1Network: ChildProcess;
  let l2Network: ChildProcess;

  // Signers
  let l1Deployer: HardhatEthersSigner;
  let l2Deployer: HardhatEthersSigner;

  // Providers 
  let l1Provider: JsonRpcProvider;
  let l2Provider: JsonRpcProvider;

  // Contracts
  let canonicalStateChain: CanonicalStateChain;
  let lightLinkPortal: LightLinkPortal;
  let l2ToL1MessagePasser: L2ToL1MessagePasser;
  let BridgeProofHelper: BridgeProofHelper;

  before(async function () {
    // Start Hardhat network instances
    const networks = await startNetworks();
    l1Network = networks.l1Network;
    l2Network = networks.l2Network;

    // Set up L1 provider and signer
    l1Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8545");
    l1Deployer = (await l1Provider.getSigner(0)) as any;

    // Set up L2 provider and signer
    l2Provider = new ethers.JsonRpcProvider("http://0.0.0.0:8546");
    l2Deployer = (await l2Provider.getSigner(0)) as any;

    // Deploy L1 contracts

    // CanonicalStateChain
    const _chain = await setupCanonicalStateChain(
      l1Deployer,
      l1Deployer.address,
    );
    canonicalStateChain = _chain.canonicalStateChain;

    // LightLinkPortal
    const lightLinkPortalDeployment = await proxyDeployAndInitialize(
      l1Deployer,
      await ethers.getContractFactory("LightLinkPortal"),
      [
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ],
    );
    lightLinkPortal = lightLinkPortalDeployment.contract as LightLinkPortal;


    // BridgeProofHelper
    const bridgeProofHelperFactory = await ethers.getContractFactory(
      "contracts/L1/test/BridgeProofHelper.sol:BridgeProofHelper",
      l1Deployer,
    );
    BridgeProofHelper = (await bridgeProofHelperFactory.deploy()) as any;
    await BridgeProofHelper.waitForDeployment();

    // Deploy L2 contracts
    const L2ToL1MessagePasserFactory = await ethers.getContractFactory(
      "contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser",
      l2Deployer,
    );
    l2ToL1MessagePasser = (await L2ToL1MessagePasserFactory.deploy()) as any;
    await l2ToL1MessagePasser.waitForDeployment();
  });

  after(async function () {
    // Kill Hardhat network instances
    l1Network.kill();
    l2Network.kill();
  });

  it("Should interact with L1 and L2", async function () {
    // Initiate withdrawal from L2
    await l2ToL1MessagePasser
      .connect(l2Deployer)
      .initiateWithdrawal(ethers.ZeroAddress, 0, "0x");

    // Push a new header to L1
    const [headerHash, header] = await pushRandomHeader(
      l1Deployer,
      canonicalStateChain,
    );

    // log headers
    console.log("Header hash:", headerHash);
    console.log("Header:", header);


  });
});
