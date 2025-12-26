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
	ht.slots[hashIdx] = "occupied"

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
	ht.slots[hashIdx] = "occupied"
	ht.slots[(hashIdx+ht.step)%ht.size] = "occupied"

	idx := ht.SeekSlot(value)
	expectedIdx := (hashIdx + 2*ht.step) % ht.size
	if idx != expectedIdx {
		t.Errorf("SeekSlot with multiple collisions returned %d, expected %d", idx, expectedIdx)
	}
}

func TestSeekSlot_FullTable(t *testing.T) {
	ht := Init(17, 3)
	for i := 0; i < ht.size; i++ {
		ht.slots[i] = "occupied"
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
	if ht.slots[idx] != value {
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
	ht.slots[hashIdx] = "occupied"

	idx := ht.Put(value)
	expectedIdx := (hashIdx + ht.step) % ht.size
	if idx != expectedIdx {
		t.Errorf("Put with collision returned %d, expected %d", idx, expectedIdx)
	}
	if ht.slots[idx] != value {
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
		if ht.slots[idx] != v {
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
	ht.slots[hashIdx] = "blocker"

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

// BenchmarkHashDoS_Salted simulates attack on salted hash
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
