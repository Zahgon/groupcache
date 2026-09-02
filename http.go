package groupcache

import (
	"bytes"
	"context"
	"net/http"
	"sync"

	"github.com/golang/groupcache/consistenthash"
	pb "github.com/golang/groupcache/groupcachepb"
)

const defaultBasePath = "/_groupcache/"

const defaultReplicas = 50

type HTTPPool struct {
	Context func(*http.Request) context.Context

	Transport func(context.Context) http.RoundTripper

	self string

	opts HTTPPoolOptions

	mu          sync.Mutex
	peers       *consistenthash.Map
	httpGetters map[string]*httpGetter
}

type HTTPPoolOptions struct {
	BasePath string

	Replicas int

	HashFn consistenthash.Hash
}

func NewHTTPPool(self string) *HTTPPool { _ = "STUB: not implemented"; return nil }

var httpPoolMade bool

func NewHTTPPoolOpts(self string, o *HTTPPoolOptions) *HTTPPool {
	_ = "STUB: not implemented"
	return nil
}

func (p *HTTPPool) Set(peers ...string) { _ = "STUB: not implemented"; return }

func (p *HTTPPool) PickPeer(key string) (ProtoGetter, bool) {
	_ = "STUB: not implemented"
	return *new(ProtoGetter), false
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type httpGetter struct {
	transport func(context.Context) http.RoundTripper
	baseURL   string
}

var bufferPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func (h *httpGetter) Get(ctx context.Context, in *pb.GetRequest, out *pb.GetResponse) error {
	_ = "STUB: not implemented"
	return nil
}
