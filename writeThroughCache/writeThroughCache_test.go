package writethroughcache

import (
	"testing"
	"time"
)

func TestWriteThroughCache(t *testing.T) {
	cache := NewWriteThroughCache(time.Second * 5)
	cache.Put("item1", "value1")
	cache.Put("item2", "value2")
	cache.Put("item3", "value3")

	// Test getting an existing item
	value, ok := cache.Get("item2")
	if !ok || value != "value2" {
		t.Error("Expected value2, got", value)
	}

	// Test getting a non-existing item
	value, ok = cache.Get("item4")
	if ok || value != nil {
		t.Error("Expected nil, got", value)
	}

	// Test updating an existing item
	cache.Put("item2", "newValue2")
	value, ok = cache.Get("item2")
	if !ok || value != "newValue2" {
		t.Error("Expected newValue2, got", value)
	}

	// Test the expiration of a key
	time.Sleep(time.Second * 6)
	value, ok = cache.Get("item2")
	if ok || value != nil {
		t.Error("Expected nil, got", value)
	}
}
