// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// NOT FOR RELEASE
// NOT IMPLEMENTED YET

import "./ChallengeBase.sol";
import "../lib/Lib_MerkleRoot.sol";

contract ChallengeL2Tx is ChallengeBase {
    enum L2TxChallengeStatus {
        None,
        Initiated,
        RootProved,
        IndexProvided,
        DefenderWon,
        ChallengerWon
    }

    struct L2TxChallenge {
        uint256 blockNum;
        bytes32 l2BlockHash;
        bytes32 l2TxRoot;
        uint32 txIndex;
        address challenger;
        address defender;
        uint256 expiry;
        L2TxChallengeStatus status;
        bool claimed;
    }

    event L2TxChallengeUpdate(
        uint256 indexed challengeIndex,
        bytes32 indexed l2BlockHash,
        bytes32 l2TxRoot,
        uint32 txIndex,
        address challenger,
        L2TxChallengeStatus indexed status
    );

    // challengeIndex => L2 tx hash
    mapping(uint256 => bytes32[]) public txHashes;

    // challengeIndex => L2TxChallenge
    uint256 l2TxChallengesIdx = 0;
    mapping(uint256 => L2TxChallenge) public l2TxChallenges;

    function challengeL2Tx(
        uint256 _rblockNum,
        bytes32 _l2BlockHash
    )
        public
        payable
        mustBeCanonical(_rblockNum)
        mustBeWithinChallengeWindow(_rblockNum)
        requireChallengeFee
        returns (uint256)
    {
        // TODO: PREVENT THE SAME CHALLENGE FROM BEING INITIATED MULTIPLE TIMES.

        // 1. load the l2 header
        IChainOracle.L2Header memory l2Header = chainOracle.getHeader(
            _l2BlockHash
        );
        require(l2Header.number > 0, "l2BlockHash does not exist");

        // 2. check that the l2BlockHash is loaded and part of the rblock
        bytes32 rblockHash = chain.chain(_rblockNum);
        require(
            chainOracle.headerToRblock(_l2BlockHash) == rblockHash,
            "l2BlockHash not part of rblock"
        );
        
        require(
            l2Header.transactionsRoot != bytes32(0),
            "l2BlockHash has no tx root"
        );

        l2TxChallenges[l2TxChallengesIdx] = L2TxChallenge({
            blockNum: _rblockNum,
            l2BlockHash: _l2BlockHash,
            l2TxRoot: l2Header.transactionsRoot,
            txIndex: 0,
            challenger: msg.sender,
            defender: address(0),
            expiry: block.timestamp + (challengePeriod / 4),
            status: L2TxChallengeStatus.Initiated,
            claimed: false
        });

        emit L2TxChallengeUpdate(
            l2TxChallengesIdx,
            _l2BlockHash,
            l2Header.transactionsRoot,
            0,
            msg.sender,
            L2TxChallengeStatus.Initiated
        );

        l2TxChallengesIdx++;
        return l2TxChallengesIdx - 1;
    }

    function defendL2TxRoot(
        uint256 challengeIndex,
        bytes32[] memory _txHashes
    ) public {
        L2TxChallenge storage challenge = l2TxChallenges[challengeIndex];
        require(
            challenge.status == L2TxChallengeStatus.Initiated,
            "challenge not initiated"
        );
        require(challenge.expiry > block.timestamp, "challenge expired");

        require(
            Lib_MerkleRoot.calculateMerkleRoot(_txHashes) == challenge.l2TxRoot,
            "invalid tx root"
        );

        challenge.status = L2TxChallengeStatus.RootProved;
        txHashes[challengeIndex] = _txHashes;
        challenge.expiry = block.timestamp + (challengePeriod / 4);
        challenge.defender = msg.sender;
        emit L2TxChallengeUpdate(
            challengeIndex,
            challenge.l2BlockHash,
            challenge.l2TxRoot,
            challenge.txIndex,
            challenge.challenger,
            L2TxChallengeStatus.RootProved
        );
    }

    function challengeL2TxIndex(
        uint256 challengeIndex,
        uint32 _txIndex
    ) public {
        L2TxChallenge storage challenge = l2TxChallenges[challengeIndex];
        require(
            challenge.status == L2TxChallengeStatus.RootProved,
            "challenge not in correct state"
        );
        require(challenge.expiry > block.timestamp, "challenge expired");
        require(_txIndex < txHashes[challengeIndex].length, "invalid tx index");

        challenge.txIndex = _txIndex;
        challenge.expiry = block.timestamp + (challengePeriod / 4);
        challenge.status = L2TxChallengeStatus.IndexProvided;

        emit L2TxChallengeUpdate(
            challengeIndex,
            challenge.l2BlockHash,
            challenge.l2TxRoot,
            _txIndex,
            challenge.challenger,
            L2TxChallengeStatus.IndexProvided
        );
    }

    function defendL2TxIndex(uint256 challengeIndex) public {
        L2TxChallenge storage challenge = l2TxChallenges[challengeIndex];
        require(
            challenge.status == L2TxChallengeStatus.IndexProvided,
            "challenge not in correct state"
        );
        require(challenge.expiry > block.timestamp, "challenge expired");

        IChainOracle.DepositTx memory _tx = chainOracle.getTransaction(
            txHashes[challengeIndex][challenge.txIndex]
        );

        require(_tx.nonce > 0, "tx not pre-submitted to chainOracle");

        challenge.status = L2TxChallengeStatus.DefenderWon;
        challenge.defender = msg.sender;
        emit L2TxChallengeUpdate(
            challengeIndex,
            challenge.l2BlockHash,
            challenge.l2TxRoot,
            challenge.txIndex,
            challenge.challenger,
            L2TxChallengeStatus.DefenderWon
        );
    }

    function settleL2TxChallenge(uint256 challengeIndex) public {
        L2TxChallenge storage challenge = l2TxChallenges[challengeIndex];
        require(challenge.expiry < block.timestamp, "challenge not expired");

        if (
            challenge.status == L2TxChallengeStatus.Initiated ||
            challenge.status == L2TxChallengeStatus.IndexProvided
        ) {
            challenge.status = L2TxChallengeStatus.ChallengerWon;
            emit L2TxChallengeUpdate(
                challengeIndex,
                challenge.l2BlockHash,
                challenge.l2TxRoot,
                challenge.txIndex,
                challenge.challenger,
                L2TxChallengeStatus.ChallengerWon
            );
            
            // rollback the tx
            chain.rollback(challenge.blockNum - 1);
        }
        // Otherwise its implied the root was proved and the defender won

        challenge.status = L2TxChallengeStatus.DefenderWon;
        emit L2TxChallengeUpdate(
            challengeIndex,
            challenge.l2BlockHash,
            challenge.l2TxRoot,
            challenge.txIndex,
            challenge.challenger,
            L2TxChallengeStatus.DefenderWon
        );
    }

    function claimL2TxChallengeReward(uint256 _challengeKey) external {
        L2TxChallenge storage challenge = l2TxChallenges[_challengeKey];
        require(challenge.claimed == false, "challenge has already been claimed");
        require(challenge.status == L2TxChallengeStatus.ChallengerWon || challenge.status == L2TxChallengeStatus.DefenderWon, "challenge is not in the correct state");

        if (challenge.status == L2TxChallengeStatus.ChallengerWon) {
            (bool success, ) = challenge.challenger.call{value: challengeFee}("");
            require(success, "failed to pay challenger");
        } else {
            (bool success, ) = defender.call{value: challengeFee}("");
            require(success, "failed to pay defender");
        }

        challenge.claimed = true;
    }
}
