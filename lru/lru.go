package lru

import (
	"container/list"
	"sync"
	"time"
)

// Cache implements a thread safe LRU with expirable entries.
type Cache[K comparable, V any] struct {
	size       int
	purgeEvery time.Duration
	ttl        time.Duration
	done       chan struct{}
	onEvicted  EvictCallback[K, V]

	sync.Mutex
	items     map[K]*list.Element
	evictList *list.List
}

// expirableEntry is used to hold a value in the evictList
type expirableEntry[K comparable, V any] struct {
	key       K
	value     V
	expiresAt time.Time
}

// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback[K comparable, V any] func(key K, value V)

// noEvictionTTL - very long ttl to prevent eviction
const noEvictionTTL = time.Hour * 24 * 365 * 10

// NewExpirableLRU returns a new cache with expirable entries.
//
// Size parameter set to 0 makes cache of unlimited size.
//
// Providing 0 TTL turns expiring off.
//
// Activates deleteExpired by purgeEvery duration.
// If MaxKeys and TTL are defined and PurgeEvery is zero, PurgeEvery will be set to 5 minutes.
func NewExpirableLRU[K comparable, V any](size int, onEvict EvictCallback[K, V], defaultTtl, purgeEvery time.Duration) *Cache[K, V] {
	if size < 0 {
		size = 0
	}
	if defaultTtl <= 0 {
		defaultTtl = noEvictionTTL
	}

	res := Cache[K, V]{
		items:      map[K]*list.Element{},
		evictList:  list.New(),
		ttl:        defaultTtl,
		purgeEvery: purgeEvery,
		size:       size,
		onEvicted:  onEvict,
		done:       make(chan struct{}),
	}

	// enable deleteExpired() running in separate goroutine for cache
	// with non-zero TTL and size defined
	if res.ttl != noEvictionTTL && (res.size > 0 || res.purgeEvery > 0) {
		if res.purgeEvery <= 0 {
			res.purgeEvery = time.Minute * 5 // non-zero purge enforced because size defined
		}
		go func(done <-chan struct{}) {
			ticker := time.NewTicker(res.purgeEvery)
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					res.Lock()
					res.deleteExpired()
					res.Unlock()
				}
			}
		}(res.done)
	}
	return &res
}

// Add adds a key and a value to the LRU interface
func (c *Cache[K, V]) Add(key K, value V) (evicted bool) {
	return c.add(key, value, c.ttl)
}

// AddWithTTL adds a key and a value with a TTL to the LRU interface
func (c *Cache[K, V]) AddWithTTL(key K, value V, ttl time.Duration) (evicted bool) {
	return c.add(key, value, ttl)
}

// add performs the actual addition to the LRU cache
func (c *Cache[K, V]) add(key K, value V, ttl time.Duration) (evicted bool) {
	c.Lock()
	defer c.Unlock()
	now := time.Now()

	// Check for existing item
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*expirableEntry[K, V]).value = value
		ent.Value.(*expirableEntry[K, V]).expiresAt = now.Add(ttl)
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
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.Lock()
	defer c.Unlock()
	if ent, ok := c.items[key]; ok {
		// Expired item check
		if time.Now().After(ent.Value.(*expirableEntry[K, V]).expiresAt) {
			return c.zeroValue(), false
		}
		c.evictList.MoveToFront(ent)
		return ent.Value.(*expirableEntry[K, V]).value, true
	}
	return c.zeroValue(), false
}

// Peek returns the key value (or undefined if not found) without updating the "recently used"-ness of the key.
func (c *Cache[K, V]) Peek(key K) (V, bool) {
	c.Lock()
	defer c.Unlock()
	if ent, ok := c.items[key]; ok {
		// Expired item check
		if time.Now().After(ent.Value.(*expirableEntry[K, V]).expiresAt) {
			return c.zeroValue(), false
		}
		return ent.Value.(*expirableEntry[K, V]).value, true
	}
	return c.zeroValue(), false
}

// GetOldest returns the oldest entry
func (c *Cache[K, V]) GetOldest() (key K, value V, ok bool) {
	c.Lock()
	defer c.Unlock()
	ent := c.evictList.Back()
	if ent != nil {
		kv := ent.Value.(*expirableEntry[K, V])
		return kv.key, kv.value, true
	}
	return c.zeroKey(), c.zeroValue(), false
}

// Contains checks if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *Cache[K, V]) Contains(key K) (ok bool) {
	c.Lock()
	defer c.Unlock()
	_, ok = c.items[key]
	return ok
}

// Remove key from the cache
func (c *Cache[K, V]) Remove(key K) bool {
	c.Lock()
	defer c.Unlock()
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache[K, V]) RemoveOldest() (key K, value V, ok bool) {
	c.Lock()
	defer c.Unlock()
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*expirableEntry[K, V])
		return kv.key, kv.value, true
	}
	return c.zeroKey(), c.zeroValue(), false
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *Cache[K, V]) Keys() []K {
	c.Lock()
	defer c.Unlock()
	return c.keys()
}

// Purge clears the cache completely.
func (c *Cache[K, V]) Purge() {
	c.Lock()
	defer c.Unlock()
	for k, v := range c.items {
		if c.onEvicted != nil {
			c.onEvicted(k, v.Value.(*expirableEntry[K, V]).value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}

// DeleteExpired clears cache of expired items
func (c *Cache[K, V]) DeleteExpired() {
	c.Lock()
	defer c.Unlock()
	c.deleteExpired()
}

// Len return count of items in cache
func (c *Cache[K, V]) Len() int {
	c.Lock()
	defer c.Unlock()
	return c.evictList.Len()
}

// Resize changes the cache size. size 0 doesn't resize the cache, as it means unlimited.
func (c *Cache[K, V]) Resize(size int) (evicted int) {
	if size <= 0 {
		return 0
	}
	c.Lock()
	defer c.Unlock()
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
func (c *Cache[K, V]) Close() {
	c.Lock()
	defer c.Unlock()
	close(c.done)
}

// removeOldest removes the oldest item from the cache. Has to be called with lock!
func (c *Cache[K, V]) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

// Keys returns a slice of the keys in the cache, from oldest to newest. Has to be called with lock!
func (c *Cache[K, V]) keys() []K {
	keys := make([]K, 0, len(c.items))
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev() {
		keys = append(keys, ent.Value.(*expirableEntry[K, V]).key)
	}
	return keys
}

// removeElement is used to remove a given list element from the cache. Has to be called with lock!
func (c *Cache[K, V]) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*expirableEntry[K, V])
	delete(c.items, kv.key)
	if c.onEvicted != nil {
		c.onEvicted(kv.key, kv.value)
	}
}

// deleteExpired deletes expired records. Has to be called with lock!
func (c *Cache[K, V]) deleteExpired() {
	for _, key := range c.keys() {
		if time.Now().After(c.items[key].Value.(*expirableEntry[K, V]).expiresAt) {
			c.removeElement(c.items[key])
			continue
		}
	}
}

func (c *Cache[K, V]) zeroKey() K {
	return zero[K]()
}

func (c *Cache[K, V]) zeroValue() V {
	return zero[V]()
}

// helper for generic default value
func zero[T any]() T {
	var result T
	return result
}
