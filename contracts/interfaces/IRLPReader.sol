pragma solidity ^0.8.0;

struct RLPItem {
    uint256 length;
    uint256 ptr;
}

interface IRLPReader {
  function toRLPItem(bytes memory _data) external pure returns (RLPItem memory);
  function readList(RLPItem memory _item) external pure returns (RLPItem[] memory);
  function readBytes32(RLPItem memory _item) external pure returns (bytes32);
  function readAddress(RLPItem memory _item) external pure returns (address);
  function readUint256(RLPItem memory _item) external pure returns (uint256);
}