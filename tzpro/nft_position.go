// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type NftPosition struct {
	Id         uint64     `json:"id"`
	Contract   string     `json:"contract"`
	Name       string     `json:"name"`
	Entity     string     `json:"entity"`
	Category   string     `json:"category"`
	SaleId     int64      `json:"sale_id"`
	SaleType   string     `json:"sale_type"`
	SaleStatus string     `json:"sale_status"`
	IsClosed   bool       `json:"is_closed"`
	Seller     string     `json:"seller"`
	Buyer      string     `json:"buyer"`
	AskPrice   tezos.Z    `json:"ask_price"`
	BidPrice   tezos.Z    `json:"bid_price"`
	MaxUnits   int64      `json:"max_units"`
	SoldUnits  int64      `json:"sold_units"`
	RoyaltyBps int64      `json:"royalty_bps"`
	OpenBlock  int64      `json:"open_block"`
	OpenTime   time.Time  `json:"open_time"`
	CloseBlock int64      `json:"close_block,omitempty"`
	CloseTime  *time.Time `json:"close_time,omitempty"`
}

type NftPositionParams = Params[NftPosition]

func NewNftPositionParams() NftPositionParams {
	return NftPositionParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListNftPositions(ctx context.Context, params NftPositionParams) ([]*NftPosition, error) {
	list := make([]*NftPosition, 0)
	u := params.WithPath("/v1/nft/positions").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListNftMarketPositions(ctx context.Context, addr tezos.Address, params NftPositionParams) ([]*NftPosition, error) {
	list := make([]*NftPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s/positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletNftPositions(ctx context.Context, addr tezos.Address, params NftPositionParams) ([]*NftPosition, error) {
	list := make([]*NftPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
