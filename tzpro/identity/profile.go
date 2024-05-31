// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"context"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type ProfileAPI interface {
	// firehose
	ListProfiles(context.Context, Query) ([]*Profile, error)
	ListEvents(context.Context, Query) ([]*ProfileEvent, error)
	ListClaims(context.Context, Query) ([]*ProfileClaim, error)
}

func NewProfileAPI(c *client.Client) ProfileAPI {
	return &profileClient{client: c}
}

var (
	ProfileRegistry         = MainnetProfileRegistry
	MainnetProfileRegistry  = MustParseAddress("KT1GBZmSxmnKJXGMdMLbugPfLyUPmuLSMwKS")
	GhostnetProfileRegistry = MustParseAddress("KT1REqKBXwULnmU6RpZxnRBUgcBmESnXhCWs")
)

type profileClient struct {
	client *client.Client
}

type Profile struct {
	Id             uint64    `json:"id"`
	Owner          string    `json:"owner"`
	Contract       Address   `json:"contract"`
	Alias          string    `json:"alias"`
	Description    string    `json:"description"`
	Logo           string    `json:"logo"`
	Website        string    `json:"website"`
	Twitter        string    `json:"twitter"`
	Ethereum       string    `json:"ethereum"`
	DomainName     string    `json:"domain_name"`
	Discord        string    `json:"discord"`
	Github         string    `json:"github"`
	FirstBlock     int64     `json:"first_block"`
	FirstTime      time.Time `json:"first_time"`
	LastBlock      int64     `json:"last_block"`
	LastTime       time.Time `json:"last_time"`
	UpdateCounter  int       `json:"update_counter"`
	AppliedCounter int       `json:"applied_counter"`
}

func (c *profileClient) ListProfiles(ctx context.Context, params Query) ([]*Profile, error) {
	list := make([]*Profile, 0)
	u := params.WithPath("/v1/profiles").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
