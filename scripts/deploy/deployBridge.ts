import { ethers } from "hardhat";
import { log } from "../lib/log";
import { proxyDeployAndInitialize } from "../lib/deploy";

// npx hardhat verify --contract contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger --network sepolia 0x719E6614628552CFE0CE83B61778b2eA8A5c2147
// npx hardhat verify --contract contracts/L1/L1StandardBridge.sol:L1StandardBridge --network sepolia 0xFe1F82B89E458F3236098cCa992209306FE45b47
// npx hardhat verify --contract contracts/L1/LightLinkPortal.sol:LightLinkPortal --network sepolia 0xc537029Daf5489d9D4B377324027a7677fEC0515

// npx hardhat verify --contract contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser --network pegasus 0x5FbDB2315678afecb367f032d93F642f64180aa3
// npx hardhat verify --contract contracts/L2/L2StandardBridge.sol:L2StandardBridge --network pegasus 0x5FbDB2315678afecb367f032d93F642f64180aa3
// npx hardhat verify --contract contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger --network pegasus 0x719E6614628552CFE0CE83B61778b2eA8A5c2147
// npx hardhat verify --contract contracts/L2/L1Block.sol:L1Block --network pegasus 0x719E6614628552CFE0CE83B61778b2eA8A5c2147

const CanonicalStateChain = "0x91C0b1164aB51c3310A7B0ceAEb208016671B1b9";
const Challenge = "0x98BAa45355ea1e6F0ABC74d6a38EBc8e82c57015";

const main = async () => {
  const l1Provider = new ethers.JsonRpcProvider(process.env.L1_RPC_URL!);
  const l1Deployer = new ethers.Wallet(
    process.env.L1_DEPLOYER_KEY!,
    l1Provider,
  );

  const l2Provider = new ethers.JsonRpcProvider(process.env.L2_RPC_URL!);
  const l2Deployer = new ethers.Wallet(
    process.env.L2_DEPLOYER_KEY!,
    l2Provider,
  );

  // Deploy L2ToL1MessagePasser contract to L2
  log("Deploying L2ToL1MessagePasser to L2...");
  const L2ToL1MessagePasser = await ethers.getContractFactory(
    "L2ToL1MessagePasser",
    l2Deployer,
  );
  const l2ToL1MessagePasser = await L2ToL1MessagePasser.deploy();
  await l2ToL1MessagePasser.waitForDeployment();

  // Deploy L1Block contract to L2
  log("Deploying L1Block to L2...");
  const L1Block = await ethers.getContractFactory("L1Block", l2Deployer);
  const l1Block = await L1Block.deploy();
  await l1Block.waitForDeployment();

  // Deploy LightLinkPortal contract to L1
  log("Deploying LightLinkPortal to L1...");
  const lightLinkPortal = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("LightLinkPortal"),
    [CanonicalStateChain, Challenge, await l1Block.getAddress()],
  );

  // Cross domain messengers
  // - Infer deployment addresses before deploying
  const l2CrossDomainMessengerAddr = ethers.getCreateAddress({
    from: l2Deployer.address,
    nonce: (await l2Provider.getTransactionCount(l2Deployer.address)) + 1,
    // +1 because implementation will be deployed first
  });
  const l1CrossDomainMessengerAddr = ethers.getCreateAddress({
    from: l1Deployer.address,
    nonce: (await l1Provider.getTransactionCount(l1Deployer.address)) + 1,
    // +1 because implementation will be deployed first
  });

  // Deploy L2CrossDomainMessenger contract to L2
  log("Deploying L2CrossDomainMessenger to L2...");
  const L2CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
    l2Deployer,
    await ethers.getContractFactory("L2CrossDomainMessenger"),
    [
      l1CrossDomainMessengerAddr,
      await l2ToL1MessagePasser.getAddress(),
      await l1Block.getAddress(),
    ],
  );

  // Deploy L1CrossDomainMessenger contract to L1
  log("Deploying L1CrossDomainMessenger to L1...");
  const L1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1CrossDomainMessenger"),
    [lightLinkPortal.address, l2CrossDomainMessengerAddr],
  );

  // Standard bridges
  // - Infer deployment addresses before deploying
  const l2StandardBridgeAddr = ethers.getCreateAddress({
    from: l2Deployer.address,
    nonce: (await l2Provider.getTransactionCount(l2Deployer.address)) + 1,
    // +1 because implementations will be deployed first
  });
  const l1StandardBridgeAddr = ethers.getCreateAddress({
    from: l1Deployer.address,
    nonce: (await l1Provider.getTransactionCount(l1Deployer.address)) + 1,
    // +1 because implementations will be deployed first
  });

  // Deploy L2StandardBridge contract to L2
  log("Deploying L2StandardBridge to L2...");
  const L2StandardBridgeDeployment = await proxyDeployAndInitialize(
    l2Deployer,
    await ethers.getContractFactory("L2StandardBridge"),
    [
      l1StandardBridgeAddr,
      l2CrossDomainMessengerAddr,
      ethers.ZeroAddress,
      await l1Block.getAddress(),
    ],
  );

  // Deploy L1StandardBridge contract to L1
  log("Deploying L1StandardBridge to L1...");
  const L1StandardBridgeDeployment = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1StandardBridge"),
    [l1CrossDomainMessengerAddr, l2StandardBridgeAddr],
  );

  log("Deployment complete!");

  // Log deployment addresses
  log(" L2ToL1MessagePasser:", `"${await l2ToL1MessagePasser.getAddress()}"`);
  log(" L1Block:", `"${await l1Block.getAddress()}"`);
  log(
    " LightLinkPortal:",
    `"${lightLinkPortal.address}"`,
    `"(impl ${lightLinkPortal.implementationAddress})"`,
  );
  log(
    " L2CrossDomainMessenger:",
    `"${L2CrossDomainMessengerDeployment.address}"`,
    `"(impl ${L2CrossDomainMessengerDeployment.implementationAddress})"`,
  );
  log(
    " L1CrossDomainMessenger:",
    `"${L1CrossDomainMessengerDeployment.address}"`,
    `"(impl ${L1CrossDomainMessengerDeployment.implementationAddress})"`,
  );
  log(
    " L2StandardBridge:",
    `"${L2StandardBridgeDeployment.address}"`,
    `"(impl ${L2StandardBridgeDeployment.implementationAddress})"`,
  );
  log(
    " L1StandardBridge:",
    `"${L1StandardBridgeDeployment.address}"`,
    `"(impl ${L1StandardBridgeDeployment.implementationAddress})"`,
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    error(error);
    process.exit(1);
  });
