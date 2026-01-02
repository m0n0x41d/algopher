package native_dict

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
