// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

func Fields(topic string) []string {
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
	"hash",
	"predecessor",
	"height",
	"cycle",
	"time",
	"solvetime",
	"version",
	"round",
	"nonce",
	"voting_period_kind",
	"baker",
	"proposer",
	"n_ops_applied",
	"n_ops_failed",
	"n_calls",
	"n_rollup_calls",
	"n_events",
	"volume",
	"fee",
	"reward",
	"deposit",
	"activated_supply",
	"burned_supply",
	"minted_supply",
	"n_accounts",
	"n_new_accounts",
	"n_new_contracts",
	"n_cleared_accounts",
	"n_funded_accounts",
	"gas_limit",
	"gas_used",
	"storage_paid",
	"lb_esc_vote",
	"lb_esc_ema",
	"protocol",
}

var ZmqRawOpColumns = []string{
	"id",
	"type",
	"hash",
	"block",
	"height",
	"cycle",
	"time",
	"op_n",
	"op_p",
	"op_c",
	"op_i",
	"status",
	"is_success",
	"is_contract",
	"is_internal",
	"is_event",
	"is_rollup",
	"counter",
	"gas_limit",
	"gas_used",
	"storage_limit",
	"storage_paid",
	"volume",
	"fee",
	"reward",
	"deposit",
	"burned",
	"sender_id",
	"sender",
	"receiver_id",
	"receiver",
	"creator_id",
	"creator",
	"baker_id",
	"baker",
	"data",
	"parameters",
	"storage",
	"big_map_diff",
	"errors",
	"entrypoint",
	"code_hash",
	"events",
	"ticket_updates",
}

var ZmqStatusColumns = []string{
	"status",
	"blocks",
	"finalized",
	"indexed",
	"progress",
}
