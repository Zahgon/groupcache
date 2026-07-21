package groupcache

import (
	"context"

	pb "github.com/golang/groupcache/groupcachepb"
)

type Context = context.Context

type ProtoGetter interface {
	Get(ctx context.Context, in *pb.GetRequest, out *pb.GetResponse) error
}

type PeerPicker interface {
	PickPeer(key string) (peer ProtoGetter, ok bool)
}

type NoPeers struct{}

func (NoPeers) PickPeer(key string) (peer ProtoGetter, ok bool) {
	_ = "STUB: not implemented"
	return *new(ProtoGetter), false
}

var (
	portPicker func(groupName string) PeerPicker
)

func RegisterPeerPicker(fn func() PeerPicker) { _ = "STUB: not implemented"; return }

func RegisterPerGroupPeerPicker(fn func(groupName string) PeerPicker) {
	_ = "STUB: not implemented"
	return
}

func getPeers(groupName string) PeerPicker { _ = "STUB: not implemented"; return *new(PeerPicker) }
