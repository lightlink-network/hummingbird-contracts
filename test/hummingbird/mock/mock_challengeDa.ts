import {
  CanonicalStateChain,
  ChallengeDataAvailability,
} from "../../../typechain-types";
import {
  SharesProofStruct,
  BinaryMerkleProofStruct,
} from "../../../typechain-types/contracts/challenge/Challenge";

interface MockData {
  rollupHeader: CanonicalStateChain.HeaderStruct;
  daProofs: {
    key: string;
    pointerIndex: number;
    shareIndex: number;
    shareProof: SharesProofStruct;
    shareToRBlockRootProof: BinaryMerkleProofStruct;
  };
}

export const challengeDAMockData: MockData = {
  daProofs: {
    key: "0x0ef104ca1c2c1d6e02a2fb58921e3615eb46fd445ba4492a3de4216ced807a15",
    pointerIndex: 0,
    shareIndex: 293,
    shareProof: {
      attestationProof: {
        proof: {
          key: "9",
          numLeaves: "300",
          sideNodes: [
            "0xa53b5edaee3f0fcd9c7c4b4611b2ae5f77895c974aa03eb24fd8c695e94a9ac4",
            "0x6f52a249d8945c5e06defc98fa14a7d69118885fdac52cfbc00d9523b4a1ece1",
            "0x0109a872b54f09567f56fdf42a0520476d36f4ee333e37d0393646a499cd48ac",
            "0x0cf80322d5fd196bcc67ff8990afb769a8fc1c7d7c1108c805d547db5dbb1ba8",
            "0xe3b4f44ad2b690745d9356e3b8e61e3563adc679bb5fa91c6f20966dd813c515",
            "0x8636d42a06095d53e7c0c2f17b389a46ccd98d82adbce5232e9258638edfeb5a",
            "0x17702c2c6dd4c009fb8396e1719dedd39fd561b0f56ea6c00374430db108e425",
            "0x708a31bc4610c9fd813a6bafaef07ac9749be0dbe4f2d20a039ac8a8cb7b2c89",
            "0xf43f97ef4b4f5914037a00c1f294297378da06484a6245efa34b8c8622a8abe0",
          ],
        },
        tuple: {
          dataRoot:
            "0xb4b5e9fd4328e422693bc1e3f8db2d191ca9a192ed16fd0ee83cc499676c2ce8",
          height: "1826709",
        },
        tupleRootNonce: "887",
      },
      data: [
        "0x00000000000000000000000000000000000000006c696768746c696e6b0100002a57f92a54f9021ff9021aa030ce6b451df3634d7c89bcb02424d120868dcd36673cb558f00c67979fc4d752a0000000000000000000000000000000000000000000000000000000000000000094dfad157b8d4e58c26bf9b947f8e75b5adbc7822ba03bd868e8b2cffeb826011714ec72345bf139ddea66e1258f4d323bcf4456d5dba00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008201f4840501a1ee83e4e1c0808466437183a00000000000",
      ],
      namespace: {
        id: "0x000000000000000000000000000000000000006c696768746c696e6b",
        version: "0x00",
      },
      rowProofs: [
        {
          key: "9",
          numLeaves: "128",
          sideNodes: [
            "0x2711abdfc68a43e19e6506fc0852a09fcc32d0ba71d8634d398b0ef422b459fd",
            "0x0a944556b1a77867317c535e265ca8106c52c4154894d8c4aea4afbd432f8193",
            "0x21ea268d3540deab8f9b2024464618610c9a7083620badcf505bda647cc8e9f8",
            "0xf75abd59b084be9cd3929e420858d8a427081095161c8328d721aa1373e12402",
            "0x2bfc87d990d8344f6efd44fcb09b46b87f9a92230d41329452efee8656c6760a",
            "0x0e5c3a72dfa64533cc06de75ddbee1e08fd518d2d8199e15a303a819aec11d34",
            "0x6797ad4fead9e5520aabc48275d213a91238e77fec144a9ddcf6f8ff8c34770c",
          ],
        },
      ],
      rowRoots: [
        {
          digest:
            "0x454734ed8e56335c69537e381ee88d4ce46d6e6d90d745ca8a3e79de0177032f",
          max: {
            id: "0x00000000000000000000000000000000000040673ebb89d8a6793268",
            version: "0x00",
          },
          min: {
            id: "0x00000000000000000000000000000000000000000ed9f671055d3d32",
            version: "0x00",
          },
        },
      ],
      shareProofs: [
        {
          beginKey: "5",
          endKey: "6",
          sideNodes: [
            {
              digest:
                "0x43c97b3f0b9126dbac3bcefa7a8bd9d7811e96bc98cbd26cfc1fa31c411d865c",
              max: {
                id: "0x00000000000000000000000000000000000000000ed9f671055d3d32",
                version: "0x00",
              },
              min: {
                id: "0x00000000000000000000000000000000000000000ed9f671055d3d32",
                version: "0x00",
              },
            },
            {
              digest:
                "0x3307bab9676706627e82394ed6b14da97e301cf245641b99ccaad7739a868c77",
              max: {
                id: "0x00000000000000000000000000000000000000000ed9f671055d3d32",
                version: "0x00",
              },
              min: {
                id: "0x00000000000000000000000000000000000000000ed9f671055d3d32",
                version: "0x00",
              },
            },
            {
              digest:
                "0x67571c897e5bf60d8a42f257e62613a658f870452ada75e77e85308a742ced43",
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
                "0xf2b395f16f2f8c79a59a58b35088764d930fffa47dfca874633b68cc04494bd9",
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
                "0xb8a15898117c90db5626a165ccaa76fba2165059a743f6ba82b4c83608a23b70",
              max: {
                id: "0x00000000000000000000000000000000000040673ebb89d8a6793268",
                version: "0x00",
              },
              min: {
                id: "0x000000000000000000000000000000000000006c696768746c696e6b",
                version: "0x00",
              },
            },
            {
              digest:
                "0xe8f140f0510ba1ead00301fe11132f5f3f68e12ed20a60aa60d50bc3386138a4",
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
    shareToRBlockRootProof: {
      key: "0",
      numLeaves: "35",
      sideNodes: [
        "0x403e05150ea66aadcda8dfee4be184f38b3ad5dcc936db808d9a5edc8437e6b8",
        "0x172522a3d6d3148902bca05f14a7abcba5da539e99c617b34dac4c460460c3a1",
        "0xc1c4e2d5dd68d985a159153b858099e2550fd40aa640e60e998ce6e0b51368fe",
        "0x97e3e0b02cdbac846e1ff4b90333c341f946646742aa2f517439e8c1474ed39f",
        "0x650270c86eeb203a475d53f125ddef660bf373403268c9fd42bb56bc549dcffa",
        "0x3fd10406a7112db1722c31e4112815bdc92caf7e55d35d218bed12c87af8de69",
      ],
    },
  },
  rollupHeader: {
    celestiaPointers: [
      {
        height: 1826709,
        shareLen: 23,
        shareStart: "293",
      },
      {
        height: 1826712,
        shareLen: 12,
        shareStart: "111",
      },
    ],
    epoch: 5902275,
    l2Height: 83993089,
    prevHash:
      "0x61b42beb89c37b9c02b1e343719335369a966dde6abb1e48e5dabd54f8300ae1",
    stateRoot:
      "0x7d0f4357f5226358b2caa95be8452edda23ee88ae40a4de9351810589c2bdd8c",
  },
};
