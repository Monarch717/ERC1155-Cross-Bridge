import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { swapERC721 } from './utils/721-agent';
import { get721Agent, get721MockToken } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);
  const agent = get721Agent(chainId);

  await swapERC721({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mockToken.address,
    recipient: signers[0].address,
    tokenId: mockToken.tokenId,
    signers,
  });
}

func.tags = ["ERC721ForwardSwap"];

export default func;
