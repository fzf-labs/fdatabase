package encoding

import (
	"encoding/json"

	"github.com/bytedance/sonic"
	"github.com/vmihailenco/msgpack/v5"
)

type API interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func NewJSON() *JSON {
	return &JSON{}
}

type JSON struct{}

func (e *JSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (e *JSON) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewSonic() *Sonic {
	return &Sonic{}
}

type Sonic struct{}

func (e *Sonic) Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}

func (e *Sonic) Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func NewMsgPack() *MsgPack {
	return &MsgPack{}
}

type MsgPack struct{}

func (e *MsgPack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (e *MsgPack) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
