package lru

import (
	"sync"
	"time"

	"github.com/ImSingee/go-ex/linkedlist"
)

// ExpirableCache implements a thread safe LRU with expirable entries.
type ExpirableCache[K comparable, V any] struct {
	size       int
	purgeEvery time.Duration
	ttl        time.Duration
	done       chan struct{}
	onEvicted  EvictCallback[K, V]

	mu        sync.Mutex
	items     map[K]*linkedlist.Element[*expirableEntry[K, V]]
	evictList *linkedlist.List[*expirableEntry[K, V]]
}

// expirableEntry is used to hold a value in the evictList
type expirableEntry[K comparable, V any] struct {
	key       K
	value     V
	expiresAt time.Time
}

func (e *expirableEntry[K, V]) Expired() bool {
	return time.Now().After(e.expiresAt)
}

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback[K comparable, V any] func(key K, value V)

// noEvictionTTL - very long ttl to prevent eviction
const noEvictionTTL = time.Hour * 24 * 365 * 10

// NewExpirable returns a new cache with expirable entries.
//
// If size <= 0 means unlimited size
// If defaultTtl <= 0 means by default there is no ttl
// If purgeEvery = 0 means 5min
//
// Please remember to call .Close if you don't use the cache anymore
func NewExpirable[K comparable, V any](size int, onEvict EvictCallback[K, V], defaultTtl, purgeEvery time.Duration) *ExpirableCache[K, V] {
	if size < 0 {
		size = 0
	}
	if defaultTtl <= 0 {
		defaultTtl = noEvictionTTL
	}
	if purgeEvery == 0 {
		purgeEvery = 5 * time.Minute
	} else if purgeEvery < 0 {
		panic("Invalid Argument: purgeEvery must be a non-negative integer")
	}

	res := ExpirableCache[K, V]{
		items:      map[K]*linkedlist.Element[*expirableEntry[K, V]]{},
		evictList:  linkedlist.New[*expirableEntry[K, V]](),
		ttl:        defaultTtl,
		purgeEvery: purgeEvery,
		size:       size,
		onEvicted:  onEvict,
		done:       make(chan struct{}),
	}

	// enable deleteExpired() running in separate goroutine for cache
	// with non-zero TTL and size defined
	go func(done <-chan struct{}) {
		ticker := time.NewTicker(res.purgeEvery)
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				res.mu.Lock()
				res.deleteExpired()
				res.mu.Unlock()
			}
		}
	}(res.done)
	return &res
}

// Add adds a key and a value to the LRU interface (ttl equals to the default)
func (c *ExpirableCache[K, V]) Add(key K, value V) (evicted bool) {
	return c.add(key, value, c.ttl)
}

// AddWithTTL adds a key and a value with a TTL to the LRU interface
func (c *ExpirableCache[K, V]) AddWithTTL(key K, value V, ttl time.Duration) (evicted bool) {
	return c.add(key, value, ttl)
}

// add performs the actual addition to the LRU cache
func (c *ExpirableCache[K, V]) add(key K, value V, ttl time.Duration) (evicted bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()

	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.value = value
		ent.Value.expiresAt = now.Add(ttl)
		return false
	}

	// Add new item
	ent := &expirableEntry[K, V]{key: key, value: value, expiresAt: now.Add(ttl)}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	// Verify size not exceeded
	if c.size > 0 && len(c.items) > c.size {
		c.removeOldest()
		return true
	}
	return false
}

// Get returns the key value
func (c *ExpirableCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent, ok := c.items[key]; ok {
		// Expired item check
		if ent.Value.Expired() {
			return c.zeroValue(), false
		}
		c.evictList.MoveToFront(ent)
		return ent.Value.value, true
	}
	return c.zeroValue(), false
}

// Peek returns the key value (or undefined if not found) without updating the "recently used"-ness of the key.
func (c *ExpirableCache[K, V]) Peek(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent, ok := c.items[key]; ok {
		// Expired item check
		if ent.Value.Expired() {
			return c.zeroValue(), false
		}
		return ent.Value.value, true
	}
	return c.zeroValue(), false
}

// PeekOldest returns the oldest entry
// This function won't change the order of the entry, if you want to move
// it the front, use GetOldest
func (c *ExpirableCache[K, V]) PeekOldest() (key K, value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ent := c.peekOldest()
	if ent == nil {
		return c.zeroKey(), c.zeroValue(), false
	} else {
		kv := ent.Value
		return kv.key, kv.value, true
	}
}

func (c *ExpirableCache[K, V]) peekOldest() (ent *linkedlist.Element[*expirableEntry[K, V]]) {
	for {
		ent = c.evictList.Back()
		if ent == nil { // no more elements
			return nil
		}
		if ent.Value.Expired() { // expired
			c.removeElement(ent)
			continue
		}

		return ent
	}
}

// GetOldest returns the oldest entry and mark it as recently-used
func (c *ExpirableCache[K, V]) GetOldest() (key K, value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ent := c.peekOldest()
	if ent == nil {
		return c.zeroKey(), c.zeroValue(), false
	} else {
		c.evictList.MoveToFront(ent)
		kv := ent.Value
		return kv.key, kv.value, true
	}
}

// Contains checks if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *ExpirableCache[K, V]) Contains(key K) (ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ent, ok := c.items[key]
	if ok {
		if ent.Value.Expired() {
			c.removeElement(ent)
			return false
		}
	}

	return ok
}

// Remove key from the cache
func (c *ExpirableCache[K, V]) Remove(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *ExpirableCache[K, V]) RemoveOldest() (key K, value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ent := c.peekOldest()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value
		return kv.key, kv.value, true
	}
	return c.zeroKey(), c.zeroValue(), false
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
// Warning: returned slice may contain expired keys,
// if you want only un-expired keys, use UnexpiredKeys
func (c *ExpirableCache[K, V]) Keys() []K {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.keys()
}

// UnexpiredKeys returns a slice of the keys in the cache, from oldest to newest.
func (c *ExpirableCache[K, V]) UnexpiredKeys() []K {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deleteExpired()
	return c.keys()
}

// Purge clears the cache completely.
func (c *ExpirableCache[K, V]) Purge() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.items {
		if c.onEvicted != nil {
			c.onEvicted(k, v.Value.value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}

// DeleteExpired clears cache of expired items
func (c *ExpirableCache[K, V]) DeleteExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deleteExpired()
}

// Len return count of items in cache
func (c *ExpirableCache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.evictList.Len()
}

// Resize changes the cache size. size 0 doesn't resize the cache, as it means unlimited.
func (c *ExpirableCache[K, V]) Resize(size int) (evicted int) {
	if size <= 0 {
		return 0
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	diff := c.evictList.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff; i++ {
		c.removeOldest()
	}
	c.size = size
	return diff
}

// Close cleans the cache and destroys running goroutines
func (c *ExpirableCache[K, V]) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.done)
}

// removeOldest removes the oldest item from the cache. Has to be called with lock!
func (c *ExpirableCache[K, V]) removeOldest() {
	ent := c.peekOldest()
	if ent != nil {
		c.removeElement(ent)
	}
}

// Keys returns a slice of the keys in the cache, from oldest to newest. Has to be called with lock!
func (c *ExpirableCache[K, V]) keys() []K {
	keys := make([]K, 0, len(c.items))
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev() {
		keys = append(keys, ent.Value.key)
	}
	return keys
}

// removeElement is used to remove a given list element from the cache. Has to be called with lock!
func (c *ExpirableCache[K, V]) removeElement(e *linkedlist.Element[*expirableEntry[K, V]]) {
	c.evictList.Remove(e)
	kv := e.Value
	delete(c.items, kv.key)
	if c.onEvicted != nil {
		c.onEvicted(kv.key, kv.value)
	}
}

// deleteExpired deletes expired records. Has to be called with lock!
func (c *ExpirableCache[K, V]) deleteExpired() {
	for _, key := range c.keys() {
		if c.items[key].Value.Expired() {
			c.removeElement(c.items[key])
			continue
		}
	}
}

func (c *ExpirableCache[K, V]) zeroKey() K {
	return zero[K]()
}

func (c *ExpirableCache[K, V]) zeroValue() V {
	return zero[V]()
}

// helper for generic default value
func zero[T any]() T {
	var result T
	return result
}
