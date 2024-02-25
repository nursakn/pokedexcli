package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex
}

func NewCache(lifetime int) *Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	cache.readLoop(lifetime)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool)  {
	c.mux.Lock()
	defer c.mux.Unlock()
	res, ok := c.cache[key]

	return res.val, ok
}

func (c *Cache) readLoop(lifetime int) {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for range ticker.C {
			for key, item := range c.cache {
				if item.createdAt.Before((time.Now().Add(-(time.Duration(lifetime) * time.Minute)))) {
					delete(c.cache, key)
				}
			}
		}
	}()
}
