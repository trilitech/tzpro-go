// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "context"
    "fmt"
    "strconv"
    "strings"
    "time"

    "blockwatch.cc/tzgo/tezos"
)

//nolint:staticcheck
type DexPool struct {
    Id              int64         `json:"id"`
    Contract        tezos.Address `json:"contract"`
    PairId          int           `json:"pair_id"`
    Creator         string        `json:"creator"`
    Name            string        `json:"name"`
    Entity          string        `json:"entity"`
    Pair            string        `json:"pair"`
    NumTokens       int           `json:"num_tokens"`
    TokenA          *Token        `json:"token_a"`
    TokenB          *Token        `json:"token_b"`
    TokenLP         *Token        `json:"token_lp"`
    FirstBlock      int64         `json:"first_block"`
    FirstTime       time.Time     `json:"first_time"`
    Tags            string        `json:"tags"`
    SupplyA         tezos.Z       `json:"supply_a"`
    SupplyB         tezos.Z       `json:"supply_b"`
    SupplyLP        tezos.Z       `json:"supply_lp"`
    LastChangeBlock int64         `json:"last_change_block"`
    LastChangeTime  time.Time     `json:"last_change_time"`
}

type DexPoolParams struct {
    Params
}

func NewDexPoolParams() DexPoolParams {
    return DexPoolParams{NewParams()}
}

func (p DexPoolParams) WithLimit(v uint) DexPoolParams {
    p.Query.Set("limit", strconv.Itoa(int(v)))
    return p
}

func (p DexPoolParams) WithOffset(v uint) DexPoolParams {
    p.Query.Set("offset", strconv.Itoa(int(v)))
    return p
}

func (p DexPoolParams) WithCursor(v uint64) DexPoolParams {
    p.Query.Set("cursor", strconv.FormatUint(v, 10))
    return p
}

func (p DexPoolParams) WithOrder(o OrderType) DexPoolParams {
    p.Query.Set("order", string(o))
    return p
}

func (p DexPoolParams) WithDesc() DexPoolParams {
    p.Query.Set("order", string(OrderDesc))
    return p
}

func (p DexPoolParams) WithAsc() DexPoolParams {
    p.Query.Set("order", string(OrderAsc))
    return p
}

func (p DexPoolParams) WithContract(c tezos.Address) DexPoolParams {
    p.Query.Set("contract", c.String())
    return p
}

func (p DexPoolParams) WithCreator(c tezos.Address) DexPoolParams {
    p.Query.Set("creator", c.String())
    return p
}

func (p DexPoolParams) WithPairId(id int) DexPoolParams {
    p.Query.Set("pair_id", strconv.Itoa(id))
    return p
}

func (p DexPoolParams) WithEntity(e string) DexPoolParams {
    p.Query.Set("entity", e)
    return p
}

func (p DexPoolParams) WithName(n string) DexPoolParams {
    p.Query.Set("name", n)
    return p
}

func (p DexPoolParams) WithPair(s string) DexPoolParams {
    p.Query.Set("pair", s)
    return p
}

func (p DexPoolParams) WithSymbol(s string) DexPoolParams {
    p.Query.Set("symbol", s)
    return p
}

func (p DexPoolParams) WithFirstBlock(height int64) DexPoolParams {
    p.Query.Set("block", strconv.FormatInt(height, 10))
    return p
}

func (p DexPoolParams) WithFirstTime(t time.Time) DexPoolParams {
    p.Query.Set("time", t.Format(time.RFC3339))
    return p
}

func (p DexPoolParams) WithTags(t ...string) DexPoolParams {
    p.Query.Set("tags", strings.Join(t, ","))
    return p
}

func (c *Client) GetDexPool(ctx context.Context, addr tezos.Address, id int, params DexPoolParams) (*DexPool, error) {
    p := &DexPool{}
    u := params.AppendQuery(fmt.Sprintf("/v1/dex/pools/%s_%d", addr, id))
    if err := c.get(ctx, u, nil, p); err != nil {
        return nil, err
    }
    return p, nil
}

func (c *Client) ListDexPools(ctx context.Context, params DexPoolParams) ([]*DexPool, error) {
    list := make([]*DexPool, 0)
    u := params.AppendQuery("/v1/dex/pools")
    if err := c.get(ctx, u, nil, &list); err != nil {
        return nil, err
    }
    return list, nil
}
