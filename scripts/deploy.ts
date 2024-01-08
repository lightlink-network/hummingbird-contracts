import { ethers } from "hardhat";
import { verify } from "../utils/verify";

async function main() {
  // Get deployer/signer account
  const [deployer] = await ethers.getSigners();
  const deployerAddr = await deployer.getAddress()
  console.log("Deploying contracts with the deployer address:", deployerAddr);

  // Build genesis header block 62207259
  const genesisHeader = {
    epoch: 0,
    l2Height: 62207259,
    prevHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: "0x3810c405e3789ebee1ffb7ab8d36debdcefb629f129bb9db00d3271148b55f83",
    stateRoot: "0x491cc7d79299f9569e4bfddef640ade68091ab486d68a08ed1477c678db34103",
    celestiaHeight: 0,
    celestiaDataRoot: "0x0000000000000000000000000000000000000000000000000000000000000000",
  };

  // Deploy CanonicalStateChain contract
  console.log("Deploying CanonicalStateChain...")
  const csc = await ethers.deployContract("CanonicalStateChain", [deployerAddr, genesisHeader]);
  await csc.waitForDeployment();
  console.log(`CanonicalStateChain deployed to ${csc.target}`);


  // Verify contract (after 3 seconds)
  console.log("All Contracts deployed successfully!")
  console.log('Waiting for 3 seconds to verify contract')

  await new Promise(resolve => setTimeout(resolve, 30000));
  await verify(csc.target.toString(), [deployerAddr, genesisHeader])
  console.log(`Verify contract at ${csc.target}`)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
