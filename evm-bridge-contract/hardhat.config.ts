import * as dotenv from "dotenv";

import {HardhatUserConfig, task} from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@nomiclabs/hardhat-waffle";
import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "solidity-coverage";
import '@openzeppelin/hardhat-upgrades';
import 'hardhat-deploy';

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.2",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  networks: {
    chain1: {
      chainId: 1000,
      url: 'http://localhost:19545',
      accounts: [process.env.LOCAL_PRIVATE_KEY_1 || '']
    },
    chain2: {
      chainId: 2000,
      url: 'http://localhost:19546',
      accounts: [process.env.LOCAL_PRIVATE_KEY_1 || '']
    },
    rinkeby: {
      chainId: 4,
      url: process.env.RINKEBY_URL || "",
      accounts:
        process.env.RINKEBY_PRIVATE_KEY !== undefined ? [process.env.RINKEBY_PRIVATE_KEY] : [],
    },
    mumbai: {
      chainId: 80001,
      url: process.env.MATIC_MUMBAI_URL || "",
      accounts:
        process.env.MATIC_MUMBAI_PRIVATE_KEY !== undefined ? [process.env.MATIC_MUMBAI_PRIVATE_KEY] : [],
    },
    fuji: {
      chainId: 43113,
      url: process.env.AVAX_FUJI_URL || "",
      accounts:
        process.env.AVAX_FUJI_PRIVATE_KEY !== undefined ? [process.env.AVAX_FUJI_PRIVATE_KEY] : [],
    },
    fantomtest: {
      chainId: 4002,
      url: process.env.FANTOM_TEST_URL || "",
      accounts:
        process.env.FANTOM_TEST_PRIVATE_KEY !== undefined ? [process.env.FANTOM_TEST_PRIVATE_KEY] : [],
    },
    bsctest: {
      chainId: 97,
      url: process.env.BSC_TEST_URL || "",
      accounts: process.env.BSC_TEST_PRIVATE_KEY !== undefined ? [process.env.BSC_TEST_PRIVATE_KEY] : [],
    },
    hardhat: {
      chainId: 31337,
      gas: 12000000,
      blockGasLimit: 0x1fffffffffffff,
      allowUnlimitedContractSize: true,
      loggingEnabled: true,
      mining: {
        auto: true,
        interval: 10000
      },
      accounts: [{
        privateKey: process.env.LOCAL_PRIVATE_KEY_1 || '',
        balance: "10000000000000000000000",
      }, {
        privateKey: process.env.LOCAL_PRIVATE_KEY_2 || '',
        balance: "10000000000000000000000",
      }, {
        privateKey: process.env.LOCAL_PRIVATE_KEY_3 || '',
        balance: "10000000000000000000000",
      }],
    },
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
    // apiKey: process.env.BSCSCAN_API_KEY,
    // apiKey: process.env.POLYGONSCAN_API_KEY,
    // apiKey: process.env.FANTOMSCAN_API_KEY,
  },
};

export default config;
