import { ethers } from "hardhat";
import { CanonicalStateChain, L2CrossDomainMessenger, L2ToL1MessagePasser, LightLinkPortal } from "../../typechain-types";
import { getWithdrawalProofs, sendMessageL2ToL1 } from "../../test/lib/bridge";

const main = async () => {

    const l1Provider = new ethers.JsonRpcProvider(process.env.L1_RPC_URL)
    const l1wallet = new ethers.Wallet(process.env.L1_DEPLOYER_KEY!, l1Provider)

    const l2Provider = new ethers.JsonRpcProvider(process.env.L2_RPC_URL)
    const l2wallet = new ethers.Wallet(process.env.L2_DEPLOYER_KEY!, l2Provider)

    // deploy pingpong contract to l1
    const PingPongFactory = await ethers.getContractFactory("PingPong");
    const pingPong = await PingPongFactory.connect(l1wallet).deploy();

    console.log("PingPong deployed to L1:", await pingPong.getAddress());

    // attach l2CrossDomainMessenger contract
    const L2CrossDomainMessengerFactory = await ethers.getContractFactory("L2CrossDomainMessenger");
    const l2CrossDomainMessenger = L2CrossDomainMessengerFactory.connect(l2wallet).attach("0x3f1DeE9c3d80A81Fb34bb90309aA04F22df2dC83") as L2CrossDomainMessenger;

    // attach L2toL1MessagePasser contract
    const L2ToL1MessagePasserFactory = await ethers.getContractFactory("L2ToL1MessagePasser");
    const l2ToL1MessagePasser = L2ToL1MessagePasserFactory.connect(l2wallet).attach("0x11725D3ad60c4F6718E02D4aCf1cd12cf0680De0") as L2ToL1MessagePasser;

    // attach LightLinkPortal contract
    const LightLinkPortalFactory = await ethers.getContractFactory("LightLinkPortal");
    const lightLinkPortal = LightLinkPortalFactory.connect(l1wallet).attach("0x597818290898e3894222A5dCeDaA5F67cF5A1123") as LightLinkPortal;

    // attach canonicalStateChain contract
    const CanonicalStateChainFactory = await ethers.getContractFactory("CanonicalStateChain");
    const canonicalStateChain = CanonicalStateChainFactory.connect(l1wallet).attach("0x7DfCDBA375aF3bCBD69A312698D10525733e5eF6") as CanonicalStateChain;

    // encode call: `ping("Hello L1!")`
    const callData = pingPong.interface.encodeFunctionData("ping", ["Hello L1!"]);

    const withdrawal = await sendMessageL2ToL1(
        l2CrossDomainMessenger,
        l2ToL1MessagePasser,
        l2wallet,
        l1Provider,
        await pingPong.getAddress(),
        callData
    )

    console.log("\n\nSent Message – L2: Tx Hash", withdrawal.sendMessageTx.hash)
    console.log("Sent Message – L2: Withdrawal Tx Hash", withdrawal.withdrawalHash)

    // Generate withdrawal proofs
    const { withdrawalProof, outputProof, outputRoot } = await getWithdrawalProofs(
        l2Provider,
        withdrawal.sendMessageTx.blockNumber ?? "latest",
        l2ToL1MessagePasser,
        withdrawal.messageSlot,
    );


    console.log("\n\nWithdrawal Tx:", JSON.stringify(withdrawal.withdrawalTx));
    console.log("Withdrawal Proof:", JSON.stringify(withdrawalProof));
    console.log("Output Proof:", JSON.stringify(outputProof));
    console.log("Output Root:", JSON.stringify(outputRoot));


    console.log("\n\nContinue after rollup block published")
    console.log("Press any key to continue\n>>>")
    await keypress()

    // send withdrawal proof to L1
    const proveTx = await lightLinkPortal
        .connect(l1wallet)
        .proveWithdrawalTransaction(
            withdrawal.withdrawalTx,
            await canonicalStateChain.chainHead(),
            outputProof,
            withdrawalProof.storageProof,
        );

    console.log("Prove Tx Hash:", proveTx.hash)
}


const keypress = async () => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve(undefined)
    }))
}


main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });