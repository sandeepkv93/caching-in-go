package lfucache

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	// Create a new LFU cache with a maximum capacity of 2 entries.
	c := NewLFUCache(2)

	// Test set and get operations.
	c.Set("foo", "bar")
	if val, ok := c.Get("foo"); !ok || val != "bar" {
		t.Error("Expected to get value 'bar' for key 'foo'")
	}

	c.Set("baz", "qux")
	if val, ok := c.Get("baz"); !ok || val != "qux" {
		t.Error("Expected to get value 'qux' for key 'baz'")
	}

	// Test eviction of least frequently used entry.
	c.Get("foo")
	c.Set("quux", "quuz")
	if _, ok := c.Get("baz"); ok {
		t.Error("Expected key 'baz' to be evicted")
	}

	// Test eviction of oldest entry.
	c.Set("corge", "grault")
	if _, ok := c.Get("foo"); ok {
		t.Error("Expected key 'foo' to be evicted")
	}

	// Test getting non-existent key.
	if _, ok := c.Get("missing"); ok {
		t.Error("Expected to get 'false' for non-existent key")
	}
}
