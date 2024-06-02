import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";
import { log } from "../lib/log";
import {
  createGenesisHeader,
  proxyDeployAndInitialize,
  getBlobstreamXAddr,
} from "../lib/deploy";

const main = async () => {
  // Log network name and chain id selected for deployment
  const chainIdHex = await network.provider.send("eth_chainId");
  const chainId = parseInt(chainIdHex, 16);
  log("Network name:", network.name);
  log("Network chain id:", chainId + "\n");

  const blobstreamXAddr = getBlobstreamXAddr(chainId);

  // Step 1. Get deployer/signer account
  const [owner, publisher] = await ethers.getSigners();
  const [ownerAddr, publisherAddr] = [
    await owner.getAddress(),
    await publisher.getAddress(),
  ];
  log("Owner address:", ownerAddr);
  log("Publisher address:", publisherAddr);
  log("DAOracle address:", blobstreamXAddr + "\n");

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
    [canonicalStateChain.address, blobstreamXAddr, rlpReaderAddr],
  );

  // Step 6. Deploy Challenge contract as a proxy
  log("Deploying Challenge...");
  const challenge = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("Challenge"),
    [canonicalStateChain.address, blobstreamXAddr, chainOracle.address],
  );

  // Step 7. Deploy CanonicalTransactionChain contract as a proxy
  log("Setting challenge contract in CanonicalStateChain...");
  await canonicalStateChain.contract.setChallengeContract(challenge.address);
  log("Setting defender in Challenge contract...");
  await challenge.contract.setDefender(publisherAddr);

  log("DONE\n");

  console.log(" canonicalStateChain:", `"${canonicalStateChain.address}"`);
  console.log(" challenge:", `"${challenge.address}"`);
  console.log(" chainOracle:", `"${chainOracle.address}"`);
  console.log(" blobstreamX:", `"${blobstreamXAddr}"`);
  console.log(" rlpReader:", `"${rlpReaderAddr}" \n`);

  /// Verify contracts
  if (chainId !== 31337 && chainId !== 1337) {
    log("Waiting for 2 min before verifying contracts..");
    await new Promise((r) => setTimeout(r, 120000));

    // Verify Implementations
    await verify(canonicalStateChain.implementationAddress);
    await verify(rlpReaderAddr);
    await verify(chainOracle.implementationAddress);
    await verify(challenge.implementationAddress);

    // Verify Proxies
    await verify(canonicalStateChain.address, [
      canonicalStateChain.implementationAddress,
      canonicalStateChain.implementation.interface.encodeFunctionData(
        "initialize",
        [publisherAddr, genesisHeader],
      ),
    ]);

    await verify(chainOracle.address, [
      chainOracle.implementationAddress,
      chainOracle.implementation.interface.encodeFunctionData("initialize", [
        canonicalStateChain.address,
        blobstreamXAddr,
        rlpReaderAddr,
      ]),
    ]);

    await verify(challenge.address, [
      challenge.implementationAddress,
      challenge.implementation.interface.encodeFunctionData("initialize", [
        canonicalStateChain.address,
        blobstreamXAddr,
        chainOracle.address,
      ]),
    ]);
  }
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
