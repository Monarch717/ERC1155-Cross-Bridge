import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { mintMockERC721 } from './utils/721-token';
import { get721MockToken, set721MockToken } from './utils/721-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get721MockToken(chainId);

  mockToken.tokenId = (mockToken.tokenId ? parseInt(mockToken.tokenId, 10) + 1 : 1).toString();

  let tokenUri = mockToken.baseUri + mockToken.tokenId;
  if (mockToken.baseUri) {
    tokenUri = mockToken.tokenId;
  }

  await mintMockERC721({
    tokenId: mockToken.tokenId,
    tokenUri,
    to: signers[0].address,
    tokenContractAddr: mockToken.address,
    signers,
  });

  set721MockToken(chainId, mockToken);
}

func.tags = ["ERC721MintMockToken"];

export default func;
