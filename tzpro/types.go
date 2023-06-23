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

	Params     = client.Params
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
	NewParams        = client.NewParams
	IsErrApi         = client.IsErrApi
	IsErrHttp        = client.IsErrHttp
	IsErrRateLimited = client.IsErrRateLimited
	ErrorStatus      = client.ErrorStatus
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
