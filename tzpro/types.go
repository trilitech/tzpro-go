// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"time"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/tzpro/index"
)

type (
	Address     = tezos.Address
	AddressType = tezos.AddressType
	Key         = tezos.Key
	Token       = tezos.Token
	Z           = tezos.Z

	Query      = client.Query
	FilterMode = client.FilterMode
	FillMode   = client.FillMode
	OrderType  = client.OrderType
	FormatType = client.FormatType
	ApiError   = client.ApiError
	ApiErrors  = client.ApiErrors
)

var (
	NewAddress       = tezos.MustParseAddress
	ParseAddress     = tezos.ParseAddress
	NewQuery         = client.NewQuery
	IsErrApi         = client.IsErrApi
	IsErrHttp        = client.IsErrHttp
	IsErrRateLimited = client.IsErrRateLimited
	ErrorStatus      = client.ErrorStatus

	NoQuery = NewQuery()
)

const (
	FillModeInvalid FillMode = ""
	FillModeNone    FillMode = "none"
	FillModeNull    FillMode = "null"
	FillModeLast    FillMode = "last"
	FillModeLinear  FillMode = "linear"
	FillModeZero    FillMode = "zero"
)

const (
	Collapse1m = time.Minute
	Collapse1h = time.Hour
	Collapse1d = 24 * time.Hour
	Collapse1w = 7 * 24 * time.Hour
)

const (
	FilterModeEqual    FilterMode = "eq"
	FilterModeNotEqual FilterMode = "ne"
	FilterModeGt       FilterMode = "gt"
	FilterModeGte      FilterMode = "gte"
	FilterModeLt       FilterMode = "lt"
	FilterModeLte      FilterMode = "lte"
	FilterModeIn       FilterMode = "in"
	FilterModeNotIn    FilterMode = "nin"
	FilterModeRange    FilterMode = "rg"
	FilterModeRegexp   FilterMode = "re"
)

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

const (
	FormatJSON FormatType = "json"
	FormatCSV  FormatType = "csv"
)

const (
	OpTypeBake                 = index.OpTypeBake
	OpTypeEndorsement          = index.OpTypeEndorsement
	OpTypeTransaction          = index.OpTypeTransaction
	OpTypeReveal               = index.OpTypeReveal
	OpTypeDelegation           = index.OpTypeDelegation
	OpTypeOrigination          = index.OpTypeOrigination
	OpTypeNonceRevelation      = index.OpTypeNonceRevelation
	OpTypeActivation           = index.OpTypeActivation
	OpTypeBallot               = index.OpTypeBallot
	OpTypeProposal             = index.OpTypeProposal
	OpTypeDoubleBaking         = index.OpTypeDoubleBaking
	OpTypeDoubleEndorsement    = index.OpTypeDoubleEndorsement
	OpTypeUnfreeze             = index.OpTypeUnfreeze
	OpTypeInvoice              = index.OpTypeInvoice
	OpTypeAirdrop              = index.OpTypeAirdrop
	OpTypeSeedSlash            = index.OpTypeSeedSlash
	OpTypeMigration            = index.OpTypeMigration
	OpTypeSubsidy              = index.OpTypeSubsidy
	OpTypeRegisterConstant     = index.OpTypeRegisterConstant
	OpTypePreendorsement       = index.OpTypePreendorsement
	OpTypeDoublePreendorsement = index.OpTypeDoublePreendorsement
	OpTypeDepositsLimit        = index.OpTypeDepositsLimit
	OpTypeDeposit              = index.OpTypeDeposit
	OpTypeBonus                = index.OpTypeBonus
	OpTypeReward               = index.OpTypeReward
	OpTypeRollupOrigination    = index.OpTypeRollupOrigination
	OpTypeRollupTransaction    = index.OpTypeRollupTransaction
	OpTypeVdfRevelation        = index.OpTypeVdfRevelation
	OpTypeIncreasePaidStorage  = index.OpTypeIncreasePaidStorage
	OpTypeDrainDelegate        = index.OpTypeDrainDelegate
	OpTypeUpdateConsensusKey   = index.OpTypeUpdateConsensusKey
	OpTypeTransferTicket       = index.OpTypeTransferTicket
	OpTypeBatch                = index.OpTypeBatch
	OpTypeInvalid              = index.OpTypeInvalid
)

func AndArg(key string, val ...any) Query {
	return NewQuery().AndArg(key, val...)
}

func AndFilter(key string, mode FilterMode, val ...any) Query {
	return NewQuery().AndFilter(key, mode, val...)
}

func AndEqual(key string, val any) Query {
	return NewQuery().AndEqual(key, val)
}

func AndNotEqual(key string, val any) Query {
	return NewQuery().AndNotEqual(key, val)
}

func AndGt(key string, val any) Query {
	return NewQuery().AndGt(key, val)
}

func AndGte(key string, val any) Query {
	return NewQuery().AndGte(key, val)
}

func AndLt(key string, val any) Query {
	return NewQuery().AndLt(key, val)
}

func AndLte(key string, val any) Query {
	return NewQuery().AndLte(key, val)
}

func AndIn(key string, val ...any) Query {
	return NewQuery().AndIn(key, val...)
}

func AndNotIn(key string, val ...any) Query {
	return NewQuery().AndNotIn(key, val...)
}

func AndRange(key string, from, to any) Query {
	return NewQuery().AndRange(key, from, to)
}

func AndRegexp(key string, re string) Query {
	return NewQuery().AndRegexp(key, re)
}

func And(key string) Query {
	return NewQuery().And(key)
}

func AndNot(key string) Query {
	return NewQuery().AndNot(key)
}

func WithLimit(v uint) Query {
	return NewQuery().WithLimit(v)
}

func WithOffset(v uint) Query {
	return NewQuery().WithOffset(v)
}

func WithCursor(v uint64) Query {
	return NewQuery().WithCursor(v)
}

func WithOrder(o OrderType) Query {
	return NewQuery().WithOrder(o)
}

func Desc() Query {
	return NewQuery().Desc()
}

func Asc() Query {
	return NewQuery().Asc()
}

func WithMeta() Query {
	return NewQuery().WithMeta()
}

func WithTags(t ...string) Query {
	return NewQuery().WithTags(t...)
}

func WithFrom(t time.Time) Query {
	return NewQuery().WithFrom(t)
}

func WithTo(t time.Time) Query {
	return NewQuery().WithTo(t)
}

func WithTimeRange(from, to time.Time) Query {
	return NewQuery().WithTimeRange(from, to)
}

func WithPrim() Query {
	return NewQuery().WithPrim()
}

func WithMerge() Query {
	return NewQuery().WithMerge()
}

func WithStorage() Query {
	return NewQuery().WithStorage()
}

func WithUnpack() Query {
	return NewQuery().WithUnpack()
}

func WithRights() Query {
	return NewQuery().WithRights()
}
