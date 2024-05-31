// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type DomainAPI interface {
	LookupByName(context.Context, string) (*Domain, error)
	LookupByAddress(context.Context, Address) (*Domain, error)

	// firehose
	ListDomains(context.Context, Query) ([]*Domain, error)
	ListEvents(context.Context, Query) ([]*DomainEvent, error)
}

func NewDomainAPI(c *client.Client) DomainAPI {
	return &domainClient{client: c}
}

var (
	DomainRegistry         = MainnetDomainRegistry
	MainnetDomainRegistry  = MustParseAddress("KT1GBZmSxmnKJXGMdMLbugPfLyUPmuLSMwKS")
	GhostnetDomainRegistry = MustParseAddress("KT1REqKBXwULnmU6RpZxnRBUgcBmESnXhCWs")
)

type domainClient struct {
	client *client.Client
}

type Domain struct {
	Id             uint64          `json:"id"`
	Domain         string          `json:"domain"`
	TokenId        int64           `json:"token_id,omitempty"`
	Level          int             `json:"level"`
	Owner          string          `json:"owner"`
	ForwardAddress string          `json:"forward_address"`
	ReverseAddress string          `json:"reverse_address"`
	Expiry         time.Time       `json:"expiry"`
	Metadata       json.RawMessage `json:"metadata,omitempty"`
	FirstBlock     int64           `json:"first_block"`
	FirstTime      time.Time       `json:"first_time"`
}

func (p Domain) Token() Token {
	return NewToken(DomainRegistry, NewZ(p.TokenId))
}

func (c *domainClient) LookupByName(ctx context.Context, name string) (*Domain, error) {
	p := &Domain{}
	u := fmt.Sprintf("/v1/domains/%s", name)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *domainClient) LookupByAddress(ctx context.Context, addr Address) (*Domain, error) {
	p := &Domain{}
	u := fmt.Sprintf("/v1/domains/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *domainClient) ListDomains(ctx context.Context, params Query) ([]*Domain, error) {
	list := make([]*Domain, 0)
	u := params.WithPath("/v1/domains").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
