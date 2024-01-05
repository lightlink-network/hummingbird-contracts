pragma solidity ^0.8.0;

interface IMipsChallenge {
    function initiateChallenge(
        uint blockNumberN,
        bytes calldata blockHeaderNp1,
        bytes32 assertionRoot,
        bytes32 finalSystemState,
        uint256 stepCount
    ) external returns (uint256);

    function callWithTrieNodes(
        address target,
        bytes calldata dat,
        bytes[] calldata nodes
    ) external;

    function isSearching(uint256 challengeId) external view returns (bool);

    function getStepNumber(uint256 challengeId) external view returns (uint256);

    function getProposedState(
        uint256 challengeId
    ) external view returns (bytes32);

    function proposeState(uint256 challengeId, bytes32 stateHash) external;

    function respondState(uint256 challengeId, bytes32 stateHash) external;

    function confirmStateTransition(uint256 challengeId) external;

    function denyStateTransition(uint256 challengeId) external;

    function withdraw() external;
}
