import { ethers } from "hardhat";

export type Header = {
  prevHash: string;
  epoch: bigint;
  l2Height: bigint;
  celestiaHeight: bigint;
  celestiaShareStart: bigint;
  celestiaShareLen: bigint;
};

export const packHeader = (h: Header) =>
  ethers.AbiCoder.defaultAbiCoder().encode(
    [
      "bytes32", // prevHash
      "uint256", // epoch
      "uint256", // l2Height
      "uint256", // celestiaHeight
      "uint256", // celestiaShareStart
      "uint256", // celestiaShareLen
    ],
    [
      h.prevHash,
      h.epoch,
      h.l2Height,
      h.celestiaHeight,
      h.celestiaShareStart,
      h.celestiaShareLen,
    ],
  );

export const hashHeader = (h: Header) => ethers.keccak256(packHeader(h));
