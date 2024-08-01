import { ethers } from "hardhat";
import { log } from "../lib/log";
import { proxyDeployAndInitialize } from "../lib/deploy";

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
  const lightLinkPortal = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("LightLinkPortal"),
    [CanonicalStateChain, Challenge, l1Deployer.address],
  );

  // Deploy L1CrossDomainMessenger contract to L1
  log("Deploying L1CrossDomainMessenger to L1...");
  const L1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
    l1Deployer,
    await ethers.getContractFactory("L1CrossDomainMessenger"),
    [lightLinkPortal.address],
  );
  const l1CrossDomainMessengerAddr = L1CrossDomainMessengerDeployment.address;

  // Deploy L1StandardBridge contract to L1
  log("Deploying L1StandardBridge to L1...");
  const L1StandardBridgeDeployment = await proxyDeployAndInitialize(
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
