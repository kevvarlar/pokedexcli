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
	CacheEntry map[string]cacheEntry
	mu         *sync.Mutex
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.CacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.CacheEntry[key]
	if !ok {
		return []byte{}, ok
	}
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for key, val := range c.CacheEntry{
			if time.Since(val.createdAt) > interval {
				delete(c.CacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		CacheEntry: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
