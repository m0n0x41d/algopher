package main

// Dynamic HashTable
const LOAD_FACTOR_THRESHOLD = 0.7

type DynamicHashTable struct {
	size  int
	step  int
	count int
	slots []*string
}

// Time: O(n) where n = sz (allocating slots)
// Space: O(n)
func InitDynamic(sz int, stp int) DynamicHashTable {
	ht := DynamicHashTable{size: sz, step: stp, count: 0, slots: nil}
	ht.slots = make([]*string, sz)
	return ht
}

// Time: O(k) where k = len(value)
// Space: O(1)
func (ht *DynamicHashTable) HashFun(value string) int {
	var hash uint = 0
	for _, x := range []byte(value) {
		hash = hash*MAGIC_NUMBER + uint(x)
	}
	return int(hash % uint(ht.size))
}

// Time: O(1) average, O(n) worst case (table nearly full)
// Space: O(1)
func (ht *DynamicHashTable) SeekSlot(value string) int {
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

// Time: O(1)
// Space: O(1)
func (ht *DynamicHashTable) loadFactor() float64 {
	return float64(ht.count) / float64(ht.size)
}

// Time: O(n) where n = current count (must rehash all elements)
// Space: O(n) for new slots array
func (ht *DynamicHashTable) resize() {
	oldSlots := ht.slots
	ht.size = nextPrime(ht.size * 2)
	ht.slots = make([]*string, ht.size)
	ht.count = 0

	for _, slot := range oldSlots {
		if slot != nil {
			ht.put(*slot)
		}
	}
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (ht *DynamicHashTable) put(value string) int {
	idx := ht.SeekSlot(value)
	if idx != -1 {
		v := value
		ht.slots[idx] = &v
		ht.count++
	}
	return idx
}

// Time: O(1) amortized (occasional O(n) resize)
// Space: O(1) amortized
func (ht *DynamicHashTable) Put(value string) int {
	if float64(ht.count+1)/float64(ht.size) > LOAD_FACTOR_THRESHOLD {
		ht.resize()
	}
	return ht.put(value)
}

// Time: O(1) average, O(n) worst case (many collisions)
// Space: O(1)
func (ht *DynamicHashTable) Find(value string) int {
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

// Time: O(1)
// Space: O(1)
func (ht *DynamicHashTable) Count() int {
	return ht.count
}

// Time: O(1)
// Space: O(1)
func (ht *DynamicHashTable) Size() int {
	return ht.size
}

// Time: O(n * sqrt(n)) worst case, but primes are dense so usually fast
// Space: O(1)
func nextPrime(n int) int {
	if n <= 2 {
		return 2
	}
	if n%2 == 0 {
		n++
	}
	for !isPrime(n) {
		n += 2
	}
	return n
}

// Time: O(sqrt(n))
// Space: O(1)
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// =====================================
// MultiHashTable uses multiple hash functions to reduce collision probability.

const NUM_HASH_FUNCTIONS = 3

type MultiHashTable struct {
	size   int
	count  int
	slots  []*string
	primes []uint // different primes for different hash functions
}

// InitMultiHash creates a new multi-hash table with given size.
// Time: O(n) where n = sz (allocating slots)
// Space: O(n)
func InitMultiHash(sz int) MultiHashTable {
	ht := MultiHashTable{
		size:   sz,
		count:  0,
		slots:  make([]*string, sz),
		primes: []uint{31, 37, 41},
	}
	return ht
}

// hashN returns hash using n-th hash function (0-indexed).
// Time: O(k) where k = len(value)
// Space: O(1)
func (ht *MultiHashTable) hashN(value string, n int) int {
	var hash uint = 0
	prime := ht.primes[n%len(ht.primes)]
	for _, x := range []byte(value) {
		hash = hash*prime + uint(x)
	}
	// add offset based on hash function index to spread values
	hash += uint(n) * 7919
	return int(hash % uint(ht.size))
}

// allHashes returns all candidate slot indices for a value.
// Time: O(h * k) where h = NUM_HASH_FUNCTIONS, k = len(value)
// Space: O(h)
func (ht *MultiHashTable) allHashes(value string) []int {
	indices := make([]int, NUM_HASH_FUNCTIONS)
	for i := 0; i < NUM_HASH_FUNCTIONS; i++ {
		indices[i] = ht.hashN(value, i)
	}
	return indices
}

// Time: O(h) average (check h candidates), O(n) worst case (table nearly full)
// Space: O(h) for candidate indices
func (ht *MultiHashTable) SeekSlot(value string) int {
	indices := ht.allHashes(value)

	// first pass: find empty slot among candidates
	for _, idx := range indices {
		if ht.slots[idx] == nil {
			return idx
		}
	}

	// all candidate slots occupied - use linear probing from first hash
	start := indices[0]
	idx := (start + 1) % ht.size
	for idx != start {
		if ht.slots[idx] == nil {
			return idx
		}
		idx = (idx + 1) % ht.size
	}

	return -1
}

// Time: O(h) average, O(n) worst case (table nearly full)
// Space: O(h) for candidate indices
func (ht *MultiHashTable) Put(value string) int {
	idx := ht.SeekSlot(value)
	if idx != -1 {
		v := value
		ht.slots[idx] = &v
		ht.count++
	}
	return idx
}

// Time: O(h) average (check h candidates first), O(n) worst case (linear fallback)
// Space: O(h) for candidate indices
func (ht *MultiHashTable) Find(value string) int {
	indices := ht.allHashes(value)

	// first check all candidate slots (fast path)
	for _, idx := range indices {
		if ht.slots[idx] != nil && *ht.slots[idx] == value {
			return idx
		}
	}

	// not in candidate slots - linear search from first hash
	// (in case it was placed via linear probing due to collision)
	start := indices[0]
	idx := start
	for {
		if ht.slots[idx] == nil {
			return -1
		}
		if *ht.slots[idx] == value {
			return idx
		}
		idx = (idx + 1) % ht.size
		if idx == start {
			return -1
		}
	}
}

// Time: O(1)
// Space: O(1)
func (ht *MultiHashTable) Count() int {
	return ht.count
}

// Time: O(1)
// Space: O(1)
func (ht *MultiHashTable) Size() int {
	return ht.size
}

// Time: O(n * h) where n = size, h = NUM_HASH_FUNCTIONS
// Space: O(h) for candidate indices per element
func (ht *MultiHashTable) Stats() (primaryHits, secondaryHits, probingHits int) {
	for i := 0; i < ht.size; i++ {
		if ht.slots[i] == nil {
			continue
		}
		value := *ht.slots[i]
		indices := ht.allHashes(value)

		found := false
		for j, idx := range indices {
			if idx == i {
				if j == 0 {
					primaryHits++
				} else {
					secondaryHits++
				}
				found = true
				break
			}
		}
		if !found {
			probingHits++
		}
	}
	return
}
