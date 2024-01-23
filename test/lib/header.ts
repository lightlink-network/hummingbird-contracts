import { ethers } from "hardhat";

export interface Header {
  epoch: bigint;
  l2Height: bigint;
  prevHash: string;
  txRoot: string;
  blockRoot: string;
  stateRoot: string;
  celestiaHeight: bigint;
  celestiaShareStart: bigint;
  celestiaShareLen: bigint;
}

export const packHeader = (h: Header) =>
  ethers.AbiCoder.defaultAbiCoder().encode(
    [
      "uint256", // epoch
      "uint256", // l2Height
      "bytes32", // prevHash
      "bytes32", // txRoot
      "bytes32", // blockRoot
      "bytes32", // stateRoot
      "uint256", // celestiaHeight
      "uint256", // celestiaShareStart
      "uint256", // celestiaShareLen
    ],
    [
      h.epoch,
      h.l2Height,
      h.prevHash,
      h.txRoot,
      h.blockRoot,
      h.stateRoot,
      h.celestiaHeight,
      h.celestiaShareStart,
      h.celestiaShareLen,
    ]
  );

export const hashHeader = (h: Header) => ethers.keccak256(packHeader(h));
