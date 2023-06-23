// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
)

type Deployment struct {
	Protocol    string `json:"protocol"`
	Version     int    `json:"version"`      // protocol version sequence on indexed chain
	StartHeight int64  `json:"start_height"` // first block on indexed chain
	EndHeight   int64  `json:"end_height"`   // last block on indexed chain or -1
}

func (c *explorerClient) ListProtocols(ctx context.Context) ([]Deployment, error) {
	protos := make([]Deployment, 0)
	if err := c.client.Get(ctx, "/explorer/protocols", nil, &protos); err != nil {
		return nil, err
	}
	return protos, nil
}
