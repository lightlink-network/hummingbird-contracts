import { ethers } from "hardhat";
import { CanonicalStateChain } from "../../typechain-types";
import { asBigInt } from "./utils";

// export type Header = {
//   prevHash: string;
//   epoch: bigint; // uint32
//   l2Height: bigint; // uint32
//   pointers: CelestiaPointer[];
// };

// export type CelestiaPointer = {
//   celestiaHeight: bigint; // uint32
//   celestiaShareStart: bigint; // uint32
//   celestiaShareLen: bigint; // uint32
// };

export type Header = CanonicalStateChain.HeaderStruct;

export const hashHeader = async (csc: CanonicalStateChain, h: Header) => {
  return await csc.hashHeader({
    prevHash: h.prevHash,
    epoch: asBigInt(h.epoch),
    l2Height: asBigInt(h.l2Height),
    pointers: h.pointers.map((p) => ({
      celestiaHeight: asBigInt(p.celestiaHeight),
      celestiaShareStart: asBigInt(p.celestiaShareStart),
      celestiaShareLen: asBigInt(p.celestiaShareLen),
    })),
  });
};
