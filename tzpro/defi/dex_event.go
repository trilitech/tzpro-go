// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type DexEvent struct {
	Id         uint64    `json:"id"`
	Contract   Address   `json:"contract"`
	PairId     int       `json:"pair_id"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Pair       string    `json:"pair"`
	Type       string    `json:"event_type"`
	VolumeA    Z         `json:"volume_a"`
	VolumeB    Z         `json:"volume_b"`
	VolumeLP   Z         `json:"volume_lp"`
	DecimalsA  int       `json:"decimals_a"`
	DecimalsB  int       `json:"decimals_b"`
	DecimalsLP int       `json:"decimals_lp"`
	SupplyA    Z         `json:"supply_a"`
	SupplyB    Z         `json:"supply_b"`
	SupplyLP   Z         `json:"supply_lp"`
	ValueUSD   float64   `json:"value_usd,string"`
	Signer     Address   `json:"signer"`
	Sender     Address   `json:"sender"`
	Receiver   Address   `json:"receiver"`
	Router     Address   `json:"router"`
	TxHash     OpHash    `json:"tx_hash"`
	TxFee      int64     `json:"tx_fee,string"`
	Block      int64     `json:"block"`
	Time       time.Time `json:"time"`
}

func (c *dexClient) ListEvents(ctx context.Context, params Query) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath("/v1/dex/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *dexClient) ListPoolEvents(ctx context.Context, addr PoolAddress, params Query) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/dex/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
