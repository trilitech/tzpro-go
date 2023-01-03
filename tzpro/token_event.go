// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"strconv"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type TokenEvent struct {
	Id        int64         `json:"id"`
	Contract  tezos.Address `json:"contract"`
	TokenId   tezos.Z       `json:"token_id"`
	TokenKind string        `json:"kind"`
	TokenType string        `json:"type"`
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

type TokenEventParams struct {
	Params
}

func NewTokenEventParams() TokenEventParams {
	return TokenEventParams{NewParams()}
}

func (p TokenEventParams) WithLimit(v uint) TokenEventParams {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p TokenEventParams) WithOffset(v uint) TokenEventParams {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p TokenEventParams) WithCursor(v uint64) TokenEventParams {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p TokenEventParams) WithOrder(o OrderType) TokenEventParams {
	p.Query.Set("order", string(o))
	return p
}

func (p TokenEventParams) WithDesc() TokenEventParams {
	p.Query.Set("order", string(OrderDesc))
	return p
}

func (p TokenEventParams) WithAsc() TokenEventParams {
	p.Query.Set("order", string(OrderAsc))
	return p
}

func (p TokenEventParams) WithToken(c tezos.Token) TokenEventParams {
	p.Query.Set("token", c.String())
	return p
}

func (p TokenEventParams) WithContract(c tezos.Address) TokenEventParams {
	p.Query.Set("contract", c.String())
	return p
}

func (p TokenEventParams) WithEventType(t string) TokenEventParams {
	p.Query.Set("event_type", t)
	return p
}

func (p TokenEventParams) WithTokenKind(k string) TokenEventParams {
	p.Query.Set("token_kind", k)
	return p
}

func (p TokenEventParams) WithSigner(c tezos.Address) TokenEventParams {
	p.Query.Set("signer", c.String())
	return p
}

func (p TokenEventParams) WithSender(c tezos.Address) TokenEventParams {
	p.Query.Set("sender", c.String())
	return p
}

func (p TokenEventParams) WithReceiver(c tezos.Address) TokenEventParams {
	p.Query.Set("receiver", c.String())
	return p
}

func (p TokenEventParams) WithTxHash(h tezos.OpHash) TokenEventParams {
	p.Query.Set("tx_hash", h.String())
	return p
}

func (c *Client) ListTokenEvents(ctx context.Context, params TokenParams) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.AppendQuery("/v1/tokens/events")
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
