import { ethers } from "hardhat";
import chalk from "chalk";
import { logProgress, timeAgo } from "./lib/utils";

const ANTI_RATE_LIMIT_DELAY = 1000;

const main = async () => {
  if (!process.env.CHALLENGE) {
    console.error("Missing challenge address");
    console.error("Define CHALLENGE in your .env file.");
    process.exit(1);
  }

  const challengeAddr = process.env.CHALLENGE!;
  const challenge = await ethers.getContractAt("Challenge", challengeAddr);
  const canonicalStateChainAddr = await challenge.chain();
  const canonicalStateChain = await ethers.getContractAt(
    "CanonicalStateChain",
    canonicalStateChainAddr,
  );

  const head = await canonicalStateChain.chainHead();
  for (let i = Number(head); i > -1; i--) {
    logProgress(`Fetching headers: ${Number(head) - i} of ${head}...`);

    const headerHash = await canonicalStateChain.chain(i);
    const header = await canonicalStateChain.headers(headerHash);
    const headerMetadata = await canonicalStateChain.headerMetadata(headerHash);
    const timestamp = Number(headerMetadata.timestamp) * 1000;

    // Anti-rate limit
    if (i % 10 === 0) {
      await new Promise((resolve) =>
        setTimeout(resolve, ANTI_RATE_LIMIT_DELAY),
      );
    }

    process.stdout.clearLine(0);
    console.log(
      `Header ${i}: ${chalk.bold(headerHash)} (${timeAgo(timestamp)} ago) - ${chalk.cyan(`L1 Epoch: ${header.epoch}`)} : ${chalk.yellow(`L2 Epoch ${header.l2Height}`)}`,
    );
  }
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
