// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

import (
	// "context"

	"blockwatch.cc/tzpro-go/internal/client"
)

type ZmqAPI interface {
	// DecodeOp(context.Context, *Message) (*Op, error)
}

func NewZmqAPI(c *client.Client) ZmqAPI {
	return &zmqClient{client: c}
}

type zmqClient struct {
	client *client.Client
}
