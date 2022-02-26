// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "fmt"
)

// Indexer operation and event type
type OpType byte

// enums are allocated in chronological order with most often used ops first
const (
    OpTypeBake                         OpType = iota // 0
    OpTypeEndorsement                                // 1
    OpTypeTransaction                                // 2
    OpTypeReveal                                     // 3
    OpTypeDelegation                                 // 4
    OpTypeOrigination                                // 5
    OpTypeSeedNonceRevelation                        // 6
    OpTypeActivateAccount                            // 7
    OpTypeBallot                                     // 8
    OpTypeProposals                                  // 9
    OpTypeDoubleBakingEvidence                       // 10
    OpTypeDoubleEndorsementEvidence                  // 11
    OpTypeUnfreeze                                   // 12 implicit event
    OpTypeInvoice                                    // 13 implicit event
    OpTypeAirdrop                                    // 14 implicit event
    OpTypeSeedSlash                                  // 15 implicit event
    OpTypeMigration                                  // 16 implicit event
    OpTypeRegisterConstant                           // 17 v011
    OpTypePreendorsement                             // 28 v012
    OpTypeDoublePreendorsementEvidence               // 19 v012
    OpTypeSetDepositsLimit                           // 20 v012
    OpTypeDeposit                                    // 21 v012 implicit event (baker deposit)
    OpTypeBonus                                      // 22 v012 implicit event (baker extra bonus)
    OpTypeReward                                     // 23 v012 implicit event (endorsement reward pay/burn)
    OpTypeBatch                        = 254         // API output only
    OpTypeInvalid                      = 255
)

var (
    opTypeStrings = map[OpType]string{
        OpTypeBake:                         "bake",
        OpTypeEndorsement:                  "endorsement",
        OpTypeTransaction:                  "transaction",
        OpTypeReveal:                       "reveal",
        OpTypeDelegation:                   "delegation",
        OpTypeOrigination:                  "origination",
        OpTypeSeedNonceRevelation:          "seed_nonce_revelation",
        OpTypeActivateAccount:              "activate_account",
        OpTypeBallot:                       "ballot",
        OpTypeProposals:                    "proposals",
        OpTypeDoubleBakingEvidence:         "double_baking_evidence",
        OpTypeDoubleEndorsementEvidence:    "double_endorsement_evidence",
        OpTypeUnfreeze:                     "unfreeze",
        OpTypeInvoice:                      "invoice",
        OpTypeAirdrop:                      "airdrop",
        OpTypeSeedSlash:                    "seed_slash",
        OpTypeMigration:                    "migration",
        OpTypeRegisterConstant:             "register_global_constant",
        OpTypePreendorsement:               "preendorsement",
        OpTypeDoublePreendorsementEvidence: "double_preendorsement_evidence",
        OpTypeSetDepositsLimit:             "set_deposits_limit",
        OpTypeDeposit:                      "deposit",
        OpTypeBonus:                        "bonus",
        OpTypeReward:                       "reward",
        OpTypeBatch:                        "batch",
        OpTypeInvalid:                      "",
    }
    opTypeReverseStrings = make(map[string]OpType)
)

func init() {
    for n, v := range opTypeStrings {
        opTypeReverseStrings[v] = n
    }
}

func (t OpType) IsValid() bool {
    return t != OpTypeInvalid
}

func (t *OpType) UnmarshalText(data []byte) error {
    v := ParseOpType(string(data))
    if !v.IsValid() {
        return fmt.Errorf("invalid operation type '%s'", string(data))
    }
    *t = v
    return nil
}

func (t *OpType) MarshalText() ([]byte, error) {
    return []byte(t.String()), nil
}

func ParseOpType(s string) OpType {
    t, ok := opTypeReverseStrings[s]
    if !ok {
        t = OpTypeInvalid
    }
    return t
}

func (t OpType) String() string {
    return opTypeStrings[t]
}
