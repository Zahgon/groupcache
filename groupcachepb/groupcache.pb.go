package groupcachepb

import (
	json "encoding/json"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type GetRequest struct {
	Group            *string `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	Key              *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetRequest) Reset()         { _ = "STUB: not implemented"; return }
func (m *GetRequest) String() string { _ = "STUB: not implemented"; return "" }
func (*GetRequest) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *GetRequest) GetGroup() string { _ = "STUB: not implemented"; return "" }

func (m *GetRequest) GetKey() string { _ = "STUB: not implemented"; return "" }

type GetResponse struct {
	Value            []byte   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	MinuteQps        *float64 `protobuf:"fixed64,2,opt,name=minute_qps" json:"minute_qps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetResponse) Reset()         { _ = "STUB: not implemented"; return }
func (m *GetResponse) String() string { _ = "STUB: not implemented"; return "" }
func (*GetResponse) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *GetResponse) GetValue() []byte { _ = "STUB: not implemented"; return nil }

func (m *GetResponse) GetMinuteQps() float64 { _ = "STUB: not implemented"; return 0 }

func init() {
}
