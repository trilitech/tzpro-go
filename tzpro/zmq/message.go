// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

import (
	"blockwatch.cc/tzpro-go/internal/client"
)

type Message struct {
	topic  string
	body   []byte
	fields []string
}

func NewMessage(topic, body []byte) *Message {
	return &Message{string(topic), body, nil}
}

func (m *Message) DecodeOpHash() (OpHash, error) {
	return ParseOpHash(string(m.body))
}

func (m *Message) DecodeBlockHash() (BlockHash, error) {
	return ParseBlockHash(string(m.body))
}

func (m *Message) DecodeOp() (*Op, error) {
	o := new(Op)
	err := client.Decode(m.body, ZmqRawOpColumns, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (m *Message) DecodeBlock() (*Block, error) {
	b := new(Block)
	err := client.Decode(m.body, ZmqRawBlockColumns, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *Message) DecodeStatus() (*Status, error) {
	s := new(Status)
	err := client.Decode(m.body, ZmqStatusColumns, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
