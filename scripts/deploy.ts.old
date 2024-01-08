import { ethers } from "hardhat";

// async function main() {
//   const currentTimestampInSeconds = Math.round(Date.now() / 1000);
//   const unlockTime = currentTimestampInSeconds + 60;

//   const lockedAmount = ethers.parseEther("0.001");

//   const lock = await ethers.deployContract("Lock", [unlockTime], {
//     value: lockedAmount,
//   });

//   await lock.waitForDeployment();

//   console.log(
//     `Lock with ${ethers.formatEther(
//       lockedAmount
//     )}ETH and unlock timestamp ${unlockTime} deployed to ${lock.target}`
//   );
// }

async function main() {
  const [deployer, publisher, otherAccount, challengeContract] =
    await ethers.getSigners();

  const CanonicalStateChain = await ethers.getContractFactory(
    "CanonicalStateChain"
  );
  const canonicalStateChain = await CanonicalStateChain.deploy(
    publisher.address,
    {
      epoch: 0,
      l2Height: 0,
      prevHash: ethers.keccak256(ethers.toUtf8Bytes("0")),
      txRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
      blockRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
      stateRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
      celestiaHeight: 0,
      celestiaDataRoot: ethers.keccak256(ethers.toUtf8Bytes("0")),
    }
  );

  console.log("Deployed to:", await canonicalStateChain.getAddress());
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
