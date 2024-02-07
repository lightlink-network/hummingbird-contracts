import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { Header, hashHeader } from "./header";
import { Contract } from "ethers";

export const setupCanonicalStateChain = async (
  signer: HardhatEthersSigner,
  publisher: string
) => {
  let genesisHeader: Header = {
    epoch: BigInt(0),
    l2Height: BigInt(1),
    prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaHeight: BigInt(0),
    celestiaShareStart: BigInt(0),
    celestiaShareLen: BigInt(0),
  };

  let genesisHash = hashHeader(genesisHeader);

  const proxyFactory: any = await ethers.getContractFactory("CoreProxy");
  const canonicalStateChainFactory = await ethers.getContractFactory(
    "CanonicalStateChain"
  );
  const canonicalStateChainImplementation =
    await canonicalStateChainFactory.deploy();

  const proxy = await proxyFactory.deploy(
    await canonicalStateChainImplementation.getAddress(),
    canonicalStateChainImplementation.interface.encodeFunctionData(
      "initialize",
      [publisher, genesisHeader]
    )
  );

  const canonicalStateChain = canonicalStateChainFactory.attach(
    await proxy.getAddress()
  ) as any;

  return { canonicalStateChain, genesisHash, genesisHeader };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: Contract
): Promise<[string, Header]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header: Header = {
    epoch: head.epoch + BigInt(1),
    l2Height: head.l2Height + BigInt(5),
    prevHash: headHash,
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaHeight: head.celestiaHeight + BigInt(5),
    celestiaShareStart: head.celestiaShareStart + BigInt(5),
    celestiaShareLen: head.celestiaShareLen + BigInt(5),
  };

  // push header

  await canonicalStateChain
    .connect(signer)
    .getFunction("pushBlock")
    .send(header);

  return [hashHeader(header), header];
};

export const makeNextBlock = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: Contract
): Promise<[Header, string]> => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header: Header = {
    epoch: head.epoch + BigInt(1),
    l2Height: head.l2Height + BigInt(1),
    prevHash: headHash,
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaHeight: head.celestiaHeight + BigInt(1),
    celestiaShareStart: head.celestiaShareStart + BigInt(1),
    celestiaShareLen: head.celestiaShareLen + BigInt(1),
  };

  return [header, hashHeader(header)];
};
