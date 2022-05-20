import path from 'path';
import fs from 'fs';

type ERC721AgentCache = {
  address: string,
  registerTxHash: string,
  registerBlockNumber: string,
};

type ERC721MockTokenCache = {
  address: string,
  baseUri: string,
  name: string,
  symbol: string,
  tokenId: string,
};

type ERC721MirroredTokenCache = {
  address: string,
  tokenId: string,
};

const erc721AgentFile = 'erc721-agent.json';
const erc721MockTokenFile = 'erc721-mock-token.json'
const erc721MirroredTokenFile = 'erc721-mirrored-token.json'

const readFileByChainId = <T>(chainId: string, fileName: string): T => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  return JSON.parse(fs.readFileSync(filePath).toString());
};

const writeFileByChainId = <T>(chainId: string, fileName: string, data: T) => {
  const filePath = path.join(__dirname, `../../chains/${chainId}/${fileName}`);
  fs.writeFileSync(filePath, JSON.stringify(data));
};

export const set721Agent = (chainId: string, data: ERC721AgentCache) => {
  writeFileByChainId<ERC721AgentCache>(chainId, erc721AgentFile, data);
};

export const get721Agent = (chainId: string): ERC721AgentCache => {
  return readFileByChainId<ERC721AgentCache>(chainId, erc721AgentFile);
};

export const set721MockToken = (chainId: string, data: ERC721MockTokenCache) => {
  writeFileByChainId<ERC721MockTokenCache>(chainId, erc721MockTokenFile, data);
};

export const get721MockToken = (chainId: string): ERC721MockTokenCache => {
  return readFileByChainId<ERC721MockTokenCache>(chainId, erc721MockTokenFile);
};

export const get721MirroredToken = (chainId: string): ERC721MirroredTokenCache => {
  return readFileByChainId<ERC721MirroredTokenCache>(chainId, erc721MirroredTokenFile);
};
