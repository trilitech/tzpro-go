// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"context"
	"encoding/json"
	"time"
)

type ProfileClaim struct {
	Id            uint64          `json:"id"`
	Url           string          `json:"url"`
	Owner         string          `json:"owner"`
	Contract      string          `json:"contract"`
	Profile       uint64          `json:"profile_id"`
	Signature     HexBytes        `json:"signature"`
	Data          json.RawMessage `json:"data,omitempty"`
	Status        string          `json:"status"`
	Counter       int             `json:"counter"`
	RevokeCounter int             `json:"revoke_counter,omitempty"`
	IsRevoked     bool            `json:"is_revoked,omitempty"`
	Block         int64           `json:"block"`
	Time          time.Time       `json:"time"`
}

func (c *profileClient) ListClaims(ctx context.Context, params Query) ([]*ProfileClaim, error) {
	list := make([]*ProfileClaim, 0)
	u := params.WithPath("/v1/profiles/claims").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
