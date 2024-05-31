// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/trilitech/tzpro-go/internal/util"
)

type Query struct {
	Server string
	Path   string
	Query  url.Values
}

func NewQuery() Query {
	return Query{
		Query: url.Values{},
	}
}

// parse from
// http://server:port/prefix
// server:port/prefix
// server/prefix
// /prefix
func ParseQuery(s string) (Query, error) {
	if !strings.HasPrefix(s, "http") {
		s = "https://" + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return NewQuery(), err
	}
	p := Query{
		Query: u.Query(),
	}
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

func (p Query) Clone() Query {
	np := Query{
		Query: url.Values{},
	}
	np.Server = p.Server
	np.Path = p.Path
	for n, v := range p.Query {
		np.Query[n] = v
	}
	return np
}

func (p Query) AndArg(key string, values ...any) Query {
	val := util.ToString(values)
	if val != "" {
		p.Query.Set(key, val)
	} else {
		p.Query.Del(key)
	}
	return p
}

func (p Query) AndFilter(key string, mode FilterMode, values ...any) Query {
	p.Query.Set(key+"."+string(mode), util.ToString(values))
	return p
}

func (p Query) AndEqual(key string, val any) Query {
	p.Query.Set(key+".eq", util.ToString(val))
	return p
}

func (p Query) AndNotEqual(key string, val any) Query {
	p.Query.Set(key+".ne", util.ToString(val))
	return p
}

func (p Query) AndGt(key string, val any) Query {
	p.Query.Set(key+".gt", util.ToString(val))
	return p
}

func (p Query) AndGte(key string, val any) Query {
	p.Query.Set(key+".gte", util.ToString(val))
	return p
}

func (p Query) AndLt(key string, val any) Query {
	p.Query.Set(key+".lt", util.ToString(val))
	return p
}

func (p Query) AndLte(key string, val any) Query {
	p.Query.Set(key+".lte", util.ToString(val))
	return p
}

func (p Query) AndIn(key string, val ...any) Query {
	p.Query.Set(key+".in", util.ToString(val))
	return p
}

func (p Query) AndNotIn(key string, val ...any) Query {
	p.Query.Set(key+".nin", util.ToString(val))
	return p
}

func (p Query) AndRange(key string, from, to any) Query {
	p.Query.Set(key+".rg", util.ToString([]any{from, to}))
	return p
}

func (p Query) AndRegexp(key string, re string) Query {
	p.Query.Set(key+".re", re)
	return p
}

func (p Query) And(key string) Query {
	p.Query.Set(key, "1")
	return p
}

func (p Query) AndNot(key string) Query {
	p.Query.Set(key, "0")
	return p
}

func (p Query) WithLimit(v uint) Query {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p Query) WithOffset(v uint) Query {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p Query) WithCursor(v uint64) Query {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p Query) WithOrder(o OrderType) Query {
	p.Query.Set("order", string(o))
	return p
}

func (p Query) Desc() Query {
	p.Query.Set("order", "desc")
	return p
}

func (p Query) Asc() Query {
	p.Query.Set("order", "asc")
	return p
}

func (p Query) WithMeta() Query {
	p.Query.Set("meta", "1")
	return p
}

func (p Query) WithTags(t ...string) Query {
	p.Query.Set("tags", strings.Join(t, ","))
	return p
}

func (p Query) WithFrom(t time.Time) Query {
	p.Query.Set("from", t.Format("2006-01-02"))
	return p
}

func (p Query) WithTo(t time.Time) Query {
	p.Query.Set("to", t.Format("2006-01-02"))
	return p
}

func (p Query) WithTimeRange(from, to time.Time) Query {
	p.Query.Set("from", from.Format("2006-01-02"))
	p.Query.Set("to", to.Format("2006-01-02"))
	return p
}

func (p Query) WithPrim() Query {
	p.Query.Set("prim", "1")
	return p
}

func (p Query) WithMerge() Query {
	p.Query.Set("merge", "1")
	return p
}

func (p Query) WithStorage() Query {
	p.Query.Set("storage", "1")
	return p
}

func (p Query) WithUnpack() Query {
	p.Query.Set("unpack", "1")
	return p
}

func (p Query) WithRights() Query {
	p.Query.Set("rights", "1")
	return p
}

func (p Query) WithPath(path string) Query {
	p.Path = path
	return p
}

func (p Query) Url() string {
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

func (p Query) Check() error {
	if p.Server == "" {
		return fmt.Errorf("empty server URL")
	}
	return nil
}
