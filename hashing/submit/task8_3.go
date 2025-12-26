package main

import (
	"fmt"
	"testing"
)

func TestHashTableInit(t *testing.T) {
	ht := Init(17, 2)
	if ht.size != 17 {
		t.Errorf("size is %d, expected 17", ht.size)
	}
	if ht.step != 2 {
		t.Errorf("step is %d, expected 2", ht.step)
	}
	if len(ht.slots) != 17 {
		t.Errorf("slots is %d, expected 17", len(ht.slots))
	}
}

func TestHashTableHashFunc(t *testing.T) {
	ht := Init(17, 2)
	idxFromHashFun := ht.HashFun("boogie-woogieasdыыfasdf")
	if idxFromHashFun < 0 || idxFromHashFun >= ht.size {
		t.Errorf("HashFun returned %d, expected 0..%d", idxFromHashFun, ht.size-1)
	}

	fmt.Println(idxFromHashFun)
}

func TestSeekSlot_EmptyTable(t *testing.T) {
	ht := Init(17, 3)
	idx := ht.SeekSlot("test")
	if idx < 0 || idx >= ht.size {
		t.Errorf("SeekSlot on empty table returned %d, expected valid index", idx)
	}
}

func TestSeekSlot_ReturnsHashIndex(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	expectedIdx := ht.HashFun(value)
	idx := ht.SeekSlot(value)
	if idx != expectedIdx {
		t.Errorf("SeekSlot returned %d, expected %d (hash index)", idx, expectedIdx)
	}
}

func TestSeekSlot_WithCollision(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	hashIdx := ht.HashFun(value)

	// occupy the hash index
	occupied := "occupied"
	ht.slots[hashIdx] = &occupied

	idx := ht.SeekSlot(value)
	expectedIdx := (hashIdx + ht.step) % ht.size
	if idx != expectedIdx {
		t.Errorf("SeekSlot with collision returned %d, expected %d", idx, expectedIdx)
	}
}

func TestSeekSlot_MultipleCollisions(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	hashIdx := ht.HashFun(value)

	// occupy first two positions in probe sequence
	occupied := "occupied"
	ht.slots[hashIdx] = &occupied
	ht.slots[(hashIdx+ht.step)%ht.size] = &occupied

	idx := ht.SeekSlot(value)
	expectedIdx := (hashIdx + 2*ht.step) % ht.size
	if idx != expectedIdx {
		t.Errorf("SeekSlot with multiple collisions returned %d, expected %d", idx, expectedIdx)
	}
}

func TestSeekSlot_FullTable(t *testing.T) {
	ht := Init(17, 3)
	occupied := "occupied"
	for i := 0; i < ht.size; i++ {
		ht.slots[i] = &occupied
	}

	idx := ht.SeekSlot("anything")
	if idx != -1 {
		t.Errorf("SeekSlot on full table returned %d, expected -1", idx)
	}
}

func TestPut_EmptyTable(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	idx := ht.Put(value)
	if idx < 0 || idx >= ht.size {
		t.Errorf("Put returned %d, expected valid index", idx)
	}
	if ht.slots[idx] == nil || *ht.slots[idx] != value {
		t.Errorf("Value not stored at index %d", idx)
	}
}

func TestPut_ReturnsHashIndex(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	expectedIdx := ht.HashFun(value)
	idx := ht.Put(value)
	if idx != expectedIdx {
		t.Errorf("Put returned %d, expected %d (hash index)", idx, expectedIdx)
	}
}

func TestPut_WithCollision(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	hashIdx := ht.HashFun(value)

	// occupy the hash index
	occupied := "occupied"
	ht.slots[hashIdx] = &occupied

	idx := ht.Put(value)
	expectedIdx := (hashIdx + ht.step) % ht.size
	if idx != expectedIdx {
		t.Errorf("Put with collision returned %d, expected %d", idx, expectedIdx)
	}
	if ht.slots[idx] == nil || *ht.slots[idx] != value {
		t.Errorf("Value not stored at index %d", idx)
	}
}

func TestPut_MultipleValues(t *testing.T) {
	ht := Init(17, 3)
	values := []string{"one", "two", "three", "four", "five"}

	for _, v := range values {
		idx := ht.Put(v)
		if idx == -1 {
			t.Errorf("Put(%q) returned -1", v)
		}
		if ht.slots[idx] == nil || *ht.slots[idx] != v {
			t.Errorf("Value %q not stored correctly", v)
		}
	}
}

func TestPut_FullTable(t *testing.T) {
	ht := Init(5, 1) // small table
	values := []string{"a", "b", "c", "d", "e"}

	for _, v := range values {
		idx := ht.Put(v)
		if idx == -1 {
			t.Errorf("Put(%q) returned -1 before table full", v)
		}
	}

	// table is full, next put should fail
	idx := ht.Put("overflow")
	if idx != -1 {
		t.Errorf("Put on full table returned %d, expected -1", idx)
	}
}

func TestFind_EmptyTable(t *testing.T) {
	ht := Init(17, 3)
	idx := ht.Find("hello")
	if idx != -1 {
		t.Errorf("Find on empty table returned %d, expected -1", idx)
	}
}

func TestFind_ExistingValue(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	putIdx := ht.Put(value)
	findIdx := ht.Find(value)
	if findIdx != putIdx {
		t.Errorf("Find returned %d, expected %d", findIdx, putIdx)
	}
}

func TestFind_NonExistingValue(t *testing.T) {
	ht := Init(17, 3)
	ht.Put("one")
	ht.Put("two")
	ht.Put("three")

	idx := ht.Find("four")
	if idx != -1 {
		t.Errorf("Find for non-existing value returned %d, expected -1", idx)
	}
}

func TestFind_AfterCollision(t *testing.T) {
	ht := Init(17, 3)
	value := "hello"
	hashIdx := ht.HashFun(value)

	// occupy the hash index first
	ht.Put("blocker")
	// force blocker to be at hashIdx by direct assignment
	blocker := "blocker"
	ht.slots[hashIdx] = &blocker

	// now put our value - it will go to a different slot
	putIdx := ht.Put(value)
	findIdx := ht.Find(value)

	if findIdx != putIdx {
		t.Errorf("Find after collision returned %d, expected %d", findIdx, putIdx)
	}
}

func TestFind_MultipleValues(t *testing.T) {
	ht := Init(17, 3)
	values := []string{"one", "two", "three", "four", "five"}
	indices := make(map[string]int)

	for _, v := range values {
		indices[v] = ht.Put(v)
	}

	for _, v := range values {
		findIdx := ht.Find(v)
		if findIdx != indices[v] {
			t.Errorf("Find(%q) returned %d, expected %d", v, findIdx, indices[v])
		}
	}
}

// generateCollidingKeys generates keys that would collide in an unsalted hash table.
// With salt, these keys will have different hashes in different table instances.
func generateCollidingKeys(count int, targetHash int, size int) []string {
	// For unsalted hash with MAGIC_NUMBER=42:
	// hash = 0*42 + byte1 = byte1
	// hash = byte1*42 + byte2
	// We generate keys that would all hash to targetHash without salt
	keys := make([]string, 0, count)

	// Simple approach: find single-char keys that hash to same value mod size
	for c := byte('A'); c <= byte('z') && len(keys) < count; c++ {
		if int(c)%size == targetHash%size {
			keys = append(keys, string(c))
		}
	}

	// If not enough, generate two-char keys
	for c1 := byte('A'); c1 <= byte('z') && len(keys) < count; c1++ {
		for c2 := byte('A'); c2 <= byte('z') && len(keys) < count; c2++ {
			hash := (uint(c1)*42 + uint(c2)) % uint(size)
			if int(hash) == targetHash%size {
				keys = append(keys, string([]byte{c1, c2}))
			}
		}
	}

	return keys
}

// TestHashDoS_Attack demonstrates HashDoS attack on unsalted table
func TestHashDoS_Attack(t *testing.T) {
	size := 17
	targetHash := 5

	// Generate keys that collide in unsalted hash
	collidingKeys := generateCollidingKeys(10, targetHash, size)

	t.Logf("Generated %d colliding keys for target hash %d", len(collidingKeys), targetHash)

	// With salt, same keys should NOT all collide
	ht := Init(size, 3)

	hashCounts := make(map[int]int)
	for _, key := range collidingKeys {
		h := ht.HashFun(key)
		hashCounts[h]++
	}

	// Count max collisions
	maxCollisions := 0
	for _, count := range hashCounts {
		if count > maxCollisions {
			maxCollisions = count
		}
	}

	t.Logf("With salt: max collisions = %d (out of %d keys)", maxCollisions, len(collidingKeys))
	t.Logf("Hash distribution: %v", hashCounts)

	// With good salt, collisions should be spread out
	// If all keys still collide, salt is not working
	if maxCollisions == len(collidingKeys) && len(collidingKeys) > 2 {
		t.Logf("Warning: all keys still collide - salt may not be effective for this key set")
	}
}

// TestSalt_DifferentTables verifies different tables have different salts
func TestSalt_DifferentTables(t *testing.T) {
	ht1 := Init(17, 3)
	ht2 := Init(17, 3)

	// Same key should hash differently in different tables (different salts)
	key := "test_key"
	hash1 := ht1.HashFun(key)
	hash2 := ht2.HashFun(key)

	// Note: there's a 1/17 chance they're equal by coincidence
	t.Logf("Table 1 salt: %d, hash(%q) = %d", ht1.salt, key, hash1)
	t.Logf("Table 2 salt: %d, hash(%q) = %d", ht2.salt, key, hash2)

	if ht1.salt == ht2.salt {
		t.Errorf("Different tables should have different salts")
	}
}

// TestSalt_ConsistentWithinTable verifies same key hashes consistently within one table
func TestSalt_ConsistentWithinTable(t *testing.T) {
	ht := Init(17, 3)
	key := "consistent_key"

	hash1 := ht.HashFun(key)
	hash2 := ht.HashFun(key)

	if hash1 != hash2 {
		t.Errorf("Same key should hash to same value within same table: %d != %d", hash1, hash2)
	}
}

// BenchmarkHashDoS_WithoutSalt simulates attack on predictable hash
func BenchmarkHashDoS_Salted(b *testing.B) {
	ht := Init(1009, 3) // prime size
	// Even with "colliding" keys, salt randomizes placement
	keys := generateCollidingKeys(100, 5, 1009)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			ht.Put(key)
		}
	}
}


// dynamicHashTable

package hashtable

import (
	"strconv"
	"testing"
)

func TestDynamicHashTableInit(t *testing.T) {
	ht := InitDynamic(17, 3)
	if ht.size != 17 {
		t.Errorf("size is %d, expected 17", ht.size)
	}
	if ht.step != 3 {
		t.Errorf("step is %d, expected 3", ht.step)
	}
	if ht.count != 0 {
		t.Errorf("count is %d, expected 0", ht.count)
	}
	if len(ht.slots) != 17 {
		t.Errorf("slots length is %d, expected 17", len(ht.slots))
	}
}

func TestDynamicHashTableHashFunc(t *testing.T) {
	ht := InitDynamic(17, 3)
	idx := ht.HashFun("hello")
	if idx < 0 || idx >= ht.size {
		t.Errorf("HashFun returned %d, expected 0..%d", idx, ht.size-1)
	}
}

func TestDynamicHashTablePut(t *testing.T) {
	ht := InitDynamic(17, 3)
	idx := ht.Put("hello")
	if idx < 0 || idx >= ht.size {
		t.Errorf("Put returned %d, expected valid index", idx)
	}
	if ht.count != 1 {
		t.Errorf("count is %d, expected 1", ht.count)
	}
	if ht.slots[idx] == nil || *ht.slots[idx] != "hello" {
		t.Errorf("value not stored correctly")
	}
}

func TestDynamicHashTableFind(t *testing.T) {
	ht := InitDynamic(17, 3)
	ht.Put("hello")
	ht.Put("world")

	idx := ht.Find("hello")
	if idx == -1 {
		t.Errorf("Find returned -1 for existing value")
	}

	idx = ht.Find("notexists")
	if idx != -1 {
		t.Errorf("Find returned %d for non-existing value, expected -1", idx)
	}
}

func TestDynamicHashTableResize(t *testing.T) {
	ht := InitDynamic(5, 1)
	initialSize := ht.size

	// fill to trigger resize (load factor > 0.7)
	values := []string{"a", "b", "c", "d"}
	for _, v := range values {
		ht.Put(v)
	}

	if ht.size <= initialSize {
		t.Errorf("size did not increase: got %d, initial was %d", ht.size, initialSize)
	}

	// verify all values still findable after resize
	for _, v := range values {
		if ht.Find(v) == -1 {
			t.Errorf("value %q not found after resize", v)
		}
	}
}

func TestDynamicHashTableResizePreservesCount(t *testing.T) {
	ht := InitDynamic(5, 1)

	values := []string{"one", "two", "three", "four"}
	for _, v := range values {
		ht.Put(v)
	}

	if ht.count != len(values) {
		t.Errorf("count is %d, expected %d", ht.count, len(values))
	}
}

func TestDynamicHashTableMultipleResizes(t *testing.T) {
	ht := InitDynamic(3, 1)

	// insert many values to trigger multiple resizes
	for i := 0; i < 50; i++ {
		v := "value" + strconv.Itoa(i)
		ht.Put(v)
	}

	if ht.count != 50 {
		t.Errorf("count is %d, expected 50", ht.count)
	}

	// verify all values findable
	for i := 0; i < 50; i++ {
		v := "value" + strconv.Itoa(i)
		if ht.Find(v) == -1 {
			t.Errorf("value %q not found after multiple resizes", v)
		}
	}
}

func TestDynamicHashTableNeverFull(t *testing.T) {
	ht := InitDynamic(3, 1)

	// should never return -1, always resizes
	for i := 0; i < 100; i++ {
		v := "v" + strconv.Itoa(i)
		idx := ht.Put(v)
		if idx == -1 {
			t.Errorf("Put returned -1 for value %q, dynamic table should never be full", v)
		}
	}
}

func TestDynamicHashTableLoadFactor(t *testing.T) {
	ht := InitDynamic(10, 1)

	// insert 7 values (70% load)
	for i := 0; i < 7; i++ {
		ht.Put("v" + strconv.Itoa(i))
	}

	// 8th insert should trigger resize
	initialSize := ht.size
	ht.Put("trigger")

	if ht.size == initialSize {
		t.Errorf("expected resize at >70%% load, size unchanged: %d", ht.size)
	}
}

func TestNextPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1, 2},
		{2, 2},
		{3, 3},
		{4, 5},
		{10, 11},
		{17, 17},
		{18, 19},
		{34, 37},
	}

	for _, tc := range tests {
		got := nextPrime(tc.input)
		if got != tc.expected {
			t.Errorf("nextPrime(%d) = %d, expected %d", tc.input, got, tc.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18}

	for _, p := range primes {
		if !isPrime(p) {
			t.Errorf("isPrime(%d) = false, expected true", p)
		}
	}

	for _, np := range notPrimes {
		if isPrime(np) {
			t.Errorf("isPrime(%d) = true, expected false", np)
		}
	}
}

// semiFunchashtable

package hashtable

import (
	"strconv"
	"testing"
)

func TestMultiHashTableInit(t *testing.T) {
	ht := InitMultiHash(17)
	if ht.size != 17 {
		t.Errorf("size is %d, expected 17", ht.size)
	}
	if ht.count != 0 {
		t.Errorf("count is %d, expected 0", ht.count)
	}
	if len(ht.slots) != 17 {
		t.Errorf("slots length is %d, expected 17", len(ht.slots))
	}
	if len(ht.primes) != 3 {
		t.Errorf("primes length is %d, expected 3", len(ht.primes))
	}
}

func TestMultiHashTableHashN(t *testing.T) {
	ht := InitMultiHash(17)
	value := "hello"

	// different hash functions should (usually) give different results
	h0 := ht.hashN(value, 0)
	h1 := ht.hashN(value, 1)
	h2 := ht.hashN(value, 2)

	// all should be valid indices
	for i, h := range []int{h0, h1, h2} {
		if h < 0 || h >= ht.size {
			t.Errorf("hashN(%q, %d) = %d, out of range [0, %d)", value, i, h, ht.size)
		}
	}

	// at least two should be different (high probability)
	if h0 == h1 && h1 == h2 {
		t.Logf("Warning: all hashes equal for %q: %d, %d, %d", value, h0, h1, h2)
	}
}

func TestMultiHashTablePut(t *testing.T) {
	ht := InitMultiHash(17)
	idx := ht.Put("hello")
	if idx < 0 || idx >= ht.size {
		t.Errorf("Put returned %d, expected valid index", idx)
	}
	if ht.count != 1 {
		t.Errorf("count is %d, expected 1", ht.count)
	}
	if ht.slots[idx] == nil || *ht.slots[idx] != "hello" {
		t.Errorf("value not stored correctly")
	}
}

func TestMultiHashTableFind(t *testing.T) {
	ht := InitMultiHash(17)
	ht.Put("hello")
	ht.Put("world")

	idx := ht.Find("hello")
	if idx == -1 {
		t.Errorf("Find returned -1 for existing value")
	}

	idx = ht.Find("notexists")
	if idx != -1 {
		t.Errorf("Find returned %d for non-existing value, expected -1", idx)
	}
}

func TestMultiHashTableFindAfterCollision(t *testing.T) {
	ht := InitMultiHash(17)

	// insert many values to create collisions
	values := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for _, v := range values {
		ht.Put(v)
	}

	// all should be findable
	for _, v := range values {
		if ht.Find(v) == -1 {
			t.Errorf("value %q not found after collisions", v)
		}
	}
}

func TestMultiHashTableFullTable(t *testing.T) {
	ht := InitMultiHash(5)
	values := []string{"a", "b", "c", "d", "e"}

	for _, v := range values {
		idx := ht.Put(v)
		if idx == -1 {
			t.Errorf("Put(%q) returned -1 before table full", v)
		}
	}

	// table full
	idx := ht.Put("overflow")
	if idx != -1 {
		t.Errorf("Put on full table returned %d, expected -1", idx)
	}
}

func TestMultiHashTableStats(t *testing.T) {
	ht := InitMultiHash(101) // larger table for better stats

	// insert values
	for i := 0; i < 50; i++ {
		ht.Put("value" + strconv.Itoa(i))
	}

	primary, secondary, probing := ht.Stats()
	total := primary + secondary + probing

	if total != 50 {
		t.Errorf("stats total is %d, expected 50", total)
	}

	t.Logf("Stats: primary=%d (%.1f%%), secondary=%d (%.1f%%), probing=%d (%.1f%%)",
		primary, float64(primary)*100/float64(total),
		secondary, float64(secondary)*100/float64(total),
		probing, float64(probing)*100/float64(total))
}

func TestMultiHashTableVsSingleHash(t *testing.T) {
	// Compare collision rates between single and multi hash

	size := 101
	numValues := 70 // ~70% load

	// Multi hash table
	multi := InitMultiHash(size)
	for i := 0; i < numValues; i++ {
		multi.Put("v" + strconv.Itoa(i))
	}
	primaryM, secondaryM, probingM := multi.Stats()

	// Single hash table (simulated - count how many would need probing)
	single := Init(size, 1)
	probingS := 0
	for i := 0; i < numValues; i++ {
		v := "v" + strconv.Itoa(i)
		hashIdx := single.HashFun(v)
		if single.slots[hashIdx] != nil {
			probingS++
		}
		single.Put(v)
	}

	t.Logf("Multi-hash: primary=%d, secondary=%d, probing=%d",
		primaryM, secondaryM, probingM)
	t.Logf("Single-hash: direct=%d, probing=%d",
		numValues-probingS, probingS)

	// Multi-hash should have fewer probing hits (more options before fallback)
	multiProbingRate := float64(probingM) / float64(numValues)
	singleProbingRate := float64(probingS) / float64(numValues)

	t.Logf("Probing rate: multi=%.1f%%, single=%.1f%%",
		multiProbingRate*100, singleProbingRate*100)
}

func BenchmarkMultiHashPut(b *testing.B) {
	ht := InitMultiHash(10007) // prime size
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ht.Put("value" + strconv.Itoa(i%5000))
	}
}

func BenchmarkSingleHashPut(b *testing.B) {
	ht := Init(10007, 3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ht.Put("value" + strconv.Itoa(i%5000))
	}
}

func BenchmarkMultiHashFind(b *testing.B) {
	ht := InitMultiHash(10007)
	for i := 0; i < 5000; i++ {
		ht.Put("value" + strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ht.Find("value" + strconv.Itoa(i%5000))
	}
}

func BenchmarkSingleHashFind(b *testing.B) {
	ht := Init(10007, 3)
	for i := 0; i < 5000; i++ {
		ht.Put("value" + strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ht.Find("value" + strconv.Itoa(i%5000))
	}
}
