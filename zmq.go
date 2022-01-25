// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"encoding/json"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
)

type ZmqMessage struct {
	topic  string
	body   []byte
	fields []string
}

func NewZmqMessage(topic, body []byte) *ZmqMessage {
	return &ZmqMessage{string(topic), body, nil}
}

func (m *ZmqMessage) GetField(name string) (string, bool) {
	return getTableColumn(m.body[1:len(m.body)-1], zmqFields(m.topic), name)
}

func (m *ZmqMessage) DecodeOpHash() (tezos.OpHash, error) {
	return tezos.ParseOpHash(string(m.body))
}

func (m *ZmqMessage) DecodeBlockHash() (tezos.BlockHash, error) {
	return tezos.ParseBlockHash(string(m.body))
}

func (m *ZmqMessage) DecodeOp() (*Op, error) {
	o := new(Op).WithColumns(ZmqRawOpColumns...)
	if err := json.Unmarshal(m.body, o); err != nil {
		return nil, err
	}
	return o, nil
}

func (m *ZmqMessage) DecodeOpWithScript(ctx context.Context, c *Client) (*Op, error) {
	o := new(Op).WithColumns(ZmqRawOpColumns...)

	// we may need contract scripts
	if is, ok := m.GetField("is_contract"); ok && is == "1" {
		recv, ok := m.GetField("receiver")
		if ok && recv != "" && recv != "null" {
			addr, err := tezos.ParseAddress(recv)
			if err != nil {
				return nil, fmt.Errorf("decode: invalid receiver address %s: %v, %#v", recv, err, string(m.body))
			}
			// load contract type info (required for decoding storage/param data)
			script, err := c.loadCachedContractScript(ctx, addr)
			if err != nil {
				return nil, err
			}
			o = o.WithScript(script)
		}
	}
	if err := json.Unmarshal(m.body, o); err != nil {
		return nil, err
	}
	return o, nil
}

func (m *ZmqMessage) DecodeBlock() (*Block, error) {
	b := new(Block).WithColumns(ZmqRawBlockColumns...)
	if err := json.Unmarshal(m.body, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (m *ZmqMessage) DecodeStatus() (*Status, error) {
	s := new(Status).WithColumns(ZmqStatusColumns...)
	if err := json.Unmarshal(m.body, s); err != nil {
		return nil, err
	}
	return s, nil
}

func zmqFields(topic string) []string {
	switch topic {
	case "raw_block", "raw_block/rollback":
		return ZmqRawBlockColumns
	case "raw_op", "raw_op/rollback":
		return ZmqRawOpColumns
	case "status":
		return ZmqStatusColumns
	default:
		return nil
	}
}

var ZmqRawBlockColumns = []string{
	"row_id",
	"parent_id",
	"hash",
	"is_orphan",
	"height",
	"cycle",
	"is_cycle_snapshot",
	"time",
	"solvetime",
	"version",
	"fitness",
	"priority",
	"nonce",
	"voting_period_kind",
	"baker_id",
	"endorsed_slots",
	"n_endorsed_slots",
	"n_ops",
	"n_ops_failed",
	"n_ops_contract",
	"n_contract_calls",
	"n_tx",
	"n_activation",
	"n_seed_nonce_revelation",
	"n_double_baking_evidence",
	"n_double_endorsement_evidence",
	"n_endorsement",
	"n_delegation",
	"n_reveal",
	"n_origination",
	"n_proposal",
	"n_ballot",
	"n_register_constant",
	"volume",
	"fee",
	"reward",
	"deposit",
	"activated_supply",
	"burned_supply",
	"n_accounts",
	"n_new_accounts",
	"n_new_contracts",
	"n_cleared_accounts",
	"n_funded_accounts",
	"gas_limit",
	"gas_used",
	"gas_price",
	"storage_size",
	"days_destroyed",
	"n_ops_implicit",
	"pct_account_reuse",
	"baker",
	"predecessor",
	"lb_esc_vote",
	"lb_esc_ema",
	"protocol",
}

var ZmqRawOpColumns = []string{
	"row_id",
	"time",
	"height",
	"cycle",
	"hash",
	"counter",
	"op_n",
	"op_c",
	"op_i",
	"op_l",
	"op_p",
	"type",
	"status",
	"is_success",
	"is_contract",
	"gas_limit",
	"gas_used",
	"gas_price",
	"storage_limit",
	"storage_size",
	"storage_paid",
	"volume",
	"fee",
	"reward",
	"deposit",
	"burned",
	"sender_id",
	"receiver_id",
	"creator_id",
	"delegate_id",
	"is_internal",
	"has_data",
	"data",
	"parameters",
	"storage",
	"big_map_diff",
	"errors",
	"days_destroyed",
	"branch_id",
	"branch_height",
	"branch_depth",
	"is_implicit",
	"entrypoint_id",
	"is_orphan",
	"sender",
	"receiver",
	"creator",
	"delegate",
	"is_batch",
	"is_sapling",
	"block_hash",
}

var ZmqStatusColumns = []string{
	"status",
	"blocks",
	"indexed",
	"progress",
}
