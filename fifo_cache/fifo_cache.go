package fifo_cache

import (
	"container/list"
	"sync"
)

// FIFOCache is a struct that implements a FIFO cache.
type FIFOCache struct {
	maxSize int
	items   map[string]*list.Element
	list    *list.List
	mu      sync.Mutex
}

// entry is a struct that holds the key-value pair in the cache.
type entry struct {
	key   string
	value interface{}
}

// NewFIFOCache returns a new instance of FIFOCache with the given maximum size.
func NewFIFOCache(maxSize int) *FIFOCache {
	return &FIFOCache{
		maxSize: maxSize,
		items:   make(map[string]*list.Element),
		list:    list.New(),
	}
}

// Put adds a key-value pair to the cache.
// If the key already exists, it updates its value.
// If the cache is full, it removes the oldest entry.
func (c *FIFOCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If the key already exists, update its value and move it to the front.
	if elem, ok := c.items[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	} // If the key doesn't exist, add it to the front of the list and add it to the items map.
	elem := c.list.PushFront(&entry{key, value})
	c.items[key] = elem

	// If the cache is full, remove the oldest entry (the one at the back of the list).
	if c.list.Len() > c.maxSize {
		c.removeOldest()
	}
}

// Get returns the value associated with the given key, and a boolean indicating if the key exists in the cache.
func (c *FIFOCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.items[key]; ok {
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// removeOldest removes the oldest entry (the one at the back of the list) from the cache.
func (c *FIFOCache) removeOldest() {
	elem := c.list.Back()
	if elem != nil {
		c.list.Remove(elem)
		entry := elem.Value.(*entry)
		delete(c.items, entry.key)
	}
}
