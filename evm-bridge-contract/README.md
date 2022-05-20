# avax-evm-compatible-bridge-contr
This project is part of Binance Smart Chain Hackathon : Build NFT Bridge Between EVM Compatible Chains hackathon .

### Prepare configuration
Please refer to example env file
```shell
$ cat .env.example > .env
```

### Install Dependencies
```shell
$ nvm use
$ npm ci
```

### Deploy local chains to test the contracts locally
Need to have 3 terminals to run the scripts simultaneously
```shell
$ npm run chain1
```
```shell
$ npm run chain2
```
After 2 chains are running, deploy ERC721SwapAgent and ERC1155SwapAgent contracts
```shell
$ npm run chain1:erc721-deploy-agent
$ npm run chain1:erc1155-deploy-agent
$ npm run chain2:erc721-deploy-agent
$ npm run chain2:erc1155-deploy-agent
```
There are other useful commands to deploy contracts, see them in the package.json

### Deploy to Rinkeby and BSC Test
```shell
$ npm run bsctest:erc721-deploy-agent
$ npm run bsctest:erc1155-deploy-agent
$ npm run rinkeby:erc721-deploy-agent
$ npm run rinkeby:erc1155-deploy-agent
```

The addresses of the deployed contracts will be in the chains folder which has chain ids as subfolders.
