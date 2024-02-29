import { ethers } from "hardhat";
import { TransactionReceipt } from "ethers";
import { CanonicalStateChain } from "../../typechain-types";
import {
  formatFixedEther,
  formatGWEI,
  logProgress,
  logTable,
  timeFormat,
} from "./lib/utils";
import chalk from "chalk";

const HEADER_SAMPLE_SIZE = 20;
const TX_SAMPLE_SIZE = 20;

const main = async () => {
  // 1. get contracts
  const challengeAddr = process.env.CHALLENGE!;
  const challenge = await ethers.getContractAt("Challenge", challengeAddr);
  const canonicalStateChainAddr = await challenge.chain();
  const canonicalStateChain = await ethers.getContractAt(
    "CanonicalStateChain",
    canonicalStateChainAddr,
  );

  // 2. get last X headers
  const headNum = await canonicalStateChain.chainHead();
  const headers: CanonicalStateChain.HeaderStructOutput[] = [];
  const metadata: { timestamp: bigint }[] = [];

  const from =
    headNum > HEADER_SAMPLE_SIZE ? Number(headNum) - HEADER_SAMPLE_SIZE : 0;
  for (let i = from + 1; i <= Number(headNum); i++) {
    logProgress(`Fetching headers: ${i - from} of ${HEADER_SAMPLE_SIZE} ...`);
    const headerHash = await canonicalStateChain.chain(BigInt(i));
    const header = await canonicalStateChain.headers(headerHash);
    const headerMetadata = await canonicalStateChain.headerMetadata(headerHash);
    headers.push(header);
    metadata.push(headerMetadata);
  }

  // 3. Get last recent transactions from provider to CSC
  const txs: TransactionReceipt[] = [];
  logProgress("Fetching logs...");
  const blockAddedLogs = await canonicalStateChain.queryFilter(
    canonicalStateChain.filters.BlockAdded(),
    (await ethers.provider.getBlockNumber()) - 20000,
  );
  for (const log of blockAddedLogs.reverse()) {
    logProgress(`Fetching txs: ${txs.length} of ${TX_SAMPLE_SIZE}...`);
    const tx = await ethers.provider.getTransactionReceipt(log.transactionHash);

    txs.push(tx!);
    if (txs.length >= TX_SAMPLE_SIZE) break;
  }

  // 4. calculate stats
  logProgress("Calculating stats...");
  const totalGasUsed = txs.reduce((acc, tx) => acc + tx.gasUsed!, 0n);
  const totalGasFees = txs.reduce(
    (acc, tx) => acc + tx.gasUsed! * tx.gasPrice!,
    0n,
  );
  const minGasFee = txs.reduce(
    (acc, tx) =>
      acc < tx.gasUsed! * tx.gasPrice! ? acc : tx.gasUsed! * tx.gasPrice!,
    txs[0].gasUsed! * txs[0].gasPrice!,
  );
  const maxGasFee = txs.reduce(
    (acc, tx) =>
      acc > tx.gasUsed! * tx.gasPrice! ? acc : tx.gasUsed! * tx.gasPrice!,
    txs[0].gasUsed! * txs[0].gasPrice!,
  );

  const avgGasUsed = totalGasUsed / BigInt(txs.length);
  const avgGasFees = totalGasFees / BigInt(txs.length);
  const avgGasPrice = totalGasFees / totalGasUsed;
  const avgTimeBetweenBlocks =
    (metadata[metadata.length - 1].timestamp - metadata[0].timestamp) /
    BigInt(metadata.length - 1); // in seconds

  const estimateDailyGasFees =
    (avgGasFees * 60n * 60n * 24n) / avgTimeBetweenBlocks;

  // e.g. l2BundleSize = header.l2Height - prevHeader.l2Height
  const totalL2BundleSize = headers.reduce((acc, header, i) => {
    if (i === 0) return 0n;
    return acc + header.l2Height - headers[i - 1].l2Height;
  }, 0n);
  const avgL2BundleSize = totalL2BundleSize / BigInt(headers.length);

  process.stdout.clearLine(0);
  console.log(`Stats:`);
  console.log(chalk.bold(`\nOf ${txs.length} transactions sampled:`));
  logTable(
    [chalk.italic, chalk.bold],
    [
      ["Avg gas used:", avgGasUsed],
      ["Avg gas fees:", formatFixedEther(avgGasFees, 4) + chalk.cyan(" ETH")],
      ["Avg gas price:", formatGWEI(avgGasPrice) + chalk.cyanBright(" GWEI")],
      ["Total gas used:", totalGasUsed],
      [
        "Total gas fees:",
        formatFixedEther(totalGasFees, 4) + chalk.cyan(" ETH"),
      ],
      ["Min gas fee:", formatFixedEther(minGasFee, 4) + chalk.cyan(" ETH")],
      ["Max gas fee:", formatFixedEther(maxGasFee, 4) + chalk.cyan(" ETH")],
    ],
  );

  console.log(chalk.bold(`\nOf ${metadata.length} headers sampled:`));
  logTable(
    [chalk.italic, chalk.bold],
    [
      [
        "Avg time between blocks:",
        timeFormat(Number(avgTimeBetweenBlocks) * 1000),
      ],
      ["Avg L2 bundle size:", avgL2BundleSize + " txs"],
      ["Total L2 blocks rolled up:", totalL2BundleSize + " txs"],
      [
        "Est. daily gas fees:",
        formatFixedEther(estimateDailyGasFees, 4) + chalk.cyan(" ETH"),
      ],
    ],
  );
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
