// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"encoding/json"
	"fmt"
)

type TokenMetadata struct {
	Id          int64           `json:"id"`
	Ledger      Address         `json:"contract"`
	TokenId     Z               `json:"token_id"`
	Kind        string          `json:"kind"`
	Status      string          `json:"status"`
	Url         string          `json:"name"`
	RetriesLeft int             `json:"retries_left"`
	Block       int64           `json:"block"`
	Data        json.RawMessage `json:"data"`
}

func (c *tokenClient) ListMetadata(ctx context.Context, params Params) ([]*TokenMetadata, error) {
	list := make([]*TokenMetadata, 0)
	u := params.WithPath("/v1/meta").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *tokenClient) GetLedgerMetadata(ctx context.Context, addr Address) (*TokenMetadata, error) {
	val := &TokenMetadata{}
	u := fmt.Sprintf("/v1/meta/%s", addr)
	if err := c.client.Get(ctx, u, nil, val); err != nil {
		return nil, err
	}
	return val, nil
}

func (c *tokenClient) GetTokenMetadata(ctx context.Context, addr TokenAddress) (*TokenMetadata, error) {
	val := &TokenMetadata{}
	u := fmt.Sprintf("/v1/meta/%s", addr)
	if err := c.client.Get(ctx, u, nil, val); err != nil {
		return nil, err
	}
	return val, nil
}
