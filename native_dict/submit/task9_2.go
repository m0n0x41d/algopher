package main

// OrderedDictionary based on ordered list with binary search.
//
// Time complexity:
// - Search (Get, IsKey): O(log n) - binary search
// - Insert (Put): O(n) - binary search + shift elements
// - Delete: O(n) - binary search + shift elements
//
// Trade-off vs hash table:
// + Guaranteed O(log n) search (no worst case O(n) on collisions)
// + Keys stored in sorted order (range queries possible)
// + No hash function selection or collision problems
// - Insert/delete O(n) instead of O(1) amortized

type OrderedDictionary[T any] struct {
	count  int
	keys   []string
	values []T
}

// Time: O(n) where n = sz
// Space: O(n)
func InitOrdered[T any](sz int) OrderedDictionary[T] {
	return OrderedDictionary[T]{
		count:  0,
		keys:   make([]string, 0, sz),
		values: make([]T, 0, sz),
	}
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) IsKey(key string) bool {
	_, found := od.binarySearch(key)
	return found
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) Get(key string) (T, error) {
	var result T
	index, found := od.binarySearch(key)
	if !found {
		return result, ErrKeyNotFound
	}
	return od.values[index], nil
}

// Time: O(n) - binary search O(log n) + shift O(n)
// Space: O(1) amortized
func (od *OrderedDictionary[T]) Put(key string, value T) {
	index, found := od.binarySearch(key)
	if found {
		od.values[index] = value
		return
	}

	od.keys = append(od.keys, "")
	od.values = append(od.values, value)

	copy(od.keys[index+1:], od.keys[index:])
	copy(od.values[index+1:], od.values[index:])

	od.keys[index] = key
	od.values[index] = value
	od.count++
}

// Time: O(n) - binary search O(log n) + shift O(n)
// Space: O(1)
func (od *OrderedDictionary[T]) Delete(key string) bool {
	index, found := od.binarySearch(key)
	if !found {
		return false
	}

	od.keys = append(od.keys[:index], od.keys[index+1:]...)
	od.values = append(od.values[:index], od.values[index+1:]...)
	od.count--
	return true
}

// Time: O(1)
// Space: O(1)
func (od *OrderedDictionary[T]) Count() int {
	return od.count
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) binarySearch(key string) (index int, found bool) {
	left, right := 0, len(od.keys)-1
	for left <= right {
		mid := left + (right-left)/2
		if od.keys[mid] == key {
			return mid, true
		}
		if od.keys[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left, false
}

// BitKeyDictionary - hash table with fixed-length bit string keys (uint64).
// Assumes the user works with known bit-length keys/identifiers.
//
// Advantages over string keys:
// - Key comparison O(1) - single CPU instruction vs byte-by-byte
// - Hash computation O(1) - XOR with salt vs iterating over string
// - Cache-friendly - 8 bytes fixed vs variable length strings
//
// Time complexity:
// - Search (Get, IsKey): O(1) average, O(n) worst case
// - Insert (Put): O(1) average, O(n) worst case
// - Delete: O(1) average, O(n) worst case

const BITKEY_STEP = 3

type BitKeyDictionary[T any] struct {
	size   int
	salt   uint64
	keys   []uint64
	values []T
	filled []bool
}

// Time: O(n) where n = sz
// Space: O(n)
func InitBitKey[T any](sz int, salt uint64) BitKeyDictionary[T] {
	return BitKeyDictionary[T]{
		size:   sz,
		salt:   salt,
		keys:   make([]uint64, sz),
		values: make([]T, sz),
		filled: make([]bool, sz),
	}
}

// Time: O(1)
// Space: O(1)
func (bd *BitKeyDictionary[T]) HashFun(key uint64) int {
	// XOR with salt for HashDoS protection
	hash := key ^ bd.salt
	return int(hash % uint64(bd.size))
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) IsKey(key uint64) bool {
	index, found := bd.findSlot(key)
	_ = index
	return found
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) Get(key uint64) (T, error) {
	var result T
	index, found := bd.findSlot(key)
	if !found {
		return result, ErrKeyNotFound
	}
	return bd.values[index], nil
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) Put(key uint64, value T) {
	index, found := bd.findSlot(key)
	if found {
		bd.values[index] = value
		return
	}

	emptyIndex := bd.seekSlot(key)
	if emptyIndex != -1 {
		bd.keys[emptyIndex] = key
		bd.filled[emptyIndex] = true
		bd.values[emptyIndex] = value
	}
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) Delete(key uint64) bool {
	index, found := bd.findSlot(key)
	if !found {
		return false
	}

	bd.filled[index] = false
	return true
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) findSlot(key uint64) (index int, found bool) {
	start := bd.HashFun(key)
	idx := start
	for {
		if !bd.filled[idx] {
			return idx, false
		}
		// XOR comparison: equal if result is 0
		if (bd.keys[idx] ^ key) == 0 {
			return idx, true
		}
		idx = (idx + BITKEY_STEP) % bd.size
		if idx == start {
			return -1, false
		}
	}
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (bd *BitKeyDictionary[T]) seekSlot(key uint64) int {
	start := bd.HashFun(key)
	idx := start
	for {
		if !bd.filled[idx] {
			return idx
		}
		idx = (idx + BITKEY_STEP) % bd.size
		if idx == start {
			return -1
		}
	}
}
