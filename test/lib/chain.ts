import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { Header, hashHeader } from "./header";
import { Contract } from "ethers";
import { CanonicalStateChain } from "../../typechain-types";
import { asBigInt } from "./utils";

export const setupCanonicalStateChain = async (
  signer: HardhatEthersSigner,
  publisher: string,
) => {
  let genesisHeader: Header = {
    prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
    epoch: BigInt(0),
    l2Height: BigInt(1),
    pointers: [
      {
        celestiaHeight: BigInt(0),
        celestiaShareStart: BigInt(0),
        celestiaShareLen: BigInt(0),
      },
    ],
  };

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
    genesisHash: await hashHeader(canonicalStateChain, genesisHeader),
    genesisHeader,
  };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: CanonicalStateChain,
): Promise<[string, Header]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = await hashHeader(canonicalStateChain, head);

  const celestiaHeight =
    head.pointers.length > 0 ? head.pointers[0].celestiaHeight : 0n;

  let header: Header = {
    epoch: asBigInt(head.epoch) + BigInt(1),
    l2Height: asBigInt(head.l2Height) + BigInt(5),
    prevHash: headHash,
    pointers: [
      {
        celestiaHeight: asBigInt(celestiaHeight) + 5n,
        celestiaShareStart: 5n,
        celestiaShareLen: 5n,
      },
    ],
  };

  // push header
  await canonicalStateChain.connect(signer).pushBlock(header);
  return [await hashHeader(canonicalStateChain, header), header];
};

export const makeNextBlock = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: CanonicalStateChain,
): Promise<[Header, string]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = await hashHeader(canonicalStateChain, head);

  const celestiaHeight =
    head.pointers.length > 0 ? head.pointers[0].celestiaHeight : 0n;

  let header: Header = {
    epoch: asBigInt(head.epoch) + 1n,
    l2Height: asBigInt(head.l2Height) + 1n,
    prevHash: headHash,
    pointers: [
      {
        celestiaHeight: asBigInt(celestiaHeight) + 1n,
        celestiaShareStart: 1n,
        celestiaShareLen: 1n,
      },
    ],
  };

  const hash = await hashHeader(canonicalStateChain, header);
  return [header, hash];
};
