package groupcache

import (
	"github.com/golang/protobuf/proto"
)

type Sink interface {
	SetString(s string) error

	SetBytes(v []byte) error

	SetProto(m proto.Message) error

	view() (ByteView, error)
}

func cloneBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func setSinkView(s Sink, v ByteView) error { _ = "STUB: not implemented"; return nil }

func StringSink(sp *string) Sink { _ = "STUB: not implemented"; return *new(Sink) }

type stringSink struct {
	sp *string
	v  ByteView
}

func (s *stringSink) view() (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (s *stringSink) SetString(v string) error { _ = "STUB: not implemented"; return nil }

func (s *stringSink) SetBytes(v []byte) error { _ = "STUB: not implemented"; return nil }

func (s *stringSink) SetProto(m proto.Message) error { _ = "STUB: not implemented"; return nil }

func ByteViewSink(dst *ByteView) Sink { _ = "STUB: not implemented"; return *new(Sink) }

type byteViewSink struct {
	dst *ByteView
}

func (s *byteViewSink) setView(v ByteView) error { _ = "STUB: not implemented"; return nil }

func (s *byteViewSink) view() (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (s *byteViewSink) SetProto(m proto.Message) error { _ = "STUB: not implemented"; return nil }

func (s *byteViewSink) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *byteViewSink) SetString(v string) error { _ = "STUB: not implemented"; return nil }

func ProtoSink(m proto.Message) Sink { _ = "STUB: not implemented"; return *new(Sink) }

type protoSink struct {
	dst proto.Message
	typ string

	v ByteView
}

func (s *protoSink) view() (ByteView, error) { _ = "STUB: not implemented"; return *new(ByteView), nil }

func (s *protoSink) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *protoSink) SetString(v string) error { _ = "STUB: not implemented"; return nil }

func (s *protoSink) SetProto(m proto.Message) error { _ = "STUB: not implemented"; return nil }

func AllocatingByteSliceSink(dst *[]byte) Sink { _ = "STUB: not implemented"; return *new(Sink) }

type allocBytesSink struct {
	dst *[]byte
	v   ByteView
}

func (s *allocBytesSink) view() (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (s *allocBytesSink) setView(v ByteView) error { _ = "STUB: not implemented"; return nil }

func (s *allocBytesSink) SetProto(m proto.Message) error { _ = "STUB: not implemented"; return nil }

func (s *allocBytesSink) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *allocBytesSink) setBytesOwned(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *allocBytesSink) SetString(v string) error { _ = "STUB: not implemented"; return nil }

func TruncatingByteSliceSink(dst *[]byte) Sink { _ = "STUB: not implemented"; return *new(Sink) }

type truncBytesSink struct {
	dst *[]byte
	v   ByteView
}

func (s *truncBytesSink) view() (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (s *truncBytesSink) SetProto(m proto.Message) error { _ = "STUB: not implemented"; return nil }

func (s *truncBytesSink) SetBytes(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *truncBytesSink) setBytesOwned(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *truncBytesSink) SetString(v string) error { _ = "STUB: not implemented"; return nil }
