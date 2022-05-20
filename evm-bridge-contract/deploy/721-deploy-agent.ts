import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { deployERC721Agent } from './utils/721-deploy';
import { set721Agent } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const agent = await deployERC721Agent({
    chainId,
    signers,
  });

  set721Agent(chainId, {
    address: agent.address,
    registerTxHash: '',
    registerBlockNumber: '',
  });
}

func.tags = ["ERC721SwapAgent"];

export default func;
