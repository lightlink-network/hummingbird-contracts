import { expect } from "chai";
import { ethers } from "hardhat";
import { startNetworks } from "../scripts/lib/startNetworks";
import { ChildProcess } from "child_process";
import { setupCanonicalStateChain, pushRandomHeader } from "./lib/chain";
import {
  CanonicalStateChain,
  LightLinkPortal,
  L2ToL1MessagePasser,
} from "../typechain-types";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

describe("Cross-chain interaction", function () {
  // Networks
  let l1Network: ChildProcess;
  let l2Network: ChildProcess;

  // Signers
  let l1Deployer: HardhatEthersSigner;
  let l2Deployer: HardhatEthersSigner;

  // Contracts
  let canonicalStateChain: CanonicalStateChain;
  let lightLinkPortal: LightLinkPortal;
  let l2ToL1MessagePasser: L2ToL1MessagePasser;

  before(async function () {
    // Start Hardhat network instances
    const networks = await startNetworks();
    l1Network = networks.l1Network;
    l2Network = networks.l2Network;

    // Set up L1 provider and signer
    const l1Provider = new ethers.JsonRpcProvider("http://localhost:8545");
    l1Deployer = (await l1Provider.getSigner(0)) as any;

    // Set up L2 provider and signer
    const l2Provider = new ethers.JsonRpcProvider("http://localhost:8546");
    l2Deployer = (await l2Provider.getSigner(0)) as any;

    // Deploy L1 contracts
    const proxyFactory = await ethers.getContractFactory(
      "CoreProxy",
      l1Deployer,
    );

    // CanonicalStateChain
    const _chain = await setupCanonicalStateChain(
      l1Deployer,
      l1Deployer.address,
    );
    canonicalStateChain = _chain.canonicalStateChain;

    // LightLinkPortal
    const lightLinkPortalFactory = await ethers.getContractFactory(
      "LightLinkPortal",
      l1Deployer,
    );
    const lightLinkPortalImplementation = await lightLinkPortalFactory.deploy();

    const lightLinkPortalProxy = await proxyFactory.deploy(
      await lightLinkPortalImplementation.getAddress(),
      lightLinkPortalImplementation.interface.encodeFunctionData("initialize", [
        await canonicalStateChain.getAddress(),
        ethers.ZeroAddress,
        ethers.ZeroAddress,
      ]),
    );

    lightLinkPortal = lightLinkPortalFactory.attach(
      await lightLinkPortalProxy.getAddress(),
    ) as LightLinkPortal;

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
  });
});
