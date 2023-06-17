// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type TokenEvent struct {
	Id        int64         `json:"id"`
	Ledger    tezos.Address `json:"contract"`
	TokenId   tezos.Z       `json:"token_id"`
	TokenKind string        `json:"token_kind"`
	TokenType string        `json:"token_type"`
	EventType string        `json:"event_type"`
	Name      string        `json:"name"`
	Symbol    string        `json:"symbol"`
	Decimals  int           `json:"decimals"`
	Signer    tezos.Address `json:"signer"`
	Sender    tezos.Address `json:"sender"`
	Receiver  tezos.Address `json:"receiver"`
	Amount    tezos.Z       `json:"amount"`
	TxHash    tezos.OpHash  `json:"tx_hash"`
	TxFee     int64         `json:"tx_fee"`
	Block     int64         `json:"block"`
	Time      time.Time     `json:"time"`
}

type TokenEventParams = Params[TokenEvent]

func NewTokenEventParams() TokenEventParams {
	return TokenEventParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListTokenEvents(ctx context.Context, params TokenEventParams) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath("/v1/tokens/events").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListLedgerEvents(ctx context.Context, addr tezos.Address, params TokenEventParams) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/ledgers/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListTokenIdEvents(ctx context.Context, addr tezos.Token, params TokenEventParams) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/tokens/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
