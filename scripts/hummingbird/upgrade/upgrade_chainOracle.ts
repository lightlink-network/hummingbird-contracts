import { ethers, network } from "hardhat";
import { verify } from "../../../utils/verify";

// Set current addresses
const chainOracleContractProxyAddr = "CONTRACT_ADDRESS";

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

  // Deploy new ChainOracle contract implementation
  console.log("Deploying ChainOracle...");
  const chainOracleFactory: any =
    await ethers.getContractFactory("ChainOracle");
  const chainOracleImplementation = await chainOracleFactory.deploy();
  await chainOracleImplementation.waitForDeployment();
  const chainOracleImplementationAddr =
    await chainOracleImplementation.getAddress();

  console.log(
    `→ ChainOracle implementation deployed to ${chainOracleImplementationAddr}`,
  );

  // Update ChainOracle contract proxy to new implementation
  console.log("Updating ChainOracle proxy to new implementation...");
  const chainOracleProxy = await ethers.getContractAt(
    "ChainOracle",
    chainOracleContractProxyAddr,
  );
  await chainOracleProxy.upgradeToAndCall(chainOracleImplementationAddr, "0x");
  console.log(
    `→ ChainOracle proxy (${chainOracleContractProxyAddr}) updated to ${chainOracleImplementationAddr}` +
      "\n",
  );

  console.log("All Contracts deployed successfully! \n");

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 60000));

  // Verify ChainOracle Implementation
  await verify(
    chainOracleImplementationAddr,
    [],
    "contracts/ChainOracle.sol:ChainOracle",
  );
  console.log(
    `Verified ChainOracle impl contract at ${chainOracleImplementationAddr}`,
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
