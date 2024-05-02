// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../contracts/challenge/ChallengeHeader.sol";

// Expose internal functions for testing

contract ChallengeHeaderTest is ChallengeHeader {
    constructor() {
        __ChallengeHeader_init();
    }

    // expose internal functions for testing
    function isHeaderValid(
        ICanonicalStateChain.Header memory _header,
        bytes32 _hash,
        uint256 _blockIndex
    ) public returns (bool) {
        return _isHeaderValid(_header, _hash, _blockIndex);
    }
}
