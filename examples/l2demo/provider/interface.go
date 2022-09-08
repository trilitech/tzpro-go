// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package provider

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
	"context"
)

type Provider interface {
	Init(ctx context.Context, c *tzpro.Client, addr tezos.Address) error
	Enabled() bool
	Name() string
	Status() tzpro.BlockId
	ConnectBlock(ctx context.Context, block *tzpro.Block) error
	DisconnectBlock(ctx context.Context, block *tzpro.Block) error
	FillHistory(ctx context.Context) error
}
