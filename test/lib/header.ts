import { ethers } from "hardhat";

export interface Header {
  epoch: number;
  l2Height: number;
  prevHash: string;
  txRoot: string;
  blockRoot: string;
  stateRoot: string;
  celestiaHeight: number;
  celestiaDataRoot: string;
}

export const hashHeader = (h: Header) =>
  ethers.AbiCoder.defaultAbiCoder().encode(
    [
      "uint256",
      "uint256",
      "bytes32",
      "bytes32",
      "bytes32",
      "bytes32",
      "uint256",
      "bytes32",
    ],
    [
      h.epoch,
      h.l2Height,
      h.prevHash,
      h.txRoot,
      h.blockRoot,
      h.stateRoot,
      h.celestiaHeight,
      h.celestiaDataRoot,
    ]
  );
