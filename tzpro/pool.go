// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"fmt"
	"strconv"
	"strings"

	"blockwatch.cc/tzgo/tezos"
)

// PoolAddress represents a specialized address that consists of
// a smart contract KT1 address and a numeric pool id.
type PoolAddress struct {
	Hash [20]byte // type is always KT1
	Id   int
}

func NewPoolAddress(contract tezos.Address, id int) (a PoolAddress) {
	copy(a.Hash[:], contract[1:])
	a.Id = id
	return
}

func ParsePoolAddress(s string) (PoolAddress, error) {
	x, y, ok := strings.Cut(s, "_")
	if !ok {
		return PoolAddress{}, fmt.Errorf("invalid pool address")
	}
	addr, err := tezos.ParseAddress(x)
	if err != nil {
		return PoolAddress{}, err
	}
	id, err := strconv.Atoi(y)
	if err != nil {
		return PoolAddress{}, err
	}
	return NewPoolAddress(addr, id), nil
}

func MustParsePoolAddress(s string) PoolAddress {
	a, err := ParsePoolAddress(s)
	if err != nil {
		panic(err)
	}
	return a
}

func (a PoolAddress) Contract() tezos.Address {
	return tezos.NewAddress(tezos.AddressTypeContract, a.Hash[:])
}

func (a PoolAddress) PoolId() int {
	return a.Id
}

func (a PoolAddress) Equal(b PoolAddress) bool {
	return a.Hash == b.Hash && a.Id == b.Id
}

func (a PoolAddress) Clone() PoolAddress {
	return PoolAddress{
		Hash: a.Hash,
		Id:   a.Id,
	}
}

func (a PoolAddress) String() string {
	return a.Contract().String() + "_" + strconv.Itoa(a.Id)
}

func (a *PoolAddress) UnmarshalText(data []byte) error {
	addr, err := ParsePoolAddress(string(data))
	if err == nil {
		*a = addr
	}
	return err
}
