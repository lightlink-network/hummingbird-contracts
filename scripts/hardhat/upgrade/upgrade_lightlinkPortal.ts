import { ethers, network } from "hardhat";
import { verify } from "../../../utils/verify";

// Set current addresses
const lightlinkPortalContractProxyAddr = "CONTRACT_ADDRESS";

const main = async () => {
  // Log network name and chain id selected for deployment
  const chainIdHex = await network.provider.send("eth_chainId");
  const chainId = parseInt(chainIdHex, 16);
  console.log("Network name:", network.name);
  console.log("Network chain id:", chainId + "\n");

  // Get deployer/signer account
  const [owner, publisher] = await ethers.getSigners();
  const ownerAddr = await owner.getAddress();
  const publisherAddr = await publisher.getAddress();

  console.log("Owner address is set to:", ownerAddr);
  console.log("Publisher address is set to:", publisherAddr);

  ///
  /// Deploy contracts
  ///

  // Deploy new LightLinkPortal contract implementation
  console.log("Deploying LightLinkPortal...");
  const lightlinkPortalFactory: any = await ethers.getContractFactory(
    "contracts/L1/LightLinkPortal.sol:LightLinkPortal",
  );
  const lightlinkPortalImplementation = await lightlinkPortalFactory.deploy();
  await lightlinkPortalImplementation.waitForDeployment();
  const lightlinkPortalImplementationAddr =
    await lightlinkPortalImplementation.getAddress();

  console.log(
    `→ LightLinkPortal implementation deployed to ${lightlinkPortalImplementationAddr}`,
  );

  // Update LightLinkPortal contract proxy to new implementation
  console.log("Updating LightLinkPortal proxy to new implementation...");
  const lightlinkPortalProxy = await ethers.getContractAt(
    "contracts/L1/LightLinkPortal.sol:LightLinkPortal",
    lightlinkPortalContractProxyAddr,
  );
  await lightlinkPortalProxy.upgradeToAndCall(
    lightlinkPortalImplementationAddr,
    "0x",
  );
  console.log(
    `→ LightLinkPortal proxy (${lightlinkPortalContractProxyAddr}) updated to ${lightlinkPortalImplementationAddr}` +
      "\n",
  );

  console.log("All Contracts deployed successfully! \n");

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 60000));

  // Verify LightLinkPortal Implementation
  await verify(
    lightlinkPortalImplementationAddr,
    [],
    "contracts/L1/LightLinkPortal.sol:LightLinkPortal",
  );
  console.log(
    `Verified LightLinkPortal impl contract at ${lightlinkPortalImplementationAddr}`,
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
