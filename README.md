# Hummingbird contracts

> [!NOTE]  
> Hummingbird is a work in progress and is not yet ready for production use.

> [!NOTE]
> Requires **Node version 18**. If you have nvm installed, run `nvm use` to switch to the correct version.

This repo contains smart contracts for the Hummingbird project. Work in progress.

Documentation & deployed contract addresses can be found [here](https://docs.lightlink.io/lightlink-protocol/achitecture-and-design/lightlink-protocol-deep-dive).

## Usage

```bash
npx hardhat gobind # updates go bindings
npx hardhat compile # compiles the contracts
npx hardhat test # runs the tests
```

We've also included some scripts for interacting with the contracts:

```bash
npx hardhat run scripts/deploy/deploy.ts --network <localhost> # deploy the contracts to your local node

npx hardhat run scripts/tools/stats.ts --network <sepolia|localhost> # prints some stats about the chain

npx hardhat run scripts/tools/check.ts --network <sepolia|localhost> # checks to publishers health and for active challenges

npx hardhat run scripts/tools/chain.ts --network <sepolia|localhost> # prints the local chain
```

## Overview

![Overview](image.png)

```
CanonicalStateChain.sol – Where the rollup chain is stored.
Challenge
  └ ChallengeManager.sol – A Collection of the challenges below that can be called to rollback the chain if needed.
  └ ChallengeBase.sol – Base contract for all challenges, allows getting and setting the challenge window, and the challenge fee.
  └ ChallengeHeader.sol – Challenge a rollup block header for validity.
  └ ChallengeDataAvailability.sol – Claim that block data is unavailable.
  └ ChallengeL2Header.sol – Challenge the L2 header is not included in a given rollup block.
```

## Running locally

Despite this being a hardhat project, we recommend using to [anvil](https://github.com/foundry-rs/foundry/tree/master/crates/anvil) run locally. There are some
compatibility issues with hardhat node and the go bindings used in with the hummingbird client.

### Starting the node

```bash
anvil --chain-id 1337 --fork-url <SEPOLIA_RPC_URL>
```

You need to provide a sepolia fork url if you want to use the chainOracle contract.

### Then deploy the contracts

```bash
npx hardhat run script/deploy.ts --network localhost
```

## Challenges Overview

Challenges allow anyone to challenge the validity of a block. If a block is found to be invalid, the chain is rolled back to the previous block.

Most challenges require a challenge fee, which is paid to the winner of the challenge. This mechanism is used to deincentivize
frivolous challenges, and reimburse the defender the cost of providing proof (GAS).

Challenges must be made within a valid challenge window. This is a time window that starts when the block is published, and ends after a certain time has passed. If a challenge is made outside of this window, it will be rejected. The window may be different for certain challenges.

### ChallengeHeader:

Contract that lets anyone challenge a block header against some basic validity checks. The following checks are made:

1.  The epoch is greater than the previous epoch.
2.  The l2Height is greater than the previous l2Height.
3.  The prevHash is the previous block hash.
4.  The bundle size is less than the max bundle size.

If any of these checks fail, the chain is rolled back to the previous block.

Just like with all challenges, the challenge window must be open, however there is no challenge fee.

### ChallengeDataAvailability

This a challenge game, anybody can challenge the DA of a block.

Once initiated the defender (block publisher) must provide proof within a shorted time window. (This proof is verified via the blobstream contract). If they fail to do so, the challenger wins the challenge and the chain is rolled back to the previous block.

The window on this challenge starts 80 mins after the block is published, and ends 1.5 days after the block is published. The delay to the start of the challenge window gives enough time for Celestia to validate the data and publish the proof. The shortened end window, gives time for subsequent challenges to after data availability is proven.

### ChallengeL2Header

This is another two party challenge game where the defender must provide
a valid L2 header to defend against a challenge.

The Challenge goes through the following steps:

1.  A challenger initiates a challenge by calling challengeL2Header with the rblock number and the number of the L2 block it should contain.
2.  The defending block publisher must provide valid L2 headers to the chainOracle for both the challenged block and the previous block.
3.  If the headers are valid, the defender wins the challenge and receives the challenge fee.
4.  Otherwise the challenge expires and the challenger wins the challenge and the block is rolled back.

### ChallengeExecution

**Incomplete**, but will call the MIPS challenge game to computer the correct stateRoot and compare it to the one in the block header. If the header is incorrect the block will be rolled back.

![Hummingbird](hb.png)

### TODO

- [ ] ChallengeExecution
- [ ] MultiChainMessenger
- [ ] Add TxQueue to CanonicalStateChain.
- [ ] ChallengeTxInclusion
