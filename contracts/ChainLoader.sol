pragma solidity ^0.8.0;

import "./lib/Lib_RLPReader.sol";
import "blobstream-contracts/src/lib/verifier/DAVerifier.sol";
import "blobstream-contracts/src/IDAOracle.sol";
import "./interfaces/ICanonicalStateChain.sol";

contract ChainLoader {

  ICanonicalStateChain public canonicalStateChain;
  IDAOracle public daOracle;

   // L2 Header
    struct L2Header {
        bytes32 parentHash;
        bytes32 uncleHash;
        address beneficiary;
        bytes32 stateRoot;
        bytes32 transactionsRoot;
        bytes32 receiptsRoot;
        bytes32 logsBloom;
        uint256 difficulty;
        uint256 number;
        uint256 gasLimit;
        uint256 gasUsed;
        uint256 timestamp;
        bytes32 extraData;
        bytes32 mixHash;
        uint256 nonce;
    }

    mapping(bytes32 => bytes[]) public shares;
    
    // loadShares loads some shares that were uploaded to the Data Availability layer. It verifies the shares
    // are included in a given rblock (bundle) and stores them in the contract.
    // @param _rblock The rblock (bundle) that the shares are related to.
    // @param _proof The proof that the shares are available and part of the rblocks dataroot commitment.
    function loadShares(
        bytes32 _rblock,
        SharesProof memory _proof
    ) public returns (bytes32) {
        // 1. Load the rblock (bundle) from the canonical state chain.
        ICanonicalStateChain.Header memory rHead = canonicalStateChain.headers(
            _rblock
        );
        require(rHead.epoch > 0, "rblock not found");
        require(rHead.celestiaHeight == _proof.attestationProof.tuple.height, "rblock height mismatch");
        
        // 2. verify shares are valid
        DAVerifier.verifySharesToDataRootTupleRoot(
            daOracle,
            _proof,
            _proof.attestationProof.tuple.dataRoot
        );

        // 3. create a share by hashing the rblock and shares
        bytes32 shareKey = ShareKey(_rblock, _proof.data);

        // 4. store the shares
        shares[shareKey] = _proof.data;
        return shareKey;
    }

    function ShareKey(bytes32 _rblock, bytes[] memory _shareData) public pure  returns (bytes32) {
      return keccak256(abi.encode(_rblock, _shareData));
    }

}