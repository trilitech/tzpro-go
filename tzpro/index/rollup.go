// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"github.com/trilitech/tzgo/micheline"
	"github.com/trilitech/tzgo/tezos"
)

type SmartRollupResult struct {
	Address          *tezos.Address               `json:"address,omitempty"`
	Size             *tezos.Z                     `json:"size,omitempty"`
	InboxLevel       int64                        `json:"inbox_level,omitempty"`
	StakedHash       *tezos.SmartRollupCommitHash `json:"staked_hash,omitempty"`
	PublishedAtLevel int64                        `json:"published_at_level,omitempty"`
	GameStatus       *struct {
		Status string         `json:"status,omitempty"`
		Kind   string         `json:"kind,omitempty"`
		Reason string         `json:"reason,omitempty"`
		Player *tezos.Address `json:"player,omitempty"`
	} `json:"game_status,omitempty"`
	Commitment *tezos.SmartRollupCommitHash `json:"commitment_hash,omitempty"`
}

type SmartRollupOriginate struct {
	PvmKind          tezos.PvmKind  `json:"pvm_kind"`
	Kernel           tezos.HexBytes `json:"kernel"`
	OriginationProof tezos.HexBytes `json:"origination_proof"`
	ParametersTy     micheline.Prim `json:"parameters_ty"`
}

type SmartRollupAddMessages struct {
	Messages []tezos.HexBytes `json:"message"`
}

// Deprecated: in v17
type SmartRollupCement struct {
	Commitment *tezos.SmartRollupCommitHash `json:"commitment,omitempty"`
}

type SmartRollupPublish struct {
	Commitment struct {
		CompressedState tezos.SmartRollupStateHash  `json:"compressed_state"`
		InboxLevel      int64                       `json:"inbox_level"`
		Predecessor     tezos.SmartRollupCommitHash `json:"predecessor"`
		NumberOfTicks   tezos.Z                     `json:"number_of_ticks"`
	} `json:"commitment"`
}

type SmartRollupRefute struct {
	Opponent   tezos.Address `json:"opponent"`
	Refutation struct {
		Kind         string                       `json:"refutation_kind"`
		PlayerHash   *tezos.SmartRollupCommitHash `json:"player_commitment_hash,omitempty"`
		OpponentHash *tezos.SmartRollupCommitHash `json:"opponent_commitment_hash,omitempty"`
		Choice       *tezos.Z                     `json:"choice,omitempty"`
		Step         *struct {
			Ticks []struct {
				State tezos.SmartRollupStateHash `json:"state"`
				Tick  tezos.Z                    `json:"tick"`
			} `json:"ticks,omitempty"`
			Proof *struct {
				PvmStep    tezos.HexBytes `json:"pvm_step,omitempty"`
				InputProof *struct {
					Kind    string         `json:"input_proof_kind"`
					Level   int64          `json:"level"`
					Counter tezos.Z        `json:"message_counter"`
					Proof   tezos.HexBytes `json:"serialized_proof"`
				} `json:"input_proof,omitempty"`
			} `json:"proof,omitempty"`
		} `json:"step,omitempty"`
	} `json:"refutation"`
}

type SmartRollupTimeout struct {
	Stakers struct {
		Alice tezos.Address `json:"alice"`
		Bob   tezos.Address `json:"bob"`
	} `json:"stakers"`
}

type SmartRollupExecuteOutboxMessage struct {
	CementedCommitment tezos.SmartRollupCommitHash `json:"cemented_commitment"`
	OutputProof        tezos.HexBytes              `json:"output_proof"`
}

type SmartRollupRecoverBond struct {
	Staker tezos.Address `json:"staker"`
}
