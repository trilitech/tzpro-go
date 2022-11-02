// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "net/http"
)

type Transport interface {
    Do(*http.Request) (*http.Response, error)
}

type defaultTransport struct {
    c *http.Client
}

func (t defaultTransport) Do(req *http.Request) (*http.Response, error) {
    return t.c.Do(req)
}
