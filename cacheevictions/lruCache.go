package cacheevictions

import (
	"container/list"
	"sync"
)

// LRUCache struct defines the fields for the LRU cache implementation
// maxSize defines the maximum number of items that the cache can hold
// items is a map that stores the items in the cache, with the key being the item's key and the value being a pointer to the item's element in the list
// list is a doubly-linked list that is used to keep track of the order of the items in the cache
// mu is a mutex to provide thread-safety for the cache
type LRUCache struct {
	maxSize int
	items   map[string]*list.Element
	list    *list.List
	mu      sync.Mutex
}

// entry struct defines the fields for an item stored in the cache
// key is the key for the item
// value is the value of the item
type entry struct {
	key   string
	value interface{}
}

// NewLRUCache function creates a new LRU cache with the specified maximum size
func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		maxSize: maxSize,
		items:   make(map[string]*list.Element),
		list:    list.New(),
	}
}

// Put method adds an item to the cache
// If the item already exists in the cache, it moves the item to the front of the list
// If the cache is full, it removes the least recently used item
func (c *LRUCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.list.MoveToFront(elem)          // Move the existing element to the front of the list
		elem.Value.(*entry).value = value // Update the value of the existing element
		return
	}
	// If the item does not exist in the cache, add it to the front of the list
	elem := c.list.PushFront(&entry{key, value})
	c.items[key] = elem

	// If the cache is full, remove the least recently used item
	if c.list.Len() > c.maxSize {
		c.removeOldest()
	}
}

// Get method retrieves an item from the cache
// If the item exists in the cache, it moves the item to the front of the list
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.items[key]; ok {
		c.list.MoveToFront(elem) // Move the element to the front of the list
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// removeOldest method removes the least recently used item from the cache
func (c *LRUCache) removeOldest() {
	elem := c.list.Back()
	if elem != nil {
		c.list.Remove(elem)
		entry := elem.Value.(*entry)
		delete(c.items, entry.key)
	}
}
