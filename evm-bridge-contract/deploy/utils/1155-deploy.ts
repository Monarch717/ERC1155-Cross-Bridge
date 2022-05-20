import { ethers, upgrades } from "hardhat";
import { MockERC1155__factory, ERC1155SwapAgent__factory, MockERC1155, ERC1155SwapAgent } from "../../typechain";
import { Awaited } from '../../types/utils';

type DeployMockERC1155Params = {
  agentAddr: string,
  uri: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

type DeployERC1155AgentParams = {
  chainId: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

export const deployMockERC1155= async (params: DeployMockERC1155Params): Promise<MockERC1155> => {
  if (!params.agentAddr) throw new Error("agentAddr is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");
  if (!params.uri) throw new Error("uri is not defined");

  const MockERC1155 = (await ethers.getContractFactory("MockERC1155", params.signers[0])) as MockERC1155__factory

  console.log(`>> Deploying MockERC1155`);

  const mockERC1155 = await MockERC1155.deploy(params.uri);

  console.log(">> MockERC1155 is deployed!");

  return mockERC1155;
} 

export const deployERC1155Agent = async (params: DeployERC1155AgentParams): Promise<ERC1155SwapAgent> => {
  if (!params.chainId) throw new Error("chainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC1155SwapAgent = (await ethers.getContractFactory("ERC1155SwapAgent", params.signers[0])) as ERC1155SwapAgent__factory;

  console.log(`>> Deploying ERC1155SwapAgent`);

  const agent = (await upgrades.deployProxy(ERC1155SwapAgent, [], { initializer: 'initialize' })) as ERC1155SwapAgent;

  await agent.deployed();

  console.log(">> ERC1155SwapAgent is deployed!");

  console.log(agent.address);

  return agent;
} 


