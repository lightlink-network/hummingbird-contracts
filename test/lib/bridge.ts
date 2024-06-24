/* Imports: External */
import { ethers, Log, EventLog, BytesLike, Signer, AddressLike, BigNumberish } from "ethers";
import { toHexString, toRpcHexString } from "@eth-optimism/core-utils";
import * as rlp from "rlp";
import { BigNumber } from '@ethersproject/bignumber';
import { Types } from "../../typechain-types/contracts/L1/LightLinkPortal";
import { L2ToL1MessagePasser, L2ToL1MessagePasserInterface, MessagePassedEvent } from "../../typechain-types/contracts/L2/L2ToL1MessagePasser";

/**
 * Fix for the case where the final proof element is less than 32 bytes and the element exists
 * inside of a branch node. Current implementation of the onchain MPT contract can't handle this
 * natively so we instead append an extra proof element to handle it instead.
 *
 * @param key Key that the proof is for.
 * @param proof Proof to potentially modify.
 * @returns Modified proof.
 */
export const maybeAddProofNode = (key: string, proof: string[]) => {
  const modifiedProof = [...proof];
  const finalProofEl = modifiedProof[modifiedProof.length - 1];
  const finalProofElDecoded = rlp.decode(finalProofEl) as any;
  if (finalProofElDecoded.length === 17) {
    for (const item of finalProofElDecoded) {
      // Find any nodes located inside of the branch node.
      if (Array.isArray(item)) {
        // Check if the key inside the node matches the key we're looking for. We remove the first
        // two characters (0x) and then we remove one more character (the first nibble) since this
        // is the identifier for the type of node we're looking at. In this case we don't actually
        // care what type of node it is because a branch node would only ever be the final proof
        // element if (1) it includes the leaf node we're looking for or (2) it stores the value
        // within itself. If (1) then this logic will work, if (2) then this won't find anything
        // and we won't append any proof elements, which is exactly what we would want.
        const suffix = toHexString(item[0]).slice(3);
        if (key.endsWith(suffix)) {
          modifiedProof.push(toHexString(rlp.encode(item)));
        }
      }
    }
  }

  // Return the modified proof.
  return modifiedProof;
};

/**
 * Generates a Merkle-Patricia trie proof for a given account and storage slot.
 *
 * @param provider RPC provider attached to an EVM-compatible chain.
 * @param blockNumber Block number to generate the proof at.
 * @param address Address to generate the proof for.
 * @param slot Storage slot to generate the proof for.
 * @returns Account proof and storage proof.
 */
export const makeStateTrieProof = async (
  provider: ethers.JsonRpcProvider,
  blockNumber: BigNumber,
  address: string,
  slot: string,
): Promise<{
  accountProof: string[];
  storageProof: string[];
  storageValue: BigNumber;
  storageRoot: string;
}> => {
  const proof = await provider.send("eth_getProof", [
    address,
    [slot],
    toRpcHexString(blockNumber),
  ]);

  proof.storageProof[0].proof = maybeAddProofNode(
    ethers.keccak256(slot),
    proof.storageProof[0].proof,
  );

  return {
    accountProof: proof.accountProof,
    storageProof: proof.storageProof[0].proof,
    storageValue: BigNumber.from(proof.storageProof[0].value),
    storageRoot: proof.storageHash,
  };
};

/**
 * Utility for hashing a message hash. This computes the storage slot
 * where the message hash will be stored in state. HashZero is used
 * because the first mapping in the contract is used.
 *
 * @param messageHash Message hash to hash.
 * @returns Hash of the given message hash.
 */
export const hashMessageHash = (messageHash: string): string => {
  const data = ethers.AbiCoder.defaultAbiCoder().encode(
    ["bytes32", "uint256"],
    [messageHash, ethers.ZeroHash],
  );
  return ethers.keccak256(data);
};


/** 
 * Parses a MessagePassed event from a log.
 * 
 * @param iface Interface to use for parsing.
 * @param log Log to parse.
 * @returns Parsed event and withdrawal transaction.
 */
export const parseMessagePassedEvent = (
  iface: L2ToL1MessagePasserInterface,
  log: Log | EventLog,
): {
  evt: MessagePassedEvent.Event;
  withdrawalTx: Types.WithdrawalTransactionStruct;
} => {
  const event = iface.parseLog({
    topics: [...log.topics],
    data: log.data,
  })!;

  return {
    evt: event as unknown as MessagePassedEvent.Event,
    withdrawalTx: {
      data: event.args.data,
      gasLimit: event.args.gasLimit,
      nonce: event.args.nonce,
      sender: event.args.sender,
      target: event.args.target,
      value: event.args.value,
    },
  };
};

// keccak256(abi.encode(_tx.nonce, _tx.sender, _tx.target, _tx.value, _tx.gasLimit, _tx.data));
export const hashWithdrawalTx = (tx: Types.WithdrawalTransactionStruct) => {
  return ethers.keccak256(
    ethers.AbiCoder.defaultAbiCoder().encode(
      ["uint256", "address", "address", "uint256", "uint256", "bytes"],
      [tx.nonce, tx.sender, tx.target, tx.value, tx.gasLimit, tx.data],
    ),
  );

}

/**
 * Initiates a withdrawal from the L2ToL1MessagePasser contract.
 * @param contract the L2ToL1MessagePasser initiate withdrawal from.
 * @param sender Signer to use for initiating withdrawal.
 * @param target Target address for withdrawal.
 * @param gasLimit Gas limit for withdrawal.
 * @param data Data for withdrawal.
 * @returns Initiate transaction, withdrawal transaction data, withdrawal hash, and message slot.
 */
export const initiateWithdraw = async (
  contract: L2ToL1MessagePasser,
  sender: Signer,
  target: AddressLike,
  gasLimit: BigNumberish,
  data: BytesLike
) => {
  const initiateTx = await contract.
    connect(sender).
    initiateWithdrawal(target, gasLimit, data);
  const initiateReceipt = await initiateTx.wait();

  const { withdrawalTx } = parseMessagePassedEvent(
    contract.interface,
    initiateReceipt!.logs[0],
  );

  const withdrawalHash = hashWithdrawalTx(withdrawalTx);
  const messageSlot = hashMessageHash(withdrawalHash);

  return { initiateTx, withdrawalTx, withdrawalHash, messageSlot };
}

/**
 * Generates a withdrawal proof for a given withdrawal transaction.
 * @param l2Provider RPC provider attached to an EVM-compatible chain.
 * @param l2BlockNumber Block number to generate the proof at.
 * @param l2ToL1MessagePasser L2ToL1MessagePasser contract to generate the proof for.
 * @param messageSlot Message slot to generate the proof for.
 * @returns Withdrawal proof and output proof.
 */
export const getWithdrawalProofs = async (
  l2Provider: ethers.JsonRpcProvider,
  l2BlockNumber: BigNumberish,
  l2ToL1MessagePasser: L2ToL1MessagePasser,
  messageSlot: string
) => {
  const l2Block = await l2Provider.send("eth_getBlockByNumber", [
    l2BlockNumber,
    false,
  ]);

  const withdrawalProof = await makeStateTrieProof(
    l2Provider,
    BigNumber.from(l2Block.number),
    await l2ToL1MessagePasser.getAddress(),
    messageSlot
  );

  const outputProof: Types.OutputRootProofStruct = {
    version: ethers.ZeroHash,
    latestBlockhash: l2Block.hash,
    stateRoot: l2Block.stateRoot,
    messagePasserStorageRoot: withdrawalProof.storageRoot,
  };

  const outputRoot = ethers.keccak256(
    ethers.AbiCoder.defaultAbiCoder().encode(
      ["bytes32", "bytes32", "bytes32", "bytes32"],
      [
        outputProof.version,
        outputProof.stateRoot,
        outputProof.messagePasserStorageRoot,
        outputProof.latestBlockhash,
      ],
    ),
  );


  return { withdrawalProof, outputProof, outputRoot }
}