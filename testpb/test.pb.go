package testpb

import (
	json "encoding/json"

	proto "github.com/golang/protobuf/proto"

	math "math"
)

var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type TestMessage struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	City             *string `protobuf:"bytes,2,opt,name=city" json:"city,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestMessage) Reset()         { _ = "STUB: not implemented"; return }
func (m *TestMessage) String() string { _ = "STUB: not implemented"; return "" }
func (*TestMessage) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *TestMessage) GetName() string { _ = "STUB: not implemented"; return "" }

func (m *TestMessage) GetCity() string { _ = "STUB: not implemented"; return "" }

type TestRequest struct {
	Lower            *string `protobuf:"bytes,1,req,name=lower" json:"lower,omitempty"`
	RepeatCount      *int32  `protobuf:"varint,2,opt,name=repeat_count,def=1" json:"repeat_count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestRequest) Reset()         { _ = "STUB: not implemented"; return }
func (m *TestRequest) String() string { _ = "STUB: not implemented"; return "" }
func (*TestRequest) ProtoMessage()    { _ = "STUB: not implemented"; return }

const Default_TestRequest_RepeatCount int32 = 1

func (m *TestRequest) GetLower() string { _ = "STUB: not implemented"; return "" }

func (m *TestRequest) GetRepeatCount() int32 { _ = "STUB: not implemented"; return 0 }

type TestResponse struct {
	Value            *string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestResponse) Reset()         { _ = "STUB: not implemented"; return }
func (m *TestResponse) String() string { _ = "STUB: not implemented"; return "" }
func (*TestResponse) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *TestResponse) GetValue() string { _ = "STUB: not implemented"; return "" }

type CacheStats struct {
	Items            *int64 `protobuf:"varint,1,opt,name=items" json:"items,omitempty"`
	Bytes            *int64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	Gets             *int64 `protobuf:"varint,3,opt,name=gets" json:"gets,omitempty"`
	Hits             *int64 `protobuf:"varint,4,opt,name=hits" json:"hits,omitempty"`
	Evicts           *int64 `protobuf:"varint,5,opt,name=evicts" json:"evicts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CacheStats) Reset()         { _ = "STUB: not implemented"; return }
func (m *CacheStats) String() string { _ = "STUB: not implemented"; return "" }
func (*CacheStats) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *CacheStats) GetItems() int64 { _ = "STUB: not implemented"; return 0 }

func (m *CacheStats) GetBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (m *CacheStats) GetGets() int64 { _ = "STUB: not implemented"; return 0 }

func (m *CacheStats) GetHits() int64 { _ = "STUB: not implemented"; return 0 }

func (m *CacheStats) GetEvicts() int64 { _ = "STUB: not implemented"; return 0 }

type StatsResponse struct {
	Gets             *int64      `protobuf:"varint,1,opt,name=gets" json:"gets,omitempty"`
	CacheHits        *int64      `protobuf:"varint,12,opt,name=cache_hits" json:"cache_hits,omitempty"`
	Fills            *int64      `protobuf:"varint,2,opt,name=fills" json:"fills,omitempty"`
	TotalAlloc       *uint64     `protobuf:"varint,3,opt,name=total_alloc" json:"total_alloc,omitempty"`
	MainCache        *CacheStats `protobuf:"bytes,4,opt,name=main_cache" json:"main_cache,omitempty"`
	HotCache         *CacheStats `protobuf:"bytes,5,opt,name=hot_cache" json:"hot_cache,omitempty"`
	ServerIn         *int64      `protobuf:"varint,6,opt,name=server_in" json:"server_in,omitempty"`
	Loads            *int64      `protobuf:"varint,8,opt,name=loads" json:"loads,omitempty"`
	PeerLoads        *int64      `protobuf:"varint,9,opt,name=peer_loads" json:"peer_loads,omitempty"`
	PeerErrors       *int64      `protobuf:"varint,10,opt,name=peer_errors" json:"peer_errors,omitempty"`
	LocalLoads       *int64      `protobuf:"varint,11,opt,name=local_loads" json:"local_loads,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *StatsResponse) Reset()         { _ = "STUB: not implemented"; return }
func (m *StatsResponse) String() string { _ = "STUB: not implemented"; return "" }
func (*StatsResponse) ProtoMessage()    { _ = "STUB: not implemented"; return }

func (m *StatsResponse) GetGets() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetCacheHits() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetFills() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetTotalAlloc() uint64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetMainCache() *CacheStats { _ = "STUB: not implemented"; return nil }

func (m *StatsResponse) GetHotCache() *CacheStats { _ = "STUB: not implemented"; return nil }

func (m *StatsResponse) GetServerIn() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetLoads() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetPeerLoads() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetPeerErrors() int64 { _ = "STUB: not implemented"; return 0 }

func (m *StatsResponse) GetLocalLoads() int64 { _ = "STUB: not implemented"; return 0 }

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()         { _ = "STUB: not implemented"; return }
func (m *Empty) String() string { _ = "STUB: not implemented"; return "" }
func (*Empty) ProtoMessage()    { _ = "STUB: not implemented"; return }

func init() {
}
