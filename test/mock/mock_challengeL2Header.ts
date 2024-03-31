import { ethers } from "hardhat";
import { CanonicalStateChain, ChainOracle } from "../../typechain-types";
import { SharesProofStruct } from "../../typechain-types/contracts/ChainOracle";

const rollupHeaders: CanonicalStateChain.HeaderStruct[] = [
  {
    epoch: BigInt(5394737),
    l2Height: BigInt(71321264),
    prevHash: "", // This will be replaced when pushed to the chain
    stateRoot:
      "0xa1baa0f693395b09151aaf19df587e7c0cfdeda557d71aa93f54ab1758f3e670",
    celestiaPointers: [{ height: 1286533n, shareStart: 64n, shareLen: 3465n }],
  },
];

const l2HeaderHashes: string[] = [
  "0x9d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315",
  "0x5d47205c2ea8c70dc35c76e9067724d2674538583d7b3762f77346487203449b",
];

const l2Headers: ChainOracle.L2HeaderStruct[] = [
  // 0x9d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315
  {
    parentHash:
      "0x59c64ea837712cb3c1a44273fb6a0eaf3a56d429d8cbf37c1bc9ba52d25be716",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0xc8ea9802ccee323836de401d50ac2baf871ed8eec8fd8785251b035efef9d6d0",
    transactionsRoot:
      "0x43c52c4f0bf3b0ef929caa8227c365cec71142e5b93dfd286627cad342c5f8b8",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: "0x1f4",
    number: "0x4403e7b",
    gasLimit: "0xe4e1c0",
    gasUsed: "0xdb93",
    timestamp: "0x65e1db2a",
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: BigInt("0xc9e41bfa0b90b3aa"),
  },
  // 0x5d47205c2ea8c70dc35c76e9067724d2674538583d7b3762f77346487203449b
  {
    parentHash:
      "0x9d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0x3df56e35990cebb521ddf415be1ca169d0486f6e9b4742f89df36c24276fbbee",
    transactionsRoot:
      "0x0cc14c6992b2960f0ce3d94e06692d512baada2d22a9a8cd03d7cefe40e90e8d",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: "0x1f4",
    number: "0x4403e7c",
    gasLimit: "0xe4e1c0",
    gasUsed: "0x64e1c",
    timestamp: "0x65e1db2b",
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: BigInt("0xc9e41bfa0b90b3aa"),
  },
];

const shareProofs: SharesProofStruct[] = [
  // 0x9d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315
  {
    attestationProof: {
      proof: {
        key: "132",
        numLeaves: "400",
        sideNodes: [
          "0x3c6ba044ba05ba8e76cc9dd5c6a1d97bfee61140efd98af5e7419b372923704b",
          "0x066e58eb3706491b556a0806d859da6f414d0b2fe58d5ef59ae5df382235df2c",
          "0x26e6c509bb6bcfce7ccb11b558cdbaad0d49f13463c5020d8fa94d3ceef3c16b",
          "0xdf2513bef49743d2c9f18c78d26475bd2159669de3b4eb67039b9064d49fe74b",
          "0x61d2f2e58ad61e40e43abcdd2946247fb9372189040ff2ffef372ddb4745cf25",
          "0x1ac7d724dae2312a05fd85c6a13db08b49a6406ca12b49ac228cbc704b3f75dd",
          "0xc987887bb57873e719394a8c5f700d71f1a0b58ba36c5ad4a70353f3c11cf650",
          "0xd599625d6619f628d78fc60d6f6e5a78ddacf05a8b2ec0775e005efecc769f51",
          "0x6c3833c8558b78e333725ba9c558e5dded06160b213bf0a3095c16bc66091cd7",
        ],
      },
      tuple: {
        dataRoot:
          "0x0d0f3d14f9b69bb2399edd18a6bc0b2fdd98c08d6c1aa76992d1d2af6f0a3bb1",
        height: "1286533",
      },
      tupleRootNonce: "3562",
    },
    data: [
      "0x00000000000000000000000000000000000000006c696768746c696e6b0083e4e1c0808465e1db2aa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f902aff9021ca059c64ea837712cb3c1a44273fb6a0eaf3a56d429d8cbf37c1bc9ba52d25be716a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba0c8ea9802ccee323836de401d50ac2baf871ed8eec8fd8785251b035efef9d6d0a043c52c4f0bf3b0ef929caa8227c365cec71142e5b93dfd286627cad342c5f8b8a00000000000000000000000000000000000000000000000000000000000000000b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      "0x00000000000000000000000000000000000000006c696768746c696e6b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48404403e7b83e4e1c082db938465e1db2aa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aaf88df88b83069b48830f424082db939402c27800f9c173e90dd6015d35989546e943afea80a487335443000000000000000000000000000000000000000000000000000000097d2ff2c0820ee9a06e3a89f403a2b04191435b7da3344c36b8c64381f19409f2a6f7c2002cc12183a01d09f03d226d8bbe5f83ed13b441a119bf6f000712709b7f08a1c56ac2ba3350c0f90958f9021da09d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03df56e35990cebb521ddf415be1ca169d0486f6e9b4742f89df36c24276fbbeea00cc14c6992b2960f0ce3d94e06692d512baada2d22a9a8cd03d7cefe40e90e8da0000000000000000000000000000000000000000000000000000000",
    ],
    namespace: {
      id: "0x000000000000000000000000000000000000006c696768746c696e6b",
      version: "0x00",
    },
    rowProofs: [
      {
        key: "17",
        numLeaves: "256",
        sideNodes: [
          "0xe52c7918cd5035d7d38b982d6f12c7fa695140c4d312cfbb8934f7ea069292cb",
          "0x442bf9f4570cc8e947687b79c281979b6cf9abd21f9d22277aa378fd51edca0e",
          "0x49d9c98413f399228fa09fae8789dc97df06b18dd0a1a2dfd95f171838239810",
          "0x0742c0aebffe6140e7abd718ef891acb458221f66221d419347a016ec5df305d",
          "0x9c355f7a9de08a1c653e0925d3a6d56133a296dcdd4a22b63848fd595a28ced0",
          "0xa2cc7c20143c24b795ef5ed02b53ccc566573a4a87a8d904c65046318e8fd383",
          "0x75d9458fdce308d14a835e3f0b47be154ad26eb386aa6fb5aae3b6745de9d5d5",
          "0x835d65543a214e98d3ec4028858997442de92edc4ded7581dee0952889d3e4b8",
        ],
      },
    ],
    rowRoots: [
      {
        digest:
          "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
        max: {
          id: "0x000000000000000000000000000000000000006c696768746c696e6b",
          version: "0x00",
        },
        min: {
          id: "0x000000000000000000000000000000000000006c696768746c696e6b",
          version: "0x00",
        },
      },
    ],
    shareProofs: [
      {
        beginKey: "6",
        endKey: "8",
        sideNodes: [
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
        ],
      },
    ],
  },

  // 0x5d47205c2ea8c70dc35c76e9067724d2674538583d7b3762f77346487203449b
  {
    attestationProof: {
      proof: {
        key: "132",
        numLeaves: "400",
        sideNodes: [
          "0x3c6ba044ba05ba8e76cc9dd5c6a1d97bfee61140efd98af5e7419b372923704b",
          "0x066e58eb3706491b556a0806d859da6f414d0b2fe58d5ef59ae5df382235df2c",
          "0x26e6c509bb6bcfce7ccb11b558cdbaad0d49f13463c5020d8fa94d3ceef3c16b",
          "0xdf2513bef49743d2c9f18c78d26475bd2159669de3b4eb67039b9064d49fe74b",
          "0x61d2f2e58ad61e40e43abcdd2946247fb9372189040ff2ffef372ddb4745cf25",
          "0x1ac7d724dae2312a05fd85c6a13db08b49a6406ca12b49ac228cbc704b3f75dd",
          "0xc987887bb57873e719394a8c5f700d71f1a0b58ba36c5ad4a70353f3c11cf650",
          "0xd599625d6619f628d78fc60d6f6e5a78ddacf05a8b2ec0775e005efecc769f51",
          "0x6c3833c8558b78e333725ba9c558e5dded06160b213bf0a3095c16bc66091cd7",
        ],
      },
      tuple: {
        dataRoot:
          "0x0d0f3d14f9b69bb2399edd18a6bc0b2fdd98c08d6c1aa76992d1d2af6f0a3bb1",
        height: "1286533",
      },
      tupleRootNonce: "3562",
    },
    data: [
      "0x00000000000000000000000000000000000000006c696768746c696e6b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48404403e7b83e4e1c082db938465e1db2aa00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aaf88df88b83069b48830f424082db939402c27800f9c173e90dd6015d35989546e943afea80a487335443000000000000000000000000000000000000000000000000000000097d2ff2c0820ee9a06e3a89f403a2b04191435b7da3344c36b8c64381f19409f2a6f7c2002cc12183a01d09f03d226d8bbe5f83ed13b441a119bf6f000712709b7f08a1c56ac2ba3350c0f90958f9021da09d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03df56e35990cebb521ddf415be1ca169d0486f6e9b4742f89df36c24276fbbeea00cc14c6992b2960f0ce3d94e06692d512baada2d22a9a8cd03d7cefe40e90e8da0000000000000000000000000000000000000000000000000000000",
      "0x00000000000000000000000000000000000000006c696768746c696e6b000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f48404403e7c83e4e1c083064e1c8465e1db2ba00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aaf90734f901ca82031a808301938a949be0838bc4926faa5f1937ab8f990ce662c8c17280b901649eaa5ba1000000000000000000000000eadd17d6a969f56a4d6ca63f2124de39c12d3cd73635653164623131633564393631316536666537626234620000000000000000000000000000000000000000000000",
    ],
    namespace: {
      id: "0x000000000000000000000000000000000000006c696768746c696e6b",
      version: "0x00",
    },
    rowProofs: [
      {
        key: "17",
        numLeaves: "256",
        sideNodes: [
          "0xe52c7918cd5035d7d38b982d6f12c7fa695140c4d312cfbb8934f7ea069292cb",
          "0x442bf9f4570cc8e947687b79c281979b6cf9abd21f9d22277aa378fd51edca0e",
          "0x49d9c98413f399228fa09fae8789dc97df06b18dd0a1a2dfd95f171838239810",
          "0x0742c0aebffe6140e7abd718ef891acb458221f66221d419347a016ec5df305d",
          "0x9c355f7a9de08a1c653e0925d3a6d56133a296dcdd4a22b63848fd595a28ced0",
          "0xa2cc7c20143c24b795ef5ed02b53ccc566573a4a87a8d904c65046318e8fd383",
          "0x75d9458fdce308d14a835e3f0b47be154ad26eb386aa6fb5aae3b6745de9d5d5",
          "0x835d65543a214e98d3ec4028858997442de92edc4ded7581dee0952889d3e4b8",
        ],
      },
    ],
    rowRoots: [
      {
        digest:
          "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
        max: {
          id: "0x000000000000000000000000000000000000006c696768746c696e6b",
          version: "0x00",
        },
        min: {
          id: "0x000000000000000000000000000000000000006c696768746c696e6b",
          version: "0x00",
        },
      },
    ],
    shareProofs: [
      {
        beginKey: "7",
        endKey: "9",
        sideNodes: [
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0x00000000000000000000000000000000000000006c696768746c696e6b000000",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
          {
            digest:
              "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
            max: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
            min: {
              id: "0x000000000000000000000000000000000000006c696768746c696e6b",
              version: "0x00",
            },
          },
        ],
      },
    ],
  },
];

const shareRanges: ChainOracle.ShareRangeStruct[][] = [
  // 0x9d1cf3aa5e493c95acf8d3795bd2d8e6d4b6f48d6ad32c0806837435294c1315
  [
    { end: 512, start: 120 },
    { end: 181, start: 30 },
  ],
  // 0x5d47205c2ea8c70dc35c76e9067724d2674538583d7b3762f77346487203449b
  [
    { end: 512, start: 328 },
    { end: 390, start: 30 },
  ],
];

export const MOCK_DATA = {
  rollupHeaders,
  l2HeaderHashes,
  l2Headers,
  shareProofs,
  shareRanges,
};
