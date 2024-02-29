import chalk from "chalk";
import { ethers } from "hardhat";

export const shouldFail = async (promise: Promise<any>, message: string) => {
  try {
    await promise;
  } catch (error: any) {
    if (error.message && error.message.includes(message)) {
      return;
    }
    console.log(" ! Expected failure but wrong error message:", error.message);
    throw error;
  }

  throw new Error("! Expected failure did not happen");
};

export const timeAgo = (date: number) => {
  var seconds = Math.floor((Date.now() - date) / 1000);

  var interval = seconds / 31536000;

  if (interval > 1) {
    return Math.floor(interval) + " years";
  }
  interval = seconds / 2592000;
  if (interval > 1) {
    return Math.floor(interval) + " months";
  }
  interval = seconds / 86400;
  if (interval > 1) {
    return Math.floor(interval) + " days";
  }
  interval = seconds / 3600;
  if (interval > 1) {
    return Math.floor(interval) + " hours";
  }
  interval = seconds / 60;
  if (interval > 1) {
    return Math.floor(interval) + " minutes";
  }
  return Math.floor(seconds) + " seconds";
};

let _flickerIdx = 0;
export const chalkFlicker = (str: string) => {
  _flickerIdx++;
  if (_flickerIdx % 2 === 0) {
    return chalk.grey(str);
  }
  return chalk.yellow(str);
};

export const logProgress = (str: string) => {
  process.stdout.clearLine(0);
  process.stdout.write(`\r` + chalkFlicker(str) + `\r`);
};

export const formatFixedEther = (wei: bigint, fractionDigits = 2) => {
  return parseFloat(ethers.formatEther(wei)).toFixed(fractionDigits);
};

export const formatGWEI = (wei: bigint) => {
  return parseFloat(ethers.formatUnits(wei, "gwei")).toFixed(2);
};
