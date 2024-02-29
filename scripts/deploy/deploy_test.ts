import { ethers, network } from "hardhat";
import { verify } from "../../utils/verify";

// Set DAOracle address
const DAOracleAddr = "0x3a5cbB6EF4756DA0b3f6DAE7aB6430fD8c46d247";

const main = async () => {
  // Log network name and chain id selected for deployment
  const chainIdHex = await network.provider.send("eth_chainId");
  const chainId = parseInt(chainIdHex, 16);
  console.log("Network name:", network.name);
  console.log("Network chain id:", chainId + "\n");

  // Get deployer/signer account
  const [owner] = await ethers.getSigners();
  const ownerAddr = await owner.getAddress();
  const publisherAddr = ownerAddr

  console.log("Owner address is set to:", ownerAddr);
  console.log("Publisher address is set to:", publisherAddr);
  console.log("DAOracle address set to:", DAOracleAddr + "\n");

  // Add new provider for pegasus rpc
  const pegasus = new ethers.JsonRpcProvider(process.env.PEGASUS_PROVIDER_URL);

  // Call pegasus rpc to get the latest blocks state root
  const latestBlock = await pegasus.provider.send("eth_getBlockByNumber", [
    "latest",
    true,
  ]);
  console.log(
    "Latest L2 block number for L1 genesis:",
    parseInt(latestBlock?.number, 16)
  );
  console.log("Latest L2 block hash for L1 genesis:", latestBlock?.hash);
  console.log(
    "Latest L2 block state root for L1 genesis:",
    latestBlock?.stateRoot + "\n"
  );

  // Build genesis header from latest L2 block
  const genesisHeader = {
    epoch: 0,
    l2Height: parseInt(latestBlock?.number, 16),
    prevHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: latestBlock?.hash,
    stateRoot: latestBlock?.stateRoot, // fix state root
    celestiaHeight: 0,
    celestiaShareStart: 0,
    celestiaShareLen: 0,
  };

  ///
  /// Deploy contracts
  ///

  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");

  // Deploy CanonicalStateChain contract as proxy
  console.log("Deploying CanonicalStateChain...");
  const canonicalStateChainFactory: any  = await ethers.getContractFactory(
    "CanonicalStateChain"
  );
  const canonicalStateChainImplementation = await canonicalStateChainFactory.deploy();
  await canonicalStateChainImplementation.waitForDeployment();
  const canonicalStateChainImplementationAddr = await canonicalStateChainImplementation.getAddress();

  const canonicalStateChainProxy = await proxyFactory.deploy(
    canonicalStateChainImplementationAddr,
    canonicalStateChainImplementation.interface.encodeFunctionData("initialize", [
      publisherAddr,
      genesisHeader
    ])
  );
  await canonicalStateChainProxy.waitForDeployment();
  const canonicalStateChain = canonicalStateChainFactory.attach(await canonicalStateChainProxy.getAddress());
  const canonicalStateChainContractAddr = await canonicalStateChain.getAddress();
  console.log(`→ CanonicalStateChain proxy deployed to ${canonicalStateChainContractAddr}`);
  console.log(
    `→ CanonicalStateChain implementation deployed to ${canonicalStateChainImplementationAddr}`
  );

  ///
  /// Verify contracts
  ///

  // Verify contract (after 1 min)
  console.log("Waiting for 1 min before verifying contracts..");
  await new Promise((r) => setTimeout(r, 120000));

  // Verify CanonicalStateChain Implementation
  await verify(
    canonicalStateChainImplementationAddr,
    [],
    "contracts/CanonicalStateChain.sol:CanonicalStateChain"
  );
  console.log(
    `Verified CanonicalStateChain impl contract at ${canonicalStateChainImplementationAddr}`
  );

  // Verify CanonicalStateChain Proxy
  await verify(
    canonicalStateChainContractAddr,
    [
      canonicalStateChainImplementationAddr,
      canonicalStateChainImplementation.interface.encodeFunctionData("initialize", [
        publisherAddr,
        genesisHeader,
      ]),
    ],
    "contracts/proxy/CoreProxy.sol:CoreProxy"
  );

};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
