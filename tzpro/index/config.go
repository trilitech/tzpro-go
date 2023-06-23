// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"strconv"
)

type Config struct {
	Name                   string  `json:"name"`
	Network                string  `json:"network"`
	Symbol                 string  `json:"symbol"`
	ChainId                string  `json:"chain_id"`
	Deployment             int     `json:"deployment"`
	Version                int     `json:"version"`
	Protocol               string  `json:"protocol"`
	StartHeight            int64   `json:"start_height"`
	EndHeight              int64   `json:"end_height"`
	Decimals               int     `json:"decimals"`
	MinimalStake           float64 `json:"minimal_stake"`
	PreservedCycles        int64   `json:"preserved_cycles"`
	MinimalBlockDelay      int     `json:"minimal_block_delay"`
	DelayIncrementPerRound int     `json:"delay_increment_per_round"`
}

func (c *explorerClient) GetConfigHead(ctx context.Context) (*Config, error) {
	config := &Config{}
	if err := c.client.Get(ctx, "/explorer/config/head", nil, config); err != nil {
		return nil, err
	}
	return config, nil
}

func (c *explorerClient) GetConfigHeight(ctx context.Context, height int64) (*Config, error) {
	config := &Config{}
	if err := c.client.Get(ctx, "/explorer/config/"+strconv.FormatInt(height, 10), nil, config); err != nil {
		return nil, err
	}
	return config, nil
}
