// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// HTTPStatus interface represents an unprocessed HTTP reply
type HTTPStatus interface {
	Request() string // e.g. GET /...
	Status() string  // e.g. "200 OK"
	StatusCode() int // e.g. 200
}

// HTTPError retains HTTP status and error response body
type HTTPError interface {
	error
	HTTPStatus
	Body() []byte
}

var (
	_ HTTPError = &ErrHttp{}
	_ HTTPError = &ErrApi{}
	_ HTTPError = &ErrRateLimited{}
)

type ErrApi struct {
	Code      int    `json:"code"`
	Status_   int    `json:"status"`
	Message   string `json:"message"`
	Scope     string `json:"scope"`
	Detail    string `json:"detail"`
	RequestId string `json:"requestId"`
	Reason    string `json:"reason"`
	Request_  string `json:"request"`
}

func (e *ErrApi) Error() string {
	s := make([]string, 0)
	if e.Status_ != 0 {
		s = append(s, "status="+strconv.Itoa(e.Status_))
	}
	if e.Code != 0 {
		s = append(s, "code="+strconv.Itoa(e.Code))
	}
	if e.Scope != "" {
		s = append(s, "scope="+e.Scope)
	}
	s = append(s, "message=\""+e.Message+"\"")
	if e.Detail != "" {
		s = append(s, "detail=\""+e.Detail+"\"")
	}
	if e.RequestId != "" {
		s = append(s, "request-id="+e.RequestId)
	}
	if e.Reason != "" {
		s = append(s, "reason=\""+e.Reason+"\"")
	}
	return strings.Join(s, " ")
}

func (e *ErrApi) UnmarshalJSON(buf []byte) error {
	if len(buf) < 2 {
		return nil
	}
	var t map[string]json.RawMessage
	if err := json.Unmarshal(buf, &t); err != nil {
		return err
	}
	// check if we have an embedded array and decode
	type alias *ErrApi
	if v, ok := t["errors"]; ok {
		var arr []alias
		if err := json.Unmarshal(v, &arr); err != nil {
			return err
		}
		if len(arr) > 0 {
			*e = ErrApi(*arr[0])
		}
		return nil
	}
	// if not, decode as single error
	return json.Unmarshal(buf, alias(e))
}

func (e *ErrApi) Request() string {
	return e.Request_
}

func (e *ErrApi) StatusCode() int {
	return e.Status_
}

func (e *ErrApi) Status() string {
	return e.Message
}

func (e *ErrApi) Body() []byte {
	buf, _ := json.Marshal(e)
	return buf
}

func IsErrApi(err error) (*ErrApi, bool) {
	e, ok := err.(*ErrApi)
	return e, ok
}

type ErrHttp struct {
	request    string
	status     string
	statusCode int
	body       []byte
	header     http.Header
}

func (e *ErrHttp) Error() string {
	return fmt.Sprintf("%d %s %s", e.statusCode, e.status, e.request)
}

func (e *ErrHttp) Request() string {
	return e.request
}

func (e *ErrHttp) Status() string {
	return e.status
}

func (e *ErrHttp) StatusCode() int {
	return e.statusCode
}

func (e *ErrHttp) Body() []byte {
	return e.body
}

func (e *ErrHttp) Decode(v interface{}) error {
	return json.Unmarshal(e.body, v)
}

func newHttpError(resp *http.Response) *ErrHttp {
	body, _ := io.ReadAll(resp.Body)
	return newHttpErrorWithBody(resp, body)
}

func newHttpErrorWithBody(resp *http.Response, body []byte) *ErrHttp {
	return &ErrHttp{
		request:    resp.Request.Method + " " + resp.Request.URL.String(),
		status:     resp.Status,
		statusCode: resp.StatusCode,
		body:       body,
		header:     mergeHeaders(make(http.Header), resp.Header, resp.Trailer),
	}
}

func IsErrHttp(err error) (HTTPError, bool) {
	e, ok := err.(HTTPError)
	return e, ok
}

type ErrRateLimited struct {
	ErrHttp
	deadline time.Time
	done     chan struct{}
}

func newRateLimitError(resp *http.Response, d time.Duration) *ErrRateLimited {
	he := newHttpError(resp)
	e := &ErrRateLimited{
		ErrHttp:  *he,
		deadline: time.Now().UTC().Add(d),
		done:     make(chan struct{}),
	}
	go e.timeout(d)
	return e
}

func (e *ErrRateLimited) timeout(d time.Duration) {
	<-time.After(d)
	close(e.done)
}

func (e *ErrRateLimited) Error() string {
	return fmt.Sprintf("%s rate limited for %s", e.ErrHttp.Error(), time.Until(e.deadline))
}

func (e *ErrRateLimited) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-e.done:
		return nil
	}
}

func (e *ErrRateLimited) Done() <-chan struct{} {
	return e.done
}

func (e *ErrRateLimited) Deadline() time.Duration {
	return e.deadline.Sub(time.Now().UTC())
}

func IsErrRateLimited(err error) (*ErrRateLimited, bool) {
	e, ok := err.(*ErrRateLimited)
	return e, ok
}

func ErrorStatus(err error) int {
	switch e := err.(type) {
	case *ErrRateLimited:
		return 427
	case *ErrHttp:
		return e.statusCode
	case *ErrApi:
		return e.StatusCode()
	default:
		return 0
	}
}

func isNetError(err error) bool {
	if err == nil {
		return false
	}
	// direct type
	switch err.(type) {
	case *net.OpError:
		return true
	case *net.DNSError:
		return true
	case *os.SyscallError:
		return true
	case *url.Error:
		return true
	}
	// wrapped
	var (
		neterr *net.OpError
		dnserr *net.DNSError
		oserr  *os.SyscallError
		urlerr *url.Error
	)
	switch {
	case errors.As(err, &neterr):
		return true
	case errors.As(err, &dnserr):
		return true
	case errors.As(err, &oserr):
		return true
	case errors.As(err, &urlerr):
		return true
	}
	return false
}
