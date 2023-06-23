// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package client

import (
	"bytes"
	"context"
	"fmt"
	// "io"
	// "net/http"
	"reflect"
	"strconv"
	"strings"

	"blockwatch.cc/tzpro-go/internal/util"
)

type Filter struct {
	Mode   FilterMode
	Column string
	Value  any
}

type FilterList []Filter

func (l *FilterList) Add(mode FilterMode, col string, val ...any) {
	*l = append(*l, Filter{
		Mode:   mode,
		Column: col,
		Value:  util.ToString(val),
	})
}

type FillMode string

type FilterMode string

type OrderType string

type FormatType string

// type TableQueryAPI interface {
// 	WithFilter(mode FilterMode, col string, val ...any) TableQueryAPI
// 	ReplaceFilter(mode FilterMode, col string, val ...any) TableQueryAPI
// 	ResetFilter() TableQueryAPI
// 	WithLimit(limit int) TableQueryAPI
// 	WithColumns(cols ...string) TableQueryAPI
// 	WithOrder(order OrderType) TableQueryAPI
// 	WithDesc() TableQueryAPI
// 	WithVerbose() TableQueryAPI
// 	WithQuiet() TableQueryAPI
// 	WithFormat(format FormatType) TableQueryAPI
// 	WithPrim() TableQueryAPI
// 	GetColumns() []string
// 	Check() error
// 	Url() string
// }

type TableQuery[T any] struct {
	Params  Params
	Table   string     // "op", "block", "chain", "flow"
	Format  FormatType // "json", "csv"
	Columns []string
	Limit   int
	Cursor  uint64
	Verbose bool
	Prim    bool
	NoFail  bool
	Filter  FilterList
	Order   OrderType // asc, desc
	// OrderBy string // column name
	// Sort string // asc/desc
	client *Client
}

func NewTableQuery[T any](c *Client, name string) *TableQuery[T] {
	var t T
	tinfo, err := getTypeInfo(t)
	if err != nil {
		panic(err)
	}
	return &TableQuery[T]{
		Params:  c.base.Clone(),
		Table:   name,
		Format:  "json",
		Limit:   DefaultLimit,
		Order:   "asc",
		Columns: tinfo.FilteredAliases(fieldFlagIgnore),
		Filter:  make(FilterList, 0),
		client:  c,
	}
}

func (q *TableQuery[T]) GetColumns() []string {
	return q.Columns
}

func (q *TableQuery[T]) WithFilter(mode FilterMode, col string, val ...any) *TableQuery[T] {
	q.Filter.Add(mode, col, val)
	return q
}

func (q *TableQuery[T]) WithEqual(col string, val any) *TableQuery[T] {
	q.Filter.Add("eq", col, val)
	return q
}

func (q *TableQuery[T]) WithNotEqual(col string, val any) *TableQuery[T] {
	q.Filter.Add("ne", col, val)
	return q
}

func (q *TableQuery[T]) WithGt(col string, val any) *TableQuery[T] {
	q.Filter.Add("gt", col, val)
	return q
}

func (q *TableQuery[T]) WithGte(col string, val any) *TableQuery[T] {
	q.Filter.Add("gte", col, val)
	return q
}

func (q *TableQuery[T]) WithLt(col string, val any) *TableQuery[T] {
	q.Filter.Add("lt", col, val)
	return q
}

func (q *TableQuery[T]) WithLte(col string, val any) *TableQuery[T] {
	q.Filter.Add("lte", col, val)
	return q
}

func (q *TableQuery[T]) WithIn(col string, val ...any) *TableQuery[T] {
	q.Filter.Add("in", col, val)
	return q
}

func (q *TableQuery[T]) WithNotIn(col string, val ...any) *TableQuery[T] {
	q.Filter.Add("nin", col, val)
	return q
}

func (q *TableQuery[T]) WithRange(col string, from, to any) *TableQuery[T] {
	q.Filter.Add("rg", col, from, to)
	return q
}

func (q *TableQuery[T]) WithRegexp(col string, re string) *TableQuery[T] {
	q.Filter.Add("re", col, re)
	return q
}

func (q *TableQuery[T]) ReplaceFilter(mode FilterMode, col string, val ...any) *TableQuery[T] {
	for i, v := range q.Filter {
		if v.Column == col {
			q.Filter[i].Mode = mode
			q.Filter[i].Value = util.ToString(val)
			return q
		}
	}
	q.Filter.Add(mode, col, val)
	return q
}

func (q *TableQuery[T]) ResetFilter() *TableQuery[T] {
	q.Filter = q.Filter[:0]
	return q
}

func (q *TableQuery[T]) WithLimit(limit int) *TableQuery[T] {
	q.Limit = limit
	return q
}

func (q *TableQuery[T]) WithColumns(cols ...string) *TableQuery[T] {
	q.Columns = cols
	return q
}

func (q *TableQuery[T]) WithOrder(order OrderType) *TableQuery[T] {
	q.Order = order
	return q
}

func (q *TableQuery[T]) WithDesc() *TableQuery[T] {
	q.Order = "desc"
	return q
}

func (q *TableQuery[T]) WithVerbose() *TableQuery[T] {
	q.Verbose = true
	return q
}

func (q *TableQuery[T]) WithQuiet() *TableQuery[T] {
	q.Verbose = false
	return q
}

func (q *TableQuery[T]) WithFormat(format FormatType) *TableQuery[T] {
	q.Format = format
	return q
}

func (q *TableQuery[T]) WithPrim() *TableQuery[T] {
	q.Prim = true
	return q
}

func (q *TableQuery[T]) WithNoFail() *TableQuery[T] {
	q.NoFail = true
	return q
}

func (q *TableQuery[T]) WithCursor(c uint64) *TableQuery[T] {
	q.Cursor = c
	return q
}

func (p TableQuery[T]) Check() error {
	if err := p.Params.Check(); err != nil {
		return err
	}
	if p.Table == "" {
		return fmt.Errorf("empty table name")
	}
	for _, v := range p.Filter {
		if v.Column == "" {
			return fmt.Errorf("empty filter column name")
		}
		if v.Mode == "" {
			return fmt.Errorf("invalid filter mode for filter column '%s'", v.Column)
		}
		if v.Value == nil {
			return fmt.Errorf("empty value for filter column '%s'", v.Column)
		}
	}
	switch p.Format {
	case "json", "csv", "":
		// OK
	default:
		return fmt.Errorf("unsupported format '%s'", p.Format)
	}
	return nil
}

func (p TableQuery[T]) Url() string {
	base := p.Params.Clone()
	if p.Cursor > 0 {
		base.Query.Set("cursor", strconv.FormatUint(p.Cursor, 10))
	}
	if p.Limit > 0 && base.Query.Get("limit") == "" {
		base.Query.Set("limit", strconv.Itoa(p.Limit))
	}
	if len(p.Columns) > 0 && base.Query.Get("columns") == "" {
		base.Query.Set("columns", strings.Join(p.Columns, ","))
	}
	if p.Verbose {
		base.Query.Set("verbose", "true")
	}
	for _, v := range p.Filter {
		base.Query.Set(v.Column+"."+string(v.Mode), util.ToString(v.Value))
	}
	base.Query.Set("order", string(p.Order))
	format := p.Format
	if format == "" {
		format = "json"
	}
	return base.WithPath("tables/" + p.Table + "." + string(format)).Url()
}

func (q TableQuery[T]) Run(ctx context.Context) (*TableQueryResult[T], error) {
	if err := q.Check(); err != nil {
		return nil, err
	}
	res := NewTableQueryResult[T](q.Columns)
	if err := q.client.Get(ctx, q.Url(), nil, res); err != nil {
		return nil, err
	}
	return res, nil
}

type TableQueryResult[T any] struct {
	rows    []T
	columns []string
}

func NewTableQueryResult[T any](cols []string) *TableQueryResult[T] {
	return &TableQueryResult[T]{
		rows:    make([]T, 0),
		columns: cols,
	}
}

func (r *TableQueryResult[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, []byte(`null`)) {
		return nil
	}
	if data[0] != '[' {
		var t T
		return fmt.Errorf("%T: expected JSON array", t)
	}
	return DecodeSlice(data, r.columns, &r.rows)
}

func (r *TableQueryResult[T]) Rows() []T {
	return r.rows
}

func (r *TableQueryResult[T]) Len() int {
	return len(r.rows)
}

func (r *TableQueryResult[T]) Cursor() uint64 {
	if len(r.rows) == 0 {
		return 0
	}
	tinfo, _ := getTypeInfo(r.rows[0])
	if !tinfo.Fields[0].ContainsFlag(fieldFlagUint64) {
		return 0
	}
	val := derefValue(reflect.ValueOf(r.rows[0]))
	return val.Field(tinfo.Fields[0].Idx[0]).Uint()
}

// func (c *Client) StreamTable(ctx context.Context, q TableQueryAPI, w io.Writer) (StreamResponse, error) {
// 	if err := q.Check(); err != nil {
// 		return StreamResponse{}, err
// 	}
// 	// call with a non-nil header to indicate we expect response headers and trailers
// 	headers := make(http.Header)
// 	// signal upstream we accept trailers (required for some proxies to forward)
// 	headers.Add("TE", "trailers")
// 	if err := c.Get(ctx, q.Url(), headers, w); err != nil {
// 		return StreamResponse{}, err
// 	}
// 	return NewStreamResponse(headers)
// }

// func getTableColumn(data []byte, columns []string, name string) (string, bool) {
// 	idx := colIndex(columns, name)
// 	if idx < 0 || len(data) < 2 {
// 		return "", false
// 	}

// 	var (
// 		skipJson int = -1
// 		skip     bool
// 		escape   bool
// 		field    int
// 		res      []byte
// 	)
// 	for _, v := range data {
// 		if field > idx {
// 			break
// 		}
// 		if field == idx && v != ',' {
// 			res = append(res, v)
// 		}
// 		if escape {
// 			escape = false
// 			continue
// 		}
// 		switch v {
// 		case '[', '{':
// 			skipJson++
// 		case ']', '}':
// 			skipJson--
// 		}
// 		switch v {
// 		case '\\':
// 			escape = true
// 		case '"':
// 			skip = !skip
// 		case ',':
// 			if !skip && skipJson == 0 {
// 				field++
// 			}
// 		}
// 	}
// 	return strings.Trim(strings.Trim(string(res), `]`), `"`), true
// }

// func colIndex(columns []string, name string) int {
// 	for i, v := range columns {
// 		if v != name {
// 			continue
// 		}
// 		return i
// 	}
// 	return -1
// }
