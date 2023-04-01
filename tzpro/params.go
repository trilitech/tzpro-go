// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Params[T any] struct {
	Server string
	Path   string
	Query  url.Values
}

func (p Params[T]) Clone() Params[T] {
	np := Params[T]{
		Query: url.Values{},
	}
	np.Server = p.Server
	np.Path = p.Path
	for n, v := range p.Query {
		np.Query[n] = v
	}
	return np
}

func (p Params[T]) WithArg(key string, values ...any) Params[T] {
	val := toString(values)
	if val != "" {
		p.Query.Set(key, val)
	} else {
		p.Query.Del(key)
	}
	return p
}

func (p Params[T]) WithFilter(key string, mode FilterMode, values ...any) Params[T] {
	p.Query.Set(key+"."+string(mode), toString(values))
	return p
}

func (p Params[T]) With(key string) Params[T] {
	p.Query.Set(key, "1")
	return p
}

func (p Params[T]) WithLimit(v uint) Params[T] {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p Params[T]) WithOffset(v uint) Params[T] {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p Params[T]) WithCursor(v uint64) Params[T] {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p Params[T]) WithOrder(o OrderType) Params[T] {
	p.Query.Set("order", string(o))
	return p
}

func (p Params[T]) WithDesc() Params[T] {
	p.Query.Set("order", string(OrderDesc))
	return p
}

func (p Params[T]) WithAsc() Params[T] {
	p.Query.Set("order", string(OrderAsc))
	return p
}

func (p Params[T]) WithMeta() Params[T] {
	p.Query.Set("meta", "1")
	return p
}

func (p Params[T]) WithTags(t ...string) Params[T] {
	p.Query.Set("tags", strings.Join(t, ","))
	return p
}

func (p Params[T]) WithFrom(t time.Time) Params[T] {
	p.Query.Set("from", t.Format("2006-01-02"))
	return p
}

func (p Params[T]) WithTo(t time.Time) Params[T] {
	p.Query.Set("to", t.Format("2006-01-02"))
	return p
}

func (p Params[T]) WithTimeRange(from, to time.Time) Params[T] {
	p.Query.Set("from", from.Format("2006-01-02"))
	p.Query.Set("to", to.Format("2006-01-02"))
	return p
}

func (p Params[T]) WithPrim() Params[T] {
	p.Query.Set("prim", "1")
	return p
}

func (p Params[T]) WithMerge() Params[T] {
	p.Query.Set("merge", "1")
	return p
}

func (p Params[T]) WithStorage() Params[T] {
	p.Query.Set("storage", "1")
	return p
}

func (p Params[T]) WithUnpack() Params[T] {
	p.Query.Set("unpack", "1")
	return p
}

func (p Params[T]) WithRights() Params[T] {
	p.Query.Set("rights", "1")
	return p
}

func (p Params[T]) WithPath(path string) Params[T] {
	p.Path = path
	return p
}

func (p Params[T]) Url() string {
	var b strings.Builder
	b.WriteString(p.Server)
	if p.Path != "" {
		b.WriteByte('/')
		b.WriteString(strings.TrimLeft(p.Path, "/"))
	}
	if len(p.Query) > 0 {
		b.WriteByte('?')
		b.WriteString(p.Query.Encode())
	}
	return b.String()
}

func (p Params[T]) Check() error {
	if p.Server == "" {
		return fmt.Errorf("empty server URL")
	}
	return nil
}

// parse from
// http://server:port/prefix
// server:port/prefix
// server/prefix
// /prefix
func (p Params[T]) Parse(s string) (Params[T], error) {
	if !strings.HasPrefix(s, "http") {
		s = "https://" + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return p, err
	}
	p.Query = u.Query()
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	if u.Path != "" {
		p.Path = strings.Trim(u.Path, "/")
	}
	u.RawQuery = ""
	u.Fragment = ""
	u.Path = ""
	u.RawPath = ""
	p.Server = u.String()
	return p, nil
}

type base struct{}
type BaseParams = Params[base]

func NewBaseParams() BaseParams {
	return BaseParams{
		Query: url.Values{},
	}
}
