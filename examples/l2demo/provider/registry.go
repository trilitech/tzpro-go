// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package provider

import (
	"blockwatch.cc/tzpro-go/tzpro"
)

func Providers() []Provider {
	return reg.asList
}

func GetProvider(name string) (Provider, bool) {
	return reg.get(name)
}

func MinStatus() tzpro.BlockId {
	id := tzpro.BlockId{}
	first := true
	for _, v := range reg.asList {
		if !v.Enabled() {
			continue
		}
		if first {
			id = v.Status()
			first = false
		} else if s := v.Status(); s.Height < id.Height {
			id = s
		}
	}
	return id
}

var reg = newRegistry()

func register(p Provider) {
	reg.register(p)
}

type registry struct {
	byName map[string]Provider
	asList []Provider
}

func newRegistry() *registry {
	return &registry{byName: make(map[string]Provider)}
}

func (r *registry) register(p Provider) {
	r.byName[p.Name()] = p
	r.asList = append(r.asList, p)
}

func (r registry) get(name string) (Provider, bool) {
	p, ok := r.byName[name]
	return p, ok
}
