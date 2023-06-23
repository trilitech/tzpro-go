// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"blockwatch.cc/tzpro-go/internal/util"
)

type Params struct {
	Server string
	Path   string
	Query  url.Values
}

func NewParams() Params {
	return Params{
		Query: url.Values{},
	}
}

// parse from
// http://server:port/prefix
// server:port/prefix
// server/prefix
// /prefix
func (p Params) Parse(s string) (Params, error) {
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

func (p Params) Clone() Params {
	np := Params{
		Query: url.Values{},
	}
	np.Server = p.Server
	np.Path = p.Path
	for n, v := range p.Query {
		np.Query[n] = v
	}
	return np
}

func (p Params) WithArg(key string, values ...any) Params {
	val := util.ToString(values)
	if val != "" {
		p.Query.Set(key, val)
	} else {
		p.Query.Del(key)
	}
	return p
}

func (p Params) WithFilter(key string, mode FilterMode, values ...any) Params {
	p.Query.Set(key+"."+string(mode), util.ToString(values))
	return p
}

func (p Params) WithEqual(key string, val any) Params {
	p.Query.Set(key+".eq", util.ToString(val))
	return p
}

func (p Params) WithNotEqual(key string, val any) Params {
	p.Query.Set(key+".ne", util.ToString(val))
	return p
}

func (p Params) WithGt(key string, val any) Params {
	p.Query.Set(key+".gt", util.ToString(val))
	return p
}

func (p Params) WithGte(key string, val any) Params {
	p.Query.Set(key+".gte", util.ToString(val))
	return p
}

func (p Params) WithLt(key string, val any) Params {
	p.Query.Set(key+".lt", util.ToString(val))
	return p
}

func (p Params) WithLte(key string, val any) Params {
	p.Query.Set(key+".lte", util.ToString(val))
	return p
}

func (p Params) WithIn(key string, val ...any) Params {
	p.Query.Set(key+".in", util.ToString(val))
	return p
}

func (p Params) WithNotIn(key string, val ...any) Params {
	p.Query.Set(key+".nin", util.ToString(val))
	return p
}

func (p Params) WithRange(key string, from, to any) Params {
	p.Query.Set(key+".rg", util.ToString([]any{from, to}))
	return p
}

func (p Params) WithRegexp(key string, re string) Params {
	p.Query.Set(key+".re", re)
	return p
}

func (p Params) With(key string) Params {
	p.Query.Set(key, "1")
	return p
}

func (p Params) WithLimit(v uint) Params {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p Params) WithOffset(v uint) Params {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p Params) WithCursor(v uint64) Params {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p Params) WithOrder(o OrderType) Params {
	p.Query.Set("order", string(o))
	return p
}

func (p Params) WithDesc() Params {
	p.Query.Set("order", "desc")
	return p
}

func (p Params) WithAsc() Params {
	p.Query.Set("order", "asc")
	return p
}

func (p Params) WithMeta() Params {
	p.Query.Set("meta", "1")
	return p
}

func (p Params) WithTags(t ...string) Params {
	p.Query.Set("tags", strings.Join(t, ","))
	return p
}

func (p Params) WithFrom(t time.Time) Params {
	p.Query.Set("from", t.Format("2006-01-02"))
	return p
}

func (p Params) WithTo(t time.Time) Params {
	p.Query.Set("to", t.Format("2006-01-02"))
	return p
}

func (p Params) WithTimeRange(from, to time.Time) Params {
	p.Query.Set("from", from.Format("2006-01-02"))
	p.Query.Set("to", to.Format("2006-01-02"))
	return p
}

func (p Params) WithPrim() Params {
	p.Query.Set("prim", "1")
	return p
}

func (p Params) WithMerge() Params {
	p.Query.Set("merge", "1")
	return p
}

func (p Params) WithStorage() Params {
	p.Query.Set("storage", "1")
	return p
}

func (p Params) WithUnpack() Params {
	p.Query.Set("unpack", "1")
	return p
}

func (p Params) WithRights() Params {
	p.Query.Set("rights", "1")
	return p
}

func (p Params) WithPath(path string) Params {
	p.Path = path
	return p
}

func (p Params) Url() string {
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

func (p Params) Check() error {
	if p.Server == "" {
		return fmt.Errorf("empty server URL")
	}
	return nil
}
