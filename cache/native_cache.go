package cache

import "errors"

var ErrKeyNotFound = errors.New("key not found in cache")

// NativeCache is a simple LFU cache based on hash table with open addressing.
// This is the "naive" implementation as suggested by the lab task.
//
// Eviction strategy: when cache is full, evicts the element with minimum hits count.
// This requires O(n) scan of hits[] array - inefficient compared to LFUCache,
// but simpler to understand and matches the original task requirements.
//
// See LFUCache in lfu_cache.go for O(1) eviction implementation.

type NativeCache[T any] struct {
	size   int
	step   int
	slots  []string
	values []T
	hits   []int
}

// Time: O(n) for slices allocation
// Space: O(n) where n = size
func InitNativeCache[T any](size int, step int) NativeCache[T] {
	return NativeCache[T]{
		size:   size,
		step:   step,
		slots:  make([]string, size),
		values: make([]T, size),
		hits:   make([]int, size),
	}
}

// Time: O(k) where k = len(key)
// Space: O(1)
func (nc *NativeCache[T]) HashFun(key string) int {
	var sum int
	for _, r := range key {
		sum += int(r)
	}
	return sum % nc.size
}

// Time: O(n) worst case (full table with collisions)
// Space: O(1)
func (nc *NativeCache[T]) seekSlot(key string) int {
	idx := nc.HashFun(key)

	for i := 0; i < nc.size; i++ {
		if nc.slots[idx] == "" || nc.slots[idx] == key {
			return idx
		}
		idx = (idx + nc.step) % nc.size
	}

	return -1
}

// Time: O(n) - linear scan of hits array
// Space: O(1)
func (nc *NativeCache[T]) findMinHitsIdx() int {
	minIdx := 0
	minHits := nc.hits[0]

	for i, h := range nc.hits {
		if h < minHits {
			minHits = h
			minIdx = i
		}
	}

	return minIdx
}

// Time: O(n) worst case
// Space: O(1)
func (nc *NativeCache[T]) Put(key string, value T) {
	idx := nc.seekSlot(key)

	if idx == -1 {
		idx = nc.findMinHitsIdx()
		nc.hits[idx] = 0
	}

	nc.slots[idx] = key
	nc.values[idx] = value
}

// Time: O(n) worst case (probing through collisions)
// Space: O(1)
func (nc *NativeCache[T]) Get(key string) (T, error) {
	idx := nc.findKey(key)

	if idx == -1 {
		var zero T
		return zero, ErrKeyNotFound
	}

	nc.hits[idx]++
	return nc.values[idx], nil
}

// Time: O(n) worst case
// Space: O(1)
func (nc *NativeCache[T]) IsKey(key string) bool {
	return nc.findKey(key) != -1
}

// Time: O(n) worst case
// Space: O(1)
func (nc *NativeCache[T]) findKey(key string) int {
	idx := nc.HashFun(key)

	for i := 0; i < nc.size; i++ {
		if nc.slots[idx] == key {
			return idx
		}
		if nc.slots[idx] == "" {
			return -1
		}
		idx = (idx + nc.step) % nc.size
	}

	return -1
}

// GetHits returns hit count for a key (for testing purposes)
// Time: O(n) worst case
// Space: O(1)
func (nc *NativeCache[T]) GetHits(key string) int {
	idx := nc.findKey(key)
	if idx == -1 {
		return -1
	}
	return nc.hits[idx]
}
