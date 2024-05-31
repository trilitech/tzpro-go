// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package ipfs

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/trilitech/tzpro-go/internal/client"
)

type IpfsAPI interface {
	GetData(context.Context, string, any) error
	GetImage(context.Context, string, string, io.Writer) error
}

func NewIpfsAPI(c *client.Client) IpfsAPI {
	return &ipfsClient{client: c}
}

type ipfsClient struct {
	client *client.Client
}

func (c *ipfsClient) GetData(ctx context.Context, uri string, val any) error {
	if strings.HasPrefix(uri, "ipfs://") {
		uri = "/ipfs/" + strings.TrimPrefix(uri, "ipfs://")
	}
	return c.client.Get(ctx, uri, nil, val)
}

func (c *ipfsClient) GetImage(ctx context.Context, uri, mime string, w io.Writer) error {
	if strings.HasPrefix(uri, "ipfs://") {
		uri = "/ipfs/" + strings.TrimPrefix(uri, "ipfs://")
	}
	h := make(http.Header)
	h.Add("Accept", mime)
	return c.client.Get(ctx, uri, h, w)
}
