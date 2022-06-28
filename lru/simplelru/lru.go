package simplelru

import (
	"errors"

	"github.com/ImSingee/go-ex/linkedlist"
)

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback[K comparable, V any] func(key K, value V)

// LRU implements a non-thread safe fixed size LRU cache
type LRU[K comparable, V any] struct {
	size      int
	evictList *linkedlist.List[*entry[K, V]]
	items     map[K]*linkedlist.Element[*entry[K, V]]
	onEvict   EvictCallback[K, V]
}

// Entry is used to hold a key-value pair
type Entry[K, V any] struct {
	Key   K
	Value V
}

// entry is used to hold a value in the evictList
type entry[K comparable, V any] struct {
	key   K
	value V
}

// NewLRU constructs an LRU of the given size
func NewLRU[K comparable, V any](size int, onEvict EvictCallback[K, V]) (*LRU[K, V], error) {
	if size <= 0 {
		return nil, errors.New("must provide a positive size")
	}
	c := &LRU[K, V]{
		size:      size,
		evictList: linkedlist.New[*entry[K, V]](),
		items:     make(map[K]*linkedlist.Element[*entry[K, V]]),
		onEvict:   onEvict,
	}
	return c, nil
}

// Purge is used to completely clear the cache.
func (c *LRU[K, V]) Purge() {
	for k, v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.Value.value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *LRU[K, V]) Add(key K, value V) (evicted bool) {
	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.value = value
		return false
	}

	// Add new item
	ent := &entry[K, V]{key, value}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size
	// Verify size not exceeded
	if evict {
		c.removeOldest()
	}
	return evict
}

// Get looks up a key's value from the cache.
func (c *LRU[K, V]) Get(key K) (value V, ok bool) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		if ent.Value == nil {
			return c.zeroValue(), false
		}
		return ent.Value.value, true
	}
	return
}

// Contains checks if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *LRU[K, V]) Contains(key K) (ok bool) {
	_, ok = c.items[key]
	return ok
}

// Peek returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *LRU[K, V]) Peek(key K) (value V, ok bool) {
	var ent *linkedlist.Element[*entry[K, V]]
	if ent, ok = c.items[key]; ok {
		return ent.Value.value, true
	}
	return c.zeroValue(), ok
}

// Remove removes the provided key from the cache, returning if the
// key was contained.
func (c *LRU[K, V]) Remove(key K) (present bool) {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *LRU[K, V]) RemoveOldest() (key K, value V, ok bool) {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value
		return kv.key, kv.value, true
	}
	return c.zeroKey(), c.zeroValue(), false
}

// GetOldest returns the oldest entry
func (c *LRU[K, V]) GetOldest() (key K, value V, ok bool) {
	ent := c.evictList.Back()
	if ent != nil {
		kv := ent.Value
		return kv.key, kv.value, true
	}
	return c.zeroKey(), c.zeroValue(), false
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *LRU[K, V]) Keys() []K {
	keys := make([]K, len(c.items))
	i := 0
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev() {
		keys[i] = ent.Value.key
		i++
	}
	return keys
}

// Items returns a slice of the key-value pairs in the cache, from oldest to newest.
func (c *LRU[K, V]) Items() []Entry[K, V] {
	items := make([]Entry[K, V], len(c.items))
	i := 0
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev() {
		items[i] = Entry[K, V]{
			Key:   ent.Value.key,
			Value: ent.Value.value,
		}
		i++
	}
	return items
}

// Len returns the number of items in the cache.
func (c *LRU[K, V]) Len() int {
	return c.evictList.Len()
}

// Resize changes the cache size.
func (c *LRU[K, V]) Resize(size int) (evicted int) {
	diff := c.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

// removeOldest removes the oldest item from the cache.
func (c *LRU[K, V]) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

// removeElement is used to remove a given list element from the cache
func (c *LRU[K, V]) removeElement(e *linkedlist.Element[*entry[K, V]]) {
	c.evictList.Remove(e)
	kv := e.Value
	delete(c.items, kv.key)
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.value)
	}
}
