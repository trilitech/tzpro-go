// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"

	"blockwatch.cc/tzgo/micheline"
)

func (c *opClient) loadScript(ctx context.Context, addr Address) (*ContractScript, error) {
	if script, ok := c.client.CacheGet(addr); ok {
		return script.(*ContractScript), nil
	}
	api := NewContractAPI(c.client)
	script, err := api.GetScript(ctx, addr, NewQuery().WithPrim())
	if err != nil {
		return nil, err
	}
	// strip code
	script.Script.Code.Code = micheline.Prim{}
	script.Script.Code.View = micheline.Prim{}
	// fill bigmap type info
	script.BigmapNames = script.Script.Bigmaps()
	script.BigmapTypes = script.Script.BigmapTypes()
	script.BigmapTypesById = make(map[int64]Type)
	for n, v := range script.BigmapTypes {
		id := script.BigmapNames[n]
		script.BigmapTypesById[id] = v
	}
	c.client.CacheAdd(addr, script)
	return script, nil
}

// func (c *Client) AddCachedScript(addr Address, script *micheline.Script) {
// 	if !addr.IsValid() || script == nil || c.cache == nil {
// 		return
// 	}
// 	eps, _ := script.Entrypoints(true)
// 	views, _ := script.Views(true, false)
// 	s := &ContractScript{
// 		Script:          script,
// 		StorageType:     script.StorageType().Typedef(""),
// 		Entrypoints:     eps,
// 		Views:           views,
// 		BigmapNames:     script.Bigmaps(),
// 		BigmapTypes:     script.BigmapTypes(),
// 		BigmapTypesById: make(map[int64]micheline.Type),
// 	}
// 	for n, v := range s.BigmapTypes {
// 		id := s.BigmapNames[n]
// 		s.BigmapTypesById[id] = v
// 	}
// 	c.cache.Add(addr.String(), s)
// }
