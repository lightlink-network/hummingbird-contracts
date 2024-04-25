import { CanonicalStateChain, ChainOracle } from "../../typechain-types";
import { SharesProofStruct } from "../../typechain-types/contracts/interfaces/IChainOracle";

interface headerData {
  header: ChainOracle.L2HeaderStruct;
  headerHash: string;
  shareProofs: SharesProofStruct;
  shareRanges: ChainOracle.ShareRangeStruct[];
  shares: string[];
}

interface MockData {
  rollupHash: string;
  rollupHeader: CanonicalStateChain.HeaderStruct;
  headers: headerData[];
}

const header1: headerData = {
  header: {
    parentHash:
      "0x7a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0x27a43969c92114604269627a6be91f26e65beb2a965526c7cb38908641821d67",
    transactionsRoot:
      "0x62968cbf343e1f02ede6ed4be2dc3ca5d6d5daa8dfdf4ef57fe88228b7022b10",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: "0x1f4",
    number: "0x49607df",
    gasLimit: "0xe4e1c0",
    gasUsed: "0x313a2",
    timestamp: "0x660d30cd",
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: "0xc9e41bfa0b90b3aa",
  },
  headerHash:
    "0xcadf01f27275f6927d028471b033909ae89a1a3261ed4a5b34a9a8a1b55bc213",
  shareProofs: {
    attestationProof: {
      proof: {
        key: "275",
        numLeaves: "300",
        sideNodes: [
          "0x72833e0a7af294800fd0c553143a553c29fc9d260f867f885f3023dcad35d118",
          "0x3ba7b012ab170b22734a0c588a4fca004c33df42cb6880882cab2a224dcdaef9",
          "0x75d81f9a0e4ba03fcdb6c9ef357ae229b47b754d9c0f732bd3691b7449f9faa0",
          "0xd879b4943d7237609648149f69d6fbd6b7773db7fefb1b432bf52719393421e9",
          "0xb387dbd207cb49b1f4168608466ece03f78abb234cb3f9571d6f06456710b91e",
          "0xb510ef0669e5d9c1b261655fd826b60d6bf1c274585b92051906a91ed7069387",
          "0xf69e6cf8fe95fe1d49343bfb0261ba2cc779dd06b2a09f1fa0b29fbb82e3d697",
        ],
      },
      tuple: {
        dataRoot:
          "0xb54f1feb1ae5b5659f21f9fb6492fbb7f1bb99c1f4c748ec84898aa62685e800",
        height: "1525175",
      },
      tupleRootNonce: "416",
    },
    data: [
      "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f90b55f9021da07a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba027a43969c92114604269627a6be91f26e65beb2a965526c7cb38908641821d67a062968cbf343e1f02ede6ed4be2dc3ca5d6d5daa8dfdf4ef57fe88228b7022b10a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
      "0x00000000000000000000000000000000000000006c696768746c696e6b000000000000008201f484049607df83e4e1c0830313a284660d30cda00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aaf90931f9092e8303d787830f424083061a8094f0ef43a073ee30239a1ecc84e1dadc8895dbc1c280b908c4fe61e35c00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000000860000000000000000000000000000000000000000000000000000000000000088000000000000000000000000000000000000000000000000000000000000008a00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000",
    ],
    namespace: {
      id: "0x000000000000000000000000000000000000006c696768746c696e6b",
      version: "0x00",
    },
    rowProofs: [
      {
        key: "1",
        numLeaves: "32",
        sideNodes: [
          "0x425bff78b7f5f6048d60b3f9eee250a8e0012c6afab338ecb0105f37f17ac1d2",
          "0x10398f5b43bcc56496b9ba89c0baf2a2cb7aec2760c3e141b1eac51e6f537f2e",
          "0xad86d83623c1bdd466fa2d053b120fef8b75d73147be687335b56897f3e1f8b2",
          "0x0ca73c0b1763796cc5bfbb9205f05819788aa4f478b7397080ffaaaa9e1f278a",
          "0x814554abc490cc61309298d97f8c280bd881782a741bf65f58ad7a5ea742257c",
        ],
      },
    ],
    rowRoots: [
      {
        digest:
          "0x3452fe6ebe4f5919b093c992264f9c1f76dd71caa6f9a9e73c08e3b8cb33e63b",
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
        beginKey: "4",
        endKey: "6",
        sideNodes: [
          {
            digest:
              "0xf83d635048a7df6c779acdb7b5f01d67df94b2faf6492670d112d734dbf3cfee",
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
              "0x2b468abed67a311eed2594f76cf1d919f4d5ec1858ff28587abf96bfb283821d",
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
              "0xb390bfd76fde7eccee1f13a77cf5c930cb700fb8c2045784c090a351f5739e7c",
            max: {
              id: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
              version: "0xff",
            },
            min: {
              id: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
              version: "0xff",
            },
          },
        ],
      },
    ],
  },
  shareRanges: [
    {
      end: 512n,
      start: 70n,
    },
    {
      end: 132n,
      start: 30n,
    },
  ],
  shares: [
    "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f90b55f9021da07a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba027a43969c92114604269627a6be91f26e65beb2a965526c7cb38908641821d67a062968cbf343e1f02ede6ed4be2dc3ca5d6d5daa8dfdf4ef57fe88228b7022b10a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "0x00000000000000000000000000000000000000006c696768746c696e6b000000000000008201f484049607df83e4e1c0830313a284660d30cda00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aaf90931f9092e8303d787830f424083061a8094f0ef43a073ee30239a1ecc84e1dadc8895dbc1c280b908c4fe61e35c00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000000860000000000000000000000000000000000000000000000000000000000000088000000000000000000000000000000000000000000000000000000000000008a00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000",
  ],
};

const header0: headerData = {
  header: {
    parentHash:
      "0x558baf050f7a9f97d698384ccaa6894d62099a3cf674f71d6ce0e442dde9cb54",
    uncleHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    beneficiary: "0xdfad157b8d4e58c26bf9b947f8e75b5adbc7822b",
    stateRoot:
      "0xc53b0ee15d38bb949f4c8225eb6e9e67ff303d4a4f53e39a4bb0af6cd45fd115",
    transactionsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    receiptsRoot:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    logsBloom:
      "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    difficulty: "0x1f4",
    number: "0x49607de",
    gasLimit: "0xe4e1c0",
    gasUsed: "0x0",
    timestamp: "0x660d30cd",
    extraData: "0x",
    mixHash:
      "0x0000000000000000000000000000000000000000000000000000000000000000",
    nonce: "0xc9e41bfa0b90b3aa",
  },
  headerHash:
    "0x7a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89",
  shareProofs: {
    attestationProof: {
      proof: {
        key: "275",
        numLeaves: "300",
        sideNodes: [
          "0x72833e0a7af294800fd0c553143a553c29fc9d260f867f885f3023dcad35d118",
          "0x3ba7b012ab170b22734a0c588a4fca004c33df42cb6880882cab2a224dcdaef9",
          "0x75d81f9a0e4ba03fcdb6c9ef357ae229b47b754d9c0f732bd3691b7449f9faa0",
          "0xd879b4943d7237609648149f69d6fbd6b7773db7fefb1b432bf52719393421e9",
          "0xb387dbd207cb49b1f4168608466ece03f78abb234cb3f9571d6f06456710b91e",
          "0xb510ef0669e5d9c1b261655fd826b60d6bf1c274585b92051906a91ed7069387",
          "0xf69e6cf8fe95fe1d49343bfb0261ba2cc779dd06b2a09f1fa0b29fbb82e3d697",
        ],
      },
      tuple: {
        dataRoot:
          "0xb54f1feb1ae5b5659f21f9fb6492fbb7f1bb99c1f4c748ec84898aa62685e800",
        height: "1525175",
      },
      tupleRootNonce: "416",
    },
    data: [
      "0x00000000000000000000000000000000000000006c696768746c696e6b00822ba0c53b0ee15d38bb949f4c8225eb6e9e67ff303d4a4f53e39a4bb0af6cd45fd115a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f484049607dd83e4e1c08084660d30cca00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa0558baf050f7a9f97d698384ccaa6894d62099a3c",
      "0x00000000000000000000000000000000000000006c696768746c696e6b00f674f71d6ce0e442dde9cb54a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba0c53b0ee15d38bb949f4c8225eb6e9e67ff303d4a4f53e39a4bb0af6cd45fd115a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f484049607de83e4e1c08084660d30cda00000000000000000000000000000000000000000000000000000000000000000a0000000000000",
      "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f90b55f9021da07a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba027a43969c92114604269627a6be91f26e65beb2a965526c7cb38908641821d67a062968cbf343e1f02ede6ed4be2dc3ca5d6d5daa8dfdf4ef57fe88228b7022b10a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    ],
    namespace: {
      id: "0x000000000000000000000000000000000000006c696768746c696e6b",
      version: "0x00",
    },
    rowProofs: [
      {
        key: "1",
        numLeaves: "32",
        sideNodes: [
          "0x425bff78b7f5f6048d60b3f9eee250a8e0012c6afab338ecb0105f37f17ac1d2",
          "0x10398f5b43bcc56496b9ba89c0baf2a2cb7aec2760c3e141b1eac51e6f537f2e",
          "0xad86d83623c1bdd466fa2d053b120fef8b75d73147be687335b56897f3e1f8b2",
          "0x0ca73c0b1763796cc5bfbb9205f05819788aa4f478b7397080ffaaaa9e1f278a",
          "0x814554abc490cc61309298d97f8c280bd881782a741bf65f58ad7a5ea742257c",
        ],
      },
    ],
    rowRoots: [
      {
        digest:
          "0x3452fe6ebe4f5919b093c992264f9c1f76dd71caa6f9a9e73c08e3b8cb33e63b",
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
        beginKey: "2",
        endKey: "5",
        sideNodes: [
          {
            digest:
              "0x547bb6ac124764ba18e2234d5eb0eff9cb30d6ebc418c3cff9a8589bec72e279",
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
              "0x96194e8cfc8ae5b9bb542b5ff12e12116df9d8b4f814a0f864275579cbd4f333",
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
              "0x2b468abed67a311eed2594f76cf1d919f4d5ec1858ff28587abf96bfb283821d",
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
              "0xb390bfd76fde7eccee1f13a77cf5c930cb700fb8c2045784c090a351f5739e7c",
            max: {
              id: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
              version: "0xff",
            },
            min: {
              id: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
              version: "0xff",
            },
          },
        ],
      },
    ],
  },
  shareRanges: [
    {
      end: 512n,
      start: 488n,
    },
    {
      end: 512n,
      start: 30n,
    },
    {
      end: 65n,
      start: 30n,
    },
  ],
  shares: [
    "0x00000000000000000000000000000000000000006c696768746c696e6b00822ba0c53b0ee15d38bb949f4c8225eb6e9e67ff303d4a4f53e39a4bb0af6cd45fd115a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f484049607dd83e4e1c08084660d30cca00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f9021ff9021aa0558baf050f7a9f97d698384ccaa6894d62099a3c",
    "0x00000000000000000000000000000000000000006c696768746c696e6b00f674f71d6ce0e442dde9cb54a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba0c53b0ee15d38bb949f4c8225eb6e9e67ff303d4a4f53e39a4bb0af6cd45fd115a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f484049607de83e4e1c08084660d30cda00000000000000000000000000000000000000000000000000000000000000000a0000000000000",
    "0x00000000000000000000000000000000000000006c696768746c696e6b00000000000000000000000000000000000000000000000000000088c9e41bfa0b90b3aac0c0f90b55f9021da07a52b3418d621ebe5598ab4aa43d84df49c5a7723198afeee2e6457b4aeeea89a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba027a43969c92114604269627a6be91f26e65beb2a965526c7cb38908641821d67a062968cbf343e1f02ede6ed4be2dc3ca5d6d5daa8dfdf4ef57fe88228b7022b10a00000000000000000000000000000000000000000000000000000000000000000b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  ],
};

export const chainOracleMockData: MockData[] = [
  {
    rollupHash:
      "0x69ebdd08e7d3a7aaefa29edf185f0f1894c6069ea012ab777d9464657cc5f03b",
    rollupHeader: {
      epoch: 29850553n,
      l2Height: 76941291n,
      prevHash:
        "0x8832c48036b142c433909b97f6d4afb316be21ead565fb9343fbb359665b9bda",
      stateRoot:
        "0x25a827ad1a5f4828337b12a979a2151f5398c4a726dc2599440d720d62400061",
      celestiaPointers: [
        {
          height: 1525175n,
          shareStart: 3n,
          shareLen: 17n,
        },
        {
          height: 1525176n,
          shareStart: 35n,
          shareLen: 17n,
        },
      ],
    },
    headers: [header0, header1],
  },
];
