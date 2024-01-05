import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { Header, hashHeader } from "./header";
import { Contract } from "ethers";

export const setupCanonicalStateChain = async (signer: HardhatEthersSigner) => {
  let genesisHeader: Header = {
    epoch: 0,
    l2Height: 0,
    prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
    txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    celestiaHeight: 0,
    celestiaDataRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
  };

  let genesisHash = hashHeader(genesisHeader);

  const CanonicalStateChain = await ethers.getContractFactory(
    "CanonicalStateChain"
  );

  const canonicalStateChain = await CanonicalStateChain.deploy(
    signer.address,
    genesisHeader
  );

  return { canonicalStateChain, genesisHash };
};

export const pushRandomHeader = async (
  signer: HardhatEthersSigner,
  canonicalStateChain: Contract
) => {
  const head: Header = await canonicalStateChain.getHead();
  const headHash = hashHeader(head);

  let header: Header = {
    epoch: head.epoch + 1,
    l2Height: head.l2Height + 5,
    prevHash: headHash,
    txRoot: ethers.hexlify(ethers.randomBytes(32)),
    blockRoot: ethers.hexlify(ethers.randomBytes(32)),
    stateRoot: ethers.hexlify(ethers.randomBytes(32)),
    celestiaHeight: head.celestiaHeight + 5,
    celestiaDataRoot: ethers.hexlify(ethers.randomBytes(32)),
  };

  await canonicalStateChain
    .connect(signer)
    .getFunction("pushHeader")
    .send(header);

  return hashHeader(header);
};
