// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"

	"github.com/trilitech/tzpro-go/internal/client"
)

type ExplorerAPI interface {
	GetStatus(context.Context) (*Status, error)
	GetTip(context.Context) (*Tip, error)
	GetConfigHead(context.Context) (*Config, error)
	GetConfigHeight(context.Context, int64) (*Config, error)
	ListProtocols(context.Context) ([]Deployment, error)
	GetElection(context.Context, int) (*Election, error)
	ListVoters(context.Context, int, int) ([]Voter, error)
	ListBallots(context.Context, int, int) (BallotList, error)

	NewChainQuery() *ChainQuery
}

func NewExplorerAPI(c *client.Client) ExplorerAPI {
	return &explorerClient{client: c}
}

type explorerClient struct {
	client *client.Client
}
