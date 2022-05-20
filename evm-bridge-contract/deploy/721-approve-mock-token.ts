import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { approveMockERC721 } from './utils/721-token';
import { get721MockToken, get721Agent } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);
  const agent = get721Agent(chainId);

  await approveMockERC721({
    tokenId: mockToken.tokenId,
    to: agent.address,
    tokenContractAddr: mockToken.address,
    signers,
  });
}

func.tags = ["ERC721ApproveMockToken"];

export default func;
