package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) cache {
	cache := cache{
		entries: map[string]cacheEntry{},
		mu:      sync.Mutex{},
	}

	return cache
}

func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return v.val, true
}

func (c *cache) reapLoop() {

}
