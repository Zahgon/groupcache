package lru

import "container/list"

type Cache struct {
	MaxEntries int

	OnEvicted func(key Key, value interface{})

	ll    *list.List
	cache map[interface{}]*list.Element
}

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

func New(maxEntries int) *Cache { _ = "STUB: not implemented"; return nil }

func (c *Cache) Add(key Key, value interface{}) { _ = "STUB: not implemented"; return }

func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Cache) Remove(key Key) { _ = "STUB: not implemented"; return }

func (c *Cache) RemoveOldest() { _ = "STUB: not implemented"; return }

func (c *Cache) removeElement(e *list.Element) { _ = "STUB: not implemented"; return }

func (c *Cache) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Cache) Clear() { _ = "STUB: not implemented"; return }
