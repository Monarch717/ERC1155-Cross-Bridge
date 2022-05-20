import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { swapERC721 } from './utils/721-agent';
import { get721Agent, get721MirroredToken } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }

  const chainId = hre.network.config.chainId?.toString() || '';
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();

  console.warn(`[${func.tags![0]}]: chains/${chainId}/erc721-mirrored-token.json must be prior created with completed details and the token should be prepared before swapping`);

  const signers = await ethers.getSigners();
  const mirrored = get721MirroredToken(chainId);
  const agent = get721Agent(chainId);

  if (!mirrored.address) {
    throw new Error("no mirrored token address");
  }
  if (!mirrored.tokenId) {
    throw new Error("no mirrored tokenId");
  }

  await swapERC721({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mirrored.address,
    recipient: signers[0].address,
    tokenId: mirrored.tokenId,
    signers,
  });
}

func.tags = ["ERC721BackwardSwap"];

export default func;
