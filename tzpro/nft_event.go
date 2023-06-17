// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type NftEvent struct {
	Id         uint64    `json:"id"`
	Contract   string    `json:"contract"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Category   string    `json:"category"`
	EventType  string    `json:"event_type"`
	SaleId     int64     `json:"sale_id"`
	SaleType   string    `json:"sale_type"`
	Collection string    `json:"collection"`
	TokenId    tezos.Z   `json:"token_id,omitempty"`
	Currency   *Token    `json:"currency,omitempty"`
	NumUnits   int64     `json:"num_units,omitempty"`
	Amount     tezos.Z   `json:"amount,omitempty"`
	Fee        tezos.Z   `json:"fee,omitempty"`
	Royalty    tezos.Z   `json:"royalty,omitempty"`
	RoyaltyBps int64     `json:"royalty_bps,omitempty"`
	Signer     string    `json:"signer"`
	Sender     string    `json:"sender"`
	TxHash     string    `json:"tx_hash"`
	TxFee      int64     `json:"tx_fee,string"`
	Block      int64     `json:"block"`
	Time       time.Time `json:"time"`
}

type NftEventParams = Params[NftEvent]

func NewNftEventParams() NftEventParams {
	return NftEventParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListNftEvents(ctx context.Context, params NftEventParams) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath("/v1/nft/events").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListNftMarketEvents(ctx context.Context, addr tezos.Address, params NftEventParams) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletNftEvents(ctx context.Context, addr tezos.Address, params NftEventParams) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
