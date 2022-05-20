package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	contractabi "github.com/synycboom/bsc-evm-compatible-bridge-core/abi"
)

type IToken interface {
	URI(opts *bind.CallOpts, tokenAddr string) (string, error)
}

type Token struct {
	client *ethclient.Client
}

func NewToken(c *ethclient.Client) *Token {
	return &Token{
		client: c,
	}
}

func (t *Token) bind(addr string) (*contractabi.ERC1155Token, error) {
	tokenAddr := common.HexToAddress(addr)
	token, err := contractabi.NewERC1155Token(tokenAddr, t.client)
	if err != nil {
		return nil, errors.Wrap(err, "[Token.bind]: failed to bind token address")
	}

	return token, nil
}

func (t *Token) URI(opts *bind.CallOpts, tokenAddr string) (string, error) {
	token, err := t.bind(tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Token.URI]: failed to bind token")
	}

	uri, err := token.Uri(opts, big.NewInt(0))
	if err != nil {
		return "", errors.Wrap(err, "[Token.URI]: failed to get token URI")
	}

	return uri, nil
}
