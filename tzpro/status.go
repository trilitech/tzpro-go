// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
)

type Status struct {
	Status    string  `json:"status"` // loading, connecting, stopping, stopped, waiting, syncing, synced, failed
	Blocks    int64   `json:"blocks"`
	Finalized int64   `json:"finalized"`
	Indexed   int64   `json:"indexed"`
	Progress  float64 `json:"progress"`

	columns []string
}

func (s *Status) WithColumns(cols ...string) *Status {
	s.columns = cols
	return s
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if len(data) == 2 {
		return nil
	}
	if data[0] == '[' {
		return s.UnmarshalJSONBrief(data)
	}
	type Alias *Status
	return json.Unmarshal(data, Alias(s))
}

func (s *Status) UnmarshalJSONBrief(data []byte) error {
	st := Status{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	unpacked := make([]interface{}, 0)
	err := dec.Decode(&unpacked)
	if err != nil {
		return err
	}
	for i, v := range s.columns {
		f := unpacked[i]
		if f == nil {
			continue
		}
		switch v {
		case "status":
			st.Status = f.(string)
		case "blocks":
			st.Blocks, err = strconv.ParseInt(f.(json.Number).String(), 10, 64)
		case "indexed":
			st.Indexed, err = strconv.ParseInt(f.(json.Number).String(), 10, 64)
		case "progress":
			st.Progress, err = f.(json.Number).Float64()
		}
		if err != nil {
			return err
		}
	}
	*s = st
	return nil
}

func (c *Client) GetStatus(ctx context.Context) (*Status, error) {
	s := &Status{}
	if err := c.get(ctx, "/explorer/status", nil, s); err != nil {
		return nil, err
	}
	return s, nil
}
