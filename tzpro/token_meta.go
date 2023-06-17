// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"encoding/json"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
)

type TokenMetadata struct {
	Id          int64           `json:"id"`
	Ledger      tezos.Address   `json:"contract"`
	TokenId     tezos.Z         `json:"token_id"`
	Kind        string          `json:"kind"`
	Status      string          `json:"status"`
	Url         string          `json:"name"`
	RetriesLeft int             `json:"retries_left"`
	Block       int64           `json:"block"`
	Data        json.RawMessage `json:"data"`
}

type TokenMetadataParams = Params[TokenMetadata]

func NewTokenMetadataParams() TokenMetadataParams {
	return TokenMetadataParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListTokenMetadata(ctx context.Context, params TokenMetadataParams) ([]*TokenMetadata, error) {
	list := make([]*TokenMetadata, 0)
	u := params.WithPath("/v1/meta").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetLedgerMetadata(ctx context.Context, addr tezos.Address, params TokenMetadataParams) ([]*TokenMetadata, error) {
	list := make([]*TokenMetadata, 0)
	u := params.WithPath(fmt.Sprintf("/v1/meta/%s", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) GetTokenMetadata(ctx context.Context, addr tezos.Token, params TokenMetadataParams) (*TokenMetadata, error) {
	val := &TokenMetadata{}
	u := params.WithPath(fmt.Sprintf("/v1/meta/%s", addr)).Url()
	if err := c.get(ctx, u, nil, &val); err != nil {
		return nil, err
	}
	return val, nil
}
