import { ethers } from "hardhat";
import { MockERC1155__factory, MirroredERC1155__factory } from "../../typechain";
import { Awaited } from '../../types/utils';

type MintMockERC1155Params = {
  ids: number[],
  amounts: number[],
  to: string,
  tokenContractAddr: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

type ApproveMockERC1155Params = {
  to: string,
  tokenContractAddr: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

type ApproveMirroredERC1155Params = {
  to: string,
  tokenContractAddr: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
}

export const approveMockERC1155 = async (params: ApproveMockERC1155Params) => {
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0])) as MockERC1155__factory;
  const mockERC1155 = await MockERC1155.attach(params.tokenContractAddr);

  await mockERC1155.setApprovalForAll(params.to, true);
  console.log(`>> Approved MockERC1155 to ${params.to}`);
};

export const approveMirroredERC1155 = async (params: ApproveMirroredERC1155Params) => {
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MirroredERC1155 = (await ethers.getContractFactory("MirroredERC1155", params.signers[0])) as MirroredERC1155__factory;
  const mirroredERC1155 = await MirroredERC1155.attach(params.tokenContractAddr);

  await mirroredERC1155.setApprovalForAll(params.to, true);
  console.log(`>> Approved MirroredERC1155 to ${params.to}`);
};

export const mintMockERC1155 = async (params: MintMockERC1155Params) => {
  if (!params.ids.length) throw new Error("ids is not defined");
  if (!params.amounts.length) throw new Error("amounts is not defined");
  if (!params.to) throw new Error("to is not defined");
  if (!params.tokenContractAddr) throw new Error("tokenContractAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0])) as MockERC1155__factory;
  const mockERC1155 = await MockERC1155.attach(params.tokenContractAddr);

  await mockERC1155.mintBatch(params.to, params.ids, params.amounts, "0x00");
  console.log(`>> Minted MockERC1155 to ${params.to}`);
}
