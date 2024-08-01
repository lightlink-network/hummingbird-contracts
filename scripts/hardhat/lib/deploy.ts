import { ethers } from "hardhat";
import { CanonicalStateChain } from "../../../typechain-types";
import { Proxy } from "../../../typechain-types/contracts/universal";

export const proxyDeployAndInitialize = async (
  signer: any,
  implementationFactory: any,
  args: any[],
) => {
  // step 1: deploy implementation contract
  const implementation = await implementationFactory.connect(signer).deploy();
  await implementation.waitForDeployment();
  const implementationAddress = await implementation.getAddress();

  // step 2: deploy proxy contract
  const proxyFactory = await ethers.getContractFactory(
    "contracts/universal/Proxy.sol:Proxy",
    signer,
  );
  const proxy: Proxy = (await proxyFactory
    .connect(signer)
    .deploy(signer.address)) as any;
  await proxy.waitForDeployment();

  await proxy.upgradeToAndCall(
    implementationAddress,
    implementation.interface.encodeFunctionData("initialize", args),
  );

  const proxyAddress = await proxy.getAddress();

  return {
    contract: await implementationFactory.attach(proxyAddress).connect(signer),
    address: proxyAddress,
    implementation,
    implementationAddress,
  };
};

/**
 * Deploy a Proxy and contract and initialize it.
 * @param signer The signer to deploy the contract.
 * @param implementationFactory The factory of the implementation contract.
 * @param args The arguments to initialize on the contract.
 * @returns The proxy and implementation contract.
 */
export const uupsProxyDeployAndInitialize = async (
  signer: any,
  implementationFactory: any,
  args: any[],
) => {
  // step 1: deploy implementation contract
  const implementation = await implementationFactory.connect(signer).deploy();
  await implementation.waitForDeployment();
  const implementationAddress = await implementation.getAddress();

  // step 2: deploy proxy contract
  const proxyFactory = await ethers.getContractFactory("CoreProxy");
  const contract = await proxyFactory
    .connect(signer)
    .deploy(
      implementationAddress,
      implementation.interface.encodeFunctionData("initialize", args),
    );
  await contract.waitForDeployment();
  const proxyAddress = await contract.getAddress();

  return {
    contract: await implementationFactory.attach(proxyAddress),
    address: proxyAddress,
    implementation,
    implementationAddress,
  };
};

export const deployAndInitialize = async (
  signer: any,
  factory: any,
  args: any[],
) => {
  const contract = await factory.connect(signer).deploy();
  await contract.waitForDeployment();
  await contract.initialize(...args);

  return {
    contract: await factory.attach(contract.address),
    address: contract.address,
  };
};

export const createGenesisHeader = async (providerRPC: string) => {
  const rpc = new ethers.JsonRpcProvider(providerRPC);
  const latestBlock = await rpc.provider.send("eth_getBlockByNumber", [
    "latest",
    true,
  ]);

  if (latestBlock == undefined) throw new Error("Failed to get latest block");

  // calculate output root:
  const versionHash = ethers.ZeroHash;
  const stateRoot = latestBlock.stateRoot;
  const withdrawalRoot = ethers.ZeroHash;
  const blockHash = latestBlock.hash;

  const outputRoot = ethers.keccak256(
    ethers.solidityPacked(
      ["bytes32", "bytes32", "bytes32", "bytes32"],
      [versionHash, stateRoot, withdrawalRoot, blockHash],
    ),
  );

  // Step 3. Build genesis header from latest L2 block
  const genesisHeader: CanonicalStateChain.HeaderStruct = {
    epoch: 0,
    l2Height: parseInt(latestBlock?.number, 16),
    prevHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    outputRoot: outputRoot,
    celestiaPointers: [],
  };

  return genesisHeader;
};

export const getBlobstreamXAddr = (chainId: any) => {
  switch (chainId) {
    case 1: // Ethereum
      if (!process.env.BLOBSTREAM_ETHEREUM) {
        throw new Error("Environment variable BLOBSTREAM_ETHEREUM is not set");
      }
      return process.env.BLOBSTREAM_ETHEREUM!;
    case 11155111: // Sepolia
      if (!process.env.BLOBSTREAM_SEPOLIA) {
        throw new Error("Environment variable BLOBSTREAM_SEPOLIA is not set");
      }
      return process.env.BLOBSTREAM_SEPOLIA!;
    default:
      // throw new Error("BlobstreamX address not found for chain id: " + chainId);
      console.warn("BlobstreamX address not found for chain id: " + chainId);
      return ethers.ZeroAddress;
  }
};
