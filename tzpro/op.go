// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"

	// "encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	// "strconv"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

var (
	ErrNoStorage    = errors.New("no storage")
	ErrNoParams     = errors.New("no parameters")
	ErrNoBigmapDiff = errors.New("no bigmap diff")
	ErrNoType       = errors.New("API type missing")
)

type Costs struct {
	Fee            float64 // the total fee paid
	Burn           float64 // total amount burned (not included in fee)
	GasUsed        int64   // gas used
	StorageUsed    int64   // new storage bytes allocated
	StorageBurn    float64 // burned for allocating new storage (not included in fee)
	AllocationBurn float64 // burned for allocating a new account (not included in fee)
}

func (x Costs) Add(y Costs) Costs {
	x.Fee += y.Fee
	x.Burn += y.Burn
	x.GasUsed += y.GasUsed
	x.StorageUsed += y.StorageUsed
	x.StorageBurn += y.StorageBurn
	x.AllocationBurn += y.AllocationBurn
	return x
}

func (x Costs) Sum() float64 {
	return x.Fee + x.Burn
}

type Op struct {
	Id           uint64          `json:"id"`
	Type         OpType          `json:"type"`
	Hash         tezos.OpHash    `json:"hash"`
	Height       int64           `json:"height"`
	Cycle        int64           `json:"cycle"`
	Timestamp    time.Time       `json:"time"`
	OpN          int             `json:"op_n"`
	OpP          int             `json:"op_p"`
	Status       tezos.OpStatus  `json:"status"`
	IsSuccess    bool            `json:"is_success"`
	IsContract   bool            `json:"is_contract"`
	IsInternal   bool            `json:"is_internal"`
	IsEvent      bool            `json:"is_event"`
	IsRollup     bool            `json:"is_rollup"`
	Counter      int64           `json:"counter"`
	GasLimit     int64           `json:"gas_limit"`
	GasUsed      int64           `json:"gas_used"`
	StorageLimit int64           `json:"storage_limit"`
	StoragePaid  int64           `json:"storage_paid"`
	Volume       float64         `json:"volume"`
	Fee          float64         `json:"fee"`
	Reward       float64         `json:"reward"`
	Deposit      float64         `json:"deposit"`
	Burned       float64         `json:"burned"`
	SenderId     uint64          `json:"sender_id"`
	ReceiverId   uint64          `json:"receiver_id"`
	CreatorId    uint64          `json:"creator_id"`
	BakerId      uint64          `json:"baker_id"`
	Data         json.RawMessage `json:"data,omitempty"`
	Parameters   json.RawMessage `json:"parameters,omitempty"`
	BigmapDiff   json.RawMessage `json:"big_map_diff,omitempty"` // transaction, origination
	StorageHash  string          `json:"storage_hash,omitempty"`
	CodeHash     string          `json:"code_hash,omitempty"`
	Errors       json.RawMessage `json:"errors,omitempty"`
	Sender       tezos.Address   `json:"sender"`
	Receiver     tezos.Address   `json:"receiver"`
	Creator      tezos.Address   `json:"creator"` // origination
	Baker        tezos.Address   `json:"baker"`   // delegation, origination
	Block        tezos.BlockHash `json:"block"`
	Entrypoint   string          `json:"entrypoint,omitempty"`

	// explorer or ZMQ APIs only
	PrevBaker     tezos.Address       `json:"previous_baker"            tzpro:"-"` // delegation
	Source        tezos.Address       `json:"source"                    tzpro:"-"` // internal operations
	Offender      tezos.Address       `json:"offender"                  tzpro:"-"` // double_x
	Accuser       tezos.Address       `json:"accuser"                   tzpro:"-"` // double_x
	Loser         tezos.Address       `json:"loser"                     tzpro:"-"` // smart rollup refutation game
	Winner        tezos.Address       `json:"winner"                    tzpro:"-"` // smart rollup refutation game
	Staker        tezos.Address       `json:"staker"                    tzpro:"-"` // smart rollup refutation game
	Storage       json.RawMessage     `json:"storage,omitempty"         tzpro:"-"` // transaction, origination
	Script        *micheline.Script   `json:"script,omitempty"          tzpro:"-"` // origination
	Power         int                 `json:"power,omitempty"           tzpro:"-"` // endorsement
	Limit         *float64            `json:"limit,omitempty"           tzpro:"-"` // set deposits limit
	Confirmations int64               `json:"confirmations"             tzpro:"-"`
	NOps          int                 `json:"n_ops,omitempty"           tzpro:"-"`
	Batch         []*Op               `json:"batch,omitempty"           tzpro:"-"`
	Internal      []*Op               `json:"internal,omitempty"        tzpro:"-"`
	Metadata      map[string]Metadata `json:"metadata,omitempty"        tzpro:"-"`
	Events        []Event             `json:"events,omitempty"          tzpro:"-"`
	TicketUpdates []TicketUpdate      `json:"ticket_updates,omitempty"  tzpro:"-"`

	param    micheline.Type           // optional, may be decoded from script
	store    micheline.Type           // optional, may be decoded from script
	eps      micheline.Entrypoints    // optional, may be decoded from script
	bigmaps  map[int64]micheline.Type // optional, may be decoded from script
	withPrim bool
	withMeta bool
	noFail   bool
	onError  int
}

func (o *Op) BlockId() BlockId {
	return BlockId{
		Height: o.Height,
		Hash:   o.Block.Clone(),
		Time:   o.Timestamp,
	}
}

func (o *Op) Content() []*Op {
	list := []*Op{o}
	if len(o.Batch) == 0 && len(o.Internal) == 0 {
		return list
	}
	if len(o.Batch) > 0 {
		list = list[:0]
		for _, v := range o.Batch {
			list = append(list, v)
			if len(v.Internal) > 0 {
				list = append(list, v.Internal...)
			}
		}
	}
	if len(o.Internal) > 0 {
		list = append(list, o.Internal...)
	}
	return list
}

func (o *Op) Addresses() *tezos.AddressSet {
	set := tezos.NewAddressSet()
	for _, op := range o.Content() {
		for _, v := range []tezos.Address{
			op.Sender,
			op.Receiver,
			op.Creator,
			op.Baker,
			op.PrevBaker,
			op.Source,
			op.Offender,
			op.Accuser,
		} {
			if v.IsValid() {
				set.AddUnique(v)
			}
		}
	}
	return set
}

func (o *Op) Cursor() uint64 {
	op := o
	if l := len(op.Batch); l > 0 {
		op = op.Batch[l-1]
	}
	if l := len(op.Internal); l > 0 {
		op = op.Internal[l-1]
	}
	return op.Id
}

func (o *Op) WithScript(s *ContractScript) *Op {
	if s != nil {
		o.param, o.store, o.eps, o.bigmaps = s.Types()
	} else {
		o.param, o.store, o.eps, o.bigmaps = micheline.Type{}, micheline.Type{}, nil, nil
	}
	return o
}

func (o *Op) WithTypes(param, store micheline.Type, eps micheline.Entrypoints, b map[int64]micheline.Type) *Op {
	o.param = param
	o.store = store
	o.eps = eps
	o.bigmaps = b
	return o
}

func (o *Op) WithPrim(b bool) *Op {
	o.withPrim = b
	return o
}

func (o *Op) WithMeta(b bool) *Op {
	o.withMeta = b
	return o
}

func (o *Op) OnError(action int) *Op {
	o.onError = action
	return o
}

func (o Op) Costs() Costs {
	storageBurn := float64(o.StoragePaid) * 0.000250
	return Costs{
		Fee:            o.Fee,
		Burn:           o.Burned,
		GasUsed:        o.GasUsed,
		StorageUsed:    o.StoragePaid,
		StorageBurn:    storageBurn,
		AllocationBurn: o.Burned - storageBurn,
	}
}

func (o Op) DecodeParams() (*ContractParameters, error) {
	if o.Parameters == nil {
		return nil, ErrNoParams
	}
	switch o.Parameters[0] {
	case '"':
		buf, err := hex.DecodeString(string(o.Parameters[1 : len(o.Parameters)-1]))
		if err != nil && !o.noFail {
			return nil, err
		}
		params := &micheline.Parameters{}
		err = params.UnmarshalBinary(buf)
		if err != nil && !o.noFail {
			return nil, err
		}
		if o.param.IsValid() {
			ep, prim, _ := params.MapEntrypoint(o.param)
			cp := &ContractParameters{
				Entrypoint: ep.Name,
			}
			cp.ContractValue.Prim = &prim
			// strip entrypoint name annot
			typ := ep.Type()
			typ.Prim.Anno = nil
			val := micheline.NewValue(typ, prim)
			val.Render = o.onError
			cp.ContractValue.Value, err = val.Map()
			if err != nil && !o.noFail {
				return nil, fmt.Errorf("op %s (%d) decoding params %s: %v", o.Hash, o.Id, string(o.Parameters), err)
			}
			return cp, nil
		} else {
			return nil, ErrNoType
		}
	case '{':
		cp := &ContractParameters{}
		err := json.Unmarshal(o.Parameters, cp)
		return cp, err
	}
	return nil, ErrNoParams
}

func (o Op) DecodeStoragePrim() (prim micheline.Prim, err error) {
	if o.Storage == nil {
		err = ErrNoStorage
		return
	}
	switch o.Storage[0] {
	case '"':
		var buf []byte
		buf, err = hex.DecodeString(string(o.Storage[1 : len(o.Storage)-1]))
		if err != nil && !o.noFail {
			return
		}
		err = prim.UnmarshalBinary(buf)
		if err != nil && !o.noFail {
			return
		}
		err = nil
		return
	default:
		cv := &ContractValue{}
		err := json.Unmarshal(o.Storage, cv)
		return *cv.Prim, err
	}
}

func (o Op) DecodeStorage() (*ContractValue, error) {
	if o.Storage == nil {
		return nil, ErrNoStorage
	}
	switch o.Storage[0] {
	case '"':
		buf, err := hex.DecodeString(string(o.Storage[1 : len(o.Storage)-1]))
		if err != nil && !o.noFail {
			return nil, err
		}
		var prim micheline.Prim
		err = prim.UnmarshalBinary(buf)
		if err != nil && !o.noFail {
			return nil, err
		}
		cv := &ContractValue{
			Prim: &prim,
		}
		if o.store.IsValid() {
			val := micheline.NewValue(o.store, prim)
			val.Render = o.onError
			cv.Value, err = val.Map()
			if err != nil && !o.noFail {
				return nil, fmt.Errorf("op %s (%d) decoding storage %s: %v", o.Hash, o.Id, string(o.Storage), err)
			}
		} else {
			return nil, ErrNoType
		}
		return cv, nil
	default:
		cv := &ContractValue{}
		err := json.Unmarshal(o.Storage, cv)
		return cv, err
	}
}

func (o Op) DecodeBigmapEvents() (micheline.BigmapEvents, error) {
	if o.BigmapDiff == nil {
		return nil, ErrNoBigmapDiff
	}
	switch o.BigmapDiff[0] {
	case '"':
		// hex encoded low-level events
		buf, err := hex.DecodeString(string(o.BigmapDiff[1 : len(o.BigmapDiff)-1]))
		if err != nil && !o.noFail {
			return nil, err
		}
		events := make(micheline.BigmapEvents, 0)
		err = events.UnmarshalBinary(buf)
		if err != nil && !o.noFail {
			return nil, err
		}
		return events, nil
	default:
		// json encoded high-level updates
		updates, err := o.DecodeBigmapUpdates()
		if err != nil {
			return nil, err
		}
		return updates.Events(), nil
	}
}

func (o Op) DecodeBigmapUpdates() (BigmapUpdates, error) {
	if o.BigmapDiff == nil {
		return nil, ErrNoBigmapDiff
	}
	switch o.BigmapDiff[0] {
	case '"':
		events, err := o.DecodeBigmapEvents()
		if err != nil {
			return nil, err
		}
		updates := make(BigmapUpdates, 0, len(events))
		if o.withPrim {
			// decode prim only
			for _, v := range events {
				upd := BigmapUpdate{
					Action:   v.Action,
					BigmapId: v.Id,
				}
				switch v.Action {
				case micheline.DiffActionAlloc, micheline.DiffActionCopy:
					kt, vt := v.KeyType.Clone(), v.ValueType.Clone()
					upd.KeyTypePrim = &kt
					upd.ValueTypePrim = &vt
				case micheline.DiffActionUpdate:
					key, val := v.Key.Clone(), v.Value.Clone()
					upd.KeyPrim, upd.ValuePrim = &key, &val
				case micheline.DiffActionRemove:
					key := v.Key.Clone()
					upd.KeyPrim = &key
				}
				updates = append(updates, upd)
			}
		} else {
			// full key/value unpack, requires script type
			for _, v := range events {
				var ktyp, vtyp micheline.Type
				if typ, ok := o.bigmaps[v.Id]; ok {
					ktyp, vtyp = typ.Left(), typ.Right()
				} else {
					ktyp = v.Key.BuildType()
				}
				upd := BigmapUpdate{
					Action:   v.Action,
					BigmapId: v.Id,
				}
				switch v.Action {
				case micheline.DiffActionAlloc, micheline.DiffActionCopy:
					// alloc/copy only
					upd.KeyType = micheline.Type{Prim: v.KeyType}.TypedefPtr("@key")
					upd.ValueType = micheline.Type{Prim: v.ValueType}.TypedefPtr("@value")
					upd.SourceId = v.SourceId
					upd.DestId = v.DestId
				default:
					// update/remove only
					if !v.Key.IsEmptyBigmap() {
						keybuf, _ := v.GetKey(ktyp).MarshalJSON()
						mk := MultiKey{}
						_ = mk.UnmarshalJSON(keybuf)
						upd.Key = mk
						upd.Hash = v.KeyHash
					}
					if o.withMeta {
						upd.Meta = &BigmapMeta{
							Contract:     o.Receiver,
							BigmapId:     v.Id,
							UpdateTime:   o.Timestamp,
							UpdateHeight: o.Height,
						}
					}
					if v.Action == micheline.DiffActionUpdate {
						// unpack value if type is known
						if vtyp.IsValid() {
							val := micheline.NewValue(vtyp, v.Value)
							val.Render = o.onError
							upd.Value, err = val.Map()
							if err != nil && !o.noFail {
								return nil, fmt.Errorf("op %s (%d) decoding bigmap %d/%s: %v", o.Hash, o.Id, v.Id, v.KeyHash, err)
							}
						}
					}
				}
				updates = append(updates, upd)
			}
		}
		return updates, nil
	default:
		// json encoded high-level updates
		var bmu BigmapUpdates
		err := json.Unmarshal(o.BigmapDiff, &bmu)
		return bmu, err
	}
}

type OpGroup []*Op

func (og OpGroup) Costs() Costs {
	var c Costs
	for _, v := range og {
		c = c.Add(v.Costs())
	}
	return c
}

type OpList struct {
	Rows     []*Op
	withPrim bool
	noFail   bool
	columns  []string
	ctx      context.Context
	client   *Client
}

func (l OpList) Len() int {
	return len(l.Rows)
}

func (l OpList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	// on table API only row_id is set
	return l.Rows[len(l.Rows)-1].Id
}

func (l *OpList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("OpList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

func (l *OpList) ResolveTypes(ctx context.Context) error {
	for _, op := range l.Rows {
		if !op.IsContract || !op.Receiver.IsContract() {
			continue
		}
		// load contract type info (required for decoding storage/param data)
		script, err := l.client.loadCachedContractScript(ctx, op.Receiver)
		if err != nil {
			return err
		}
		op.WithScript(script)
		op.noFail = l.noFail
		op.withPrim = l.withPrim
	}
	return nil
}

type OpQuery struct {
	tableQuery
	NoFail bool
}

func (q OpQuery) WithNoFail() OpQuery {
	q.NoFail = true
	return q
}

func (c *Client) NewOpQuery() OpQuery {
	return OpQuery{c.newTableQuery("op", &Op{}), false}
}

func (q OpQuery) Run(ctx context.Context) (*OpList, error) {
	result := &OpList{
		columns:  q.Columns,
		ctx:      ctx,
		client:   q.client,
		withPrim: q.Prim,
		noFail:   q.NoFail,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

type OpParams = Params[Op]

func NewOpParams() OpParams {
	return OpParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) GetOp(ctx context.Context, hash tezos.OpHash, params OpParams) (OpGroup, error) {
	o := make(OpGroup, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/op/%s", hash)).Url()
	if err := c.get(ctx, u, nil, &o); err != nil {
		return nil, err
	}
	return o, nil
}

func (c *Client) GetBlockOps(ctx context.Context, hash tezos.BlockHash, params OpParams) ([]*Op, error) {
	ops := make([]*Op, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/block/%s/operations", hash)).Url()
	if err := c.get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}
