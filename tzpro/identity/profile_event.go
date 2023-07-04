// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"context"
	"encoding/json"
	"time"
)

type ProfileEvent struct {
	Id           uint64          `json:"id"`
	Type         string          `json:"event_type"`
	Owner        string          `json:"owner"`
	Contract     string          `json:"contract"`
	Claims       json.RawMessage `json:"claims"`
	IsRevocation bool            `json:"is_revocation"`
	Signer       string          `json:"signer"`
	Sender       string          `json:"sender"`
	TxHash       OpHash          `json:"tx_hash"`
	TxFee        int64           `json:"tx_fee,string"`
	Block        int64           `json:"block"`
	Time         time.Time       `json:"time"`
}

func (c *profileClient) ListEvents(ctx context.Context, params Query) ([]*ProfileEvent, error) {
	list := make([]*ProfileEvent, 0)
	u := params.WithPath("/v1/profiles/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
