import { ethers } from "hardhat";
import { challenge } from "../../typechain-types/contracts";
import { chalkFlicker, logProgress, shouldFail, timeAgo } from "./lib/utils";
import chalk from "chalk";

const DACHALLENGE_INITIATED = 1;
const DACHALLENGE_SUCCEEDED = 2;
const DACHALLENGE_FAILED = 3;

export const main = async () => {
  if (!process.env.CHALLENGE) {
    console.error("Missing challenge address");
    console.error("Define CHALLENGE in your .env file.");
    process.exit(1);
  }

  // 1. get contracts
  const challengeAddr = process.env.CHALLENGE;
  const challenge = await ethers.getContractAt("Challenge", challengeAddr);
  const canonicalStateChainAddr = await challenge.chain();
  const canonicalStateChain = await ethers.getContractAt(
    "CanonicalStateChain",
    canonicalStateChainAddr,
  );

  // 2. get chain head
  const headNum = await canonicalStateChain.chainHead();
  const headHash = await canonicalStateChain.chain(headNum);

  // 3. check publisher health
  console.log(chalk.bold(`Checking publisher health:`));
  const headMetadata = await canonicalStateChain.headerMetadata(headHash);
  const oneDayAgo = Math.floor(Date.now() / 1000) - 60 * 60 * 24;
  if (headMetadata.timestamp < oneDayAgo) {
    console.error("❌ Publisher is unhealthy, last publish was over 1 day ago");
    process.exit(1);
  }
  // - check publisher balance
  const publisherBalance = await ethers.provider.getBalance(
    headMetadata.publisher,
  );
  if (publisherBalance < ethers.parseEther("0.1")) {
    console.error("❌ Publisher is unhealthy, balance is too low");
    process.exit(1);
  }
  const bal = parseFloat(ethers.formatEther(publisherBalance)).toFixed(2);
  console.log(
    chalk.greenBright(
      `✔️ Publisher is healthy, has ${bal} ETH and published ${timeAgo(Number(headMetadata.timestamp) * 1000)} ago`,
    ),
  );

  // 4. try to invalidate all headers
  // - This should fail
  console.log(chalk.bold(`Checking recent rollup block headers are valid:`));
  for (let i = 0; i < Math.min(Number(headNum), 20); i++) {
    const targetNum = headNum - BigInt(i);
    logProgress(` > Attempting to invalidate header ${targetNum}...`);
    await shouldFail(challenge.invalidateHeader(targetNum), "header is valid");
  }
  process.stdout.clearLine(0);
  console.log(chalk.greenBright("✔️ All Recent headers are valid"));

  // 5. check for recent da challenges

  console.log(chalk.bold(`Checking recent DA challenges:`));
  const daChallengeLogs = await challenge.queryFilter(
    challenge.filters.ChallengeDAUpdate,
    (await ethers.provider.getBlockNumber()) - 49990,
  );
  const uniqueChallengeKeys = new Set(
    daChallengeLogs.map((log) => log.args?._blockHash),
  );

  let [succeeded, failed, pending] = [0, 0, 0];
  for (const challengeKey of uniqueChallengeKeys) {
    const daChallenge = await challenge.daChallenges(challengeKey);
    // sleep for a bit to avoid rate limiting
    await new Promise((resolve) => setTimeout(resolve, 100));

    switch (Number(daChallenge.status)) {
      case DACHALLENGE_INITIATED:
        console.warn("⚠️ Pending DA Challenge found – key:", challengeKey);
        pending++;
        break;
      case DACHALLENGE_SUCCEEDED:
        console.error("❌ Succeeded DA Challenge found – key:", challengeKey);
        succeeded++;
        break;
      case DACHALLENGE_FAILED:
        logProgress(`> Failed DA Challenge found – key: ${challengeKey}!`);
        failed++;
        break;
      default:
        throw new Error("Unknown DA Challenge status");
    }
  }
  process.stdout.clearLine(0);
  if (failed > 0)
    console.log(
      chalk.greenBright(
        `✔️ ${failed}/${uniqueChallengeKeys.size} DA challenges defeated.`,
      ),
    );
  if (succeeded > 0)
    console.log(
      `❌ ${succeeded}/${uniqueChallengeKeys.size} DA challenges succeeded.`,
    );
  if (pending > 0)
    console.log(
      `⚠️ ${pending}/${uniqueChallengeKeys.size} DA challenges pending.`,
    );
  if (uniqueChallengeKeys.size === 0)
    console.log(chalk.greenBright("✅ No DA challenges found."));
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
