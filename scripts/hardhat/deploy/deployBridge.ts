import { ethers } from "hardhat";
import { log } from "../lib/log";
import { uupsProxyDeployAndInitialize } from "../lib/deploy";

// npx hardhat verify --contract contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger --network sepolia 0x719E6614628552CFE0CE83B61778b2eA8A5c2147
// npx hardhat verify --contract contracts/L1/L1StandardBridge.sol:L1StandardBridge --network sepolia 0xFe1F82B89E458F3236098cCa992209306FE45b47
// npx hardhat verify --contract contracts/L1/LightLinkPortal.sol:LightLinkPortal --network sepolia 0xc537029Daf5489d9D4B377324027a7677fEC0515

// npx hardhat verify --contract contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser --network pegasus 0x5FbDB2315678afecb367f032d93F642f64180aa3
// npx hardhat verify --contract contracts/L2/L2StandardBridge.sol:L2StandardBridge --network pegasus 0x5FbDB2315678afecb367f032d93F642f64180aa3
// npx hardhat verify --contract contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger --network pegasus 0x719E6614628552CFE0CE83B61778b2eA8A5c2147
// npx hardhat verify --contract contracts/L2/L1Block.sol:L1Block --network pegasus 0x719E6614628552CFE0CE83B61778b2eA8A5c2147

const CanonicalStateChain =
  process.env.CANONICAL_STATE_CHAIN_ADDR ?? ethers.ZeroAddress;
const Challenge = process.env.CHALLENGE_ADDR ?? ethers.ZeroAddress;
const HARDHAT_CHAIN_ID = 31337n;

const main = async () => {
  const l1Provider = ethers.provider;
  const l1Deployer = (await ethers.getSigners())[0];

  // Check CSC and Chalenge contracts exists
  if ((await l1Provider.getNetwork()).chainId != HARDHAT_CHAIN_ID) {
    if (!(await l1Provider.getCode(CanonicalStateChain)).length)
      throw new Error("CanonicalStateChain contract not deployed");
    if (!(await l1Provider.getCode(Challenge)).length)
      throw new Error("Challenge contract not deployed");
  }

  // Deploy LightLinkPortal contract to L1
  log("Deploying LightLinkPortal to L1...");
  const lightLinkPortal = await uupsProxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("LightLinkPortal"),
    [CanonicalStateChain, Challenge],
  );

  // Deploy L1CrossDomainMessenger contract to L1
  log("Deploying L1CrossDomainMessenger to L1...");
  const L1CrossDomainMessengerDeployment = await uupsProxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1CrossDomainMessenger"),
    [lightLinkPortal.address],
  );
  const l1CrossDomainMessengerAddr = L1CrossDomainMessengerDeployment.address;

  // Deploy L1StandardBridge contract to L1
  log("Deploying L1StandardBridge to L1...");
  const L1StandardBridgeDeployment = await uupsProxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1StandardBridge"),
    [l1CrossDomainMessengerAddr],
  );

  log("Deployment complete!");

  // Log deployment addresses
  log(
    " LightLinkPortal:",
    `"${lightLinkPortal.address}"`,
    `"(impl ${lightLinkPortal.implementationAddress})"`,
  );
  log(
    " L1CrossDomainMessenger:",
    `"${L1CrossDomainMessengerDeployment.address}"`,
    `"(impl ${L1CrossDomainMessengerDeployment.implementationAddress})"`,
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
    console.error(error);
    process.exit(1);
  });
