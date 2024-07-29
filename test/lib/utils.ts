import {
  MessagePassedEvent,
  L2ToL1MessagePasserInterface,
} from "../../typechain-types/contracts/L2/L2ToL1MessagePasser";
import { Log, EventLog } from "ethers";
import { Types } from "../../typechain-types/contracts/L1/test/BridgeProofHelper";

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
