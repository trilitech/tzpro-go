// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type NftTrade struct {
	Id       uint64 `json:"id"`
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Entity   string `json:"entity"`
	Category string `json:"category"`

	SaleId   int64  `json:"sale_id"`
	SaleType string `json:"sale_type"`

	Collection tezos.Address `json:"collection"`
	TokenId    tezos.Z       `json:"token_id"`
	NumUnits   int64         `json:"num_units"`
	Currency   *Token        `json:"currency"`
	Price      tezos.Z       `json:"price"`
	Fee        tezos.Z       `json:"fee"`
	Royalty    tezos.Z       `json:"royalty"`
	Seller     tezos.Address `json:"seller"`
	Buyer      tezos.Address `json:"buyer"`
	TxHash     tezos.OpHash  `json:"tx_hash"`
	TxFee      int64         `json:"tx_fee"`
	Block      int64         `json:"block"`
	Time       time.Time     `json:"time"`
}

type NftTradeParams = Params[NftTrade]

func NewNftTradeParams() NftTradeParams {
	return NftTradeParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListNftTrades(ctx context.Context, params NftTradeParams) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath("/v1/nft/trades").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListNftMarketTrades(ctx context.Context, addr tezos.Address, params NftTradeParams) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s/trades", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletNftTrades(ctx context.Context, addr tezos.Address, params NftTradeParams) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_trades", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
