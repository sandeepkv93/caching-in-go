package fifo_cache

import (
	"testing"
)

func TestFIFOCache(t *testing.T) {
	// create a new cache with a maximum size of 3
	cache := NewFIFOCache(3)

	// test adding elements to the cache
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// test getting elements from the cache
	val, ok := cache.Get("a")
	if !ok || val != 1 {
		t.Error("Expected value 1, got", val)
	}
	val, ok = cache.Get("b")
	if !ok || val != 2 {
		t.Error("Expected value 2, got", val)
	}
	val, ok = cache.Get("c")
	if !ok || val != 3 {
		t.Error("Expected value 3, got", val)
	}

	// test updating elements in the cache
	cache.Put("a", 4)
	val, ok = cache.Get("a")
	if !ok || val != 4 {
		t.Error("Expected value 4, got", val)
	}

	// test removing the oldest element
	cache.Put("d", 4)
	_, ok = cache.Get("b")
	if ok {
		t.Error("Expected b to be removed, but it still exists")
	}

	// test getting a non-existing element
	val, ok = cache.Get("e")
	if ok || val != nil {
		t.Error("Expected value to be nil, got", val)
	}
}
