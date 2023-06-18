// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
)

type Status struct {
	Status    string  `json:"status"` // loading, connecting, stopping, stopped, waiting, syncing, synced, failed
	Blocks    int64   `json:"blocks"`
	Finalized int64   `json:"finalized"`
	Indexed   int64   `json:"indexed"`
	Progress  float64 `json:"progress"`
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if len(data) == 2 {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("Status: expected JSON array")
	}
	return Decode(data, nil, s)
}

func (c *Client) GetStatus(ctx context.Context) (*Status, error) {
	s := &Status{}
	if err := c.get(ctx, "/explorer/status", nil, s); err != nil {
		return nil, err
	}
	return s, nil
}
