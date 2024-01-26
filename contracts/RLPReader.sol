pragma solidity 0.7.6;
pragma abicoder v2;

import "./lib/Lib_RLPReader.sol";

contract RLPReader {
  function toRLPItem(bytes memory _data) external pure returns (Lib_RLPReader.RLPItem memory) {
    return Lib_RLPReader.toRlpItem(_data);
  }

  function toRLPList(Lib_RLPReader.RLPItem memory _item) external pure returns (Lib_RLPReader.RLPItem[] memory) {
    return Lib_RLPReader.toList(_item);
  }
  
  function readBytes32(Lib_RLPReader.RLPItem memory _item) external pure returns (bytes32) {
    return bytesToBytes32(Lib_RLPReader.toRlpBytes(_item));
  }

  function readAddress(Lib_RLPReader.RLPItem memory _item) external pure returns (address) {
    return Lib_RLPReader.toAddress(_item);
  }

  function readUint256(Lib_RLPReader.RLPItem memory _item) external pure returns (uint256) {
    return Lib_RLPReader.toUint(_item);
  }

  function bytesToBytes32(bytes memory source) internal pure returns (bytes32 result) {
    // Copy the bytes data into the bytes32 variable
    assembly {
        result := mload(add(source, 32))
    }

    return result;
  }
}
