import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { Header, hashHeader } from "./header";
import { Contract } from "ethers";
import { CanonicalStateChain } from "../../typechain-types";

export const setupCanonicalStateChain = async (
  signer: HardhatEthersSigner,
  publisher: string,
) => {
  let genesisHeader: Header = {
    prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
    epoch: BigInt(0),
    l2Height: BigInt(1),
    celestiaHeight: BigInt(0),
    celestiaShareStart: BigInt(0),
    celestiaShareLen: BigInt(0),
  };

  let genesisHash = hashHeader(genesisHeader);

  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
  const canonicalStateChainFactory = await ethers.getContractFactory(
    "CanonicalStateChain",
  );
  const canonicalStateChainImplementation =
    await canonicalStateChainFactory.deploy();

  const proxy = await proxyFactory.deploy(
    await canonicalStateChainImplementation.getAddress(),
    canonicalStateChainImplementation.interface.encodeFunctionData(
      "initialize",
      [publisher, genesisHeader],
    ),
  );

  const canonicalStateChain = canonicalStateChainFactory.attach(
    await proxy.getAddress(),
  ) as any;

  return {
    canonicalStateChain: canonicalStateChain as CanonicalStateChain,
    genesisHash,
    genesisHeader,
  };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: CanonicalStateChain,
): Promise<[string, Header]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header: Header = {
    epoch: asBigInt(head.epoch) + BigInt(1),
    l2Height: asBigInt(head.l2Height) + BigInt(5),
    prevHash: headHash,
    celestiaHeight: asBigInt(head.celestiaHeight) + BigInt(5),
    celestiaShareStart: asBigInt(head.celestiaShareStart) + BigInt(5),
    celestiaShareLen: asBigInt(head.celestiaShareLen) + BigInt(5),
  };

  // push header
  await canonicalStateChain.connect(signer).pushBlock(header);
  return [hashHeader(header), header];
};

export const makeNextBlock = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: Contract,
): Promise<[Header, string]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header: Header = {
    epoch: asBigInt(head.epoch) + BigInt(1),
    l2Height: asBigInt(head.l2Height) + BigInt(1),
    prevHash: headHash,
    celestiaHeight: asBigInt(head.celestiaHeight) + BigInt(1),
    celestiaShareStart: asBigInt(head.celestiaShareStart) + BigInt(1),
    celestiaShareLen: asBigInt(head.celestiaShareLen) + BigInt(1),
  };

  return [header, hashHeader(header)];
};

const asBigInt = (n: number | string | bigint) => {
  if (typeof n === "bigint") {
    return n;
  }
  return BigInt(n);
};
