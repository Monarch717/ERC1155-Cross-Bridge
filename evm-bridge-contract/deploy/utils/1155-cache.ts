import path from 'path';
import fs from 'fs';

type ERC1155AgentCache = {
  address: string,
  registerTxHash: string,
  registerBlockNumber: string,
};

type ERC1155MockTokenCache = {
  address: string,
  uri: string,
  ids: number[],
  amounts: number[],
};

type ERC1155MirroredTokenCache = {
  address: string,
  ids: number[],
  amounts: number[],
};

const erc1155AgentFile = 'erc1155-agent.json';
const erc1155MockTokenFile = 'erc1155-mock-token.json'
const erc1155MirroredTokenFile = 'erc1155-mirrored-token.json'

const readFileByChainId = <T>(chainId: string, fileName: string): T => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  return JSON.parse(fs.readFileSync(filePath).toString());
};

const writeFileByChainId = <T>(chainId: string, fileName: string, data: T) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  fs.writeFileSync(filePath, JSON.stringify(data));
};

export const set1155Agent = (chainId: string, data: ERC1155AgentCache) => {
  writeFileByChainId<ERC1155AgentCache>(chainId, erc1155AgentFile, data);
};

export const get1155Agent = (chainId: string): ERC1155AgentCache => {
  return readFileByChainId<ERC1155AgentCache>(chainId, erc1155AgentFile);
};

export const set1155MockToken = (chainId: string, data: ERC1155MockTokenCache) => {
  writeFileByChainId<ERC1155MockTokenCache>(chainId, erc1155MockTokenFile, data);
};

export const get1155MockToken = (chainId: string): ERC1155MockTokenCache => {
  return readFileByChainId<ERC1155MockTokenCache>(chainId, erc1155MockTokenFile);
};

export const get1155MirroredToken = (chainId: string): ERC1155MirroredTokenCache => {
  return readFileByChainId<ERC1155MirroredTokenCache>(chainId, erc1155MirroredTokenFile);
};
