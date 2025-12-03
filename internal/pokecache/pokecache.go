package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	out := &Cache{cacheMap: map[string]cacheEntry{}}
	go out.reapLoop(interval)
	return out
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{time.Now(), val}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()
	for range tick.C {
		c.mu.Lock()
		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
