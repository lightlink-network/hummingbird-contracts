pragma solidity ^0.8.0;

import "./lib/Lib_RLPReader.sol";

contract RLPReader {
  function toRLPItem(bytes memory _data) external pure returns (Lib_RLPReader.RLPItem memory) {
    return Lib_RLPReader.toRLPItem(_data);
  }

  function toRLPList(Lib_RLPReader.RLPItem memory _item) external pure returns (Lib_RLPReader.RLPItem[] memory) {
    return Lib_RLPReader.readList(_item);
  }
  
  function readBytes32(Lib_RLPReader.RLPItem memory _item) external pure returns (bytes32) {
    return Lib_RLPReader.readBytes32(_item);
  }

  function readAddress(Lib_RLPReader.RLPItem memory _item) external pure returns (address) {
    return Lib_RLPReader.readAddress(_item);
  }

  function readUint256(Lib_RLPReader.RLPItem memory _item) external pure returns (uint256) {
    return Lib_RLPReader.readUint256(_item);
  }
}