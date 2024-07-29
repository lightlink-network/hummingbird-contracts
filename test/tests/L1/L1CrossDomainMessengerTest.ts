import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { pushRandomHeader, setupCanonicalStateChain } from "../../lib/chain";
import {
  LightLinkPortal,
  L1CrossDomainMessenger,
} from "../../../typechain-types";
import { proxyDeployAndInitialize } from "../../../scripts/hardhat/lib/deploy";
import { l1 } from "../../../typechain-types/contracts";

describe("L1CrossDomainMessenger", function () {
  let lightLinkPortal: LightLinkPortal;
  let l1CrossDomainMessenger: L1CrossDomainMessenger;
  let owner: HardhatEthersSigner, _chain: any;

  beforeEach(async function () {
    [owner] = await ethers.getSigners();

    // - LightLinkPortal
    const lightLinkPortalDeployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("LightLinkPortal"),
      [ethers.ZeroAddress, ethers.ZeroAddress],
    );
    lightLinkPortal = lightLinkPortalDeployment.contract as LightLinkPortal;

    // - L1CrossDomainMessenger
    const l1CrossDomainMessengerDeployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("L1CrossDomainMessenger"),
      [lightLinkPortalDeployment.address],
    );

    l1CrossDomainMessenger =
      l1CrossDomainMessengerDeployment.contract as L1CrossDomainMessenger;
  });

  describe("Deployment", function () {
    it("Should set the right LightLinkPortal", async function () {
      expect(await l1CrossDomainMessenger.PORTAL()).to.equal(
        await lightLinkPortal.getAddress(),
      );
    });

    it("Should not be allowed to initialize twice", async function () {
      await expect(
        l1CrossDomainMessenger
          .connect(owner)
          .getFunction("initialize")
          .send(await lightLinkPortal.getAddress()),
      ).to.be.revertedWithCustomError(
        l1CrossDomainMessenger,
        "InvalidInitialization",
      );
    });
  });

  // Internal
  //   describe("gasPayingToken", function () {
  //     it("Should return ETHER address and decimals 18", async function () {
  //        expect( await l1CrossDomainMessenger.
  //         l1CrossDomainMessenger.
  //           .connect(owner)
  //           .getFunction("gasPayingToken")
  //           .send(),
  //       ).to.be.equal("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE");
  //     });
  //   });
});
