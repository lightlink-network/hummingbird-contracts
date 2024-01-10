import { ethers } from "hardhat";

export interface Header {
  epoch: bigint;
  l2Height: bigint;
  prevHash: string;
  txRoot: string;
  blockRoot: string;
  stateRoot: string;
  celestiaHeight: bigint;
  celestiaDataRoot: string;
}

export const packHeader = (h: Header) =>
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

export const hashHeader = (h: Header) => ethers.keccak256(packHeader(h));
