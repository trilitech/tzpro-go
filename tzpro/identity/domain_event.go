// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"context"
	"encoding/json"
	"time"
)

type DomainEvent struct {
	Id             uint64          `json:"id"`
	Type           string          `json:"event_type"`
	Domain         string          `json:"domain"`
	TokenId        int64           `json:"token_id"`
	Level          int             `json:"level"`
	Owner          string          `json:"owner"`
	ForwardAddress string          `json:"forward_address"`
	ReverseAddress string          `json:"reverse_address"`
	Expiry         time.Time       `json:"expiry"`
	Metadata       json.RawMessage `json:"metadata"`
	Signer         string          `json:"signer"`
	Sender         string          `json:"sender"`
	TxHash         OpHash          `json:"tx_hash"`
	TxFee          int64           `json:"tx_fee,string"`
	Block          int64           `json:"block"`
	Time           time.Time       `json:"time"`
}

func (c *domainClient) ListEvents(ctx context.Context, params Query) ([]*DomainEvent, error) {
	list := make([]*DomainEvent, 0)
	u := params.WithPath("/v1/domains/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
