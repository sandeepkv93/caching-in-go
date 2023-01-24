package lfu_cache

import (
	"container/list"
	"sync"
)

// LFUCache is a thread-safe, least frequently used (LFU) cache.
type LFUCache struct {
	// maxEntries is the maximum number of entries the cache can hold.
	maxEntries int

	// evictionList is a list of eviction nodes, with the most recently
	// accessed entry at the front of the list and the least recently
	// accessed entry at the back of the list.
	evictionList *list.List

	// items is a map of keys to eviction nodes.
	items map[string]*list.Element

	// lock synchronizes access to the cache.
	lock sync.RWMutex
}

// entry is an entry in the cache.
type entry struct {
	key   string
	value interface{}
	count int
}

// NewLFUCache returns a new LFU cache with the given maximum number of entries.
func NewLFUCache(maxEntries int) *LFUCache {
	return &LFUCache{
		maxEntries:   maxEntries,
		evictionList: list.New(),
		items:        make(map[string]*list.Element),
	}
}

// Set sets the value for the given key.
func (c *LFUCache) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// Check if the key is already in the cache.
	if ent, ok := c.items[key]; ok {
		c.evictionList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		ent.Value.(*entry).count++
		return
	}

	// If the cache is full, remove the least frequently used entry.
	if c.evictionList.Len() >= c.maxEntries {
		ent := c.evictionList.Back()
		c.removeElement(ent)
	}

	// Add the new entry to the front of the list and the map.
	ent := c.evictionList.PushFront(&entry{key: key, value: value, count: 1})
	c.items[key] = ent
}

// Get gets the value for the given key.
func (c *LFUCache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	// Check if the key is in the cache.
	if ent, ok := c.items[key]; ok {
		c.evictionList.MoveToFront(ent)
		ent.Value.(*entry).count++
		return ent.Value.(*entry).value, true
	}

	return nil, false
}

// removeElement removes the given list element from the cache.
func (c *LFUCache) removeElement(e *list.Element) {
	c.evictionList.Remove(e)
	delete(c.items, e.Value.(*entry).key)
}
