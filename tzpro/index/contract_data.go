// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/trilitech/tzpro-go/internal/util"
)

type ContractParameters struct {
	ContractValue                 // contract
	Entrypoint    string          `json:"entrypoint,omitempty"` // contract
	L2Address     *Address        `json:"l2_address,omitempty"` // rollup
	Kind          string          `json:"kind,omitempty"`       // rollup
	Method        string          `json:"method,omitempty"`     // rollup
	Args          json.RawMessage `json:"args,omitempty"`       // rollup
	Result        json.RawMessage `json:"result,omitempty"`     // rollup
}

type ContractScript struct {
	Script          *Script          `json:"script,omitempty"`
	StorageType     Typedef          `json:"storage_type"`
	Entrypoints     Entrypoints      `json:"entrypoints"`
	Views           Views            `json:"views,omitempty"`
	BigmapNames     map[string]int64 `json:"bigmaps,omitempty"`
	BigmapTypes     map[string]Type  `json:"bigmap_types,omitempty"`
	BigmapTypesById map[int64]Type   `json:"-"`
}

func (s ContractScript) Types() (param, store Type, eps Entrypoints, bigmaps map[int64]Type) {
	param = s.Script.ParamType()
	store = s.Script.StorageType()
	eps, _ = s.Script.Entrypoints(true)
	bigmaps = s.BigmapTypesById
	return
}

type ContractValue struct {
	Value any   `json:"value,omitempty"`
	Prim  *Prim `json:"prim,omitempty"`
}

func (v ContractValue) IsPrim() bool {
	if v.Value == nil {
		return false
	}
	if m, ok := v.Value.(map[string]any); !ok {
		return false
	} else {
		_, ok := m["prim"]
		return ok
	}
}

func (v ContractValue) AsPrim() (Prim, bool) {
	if v.Prim.IsValid() {
		return *v.Prim, true
	}

	if v.IsPrim() {
		buf, _ := json.Marshal(v.Value)
		p := Prim{}
		err := p.UnmarshalJSON(buf)
		return p, err == nil
	}

	return Prim{}, false
}

func (v ContractValue) Has(path string) bool {
	return util.HasPath(v.Value, path)
}

func (v ContractValue) GetString(path string) (string, bool) {
	return util.GetPathString(v.Value, path)
}

func (v ContractValue) GetInt64(path string) (int64, bool) {
	return util.GetPathInt64(v.Value, path)
}

func (v ContractValue) GetBig(path string) (*big.Int, bool) {
	return util.GetPathBig(v.Value, path)
}

func (v ContractValue) GetZ(path string) (Z, bool) {
	return util.GetPathZ(v.Value, path)
}

func (v ContractValue) GetTime(path string) (time.Time, bool) {
	return util.GetPathTime(v.Value, path)
}

func (v ContractValue) GetAddress(path string) (Address, bool) {
	return util.GetPathAddress(v.Value, path)
}

func (v ContractValue) GetValue(path string) (interface{}, bool) {
	return util.GetPathValue(v.Value, path)
}

func (v ContractValue) Walk(path string, fn util.ValueWalkerFunc) error {
	val := v.Value
	if len(path) > 0 {
		var ok bool
		val, ok = util.GetPathValue(val, path)
		if !ok {
			return nil
		}
	}
	return util.WalkValueMap(path, val, fn)
}

func (v ContractValue) Unmarshal(val interface{}) error {
	buf, _ := json.Marshal(v.Value)
	return json.Unmarshal(buf, val)
}
