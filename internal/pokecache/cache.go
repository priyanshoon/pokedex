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
	mu       sync.Mutex
	interval time.Duration
	data     map[string]cacheEntry
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value, ok := cache.data[key]
	if !ok {
		return value.val, false
	}

	return value.val, true
}

func (cache *Cache) reapLoop() {
	// to be implemented
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		data:     make(map[string]cacheEntry),
		interval: ttl,
	}
}
