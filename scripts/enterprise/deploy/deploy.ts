import { ethers, network } from "hardhat";
import { verify } from "../../../utils/verify";
import { log } from "../../universal/log";
import { proxyDeployAndInitialize } from "../../universal/utils";

const main = async () => {
  // Log network name and chain id selected for deployment
  const chainIdHex = await network.provider.send("eth_chainId");
  const chainId = parseInt(chainIdHex, 16);
  log("Network name:", network.name);
  log("Network chain id:", chainId + "\n");

  // Step 1. Get deployer/signer account
  const [owner, reserve] = await ethers.getSigners();
  const [ownerAddr] = [await owner.getAddress()];
  log("Owner address:", ownerAddr);

  // Step 2. Deploy CanonicalStateChain contract as proxy
  log("Deploying EnterprisePortal...");
  const enterprisePortal = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("EnterprisePortal"),
    [],
  );

  // Step 3. Deploy LLERC20 contract
  log("Deploying LLERC20...");
  const llERC20 = await ethers.getContractFactory("LLERC20");
  const llToken = await llERC20.deploy("1000000000000000000");
  await llToken.waitForDeployment();
  const llTokenAddr = await llToken.getAddress();

  // Step 4. Deploy EnterpriseGasStation contract as a proxy
  log("Deploying EnterpriseGasStation...");
  const enterpriseGasStation = await proxyDeployAndInitialize(
    owner,
    await ethers.getContractFactory("EnterpriseGasStation"),
    [enterprisePortal.address, llTokenAddr, ownerAddr],
  );

  log("DONE\n");

  console.log(" enterprisePortal:", `"${enterprisePortal.address}"`);
  console.log(" llERC20:", `"${llTokenAddr}"`);
  console.log(" enterpriseGasStation:", `"${enterpriseGasStation.address}"`);

  /// Verify contracts
  if (chainId !== 31337 && chainId !== 1337) {
    log("Waiting for 2 min before verifying contracts..");
    await new Promise((r) => setTimeout(r, 120000));

    // Verify Implementations
    await verify(enterprisePortal.implementationAddress);
    await verify(llTokenAddr);
    await verify(enterpriseGasStation.implementationAddress);

    // Verify Proxies
    await verify(enterprisePortal.address, [
      enterprisePortal.implementationAddress,
      enterprisePortal.implementation.interface.encodeFunctionData(
        "initialize",
        [],
      ),
    ]);

    await verify(enterpriseGasStation.address, [
      enterpriseGasStation.implementationAddress,
      enterpriseGasStation.implementation.interface.encodeFunctionData(
        "initialize",
        [enterprisePortal.address, llTokenAddr, ownerAddr],
      ),
    ]);
  }
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
