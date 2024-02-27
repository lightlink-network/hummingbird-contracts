import { ethers } from "hardhat";
import { CanonicalStateChain, ChainOracle } from "../../typechain-types";
import { SharesProofStruct } from "../../typechain-types/contracts/ChainOracle";
import { Header } from "../lib/header";

/*
Epoch: 5329636
L2Height: 69484106
PrevHash: 0xa1d0b9e0de4ccfd612f75f1dab9b1fedc60491e654a5dfd9a0292a4459369d41
TxRoot: 0x6f952897d5c682e66526f0e42a37c2a9aeaebeeaa68966a10d27815f08088607
BlockRoot: 0xbc8fe206922ff2d655c15231f8e2209d5476df916d3dce282795efb6f5722bda
StateRoot: 0xf1b3254a4dd4e47a6215b7ba699c42ec6106e63cbacbd36ca3ac064961cf1faa
CelestiaHeight: 1215624
CelestiaShareStart: 256
CelestiaShareLen: 3462
*/

const rollupHeaders: CanonicalStateChain.HeaderStruct[] = [
  {
    epoch: BigInt(5329636),
    l2Height: BigInt(69484106),
    prevHash: "", // This will be replaced when pushed to the chain
    txRoot:
      "0x6f952897d5c682e66526f0e42a37c2a9aeaebeeaa68966a10d27815f08088607",
    blockRoot:
      "0xbc8fe206922ff2d655c15231f8e2209d5476df916d3dce282795efb6f5722bda",
    stateRoot:
      "0xf1b3254a4dd4e47a6215b7ba699c42ec6106e63cbacbd36ca3ac064961cf1faa",
    celestiaHeight: BigInt(1215624),
    celestiaShareStart: BigInt(256),
    celestiaShareLen: BigInt(3462),
  },
];

const l2HeaderHashes: string[] = [
  "0x52c2fd04cfb7cbcea2b54d1d9bd5ba95a3cf08e5bf9166773cbef21f2323c58e",
  "0x4a63bb258a8b7ecde64390ae55072309cfa5894682f84b4ee6dade945e47524f",
];

const l2Headers: ChainOracle.L2HeaderStruct[] = [
  {
    parentHash:
      "0x7eb362194afa8eb951c642f4d4c99358d604646dd3982bf5d4402cf150353f0b",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0x4c6bc8a737c95ec9b2a53ac6fbb990bce84c250d2787e34e144da1a015868d46",
    transactionsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: BigInt(500),
    number: BigInt(69483210),
    gasLimit: BigInt(15000000),
    gasUsed: BigInt(0),
    timestamp: BigInt(1708377995),
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: BigInt("0xc9e41bfa0b90b3aa"),
  },
  {
    parentHash:
      "0x52c2fd04cfb7cbcea2b54d1d9bd5ba95a3cf08e5bf9166773cbef21f2323c58e",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0x168fdbcb48457df2f0f1ec9f0fb8587f092e02a67628b7f381998c762eb8276c",
    transactionsRoot:
      "0x6c87336806abadbe78ba6f77a8588bccd14134618a5a0fc1444c446f23e5ce24",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: BigInt(500),
    number: BigInt(69483211),
    gasLimit: BigInt(15000000),
    gasUsed: BigInt(206612),
    timestamp: BigInt(1708377996),
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: BigInt("0xc9e41bfa0b90b3aa"),
  },
];

const shareProofs: SharesProofStruct[] = [
  // 0x52c2fd04cfb7cbcea2b54d1d9bd5ba95a3cf08e5bf9166773cbef21f2323c58e
  {
    data: [
      Buffer.from(
        "AAAAAAAAAAAAAAAAAAAAAAAAAABsaWdodGxpbmsAC5CzqsDA+QIf+QIaoBeLsdcTSIFlaxX/NtgvbiY/ivgIYn0nvBSjz3JceR5AoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlN+tFXuNTljCa/m5R/jnW1rbx4IroJOepQQbtPTZVm+9YSw7BocT8b/sSAImY63Q/1FAm1WzoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIIB9IQEJDoog+ThwICEZdPHOqAAAAAAAAA=",
        "base64",
      ),
      Buffer.from(
        "AAAAAAAAAAAAAAAAAAAAAAAAAABsaWdodGxpbmsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIyeQb+guQs6rAwPkCH/kCGqAXjyMTlzWXrODNBxdGlaBZKsvZgy4LhxzjkdvxzqP2mqAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJTfrRV7jU5Ywmv5uUf451ta28eCK6CTnqUEG7T02VZvvWEsOwaHE/G/7EgCJmOt0P9RQJtVs6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "base64",
      ),
    ],
    shareProofs: [
      {
        beginKey: 4,
        endKey: 6,
        sideNodes: [
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
              255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
              255, 255, 255, 255, 255, 255,
            ]),
          },
        ],
      },
    ],
    namespace: {
      version: new Uint8Array([0]),
      id: new Uint8Array([
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105, 103,
        104, 116, 108, 105, 110, 107,
      ]),
    },
    rowRoots: [
      {
        min: {
          version: new Uint8Array([0]),
          id: new Uint8Array([
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
            103, 104, 116, 108, 105, 110, 107,
          ]),
        },
        max: {
          version: new Uint8Array([0]),
          id: new Uint8Array([
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
            103, 104, 116, 108, 105, 110, 107,
          ]),
        },
        digest: new Uint8Array([
          0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
          103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
        ]),
      },
    ],
    rowProofs: [
      {
        sideNodes: [
          new Uint8Array([
            203, 127, 36, 204, 55, 238, 20, 207, 232, 182, 20, 179, 166, 106,
            186, 133, 65, 17, 179, 254, 15, 153, 245, 147, 71, 188, 28, 152,
            129, 102, 118, 13,
          ]),
          new Uint8Array([
            114, 18, 74, 162, 231, 88, 150, 155, 51, 78, 64, 219, 118, 158, 199,
            87, 157, 118, 181, 158, 175, 160, 221, 176, 247, 172, 20, 148, 193,
            24, 174, 245,
          ]),
          new Uint8Array([
            169, 62, 204, 43, 160, 194, 121, 168, 116, 42, 249, 206, 13, 154,
            142, 255, 158, 126, 118, 224, 247, 117, 111, 79, 7, 167, 4, 138, 15,
            82, 71, 184,
          ]),
          new Uint8Array([
            216, 25, 139, 170, 146, 41, 202, 122, 106, 165, 201, 174, 182, 41,
            22, 159, 194, 13, 182, 97, 71, 24, 4, 219, 150, 59, 110, 98, 203,
            171, 129, 94,
          ]),
          new Uint8Array([
            70, 17, 32, 130, 228, 0, 72, 120, 40, 46, 94, 237, 60, 221, 222,
            130, 208, 137, 187, 145, 74, 152, 163, 6, 148, 112, 190, 42, 169,
            169, 77, 186,
          ]),
          new Uint8Array([
            14, 146, 27, 77, 228, 64, 249, 206, 244, 6, 179, 204, 41, 210, 187,
            196, 31, 181, 237, 129, 97, 136, 6, 135, 28, 175, 60, 153, 91, 65,
            128, 24,
          ]),
          new Uint8Array([
            6, 146, 134, 122, 130, 228, 68, 202, 204, 171, 210, 181, 245, 60,
            204, 69, 189, 170, 90, 76, 14, 189, 168, 78, 186, 144, 178, 76, 160,
            126, 117, 174,
          ]),
          new Uint8Array([
            140, 66, 35, 195, 166, 33, 48, 120, 45, 213, 132, 209, 22, 140, 128,
            54, 4, 130, 147, 245, 29, 51, 104, 183, 146, 46, 203, 25, 181, 189,
            60, 161,
          ]),
        ],
        key: 39,
        numLeaves: 256,
      },
    ],
    attestationProof: {
      tupleRootNonce: 3383,
      tuple: {
        height: 1215624,
        dataRoot: new Uint8Array([
          69, 207, 56, 157, 215, 169, 201, 84, 110, 149, 228, 243, 38, 249, 107,
          40, 240, 138, 215, 221, 91, 153, 96, 142, 215, 190, 72, 98, 98, 183,
          251, 22,
        ]),
      },
      proof: {
        sideNodes: [
          new Uint8Array([
            238, 112, 187, 95, 254, 122, 188, 44, 151, 221, 82, 132, 217, 153,
            247, 122, 209, 45, 229, 229, 101, 210, 115, 225, 182, 100, 220, 71,
            160, 49, 219, 82,
          ]),
          new Uint8Array([
            41, 178, 137, 107, 95, 47, 61, 150, 156, 46, 84, 77, 243, 208, 82,
            209, 80, 142, 247, 95, 155, 239, 147, 246, 166, 22, 148, 139, 33,
            239, 206, 72,
          ]),
          new Uint8Array([
            108, 20, 31, 228, 135, 184, 78, 33, 113, 130, 104, 139, 150, 94,
            250, 34, 195, 186, 185, 29, 96, 184, 5, 201, 157, 183, 119, 243, 61,
            165, 119, 162,
          ]),
          new Uint8Array([
            61, 12, 254, 109, 182, 152, 247, 238, 136, 122, 194, 137, 219, 164,
            197, 61, 91, 11, 170, 209, 115, 62, 3, 73, 74, 113, 58, 117, 125,
            16, 9, 157,
          ]),
          new Uint8Array([
            215, 167, 20, 40, 195, 168, 241, 214, 121, 140, 158, 7, 219, 237,
            177, 116, 102, 226, 52, 197, 186, 162, 150, 141, 185, 166, 187, 252,
            16, 178, 248, 48,
          ]),
          new Uint8Array([
            138, 8, 15, 95, 176, 175, 170, 18, 16, 248, 114, 107, 122, 244, 63,
            21, 78, 52, 154, 246, 141, 156, 108, 147, 187, 102, 129, 68, 170, 4,
            235, 91,
          ]),
          new Uint8Array([
            214, 117, 31, 212, 100, 239, 146, 4, 63, 3, 214, 226, 90, 226, 126,
            120, 106, 180, 10, 211, 3, 171, 225, 226, 159, 147, 22, 89, 246, 10,
            1, 227,
          ]),
          new Uint8Array([
            244, 228, 147, 22, 2, 149, 21, 139, 149, 194, 174, 151, 238, 255, 9,
            214, 128, 13, 249, 79, 146, 90, 120, 206, 111, 230, 247, 67, 29,
            126, 116, 65,
          ]),
          new Uint8Array([
            63, 10, 236, 59, 112, 243, 84, 242, 178, 204, 173, 187, 135, 52,
            181, 168, 108, 129, 163, 43, 241, 230, 37, 53, 214, 174, 54, 6, 238,
            27, 48, 6,
          ]),
        ],
        key: 23,
        numLeaves: 400,
      },
    },
  },

  // 0x4a63bb258a8b7ecde64390ae55072309cfa5894682f84b4ee6dade945e47524f
  {
    data: [
      Buffer.from(
        "AAAAAAAAAAAAAAAAAAAAAAAAAABsaWdodGxpbmsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACIyeQb+guQs6rAwPkCH/kCGqAXjyMTlzWXrODNBxdGlaBZKsvZgy4LhxzjkdvxzqP2mqAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJTfrRV7jU5Ywmv5uUf451ta28eCK6CTnqUEG7T02VZvvWEsOwaHE/G/7EgCJmOt0P9RQJtVs6AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALkBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "base64",
      ),
      Buffer.from(
        "AAAAAAAAAAAAAAAAAAAAAAAAAABsaWdodGxpbmsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAggH0hAQkOimD5OHAgIRl08c6oAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiMnkG/oLkLOqwMD5Ah/5AhqgnZAKAgneTE4DphKu2cMPg+UqP2d8KQJqq0MO0+RXuyigAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACU360Ve41OWMJr+blH+OdbWtvHgiugk56lBBu09NlWb71hLDsGhxPxv+xIAiZjrdD/UUCbVbOgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC5AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
        "base64",
      ),
    ],
    shareProofs: [
      {
        beginKey: 5,
        endKey: 7,
        sideNodes: [
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
              105, 103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
            ]),
          },
          {
            min: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            max: {
              version: new Uint8Array([0]),
              id: new Uint8Array([
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
                105, 103, 104, 116, 108, 105, 110, 107,
              ]),
            },
            digest: new Uint8Array([
              255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
              255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
              255, 255, 255, 255, 255, 255,
            ]),
          },
        ],
      },
    ],
    namespace: {
      version: new Uint8Array([0]),
      id: new Uint8Array([
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105, 103,
        104, 116, 108, 105, 110, 107,
      ]),
    },
    rowRoots: [
      {
        min: {
          version: new Uint8Array([0]),
          id: new Uint8Array([
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
            103, 104, 116, 108, 105, 110, 107,
          ]),
        },
        max: {
          version: new Uint8Array([0]),
          id: new Uint8Array([
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
            103, 104, 116, 108, 105, 110, 107,
          ]),
        },
        digest: new Uint8Array([
          0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 108, 105,
          103, 104, 116, 108, 105, 110, 107, 0, 0, 0,
        ]),
      },
    ],
    rowProofs: [
      {
        sideNodes: [
          new Uint8Array([
            203, 127, 36, 204, 55, 238, 20, 207, 232, 182, 20, 179, 166, 106,
            186, 133, 65, 17, 179, 254, 15, 153, 245, 147, 71, 188, 28, 152,
            129, 102, 118, 13,
          ]),
          new Uint8Array([
            114, 18, 74, 162, 231, 88, 150, 155, 51, 78, 64, 219, 118, 158, 199,
            87, 157, 118, 181, 158, 175, 160, 221, 176, 247, 172, 20, 148, 193,
            24, 174, 245,
          ]),
          new Uint8Array([
            169, 62, 204, 43, 160, 194, 121, 168, 116, 42, 249, 206, 13, 154,
            142, 255, 158, 126, 118, 224, 247, 117, 111, 79, 7, 167, 4, 138, 15,
            82, 71, 184,
          ]),
          new Uint8Array([
            216, 25, 139, 170, 146, 41, 202, 122, 106, 165, 201, 174, 182, 41,
            22, 159, 194, 13, 182, 97, 71, 24, 4, 219, 150, 59, 110, 98, 203,
            171, 129, 94,
          ]),
          new Uint8Array([
            70, 17, 32, 130, 228, 0, 72, 120, 40, 46, 94, 237, 60, 221, 222,
            130, 208, 137, 187, 145, 74, 152, 163, 6, 148, 112, 190, 42, 169,
            169, 77, 186,
          ]),
          new Uint8Array([
            14, 146, 27, 77, 228, 64, 249, 206, 244, 6, 179, 204, 41, 210, 187,
            196, 31, 181, 237, 129, 97, 136, 6, 135, 28, 175, 60, 153, 91, 65,
            128, 24,
          ]),
          new Uint8Array([
            6, 146, 134, 122, 130, 228, 68, 202, 204, 171, 210, 181, 245, 60,
            204, 69, 189, 170, 90, 76, 14, 189, 168, 78, 186, 144, 178, 76, 160,
            126, 117, 174,
          ]),
          new Uint8Array([
            140, 66, 35, 195, 166, 33, 48, 120, 45, 213, 132, 209, 22, 140, 128,
            54, 4, 130, 147, 245, 29, 51, 104, 183, 146, 46, 203, 25, 181, 189,
            60, 161,
          ]),
        ],
        key: 39,
        numLeaves: 256,
      },
    ],
    attestationProof: {
      tupleRootNonce: 3383,
      tuple: {
        height: 1215624,
        dataRoot: new Uint8Array([
          69, 207, 56, 157, 215, 169, 201, 84, 110, 149, 228, 243, 38, 249, 107,
          40, 240, 138, 215, 221, 91, 153, 96, 142, 215, 190, 72, 98, 98, 183,
          251, 22,
        ]),
      },
      proof: {
        sideNodes: [
          new Uint8Array([
            238, 112, 187, 95, 254, 122, 188, 44, 151, 221, 82, 132, 217, 153,
            247, 122, 209, 45, 229, 229, 101, 210, 115, 225, 182, 100, 220, 71,
            160, 49, 219, 82,
          ]),
          new Uint8Array([
            41, 178, 137, 107, 95, 47, 61, 150, 156, 46, 84, 77, 243, 208, 82,
            209, 80, 142, 247, 95, 155, 239, 147, 246, 166, 22, 148, 139, 33,
            239, 206, 72,
          ]),
          new Uint8Array([
            108, 20, 31, 228, 135, 184, 78, 33, 113, 130, 104, 139, 150, 94,
            250, 34, 195, 186, 185, 29, 96, 184, 5, 201, 157, 183, 119, 243, 61,
            165, 119, 162,
          ]),
          new Uint8Array([
            61, 12, 254, 109, 182, 152, 247, 238, 136, 122, 194, 137, 219, 164,
            197, 61, 91, 11, 170, 209, 115, 62, 3, 73, 74, 113, 58, 117, 125,
            16, 9, 157,
          ]),
          new Uint8Array([
            215, 167, 20, 40, 195, 168, 241, 214, 121, 140, 158, 7, 219, 237,
            177, 116, 102, 226, 52, 197, 186, 162, 150, 141, 185, 166, 187, 252,
            16, 178, 248, 48,
          ]),
          new Uint8Array([
            138, 8, 15, 95, 176, 175, 170, 18, 16, 248, 114, 107, 122, 244, 63,
            21, 78, 52, 154, 246, 141, 156, 108, 147, 187, 102, 129, 68, 170, 4,
            235, 91,
          ]),
          new Uint8Array([
            214, 117, 31, 212, 100, 239, 146, 4, 63, 3, 214, 226, 90, 226, 126,
            120, 106, 180, 10, 211, 3, 171, 225, 226, 159, 147, 22, 89, 246, 10,
            1, 227,
          ]),
          new Uint8Array([
            244, 228, 147, 22, 2, 149, 21, 139, 149, 194, 174, 151, 238, 255, 9,
            214, 128, 13, 249, 79, 146, 90, 120, 206, 111, 230, 247, 67, 29,
            126, 116, 65,
          ]),
          new Uint8Array([
            63, 10, 236, 59, 112, 243, 84, 242, 178, 204, 173, 187, 135, 52,
            181, 168, 108, 129, 163, 43, 241, 230, 37, 53, 214, 174, 54, 6, 238,
            27, 48, 6,
          ]),
        ],
        key: 23,
        numLeaves: 400,
      },
    },
  },
];

const shareRanges: ChainOracle.ShareRangeStruct[][] = [
  [
    { start: 83, end: 512 },
    { start: 30, end: 110 },
  ],
  [
    { start: 115, end: 512 },
    { start: 30, end: 145 },
  ],
];

export const MOCK_DATA = {
  rollupHeaders,
  l2HeaderHashes,
  l2Headers,
  shareProofs,
  shareRanges,
};
