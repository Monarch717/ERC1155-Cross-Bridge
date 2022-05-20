import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { swapERC1155 } from './utils/1155-agent';
import { get1155Agent, get1155MockToken } from './utils/1155-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (!process.env.DST_CHAIN_ID) {
    throw new Error("no DST_CHAIN_ID");
  }
  const dstChainId = parseInt(process.env.DST_CHAIN_ID, 10).toString();
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);
  const agent = get1155Agent(chainId);

  await swapERC1155({
    agentAddr: agent.address,
    dstChainId,
    tokenAddr: mockToken.address,
    recipient: signers[0].address,
    ids: mockToken.ids,
    amounts: mockToken.amounts,
    signers,
  });
}

func.tags = ["ERC1155ForwardSwap"];

export default func;
