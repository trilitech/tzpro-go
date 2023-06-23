// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"fmt"
	"time"
)

type TokenEvent struct {
	Id        int64     `json:"id"`
	Ledger    Address   `json:"contract"`
	TokenId   Z         `json:"token_id"`
	TokenKind string    `json:"token_kind"`
	TokenType string    `json:"token_type"`
	EventType string    `json:"event_type"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Decimals  int       `json:"decimals"`
	Signer    Address   `json:"signer"`
	Sender    Address   `json:"sender"`
	Receiver  Address   `json:"receiver"`
	Amount    Z         `json:"amount"`
	TxHash    OpHash    `json:"tx_hash"`
	TxFee     int64     `json:"tx_fee,string"`
	Block     int64     `json:"block"`
	Time      time.Time `json:"time"`
}

func (c *tokenClient) ListLedgerEvents(ctx context.Context, addr Address, params Params) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/ledgers/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *tokenClient) ListTokenEvents(ctx context.Context, addr TokenAddress, params Params) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/tokens/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
