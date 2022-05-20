import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { deployMockERC721 } from './utils/721-deploy';
import { set721MockToken } from './utils/721-cache';

const BASE_URI = process.env.ERC721_MOCK_TOKEN_BASE_URL || '';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const baseUri = process.env.BASE_URI === 'true' ? BASE_URI : '';
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = await deployMockERC721({
    baseUri,
    signers,
  });

  set721MockToken(chainId, {
    address: mockToken.address,
    baseUri,
    symbol: await mockToken.symbol(),
    name: await mockToken.name(),
    tokenId: '',
  });
}

func.tags = ["ERC721MockToken"];

export default func;
