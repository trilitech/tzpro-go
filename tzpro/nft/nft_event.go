// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"context"
	"fmt"
	"time"
)

type NftEvent struct {
	Id         uint64    `json:"id"`
	Contract   Address   `json:"contract"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Category   string    `json:"category"`
	EventType  string    `json:"event_type"`
	SaleId     int64     `json:"sale_id"`
	SaleType   string    `json:"sale_type"`
	Collection Address   `json:"collection"`
	TokenId    Z         `json:"token_id,omitempty"`
	Currency   *Token    `json:"currency,omitempty"`
	NumUnits   int64     `json:"num_units,omitempty"`
	Amount     Z         `json:"amount,omitempty"`
	Fee        Z         `json:"fee,omitempty"`
	Royalty    Z         `json:"royalty,omitempty"`
	RoyaltyBps int64     `json:"royalty_bps,omitempty"`
	Signer     Address   `json:"signer"`
	Sender     Address   `json:"sender"`
	TxHash     OpHash    `json:"tx_hash"`
	TxFee      int64     `json:"tx_fee,string"`
	Block      int64     `json:"block"`
	Time       time.Time `json:"time"`
}

func (c *nftClient) ListEvents(ctx context.Context, params Query) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath("/v1/nft/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *nftClient) ListMarketEvents(ctx context.Context, addr Address, params Query) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
