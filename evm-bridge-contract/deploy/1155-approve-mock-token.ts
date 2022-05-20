import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { approveMockERC1155 } from './utils/1155-token';
import { get1155MockToken, get1155Agent } from './utils/1155-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);
  const agent = get1155Agent(chainId);

  await approveMockERC1155({
    to: agent.address,
    tokenContractAddr: mockToken.address,
    signers,
  });
}

func.tags = ["ERC1155ApproveMockToken"];

export default func;
