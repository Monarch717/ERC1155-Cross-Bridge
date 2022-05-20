import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { deployERC1155Agent } from './utils/1155-deploy';
import { set1155Agent } from './utils/1155-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const agent = await deployERC1155Agent({
    chainId,
    signers,
  });

  set1155Agent(chainId, {
    address: agent.address,
    registerTxHash: '',
    registerBlockNumber: '',
  });
}

func.tags = ["ERC1155SwapAgent"];

export default func;
