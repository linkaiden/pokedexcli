package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	mu           sync.Mutex
	data         map[string]cacheEntry
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Create a new Cache -
func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:         map[string]cacheEntry{},
		reapInterval: interval,
	}
	go cache.reapLoop()
	return &cache
}

// Add to Cache -
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

// Get value from Cache -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)

	for range ticker.C {
		c.reap(time.Now().UTC(), c.reapInterval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.data {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.data, key)
		}
	}
}
