package token

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	contractabi "github.com/synycboom/bsc-evm-compatible-bridge-core/abi"
)

type IToken interface {
	TokenURI(opts *bind.CallOpts, tokenAddr string, tokenId *big.Int) (string, error)
	BaseURI(opts *bind.CallOpts, tokenAddr string) (string, error)
}

type Token struct {
	client *ethclient.Client
	mutex  sync.RWMutex
}

func NewToken(c *ethclient.Client) *Token {
	return &Token{
		client: c,
	}
}

func (t *Token) bind(addr string) (*contractabi.ERC721Token, error) {
	tokenAddr := common.HexToAddress(addr)
	token, err := contractabi.NewERC721Token(tokenAddr, t.client)
	if err != nil {
		return nil, errors.Wrap(err, "[Token.bind]: failed to bind token address")
	}

	return token, nil
}

func (t *Token) TokenURI(opts *bind.CallOpts, tokenAddr string, tokenId *big.Int) (string, error) {
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.TokenURI]: failed to bind token")
	}

	uri, err := token.TokenURI(opts, tokenId)
	if err != nil {
		return "", errors.Wrap(err, "[Token.TokenURI]: failed to get token URI")
	}

	return uri, nil
}

func (t *Token) BaseURI(opts *bind.CallOpts, tokenAddr string) (string, error) {
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.BaseURI]: failed to bind token")
	}

	uri, err := token.BaseURI(opts)
	if err != nil {
		return "", errors.Wrap(err, "[Token.BaseURI]: failed to get token base URI")
	}

	return uri, nil
}
