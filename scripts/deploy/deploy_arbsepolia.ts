import { ethers, network } from "hardhat";
import { ZeroAddress } from "ethers";
import { verify } from "../../utils/verify";
import { log } from "../lib/log";
import { createGenesisHeader, proxyDeployAndInitialize } from "../lib/deploy";

// Set DAOracle address
const DAOracleAddr = "0xc3e209eb245Fd59c8586777b499d6A665DF3ABD2";

const main = async () => {
  // Step 1. Get deployer/signer account
  const [owner, publisher] = await ethers.getSigners();
  const [ownerAddr, publisherAddr] = [
    await owner.getAddress(),
    await publisher.getAddress(),
  ];
  log("Owner address:", ownerAddr);
  log("Publisher address:", publisherAddr);
  log("DAOracle address:", DAOracleAddr + "\n");

  // Step 2. Fetch latests l2 block from Pegasus
  const genesisHeader = await createGenesisHeader(
    process.env.PEGASUS_PROVIDER_URL!,
  );

  // Step 3. Deploy CanonicalStateChain contract as proxy
  log("Deploying CanonicalStateChain...");
  const canonicalStateChain = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("CanonicalStateChain"),
    [publisherAddr, genesisHeader],
  );

  // Step 4. Deploy required RLPReader lib
  log("Deploying RLPReader...");
  const RLPReader = await ethers.getContractFactory("RLPReader");
  const rlpReader = await RLPReader.deploy();
  await rlpReader.waitForDeployment();
  const rlpReaderAddr = await rlpReader.getAddress();

  // Step 5. Deploying ChainOracle contract as a proxy
  log("Deploying ChainOracle...");
  const chainOracle = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("ChainOracle"),
    [canonicalStateChain.implementationAddress, DAOracleAddr, rlpReaderAddr],
  );

  // Step 6. Deploy Challenge contract as a proxy
  log("Deploying Challenge...");
  const challenge = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("Challenge"),
    [
      ZeroAddress, // treasury
      canonicalStateChain.address,
      DAOracleAddr,
      ZeroAddress,
      chainOracle.address,
    ],
  );

  // Step 7. Deploy CanonicalTransactionChain contract as a proxy
  log("Setting challenge contract in CanonicalStateChain...");
  await canonicalStateChain.contract.setChallengeContract(challenge.address);
  log("Setting defender in Challenge contract...");
  await challenge.contract.setDefender(publisherAddr);

  log("DONE\n");

  console.log(" canonicalStateChain:", canonicalStateChain.address);
  console.log(" chainOracle:", chainOracle.address);
  console.log(" challenge:", challenge.address);
  console.log(" rlpReader:", rlpReaderAddr);
  console.log(" daOracle:", DAOracleAddr);
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
