package client

import (
	"context"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	corecommon "github.com/synycboom/bsc-evm-compatible-bridge-core/common"
)

type ETHClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
}

type Client struct {
	client *ethclient.Client
	mutex  sync.RWMutex
}

func NewClient(c *ethclient.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	h, err := c.client.HeaderByNumber(ctx, number)
	if err != nil && strings.Contains(err.Error(), "not found") {
		return nil, corecommon.ErrBlockNotFound
	}

	return h, err
}

func (c *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.client.PendingNonceAt(ctx, account)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.client.SuggestGasPrice(ctx)
}

func (c *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.client.TransactionByHash(ctx, hash)
}

func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.client.TransactionReceipt(ctx, txHash)
}

func (c *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.client.EstimateGas(ctx, msg)
}
