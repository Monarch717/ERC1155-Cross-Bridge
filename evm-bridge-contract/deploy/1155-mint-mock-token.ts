import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { ethers } from "hardhat";
import { mintMockERC1155 } from './utils/1155-token';
import { get1155MockToken, set1155MockToken } from './utils/1155-cache';

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const signers = await ethers.getSigners();
  const chainId = hre.network.config.chainId?.toString() || '';
  const mockToken = get1155MockToken(chainId);

  let nextId = 1;
  if (mockToken.ids.length) {
    const lastIdx = mockToken.ids.length - 1;
    nextId = mockToken.ids[lastIdx] + 1;
  }

  mockToken.ids = [...mockToken.ids, nextId];
  mockToken.amounts = [...mockToken.amounts, 1];
  await mintMockERC1155({
    ids: mockToken.ids,
    amounts: mockToken.amounts,
    to: signers[0].address,
    tokenContractAddr: mockToken.address,
    signers,
  });

  set1155MockToken(chainId, mockToken);
}

func.tags = ["ERC1155MintMockToken"];

export default func;
