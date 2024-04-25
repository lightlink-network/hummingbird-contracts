import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { Contract, toBigInt } from "ethers";
import {
  CanonicalStateChain__factory,
  CanonicalStateChain,
} from "../../typechain-types";
import { proxyDeployAndInitialize } from "../../scripts/lib/deploy";

type Header = CanonicalStateChain.HeaderStruct;

export const setupCanonicalStateChain = async (
  signer: HardhatEthersSigner,
  publisher: string,
) => {
  let genesisHeader: Header = {
    epoch: BigInt(0),
    l2Height: BigInt(1),
    prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    shareRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaPointers: [],
  };

  const deployed = await proxyDeployAndInitialize(
    signer,
    await ethers.getContractFactory("CanonicalStateChain"),
    [publisher, genesisHeader],
  );

  const canonicalStateChain = CanonicalStateChain__factory.connect(
    deployed.address,
    signer,
  );

  return {
    canonicalStateChain,
    genesisHash: await canonicalStateChain.chain(0),
    genesisHeader,
  };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: CanonicalStateChain,
): Promise<[string, Header]> => {
  const [header, headerHash] = await makeNextBlock(signer, canonicalStateChain);

  // push header
  await canonicalStateChain
    .connect(signer)
    .getFunction("pushBlock")
    .send(header);

  return [headerHash, header];
};

export const makeNextBlock = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: CanonicalStateChain,
): Promise<[Header, string]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headNum = await canonicalStateChain.chainHead();
  const headHash = await canonicalStateChain.chain(headNum);

  let lastPointerHeight = 1n;
  if (head.celestiaPointers.length > 0) {
    lastPointerHeight = toBigInt(
      head.celestiaPointers[head.celestiaPointers.length - 1].height,
    );
  }

  let header: Header = {
    epoch: toBigInt(head.epoch) + 1n,
    l2Height: toBigInt(head.l2Height) + 5n,
    prevHash: headHash,
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    shareRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaPointers: [
      {
        height: lastPointerHeight + 1n,
        shareStart: 1n,
        shareLen: 1n,
      },
    ],
  };

  // get header hash
  const headerHash = await canonicalStateChain.calculateHeaderHash(header);
  return [header, headerHash];
};
