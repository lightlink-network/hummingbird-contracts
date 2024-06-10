import { ethers } from "hardhat";
import { expect } from "chai";
import type { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  EnterprisePortal,
  EnterpriseGasStation,
  EnterpriseGasStation__factory,
} from "../../typechain-types";
import { proxyDeployAndInitialize } from "../../scripts/universal/utils";

describe("EnterpriseGasStation", function () {
  let owner: HardhatEthersSigner,
    publisher: HardhatEthersSigner,
    otherAccount: HardhatEthersSigner;

  let portal: EnterprisePortal;
  let gasStation: EnterpriseGasStation;
  let llToken: any;

  beforeEach(async function () {
    [owner, publisher, otherAccount] = await ethers.getSigners();

    // Deploy Enterprise Portal
    const portal = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("EnterprisePortal"),
      [],
    );

    // Deploy LL Token without proxy
    const llTokenFactory: any = await ethers.getContractFactory("LLERC20");
    const llToken = await llTokenFactory
      .connect(owner)
      .deploy("1000000000000000000");
    await llToken.waitForDeployment();

    // Deploy Enterprise Gas Station
    const deployment = await proxyDeployAndInitialize(
      owner,
      await ethers.getContractFactory("EnterpriseGasStation"),
      [
        portal.address,
        await llToken.getAddress(),
        ethers.ZeroAddress, // reserve wallet
      ],
    );

    gasStation = EnterpriseGasStation__factory.connect(
      deployment.address,
      owner,
    );
  });

  describe("initialize", function () {
    it("happy path", async function () {
      expect(await gasStation.reserveWallet()).to.equal(ethers.ZeroAddress);
    });
  });

  describe("addGasPlan", function () {
    it("should allow the owner to add a new gas plan", async function () {
      await gasStation.addGasPlan(100, 1000, true);
      const plan = await gasStation.gasPlans(0);
      expect(plan.costInTokens).to.equal(100);
      expect(plan.gasAmount).to.equal(1000);
      expect(plan.isActive).to.be.true;
    });

    it("should emit GasPlanAdded event when a new gas plan is added", async function () {
      await expect(gasStation.addGasPlan(100, 1000, true))
        .to.emit(gasStation, "GasPlanAdded")
        .withArgs(0, 100, 1000, true);
    });
  });

  describe("setGasPlanStatus", function () {
    beforeEach(async function () {
      await gasStation.addGasPlan(100, 1000, true);
    });

    it("should allow the owner to update the status of an existing gas plan", async function () {
      await gasStation.setGasPlanStatus(0, false);
      const plan = await gasStation.gasPlans(0);
      expect(plan.isActive).to.be.false;
    });

    it("should emit GasPlanStatusUpdated event when the status of a gas plan is updated", async function () {
      await expect(gasStation.setGasPlanStatus(0, false))
        .to.emit(gasStation, "GasPlanStatusUpdated")
        .withArgs(0, false);
    });

    it("should revert if the gas plan does not exist", async function () {
      await expect(gasStation.setGasPlanStatus(1, false)).to.be.revertedWith(
        "Gas plan does not exist",
      );
    });
  });
});
