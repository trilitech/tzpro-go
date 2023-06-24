// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"errors"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
)

type (
	Query = client.Query

	OpHash       = tezos.OpHash
	OpStatus     = tezos.OpStatus
	BlockHash    = tezos.BlockHash
	ExprHash     = tezos.ExprHash
	ProtocolHash = tezos.ProtocolHash
	ChainIdHash  = tezos.ChainIdHash
	Address      = tezos.Address
	AddressType  = tezos.AddressType
	AddressSet   = tezos.AddressSet
	RightType    = tezos.RightType
	Key          = tezos.Key
	Token        = tezos.Token
	Z            = tezos.Z

	Script       = micheline.Script
	Prim         = micheline.Prim
	Value        = micheline.Value
	Type         = micheline.Type
	Typedef      = micheline.Typedef
	BigmapKey    = micheline.Key
	Views        = micheline.Views
	DiffAction   = micheline.DiffAction
	Entrypoints  = micheline.Entrypoints
	Parameters   = micheline.Parameters
	BigmapEvents = micheline.BigmapEvents
	BigmapEvent  = micheline.BigmapEvent
)

var (
	NewQuery           = client.NewQuery
	NewAddressSet      = tezos.NewAddressSet
	NewValue           = micheline.NewValue
	NewType            = micheline.NewType
	NewKey             = micheline.NewKey
	DiffActionAlloc    = micheline.DiffActionAlloc
	DiffActionCopy     = micheline.DiffActionCopy
	DiffActionUpdate   = micheline.DiffActionUpdate
	DiffActionRemove   = micheline.DiffActionRemove
	RightTypeBaking    = tezos.RightTypeBaking
	RightTypeEndorsing = tezos.RightTypeEndorsing
)

var (
	ErrNoStorage    = errors.New("no storage")
	ErrNoParams     = errors.New("no parameters")
	ErrNoBigmapDiff = errors.New("no bigmap diff")
	ErrNoType       = errors.New("API type missing")
)
