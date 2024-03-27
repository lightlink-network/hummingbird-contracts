import { ethers, config } from "hardhat";
import { Wallet, HDNodeWallet, Mnemonic } from "ethers";

const generateAccounts = (
  mnemonic: string,
  initialIndex = 0,
  count = 20,
  path = "m/44'/60'/0'/0",
) => {
  const accounts = [];

  const mn = Mnemonic.fromPhrase(mnemonic);
  for (let i = initialIndex; i < initialIndex + count; i++) {
    const walletPath = `${path}/${i}`;
    const wallet = HDNodeWallet.fromMnemonic(mn, walletPath);
    accounts.push({
      path: walletPath,
      address: wallet.address,
      privateKey: wallet.privateKey,
    });
  }

  return accounts;
};

const main = async () => {
  const accounts = config.networks.hardhat.accounts as any;
  generateAccounts(
    accounts.mnemonic,
    accounts.initialIndex,
    accounts.count,
    accounts.path,
  ).forEach((account) => {
    console.log(`Path: ${account.path}`);
    console.log(`Address: ${account.address}`);
    console.log(`Private Key: ${account.privateKey}`);
    console.log();
  });
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
