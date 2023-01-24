package fifo_cache

import "testing"

func TestFIFOCache(t *testing.T) {
	cache := NewFIFOCache(3)

	// Test adding elements to the cache
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Test getting elements from the cache
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

	// Test updating elements in the cache
	cache.Put("a", 4)
	val, ok = cache.Get("a")
	if !ok || val != 4 {
		t.Error("Expected value 4, got", val)
	}

	// Test removing oldest element from the cache
	cache.Put("d", 4)
	_, ok = cache.Get("a")
	if ok {
		t.Error("Expected a to be removed, but it still exists")
	}
}
