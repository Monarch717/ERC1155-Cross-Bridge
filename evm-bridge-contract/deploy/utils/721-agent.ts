import { randomBytes } from 'crypto';
import { ethers } from "hardhat";
import { MockERC721__factory, ERC721SwapAgent__factory } from "../../typechain";
import { Awaited } from '../../types/utils';

type RegisterMockERC721Params = {
  agentAddr: string,
  dstChainId: string,
  tokenContractAddr: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
}

type CreateSwapPairERC721Params = {
  agentAddr: string,
  fromChainId: string,
  tokenContractAddr: string,
  tokenName: string,
  tokenSymbol: string,
  tokenBaseUri: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
}

type SwapERC721Params = {
  agentAddr: string,
  tokenAddr: string,
  recipient: string,
  tokenId: string,
  dstChainId: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
}

type FillERC721Params = {
  agentAddr: string,
  fromTokenAddr: string,
  fromChainId: string,
  recipient: string,
  tokenId: string,
  tokenURI: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
}

export const registerMockERC721 = async (params: RegisterMockERC721Params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC721 = (await ethers.getContractFactory("MockERC721", params.signers[0])) as MockERC721__factory;
  const mockERC721 = await MockERC721.attach(params.tokenContractAddr);

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0])) as ERC721SwapAgent__factory;
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  const receipt = await erc721SwapAgent.registerSwapPair(mockERC721.address, params.dstChainId);
  console.log(">> Registered!!")

  return receipt;
}

export const createSwapPairERC721 = async (params: CreateSwapPairERC721Params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.tokenName) throw new Error("tokenName is not defined");
  if (!params.tokenSymbol) throw new Error("tokenSymbol is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0])) as ERC721SwapAgent__factory;
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.createSwapPair(
    randomBytes(32),
    params.tokenContractAddr,
    params.fromChainId,
    params.tokenBaseUri,
    params.tokenName,
    params.tokenSymbol,
  );

  console.log(">> SwapPair created !!");
};

export const swapERC721 = async (params: SwapERC721Params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.tokenAddr) throw new Error("tokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.dstChainId) throw new Error("dstChainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0])) as ERC721SwapAgent__factory;
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.swap(
    params.tokenAddr,
    params.recipient,
    params.tokenId,
    params.dstChainId,
  );

  console.log(">> ERC721 swaped !!");
};

export const fillERC721 = async (params: FillERC721Params) => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.fromTokenAddr) throw new Error("fromTokenAddr is not defined");
  if (!params.recipient) throw new Error("recipient is not defined");
  if (!params.tokenId) throw new Error("tokenId is not defined");
  if (!params.fromChainId) throw new Error("fromChainId is not defined");
  if (!params.tokenURI) throw new Error("tokenURI is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0])) as ERC721SwapAgent__factory;
  const erc721SwapAgent = await ERC721SwapAgent.attach(params.agentAddr);

  await erc721SwapAgent.fill(
    randomBytes(32),
    params.fromTokenAddr,
    params.recipient,
    params.fromChainId,
    params.tokenId,
    params.tokenURI,
  );

  console.log(">> ERC721 swaped !!");
};
