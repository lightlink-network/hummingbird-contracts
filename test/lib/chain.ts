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
    celestiaDataRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaTxHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
  };

  let genesisHash = hashHeader(genesisHeader);

  const CanonicalStateChain = await ethers.getContractFactory(
    "CanonicalStateChain"
  );

  const canonicalStateChain = await CanonicalStateChain.deploy(
    publisher,
    genesisHeader
  );

  return { canonicalStateChain, genesisHash, genesisHeader };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: Contract
) => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header = {
    epoch: head.epoch + BigInt(1),
    l2Height: head.l2Height + BigInt(5),
    prevHash: headHash,
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaHeight: head.celestiaHeight + BigInt(5),
    celestiaDataRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaTxHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
  };

  // push header

  await canonicalStateChain
    .connect(signer)
    .getFunction("pushBlock")
    .send(header);

  return hashHeader(header);
};
