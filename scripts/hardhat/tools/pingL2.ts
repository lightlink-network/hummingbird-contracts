import { ethers } from "hardhat";
import { L1CrossDomainMessenger } from "../../../typechain-types";

const main = async () => {
  const l1Provider = new ethers.JsonRpcProvider(process.env.L1_RPC_URL);
  const l1wallet = new ethers.Wallet(process.env.L1_DEPLOYER_KEY!, l1Provider);

  const l2Provider = new ethers.JsonRpcProvider(process.env.L2_RPC_URL);
  const l2wallet = new ethers.Wallet(process.env.L2_DEPLOYER_KEY!, l2Provider);

  // deploy pingpong contract to l2
  const PingPongFactory = await ethers.getContractFactory("PingPong");
  const pingPong = await PingPongFactory.connect(l2wallet).deploy();

  // attach l1CrossDomainMessenger contract
  const L1CrossDomainMessengerFactory = await ethers.getContractFactory(
    "L1CrossDomainMessenger",
  );
  const l1CrossDomainMessenger = (await L1CrossDomainMessengerFactory.connect(
    l1wallet,
  ).attach(
    "0x68B1D87F95878fE05B998F19b66F4baba5De1aed",
  )) as L1CrossDomainMessenger;

  // encode call: `ping("Hello L2!")`
  const callData = pingPong.interface.encodeFunctionData("ping", ["Hello L2!"]);

  // estimate gas
  const gasEstimate = await l2Provider.estimateGas({
    to: await pingPong.getAddress(),
    nonce: await l2Provider.getTransactionCount(l2wallet.address),
    from: l2wallet.address,
    data: callData,
  });

  const msgTx = await l1CrossDomainMessenger
    .connect(l1wallet)
    .sendMessage(await pingPong.getAddress(), callData, gasEstimate, {
      nonce: await l1Provider.getTransactionCount(l1wallet.address),
    });

  console.log("Message sent to L2:", msgTx.hash);
};

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
