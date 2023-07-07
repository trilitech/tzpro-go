// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"context"
	"fmt"
	"time"
)

type NftTrade struct {
	Id         uint64    `json:"id"`
	Contract   Address   `json:"contract"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Category   string    `json:"category"`
	SaleId     int64     `json:"sale_id"`
	SaleType   string    `json:"sale_type"`
	Collection Address   `json:"collection"`
	TokenId    Z         `json:"token_id"`
	NumUnits   int64     `json:"num_units"`
	Currency   *Token    `json:"currency"`
	Price      Z         `json:"price"`
	Fee        Z         `json:"fee"`
	Royalty    Z         `json:"royalty"`
	Seller     Address   `json:"seller"`
	Buyer      Address   `json:"buyer"`
	TxHash     OpHash    `json:"tx_hash"`
	TxFee      int64     `json:"tx_fee,string"`
	Block      int64     `json:"block"`
	Time       time.Time `json:"time"`
}

func (c *nftClient) ListTrades(ctx context.Context, params Query) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath("/v1/nft/trades").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *nftClient) ListMarketTrades(ctx context.Context, addr Address, params Query) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s/trades", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
