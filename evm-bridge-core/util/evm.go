package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	contractabi "github.com/synycboom/bsc-evm-compatible-bridge-core/abi"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/client"
)

var (
	err            error
	erc721AgentABI abi.ABI
)

func init() {
	erc721AgentABI, err = abi.JSON(strings.NewReader(contractabi.ERC721SwapAgentMetaData.ABI))
	if err != nil {
		panic(errors.Wrap(err, "[init]: failed to parse ERC721 swap agent abi"))
	}
}

func StrToBigInt(val string) *big.Int {
	integer, err := strconv.Atoi(val)
	if err != nil {
		return big.NewInt(0)
	}

	return big.NewInt(int64(integer))
}

func BigIntSliceToStrSlice(vv []*big.Int) []string {
	ss := make([]string, len(vv))
	for idx, v := range vv {
		ss[idx] = v.String()
	}

	return ss
}

func StrSliceToBigIntSlice(ss []string) []*big.Int {
	vv := make([]*big.Int, len(ss))
	for idx, s := range ss {
		vv[idx] = StrToBigInt(s)
	}

	return vv
}

func TxOpts(ctx context.Context, ethClient client.ETHClient, privateKey string, chainID *big.Int) (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "[TxOpts]: unable to load a private key")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "[TxOpts]: unable to create a new auth")
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, errors.Wrap(err, "[TxOpts]: failed to get nonce")
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "[TxOpts]: failed to gas price")
	}

	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Context = ctx
	txOpts.GasPrice = gasPrice
	txOpts.GasLimit = 0

	return txOpts, nil
}

// func CallerOpts(ctx context.Context, ethClient client.ETHClient) (*bind.CallOpts, error) {

// }

func BuildKeys(privateKeyStr string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	priKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyStr, "0x"))
	if err != nil {
		return nil, nil, err
	}
	publicKey, ok := priKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("get public key error")
	}
	return priKey, publicKey, nil
}

func PackERC721AgentInput(method string, params ...interface{}) ([]byte, error) {
	input, err := erc721AgentABI.Pack(method, params...)
	if err != nil {
		return nil, errors.Wrap(err, "[PackERC721AgentInput]: failed to pack input")
	}

	return input, nil
}

func BuildSignedTransaction(
	contract common.Address,
	ethClient *ethclient.Client,
	chainID *big.Int,
	txInput []byte,
	privateKey *ecdsa.PrivateKey,
) (*types.Transaction, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "[BuildSignedTransaction]: failed to create new transactor")
	}

	nonce, err := ethClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, errors.Wrap(err, "[BuildSignedTransaction]: failed to get nonce")
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "[BuildSignedTransaction]: failed to gas price")
	}

	value := big.NewInt(0)
	gasLimit := uint64(2500000)
	// msg := ethereum.CallMsg{From: txOpts.From, To: &contract, GasPrice: gasPrice, Value: value, Data: txInput}
	// gasLimit, err := ethClient.EstimateGas(context.Background(), msg)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "[BuildSignedTransaction]: failed to estimate gas needed")
	// }

	rawTx := types.NewTx(&types.AccessListTx{
		ChainID:  chainID,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &contract,
		Value:    value,
		Data:     txInput,
	})
	signedTx, err := txOpts.Signer(txOpts.From, rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "[BuildSignedTransaction]: failed to sign tx")
	}

	return signedTx, nil
}
