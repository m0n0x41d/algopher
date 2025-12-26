package hashtable

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
