import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

// Deploys a new RLPReader contract and sets it on the ChainOracle contract

const chainOracleProxyAddr = "CONTRACT_ADDRESS";

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

  // Deploy required RLPReader lib
  console.log("Deploying RLPReader...");
  const RLPReader = await ethers.getContractFactory("RLPReader");
  const rlpReader = await RLPReader.deploy();
  await rlpReader.waitForDeployment();
  const rlpReaderAddr = await rlpReader.getAddress();
  console.log(`→ RLPReader deployed to ${rlpReaderAddr}`);

  // Set chainOracle.setRLPReader(rlpReaderAddr)
  console.log("Setting RLPReader to ChainOracle...");
  const chainOracleProxy = await ethers.getContractAt(
    "ChainOracle",
    chainOracleProxyAddr,
  );
  await chainOracleProxy.setRLPReader(rlpReaderAddr);
  console.log(
    `→ RLPReader set on ChainOracle.setRLPReader(rlpReaderAddr) (${chainOracleProxyAddr})` +
      "\n",
  );

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 60000));

  // Verify RLPReader
  await verify(rlpReaderAddr, [], "contracts/RLPReader.sol:RLPReader");
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
