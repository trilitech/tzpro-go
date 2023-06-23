// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/internal/util"
)

type Constant struct {
	RowId       uint64          `json:"row_id"`
	Address     ExprHash        `json:"address"`
	CreatorId   uint64          `json:"creator_id"`
	Creator     Address         `json:"creator"`
	Height      int64           `json:"height"`
	Time        time.Time       `json:"time"`
	StorageSize int64           `json:"storage_size"`
	Value       Prim            `json:"value"          tzpro:",hex"`
	Features    util.StringList `json:"features"`
}

type ConstantQuery = client.TableQuery[*Constant]

func (a contractClient) NewConstantQuery() *ConstantQuery {
	return client.NewTableQuery[*Constant](a.client, "constant")
}

func (c *contractClient) GetConstant(ctx context.Context, addr ExprHash, params Params) (*Constant, error) {
	cc := &Constant{}
	u := params.WithPath(fmt.Sprintf("/explorer/constant/%s", addr)).Url()
	if err := c.client.Get(ctx, u, nil, cc); err != nil {
		return nil, err
	}
	return cc, nil
}
