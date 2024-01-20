// SPDX-License-Identifier: UNLICENSED
// LightLink Hummingbird v0.0.1

pragma solidity ^0.8.0;

import "./ChallengeBase.sol";

// ChallengeHeader lets anyone challenge a block header against some basic validity checks.
// If the header is invalid, the chain is rolled back to the previous block.
// Note: This challenge is free and has no payout.
//
// The challenge is made in a single step by calling invalidateHeader. This function directly checks
// the validity of the header without requiring the defender to respond.
//
// The following checks are made:
// 1. The epoch is greater than the previous epoch.
// 2. The l2Height is greater than the previous l2Height.
// 3. The prevHash is the previous block hash.
// 4. The bundle size is less than the max bundle size.
//
// If any of these checks fail, the chain is rolled back to the previous block.
// Just like with all challenges, the challenge window must be open.

// no constructor
abstract contract ChallengeHeader is ChallengeBase {
    uint256 MAX_BUNDLESIZE;
    uint256 CHALLENGE_PAYOUT;

    enum InvalidHeaderReason {
        InvalidEpoch,
        InvalidL2Height,
        InvalidPrevHash,
        InvalidBundleSize
    }

    event InvalidHeader(
        uint256 indexed _blockIndex,
        bytes32 indexed _hash,
        InvalidHeaderReason indexed _reason
    );

    function __ChallengeHeader_init() internal {
        MAX_BUNDLESIZE = 1000;
        CHALLENGE_PAYOUT = 0.2e18;
    }

    // invalidateHeader challenges a block header by checking that the header is valid.
    // It has no payout.
    function invalidateHeader(
        uint256 _blockIndex
    )
        external
        mustBeCanonical(_blockIndex)
        mustBeWithinChallengeWindow(_blockIndex)
    {
        bytes32 _hash = chain.chain(_blockIndex);
        ICanonicalStateChain.Header memory header = chain.headers(_hash);

        // check header validity.
        require(!_isHeaderValid(header, _hash), "header is valid");

        // rollback the chain.
        chain.rollback(_blockIndex - 1);
    }

    function _isHeaderValid(
        ICanonicalStateChain.Header memory _header,
        bytes32 _hash
    ) internal returns (bool) {
        // check that the blocks epoch is greater than the previous epoch.
        if (_header.epoch <= chain.headers(_header.prevHash).epoch) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidEpoch
            );
            return false;
        }

        // check that the l2 height is greater than the previous l2 height.
        if (_header.l2Height <= chain.headers(_header.prevHash).l2Height) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidL2Height
            );
            return false;
        }

        // check that the prevHash is the previous block hash.
        if (_header.prevHash != chain.chain(chain.chainHead() - 1)) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidPrevHash
            );
            return false;
        }

        // check that the bundle size is less than the max bundle size.
        if (
            _header.l2Height - chain.headers(_header.prevHash).l2Height >
            MAX_BUNDLESIZE
        ) {
            emit InvalidHeader(
                _header.epoch,
                _hash,
                InvalidHeaderReason.InvalidBundleSize
            );
            return false;
        }

        return true;
    }

    // gap
    uint256[50] private __gap;
}
