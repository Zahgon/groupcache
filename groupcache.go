package groupcache

import (
	"context"
	"math/rand"
	"sync"

	"github.com/golang/groupcache/lru"
)

type Getter interface {
	Get(ctx context.Context, key string, dest Sink) error
}

type GetterFunc func(ctx context.Context, key string, dest Sink) error

func (f GetterFunc) Get(ctx context.Context, key string, dest Sink) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)

	initPeerServerOnce sync.Once
	initPeerServer     func()
)

func GetGroup(name string) *Group { _ = "STUB: not implemented"; return nil }

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	_ = "STUB: not implemented"
	return nil
}

func newGroup(name string, cacheBytes int64, getter Getter, peers PeerPicker) *Group {
	_ = "STUB: not implemented"
	return nil
}

var newGroupHook func(*Group)

func RegisterNewGroupHook(fn func(*Group)) { _ = "STUB: not implemented"; return }

func RegisterServerStart(fn func()) { _ = "STUB: not implemented"; return }

func callInitPeerServer() { _ = "STUB: not implemented"; return }

type Group struct {
	name       string
	getter     Getter
	peersOnce  sync.Once
	peers      PeerPicker
	cacheBytes int64

	mainCache cache

	hotCache cache

	loadGroup flightGroup

	_ int32

	Stats Stats

	rand *rand.Rand
}

type flightGroup interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
}

type Stats struct {
	Gets           AtomicInt
	CacheHits      AtomicInt
	PeerLoads      AtomicInt
	PeerErrors     AtomicInt
	Loads          AtomicInt
	LoadsDeduped   AtomicInt
	LocalLoads     AtomicInt
	LocalLoadErrs  AtomicInt
	ServerRequests AtomicInt
}

func (g *Group) Name() string { _ = "STUB: not implemented"; return "" }

func (g *Group) initPeers() { _ = "STUB: not implemented"; return }

func (g *Group) Get(ctx context.Context, key string, dest Sink) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *Group) load(ctx context.Context, key string, dest Sink) (value ByteView, destPopulated bool, err error) {
	_ = "STUB: not implemented"
	return *new(ByteView), false, nil
}

func (g *Group) getLocally(ctx context.Context, key string, dest Sink) (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (g *Group) getFromPeer(ctx context.Context, peer ProtoGetter, key string) (ByteView, error) {
	_ = "STUB: not implemented"
	return *new(ByteView), nil
}

func (g *Group) lookupCache(key string) (value ByteView, ok bool) {
	_ = "STUB: not implemented"
	return *new(ByteView), false
}

func (g *Group) populateCache(key string, value ByteView, cache *cache) {
	_ = "STUB: not implemented"
	return
}

type CacheType int

const (
	MainCache CacheType = iota + 1

	HotCache
)

func (g *Group) CacheStats(which CacheType) CacheStats {
	_ = "STUB: not implemented"
	return *new(CacheStats)
}

type cache struct {
	mu         sync.RWMutex
	nbytes     int64
	lru        *lru.Cache
	nhit, nget int64
	nevict     int64
}

func (c *cache) stats() CacheStats { _ = "STUB: not implemented"; return *new(CacheStats) }

func (c *cache) add(key string, value ByteView) { _ = "STUB: not implemented"; return }

func (c *cache) get(key string) (value ByteView, ok bool) {
	_ = "STUB: not implemented"
	return *new(ByteView), false
}

func (c *cache) removeOldest() { _ = "STUB: not implemented"; return }

func (c *cache) bytes() int64 { _ = "STUB: not implemented"; return 0 }

func (c *cache) items() int64 { _ = "STUB: not implemented"; return 0 }

func (c *cache) itemsLocked() int64 { _ = "STUB: not implemented"; return 0 }

type AtomicInt int64

func (i *AtomicInt) Add(n int64) { _ = "STUB: not implemented"; return }

func (i *AtomicInt) Get() int64 { _ = "STUB: not implemented"; return 0 }

func (i *AtomicInt) String() string { _ = "STUB: not implemented"; return "" }

type CacheStats struct {
	Bytes     int64
	Items     int64
	Gets      int64
	Hits      int64
	Evictions int64
}
