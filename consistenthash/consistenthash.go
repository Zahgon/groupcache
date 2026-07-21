package consistenthash

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map { _ = "STUB: not implemented"; return nil }

func (m *Map) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m *Map) Add(keys ...string) { _ = "STUB: not implemented"; return }

func (m *Map) Get(key string) string { _ = "STUB: not implemented"; return "" }
