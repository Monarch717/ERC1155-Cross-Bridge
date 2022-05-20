import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { approveMirroredERC1155 } from './utils/1155-token';
import { get1155Agent, get1155MirroredToken } from './utils/1155-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const chainId = hre.network.config.chainId?.toString() || '';

  console.warn(`[${func.tags![0]}]: chains/${chainId}/erc1155-mirrored-token.json must be prior created with completed details and the token should be prepared before swapping`);

  const signers = await ethers.getSigners();
  const mirrroredToken = get1155MirroredToken(chainId);
  const agent = get1155Agent(chainId);

  await approveMirroredERC1155({
    to: agent.address,
    tokenContractAddr: mirrroredToken.address,
    signers,
  });
}

func.tags = ["ERC1155ApproveMirroredToken"];

export default func;
