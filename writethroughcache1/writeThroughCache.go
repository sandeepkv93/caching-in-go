package writethroughcache

import (
	"sync"
	"time"
)

// WriteThroughCache struct defines the fields for the Write-Through cache implementation
// cache is the in-memory cache that stores the data
// store is the backing store where the data will be permanently stored
// mu is a mutex to provide thread-safety for the cache
// ttl is the time duration for the key to expire
// ticker is a ticker which runs in background to expire the key
type WriteThroughCache struct {
	cache  map[string]interface{}
	store  map[string]interface{}
	mu     sync.RWMutex
	ttl    time.Duration
	ticker *time.Ticker
}

// NewWriteThroughCache function creates a new Write-Through cache with the specified time to live (ttl)
func NewWriteThroughCache(ttl time.Duration) *WriteThroughCache {
	cache := &WriteThroughCache{
		cache: make(map[string]interface{}),
		store: make(map[string]interface{}),
		ttl:   ttl,
	}
	// Create a ticker to expire the key
	cache.ticker = time.NewTicker(ttl)
	go cache.expire()
	return cache
}

// Put method adds an item to the cache and the backing store
func (c *WriteThroughCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
	c.store[key] = value
}

// Get method retrieves an item from the cache
func (c *WriteThroughCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.cache[key]
	return value, ok
}

// expire method runs on a ticker and expire the key after ttl
func (c *WriteThroughCache) expire() {
	for {
		select {
		case <-c.ticker.C:
			c.mu.Lock()
			for key := range c.cache {
				delete(c.cache, key)
				delete(c.store, key)
			}
			c.mu.Unlock()
		}
	}
}
