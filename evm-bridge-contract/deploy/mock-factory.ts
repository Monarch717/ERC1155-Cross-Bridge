import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { MockFactory__factory } from "../typechain";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();

  const MockFactory = (await ethers.getContractFactory("MockFactory", signers[0])) as MockFactory__factory
  const mockFactory = await MockFactory.deploy();

  console.log(`>> MockFactory Address: ${mockFactory.address}`);
}

func.tags = ["MockFactory"];

export default func;
