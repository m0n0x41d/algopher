package hashtable

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

var _ = os.Args
var _ = strconv.Itoa(0)

// The answer to the ultimate question of life, the universe, and everything
const MAGIC_NUMBER = 42

type HashTable struct {
	size  int
	step  int
	salt  uint // random salt for HashDoS protection
	slots []*string
}

// Generates random salt for HashDoS protection.
// Time: O(n) where n = sz (allocating slots)
// Space: O(n)
func Init(sz int, stp int) HashTable {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ht := HashTable{
		size:  sz,
		step:  stp,
		salt:  uint(r.Uint64()),
		slots: nil,
	}
	ht.slots = make([]*string, sz)
	return ht
}

// Time: O(k) where k = len(value)
// Space: O(1)
func (ht *HashTable) HashFun(value string) int {
	hash := ht.salt
	for _, x := range []byte(value) {
		hash = hash*MAGIC_NUMBER + uint(x)
	}
	return int(hash % uint(ht.size))
}

// Time: O(1) average, O(n) worst case (table nearly full)
// Space: O(1)
func (ht *HashTable) SeekSlot(value string) int {
	start := ht.HashFun(value)
	idx := start
	for {
		if ht.slots[idx] == nil {
			return idx
		}
		idx = (idx + ht.step) % ht.size
		if idx == start {
			return -1
		}
	}
}

// Time: O(1) average, O(n) worst case (table nearly full)
// Space: O(1)
// Returns slot index or -1 if table is full.
func (ht *HashTable) Put(value string) int {
	slotCandidate := ht.SeekSlot(value)
	if slotCandidate > -1 {
		ht.slots[slotCandidate] = &value
		return slotCandidate
	}
	return -1
}

// Time: O(1) average, O(n) worst case (many collisions)
// Space: O(1)
func (ht *HashTable) Find(value string) int {
	start := ht.HashFun(value)
	idx := start
	for {
		if ht.slots[idx] == nil {
			return -1
		}
		if *ht.slots[idx] == value {
			return idx
		}
		idx = (idx + ht.step) % ht.size
		if idx == start {
			return -1
		}
	}
}
