package wallet

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	conn *ethclient.Client
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect(url string) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("dail to %s err:%v\n", url, err)
		return
	}

	c.conn = client
}

func (c *Client) BalanceAt(address common.Address, blockNumber *big.Int) (*big.Int, error) {
	balance, err := c.conn.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		log.Println("Get balance err:", err)
		return nil, err
	}
	return balance, nil
}
