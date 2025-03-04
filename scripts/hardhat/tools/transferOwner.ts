import { ethers } from "hardhat";
import { log } from "../lib/log";
import * as readline from "readline";
import { Contract, Wallet } from "ethers";

// Environment variables
const SAFE_WALLET_ADDRESS = process.env.SAFE_WALLET_ADDRESS || "";
const CURRENT_ADMIN_KEY = process.env.CURRENT_ADMIN_KEY || "";

// Contract addresses
const LIGHT_LINK_PORTAL_ADDR = process.env.LIGHT_LINK_PORTAL_ADDR;
const CHAIN_ORACLE_ADDR = process.env.CHAIN_ORACLE_ADDR;
const CANONICAL_STATE_CHAIN_ADDR = process.env.CANONICAL_STATE_CHAIN_ADDR;
const SYSTEM_CONFIG_ADDR = process.env.SYSTEM_CONFIG_ADDR;
const CHALLENGE_ADDR = process.env.CHALLENGE_ADDR;
const L1_CROSS_DOMAIN_MESSENGER_ADDR =
  process.env.L1_CROSS_DOMAIN_MESSENGER_ADDR;
const L1_STANDARD_BRIDGE_ADDR = process.env.L1_STANDARD_BRIDGE_ADDR;

// Create wallet from private key
const adminWallet = CURRENT_ADMIN_KEY ? new Wallet(CURRENT_ADMIN_KEY) : null;

// Create readline interface for user confirmation
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const askForConfirmation = async (message: string): Promise<boolean> => {
  return new Promise((resolve) => {
    rl.question(`${message} (y/n): `, (answer) => {
      resolve(answer.toLowerCase() === "y" || answer.toLowerCase() === "yes");
    });
  });
};

const logEnvironmentVariables = () => {
  log("=== ENVIRONMENT VARIABLES ===");
  log(`SAFE_WALLET_ADDRESS: ${SAFE_WALLET_ADDRESS || "NOT SET"}`);
  log(`CURRENT_ADMIN_KEY: ${adminWallet?.address || "NOT SET"}`);
  log(`LIGHT_LINK_PORTAL_ADDR: ${LIGHT_LINK_PORTAL_ADDR || "NOT SET"}`);
  log(`CHAIN_ORACLE_ADDR: ${CHAIN_ORACLE_ADDR || "NOT SET"}`);
  log(`CANONICAL_STATE_CHAIN_ADDR: ${CANONICAL_STATE_CHAIN_ADDR || "NOT SET"}`);
  log(`SYSTEM_CONFIG_ADDR: ${SYSTEM_CONFIG_ADDR || "NOT SET"}`);
  log(`CHALLENGE_ADDR: ${CHALLENGE_ADDR || "NOT SET"}`);
  log(
    `L1_CROSS_DOMAIN_MESSENGER_ADDR: ${L1_CROSS_DOMAIN_MESSENGER_ADDR || "NOT SET"}`,
  );
  log(`L1_STANDARD_BRIDGE_ADDR: ${L1_STANDARD_BRIDGE_ADDR || "NOT SET"}`);
  log("=============================");
};

const transferOwnership = async (
  contract: Contract,
  contractName: string,
  signer: Wallet,
  newOwner: string,
) => {
  try {
    const currentOwner = await contract.owner();
    log(`\nCurrent ${contractName} owner: ${currentOwner}`);

    if (currentOwner.toLowerCase() === newOwner.toLowerCase()) {
      log(`${contractName} already owned by Safe wallet`);
      return;
    }

    if (currentOwner.toLowerCase() !== signer.address.toLowerCase()) {
      log(
        `Warning: Signer (${signer.address}) is not the current owner of ${contractName} (${currentOwner})`,
      );
      return;
    }

    const tx = await contract.transferOwnership(newOwner);
    await tx.wait();
    log(`${contractName} ownership transferred to ${newOwner}, tx: ${tx.hash}`);
  } catch (error) {
    log(`Error transferring ${contractName} ownership: ${error}`);
  }
};

const main = async () => {
  logEnvironmentVariables();

  if (!SAFE_WALLET_ADDRESS || !CURRENT_ADMIN_KEY || !adminWallet) {
    throw new Error("Required environment variables are not set");
  }

  // Validate contract addresses
  if (!LIGHT_LINK_PORTAL_ADDR)
    throw new Error("Missing LIGHT_LINK_PORTAL_ADDR");
  if (!CHAIN_ORACLE_ADDR) throw new Error("Missing CHAIN_ORACLE_ADDR");
  if (!CANONICAL_STATE_CHAIN_ADDR)
    throw new Error("Missing CANONICAL_STATE_CHAIN_ADDR");
  if (!SYSTEM_CONFIG_ADDR) throw new Error("Missing SYSTEM_CONFIG_ADDR");
  if (!CHALLENGE_ADDR) throw new Error("Missing CHALLENGE_ADDR");
  if (!L1_CROSS_DOMAIN_MESSENGER_ADDR)
    throw new Error("Missing L1_CROSS_DOMAIN_MESSENGER_ADDR");
  if (!L1_STANDARD_BRIDGE_ADDR)
    throw new Error("Missing L1_STANDARD_BRIDGE_ADDR");

  const confirmed = await askForConfirmation(
    `Are you sure you want to transfer ownership of all contracts to ${SAFE_WALLET_ADDRESS}?`,
  );

  if (!confirmed) {
    log("Operation cancelled by user.");
    rl.close();
    return;
  }

  log(`Transferring ownership to Safe wallet: ${SAFE_WALLET_ADDRESS}`);
  const signer = adminWallet.connect(ethers.provider);
  log(`Using signer: ${signer.address}`);

  // Transfer ownership for each contract
  const lightLinkPortal = (await ethers.getContractAt(
    "LightLinkPortal",
    LIGHT_LINK_PORTAL_ADDR,
    signer,
  )) as unknown as Contract;
  await transferOwnership(
    lightLinkPortal,
    "LightLinkPortal",
    signer,
    SAFE_WALLET_ADDRESS,
  );

  const chainOracle = (await ethers.getContractAt(
    "ChainOracle",
    CHAIN_ORACLE_ADDR,
    signer,
  )) as unknown as Contract;
  await transferOwnership(
    chainOracle,
    "ChainOracle",
    signer,
    SAFE_WALLET_ADDRESS,
  );

  const canonicalStateChain = (await ethers.getContractAt(
    "CanonicalStateChain",
    CANONICAL_STATE_CHAIN_ADDR,
    signer,
  )) as unknown as Contract;
  await transferOwnership(
    canonicalStateChain,
    "CanonicalStateChain",
    signer,
    SAFE_WALLET_ADDRESS,
  );

  const systemConfig = (await ethers.getContractAt(
    "SystemConfig",
    SYSTEM_CONFIG_ADDR,
    signer,
  )) as unknown as Contract;
  await transferOwnership(
    systemConfig,
    "SystemConfig",
    signer,
    SAFE_WALLET_ADDRESS,
  );

  const challenge = (await ethers.getContractAt(
    "Challenge",
    CHALLENGE_ADDR,
    signer,
  )) as unknown as Contract;
  await transferOwnership(challenge, "Challenge", signer, SAFE_WALLET_ADDRESS);

  // Transfer proxy admin ownership for op contracts (LightLinkPortal, L1CrossDomainMessenger, L1StandardBridge, SystemConfig)
  const PROXY_ADMIN_SLOT =
    "0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103";

  const proxyContracts = [
    { name: "LightLinkPortal", address: LIGHT_LINK_PORTAL_ADDR },
    { name: "L1CrossDomainMessenger", address: L1_CROSS_DOMAIN_MESSENGER_ADDR },
    { name: "L1StandardBridge", address: L1_STANDARD_BRIDGE_ADDR },
    { name: "SystemConfig", address: SYSTEM_CONFIG_ADDR },
  ];

  for (const contract of proxyContracts) {
    const currentAdmin = await ethers.provider.getStorage(
      contract.address,
      PROXY_ADMIN_SLOT,
    );
    // Remove padding zeros and convert to address
    const adminAddress = "0x" + currentAdmin.slice(26);
    log(`\nCurrent ${contract.name} proxy admin: ${adminAddress}`);

    if (adminAddress.toLowerCase() === SAFE_WALLET_ADDRESS.toLowerCase()) {
      log(`${contract.name} proxy admin already owned by Safe wallet`);
    } else if (adminAddress.toLowerCase() !== signer.address.toLowerCase()) {
      log(
        `Warning: Signer (${signer.address}) is not the current proxy admin (${adminAddress}) for ${contract.name}`,
      );
    } else {
      const proxyAdmin = (await ethers.getContractAt(
        "contracts/universal/Proxy.sol:Proxy",
        contract.address,
        signer,
      )) as unknown as Contract;

      const tx = await proxyAdmin.changeAdmin(SAFE_WALLET_ADDRESS);
      await tx.wait();
      log(
        `${contract.name} proxy admin transferred to ${SAFE_WALLET_ADDRESS}, tx: ${tx.hash}`,
      );
    }
  }

  log("Ownership transfer complete!");
  rl.close();
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    rl.close();
    process.exit(1);
  });
