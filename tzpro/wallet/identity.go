// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"context"
	"fmt"
)

func (c *walletClient) ListDomains(ctx context.Context, addr Address, params Query) ([]*Domain, error) {
	list := make([]*Domain, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/domains", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListDomainEvents(ctx context.Context, addr Address, params Query) ([]*DomainEvent, error) {
	list := make([]*DomainEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/domain_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) GetProfile(ctx context.Context, addr Address) (*Profile, error) {
	p := &Profile{}
	u := fmt.Sprintf("/v1/wallets/%s/profile", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *walletClient) ListProfileEvents(ctx context.Context, addr Address, params Query) ([]*ProfileEvent, error) {
	list := make([]*ProfileEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/profile_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListProfileClaims(ctx context.Context, addr Address, params Query) ([]*ProfileClaim, error) {
	list := make([]*ProfileClaim, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/profile_claims", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
