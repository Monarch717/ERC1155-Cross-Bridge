import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { approveMirroredERC721 } from './utils/721-token';
import { get721Agent, get721MirroredToken } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const chainId = hre.network.config.chainId?.toString() || '';

  console.warn(`[${func.tags![0]}]: chains/${chainId}/erc721-mirrored-token.json must be prior created with completed details and the token should be prepared before swapping`);

  const signers = await ethers.getSigners();
  const mirrroredToken = get721MirroredToken(chainId);
  const agent = get721Agent(chainId);

  await approveMirroredERC721({
    tokenId: mirrroredToken.tokenId,
    to: agent.address,
    tokenContractAddr: mirrroredToken.address,
    signers,
  });
}

func.tags = ["ERC721ApproveMirroredToken"];

export default func;
