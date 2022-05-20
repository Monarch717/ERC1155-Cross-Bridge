import { ethers, upgrades } from "hardhat";
import { MockERC721__factory, ERC721SwapAgent__factory, MockERC721, ERC721SwapAgent, MockERC721NoBaseURI, MockERC721NoBaseURI__factory } from "../../typechain";
import { Awaited } from '../../types/utils';

type DeployMockERC721Params = {
  baseUri: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

type DeployERC721AgentParams = {
  chainId: string,
  signers: Awaited<ReturnType<typeof ethers.getSigners>>,
};

export const deployMockERC721= async (params: DeployMockERC721Params): Promise<MockERC721 | MockERC721NoBaseURI> => {
  if (!params.signers.length) throw new Error("signers is not defined");

  if (params.baseUri) {
    const MockERC721 = (await ethers.getContractFactory("MockERC721", params.signers[0])) as MockERC721__factory

    console.log(`>> Deploying MockERC721`);

    const mockERC721 = await MockERC721.deploy('Mock721', 'M721');
    console.log(">> MockERC721 is deployed!");

    await mockERC721.setBaseURI(params.baseUri);
    console.log(`>> Set MockERC721 baseURI to ${params.baseUri}`);

    return mockERC721;
  }


  const MockERC721NoBaseURI = (await ethers.getContractFactory("MockERC721NoBaseURI", params.signers[0])) as MockERC721NoBaseURI__factory

  console.log(`>> Deploying MockERC721NoBaseURI`);

  const mockERC721NoBaseURI = await MockERC721NoBaseURI.deploy('MockERC721NoBaseURI', 'M721N');

  console.log(">> MockERC721NoBaseURI is deployed!");

  return mockERC721NoBaseURI;
} 

export const deployERC721Agent = async (params: DeployERC721AgentParams): Promise<ERC721SwapAgent> => {
  if (!params.chainId) throw new Error("chainId is not defined");
  if (!params.signers.length) throw new Error("signers is not defined");

  const ERC721SwapAgent = (await ethers.getContractFactory("ERC721SwapAgent", params.signers[0])) as ERC721SwapAgent__factory;

  console.log(`>> Deploying ERC721SwapAgent`);

  const agent = (await upgrades.deployProxy(ERC721SwapAgent, [], { initializer: 'initialize' })) as ERC721SwapAgent;

  await agent.deployed();

  console.log(">> ERC721SwapAgent is deployed!");

  console.log(agent.address);

  return agent;
} 


