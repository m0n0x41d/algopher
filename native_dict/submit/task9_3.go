package main

import (
	"testing"
)

func TestInit(t *testing.T) {
	nd := Init[int](17)
	if nd.size != 17 {
		t.Errorf("size is %d, expected 17", nd.size)
	}
	if len(nd.slots) != 17 {
		t.Errorf("slots len is %d, expected 17", len(nd.slots))
	}
	if len(nd.values) != 17 {
		t.Errorf("values len is %d, expected 17", len(nd.values))
	}
	if len(nd.filled) != 17 {
		t.Errorf("filled len is %d, expected 17", len(nd.filled))
	}
}

func TestHashFun(t *testing.T) {
	nd := Init[string](17)
	idx := nd.HashFun("test-key")
	if idx < 0 || idx >= nd.size {
		t.Errorf("HashFun returned %d, expected 0..%d", idx, nd.size-1)
	}
}

func TestHashFun_Consistent(t *testing.T) {
	nd := Init[int](17)
	key := "consistent_key"
	hash1 := nd.HashFun(key)
	hash2 := nd.HashFun(key)
	if hash1 != hash2 {
		t.Errorf("Same key should hash to same value: %d != %d", hash1, hash2)
	}
}

func TestHashFun_DifferentSalts(t *testing.T) {
	nd1 := Init[int](17)
	nd2 := Init[int](17)
	if nd1.salt == nd2.salt {
		t.Errorf("Different tables should have different salts")
	}
}

// === Put tests ===

func TestPut_NewKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("key1", 100)

	if !nd.IsKey("key1") {
		t.Error("key1 should exist after Put")
	}
	val, err := nd.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 100 {
		t.Errorf("Get returned %d, expected 100", val)
	}
}

func TestPut_UpdateExistingKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("key1", 100)
	nd.Put("key1", 200)

	val, err := nd.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 200 {
		t.Errorf("Get returned %d, expected 200 after update", val)
	}
}

func TestPut_MultipleKeys(t *testing.T) {
	nd := Init[string](17)
	nd.Put("name", "Alice")
	nd.Put("city", "Berlin")
	nd.Put("lang", "Go")

	tests := []struct {
		key      string
		expected string
	}{
		{"name", "Alice"},
		{"city", "Berlin"},
		{"lang", "Go"},
	}

	for _, tt := range tests {
		val, err := nd.Get(tt.key)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", tt.key, err)
		}
		if val != tt.expected {
			t.Errorf("Get(%q) = %q, expected %q", tt.key, val, tt.expected)
		}
	}
}

func TestPut_WithCollision(t *testing.T) {
	nd := Init[int](17)

	// fill many keys to force collisions
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		nd.Put(key, i*10)
	}

	// verify all values are retrievable
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		val, err := nd.Get(key)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", key, err)
		}
		if val != i*10 {
			t.Errorf("Get(%q) = %d, expected %d", key, val, i*10)
		}
	}
}

// === IsKey tests ===

func TestIsKey_ExistingKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("exists", 42)

	if !nd.IsKey("exists") {
		t.Error("IsKey should return true for existing key")
	}
}

func TestIsKey_NonExistingKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("exists", 42)

	if nd.IsKey("not_exists") {
		t.Error("IsKey should return false for non-existing key")
	}
}

func TestIsKey_EmptyTable(t *testing.T) {
	nd := Init[int](17)

	if nd.IsKey("any") {
		t.Error("IsKey should return false on empty table")
	}
}

// === Get tests ===

func TestGet_ExistingKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("key", 999)

	val, err := nd.Get("key")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 999 {
		t.Errorf("Get returned %d, expected 999", val)
	}
}

func TestGet_NonExistingKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("key", 999)

	_, err := nd.Get("missing")
	if err != ErrKeyNotFound {
		t.Errorf("Get should return ErrKeyNotFound, got: %v", err)
	}
}

func TestGet_EmptyTable(t *testing.T) {
	nd := Init[int](17)

	_, err := nd.Get("any")
	if err != ErrKeyNotFound {
		t.Errorf("Get on empty table should return ErrKeyNotFound, got: %v", err)
	}
}

func TestGet_AfterUpdate(t *testing.T) {
	nd := Init[string](17)
	nd.Put("status", "pending")
	nd.Put("status", "done")

	val, err := nd.Get("status")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != "done" {
		t.Errorf("Get returned %q, expected 'done'", val)
	}
}

// === Edge cases ===

func TestEmptyStringKey(t *testing.T) {
	nd := Init[int](17)
	nd.Put("", 123)

	if !nd.IsKey("") {
		t.Error("Empty string should be valid key")
	}
	val, err := nd.Get("")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 123 {
		t.Errorf("Get returned %d, expected 123", val)
	}
}

func TestFullTable(t *testing.T) {
	nd := Init[int](5)

	for i := 0; i < 5; i++ {
		nd.Put(string(rune('a'+i)), i)
	}

	// all 5 keys should be retrievable
	for i := 0; i < 5; i++ {
		key := string(rune('a' + i))
		if !nd.IsKey(key) {
			t.Errorf("key %q should exist", key)
		}
	}

	// adding to full table - silently ignored
	nd.Put("overflow", 999)
	if nd.IsKey("overflow") {
		t.Error("overflow key should not exist in full table")
	}
}

func TestGenericTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	nd := Init[Person](17)
	nd.Put("alice", Person{Name: "Alice", Age: 30})
	nd.Put("bob", Person{Name: "Bob", Age: 25})

	alice, err := nd.Get("alice")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if alice.Name != "Alice" || alice.Age != 30 {
		t.Errorf("Got %+v, expected Alice/30", alice)
	}
}

// === OrderedDictionary tests ===

func TestOrderedInit(t *testing.T) {
	od := InitOrdered[int](17)
	if od.count != 0 {
		t.Errorf("count is %d, expected 0", od.count)
	}
	if od.Count() != 0 {
		t.Errorf("Count() is %d, expected 0", od.Count())
	}
}

func TestOrderedPut_NewKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key1", 100)

	if !od.IsKey("key1") {
		t.Error("key1 should exist after Put")
	}
	val, err := od.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 100 {
		t.Errorf("Get returned %d, expected 100", val)
	}
	if od.Count() != 1 {
		t.Errorf("Count() is %d, expected 1", od.Count())
	}
}

func TestOrderedPut_UpdateExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key1", 100)
	od.Put("key1", 200)

	val, err := od.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 200 {
		t.Errorf("Get returned %d, expected 200 after update", val)
	}
	if od.Count() != 1 {
		t.Errorf("Count() is %d, expected 1 (update should not increase count)", od.Count())
	}
}

func TestOrderedPut_MultipleKeys(t *testing.T) {
	od := InitOrdered[string](17)
	od.Put("name", "Alice")
	od.Put("city", "Berlin")
	od.Put("lang", "Go")

	tests := []struct {
		key      string
		expected string
	}{
		{"name", "Alice"},
		{"city", "Berlin"},
		{"lang", "Go"},
	}

	for _, tt := range tests {
		val, err := od.Get(tt.key)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", tt.key, err)
		}
		if val != tt.expected {
			t.Errorf("Get(%q) = %q, expected %q", tt.key, val, tt.expected)
		}
	}
}

func TestOrderedPut_MaintainsOrder(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("delta", 4)
	od.Put("alpha", 1)
	od.Put("charlie", 3)
	od.Put("bravo", 2)

	expected := []string{"alpha", "bravo", "charlie", "delta"}
	for i, key := range expected {
		if od.keys[i] != key {
			t.Errorf("keys[%d] = %q, expected %q", i, od.keys[i], key)
		}
	}
}

func TestOrderedIsKey_ExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("exists", 42)

	if !od.IsKey("exists") {
		t.Error("IsKey should return true for existing key")
	}
}

func TestOrderedIsKey_NonExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("exists", 42)

	if od.IsKey("not_exists") {
		t.Error("IsKey should return false for non-existing key")
	}
}

func TestOrderedIsKey_EmptyDict(t *testing.T) {
	od := InitOrdered[int](17)

	if od.IsKey("any") {
		t.Error("IsKey should return false on empty dict")
	}
}

func TestOrderedGet_ExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key", 999)

	val, err := od.Get("key")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 999 {
		t.Errorf("Get returned %d, expected 999", val)
	}
}

func TestOrderedGet_NonExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key", 999)

	_, err := od.Get("missing")
	if err != ErrKeyNotFound {
		t.Errorf("Get should return ErrKeyNotFound, got: %v", err)
	}
}

func TestOrderedGet_EmptyDict(t *testing.T) {
	od := InitOrdered[int](17)

	_, err := od.Get("any")
	if err != ErrKeyNotFound {
		t.Errorf("Get on empty dict should return ErrKeyNotFound, got: %v", err)
	}
}

func TestOrderedGet_AfterUpdate(t *testing.T) {
	od := InitOrdered[string](17)
	od.Put("status", "pending")
	od.Put("status", "done")

	val, err := od.Get("status")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != "done" {
		t.Errorf("Get returned %q, expected 'done'", val)
	}
}

func TestOrderedDelete_ExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key", 42)

	deleted := od.Delete("key")
	if !deleted {
		t.Error("Delete should return true for existing key")
	}
	if od.IsKey("key") {
		t.Error("key should not exist after Delete")
	}
	if od.Count() != 0 {
		t.Errorf("Count() is %d, expected 0 after Delete", od.Count())
	}
}

func TestOrderedDelete_NonExistingKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("key", 42)

	deleted := od.Delete("missing")
	if deleted {
		t.Error("Delete should return false for non-existing key")
	}
	if od.Count() != 1 {
		t.Errorf("Count() is %d, expected 1", od.Count())
	}
}

func TestOrderedDelete_MaintainsOrder(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("alpha", 1)
	od.Put("bravo", 2)
	od.Put("charlie", 3)
	od.Put("delta", 4)

	od.Delete("bravo")

	expected := []string{"alpha", "charlie", "delta"}
	if len(od.keys) != 3 {
		t.Errorf("len(keys) = %d, expected 3", len(od.keys))
	}
	for i, key := range expected {
		if od.keys[i] != key {
			t.Errorf("keys[%d] = %q, expected %q", i, od.keys[i], key)
		}
	}
}

func TestOrderedEmptyStringKey(t *testing.T) {
	od := InitOrdered[int](17)
	od.Put("", 123)

	if !od.IsKey("") {
		t.Error("Empty string should be valid key")
	}
	val, err := od.Get("")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 123 {
		t.Errorf("Get returned %d, expected 123", val)
	}
}

func TestOrderedGenericTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	od := InitOrdered[Person](17)
	od.Put("alice", Person{Name: "Alice", Age: 30})
	od.Put("bob", Person{Name: "Bob", Age: 25})

	alice, err := od.Get("alice")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if alice.Name != "Alice" || alice.Age != 30 {
		t.Errorf("Got %+v, expected Alice/30", alice)
	}
}

func TestOrderedManyKeys(t *testing.T) {
	od := InitOrdered[int](100)

	for i := 0; i < 100; i++ {
		key := string(rune('a' + (i % 26)))
		key += string(rune('0' + (i / 26)))
		od.Put(key, i)
	}

	if od.Count() != 100 {
		t.Errorf("Count() is %d, expected 100", od.Count())
	}

	for i := 1; i < len(od.keys); i++ {
		if od.keys[i-1] >= od.keys[i] {
			t.Errorf("keys not sorted: keys[%d]=%q >= keys[%d]=%q", i-1, od.keys[i-1], i, od.keys[i])
		}
	}
}

// === BitKeyDictionary tests ===

func TestBitKeyInit(t *testing.T) {
	bd := InitBitKey[int](17, 0xDEADBEEF)
	if bd.size != 17 {
		t.Errorf("size is %d, expected 17", bd.size)
	}
	if len(bd.keys) != 17 {
		t.Errorf("keys len is %d, expected 17", len(bd.keys))
	}
	if len(bd.values) != 17 {
		t.Errorf("values len is %d, expected 17", len(bd.values))
	}
	if bd.salt != 0xDEADBEEF {
		t.Errorf("salt is %d, expected 0xDEADBEEF", bd.salt)
	}
}

func TestBitKeyHashFun(t *testing.T) {
	bd := InitBitKey[int](17, 0x12345678)
	idx := bd.HashFun(0xCAFEBABE)
	if idx < 0 || idx >= bd.size {
		t.Errorf("HashFun returned %d, expected 0..%d", idx, bd.size-1)
	}
}

func TestBitKeyHashFun_Consistent(t *testing.T) {
	bd := InitBitKey[int](17, 0x12345678)
	key := uint64(0xCAFEBABE)
	hash1 := bd.HashFun(key)
	hash2 := bd.HashFun(key)
	if hash1 != hash2 {
		t.Errorf("Same key should hash to same value: %d != %d", hash1, hash2)
	}
}

func TestBitKeyHashFun_DifferentSalts(t *testing.T) {
	bd1 := InitBitKey[int](17, 0x11111111)
	bd2 := InitBitKey[int](17, 0x22222222)
	key := uint64(0xCAFEBABE)
	hash1 := bd1.HashFun(key)
	hash2 := bd2.HashFun(key)
	if hash1 == hash2 {
		t.Error("Different salts should produce different hashes")
	}
}

func TestBitKeyPut_NewKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0x1234, 100)

	if !bd.IsKey(0x1234) {
		t.Error("key should exist after Put")
	}
	val, err := bd.Get(0x1234)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 100 {
		t.Errorf("Get returned %d, expected 100", val)
	}
}

func TestBitKeyPut_UpdateExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0x1234, 100)
	bd.Put(0x1234, 200)

	val, err := bd.Get(0x1234)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 200 {
		t.Errorf("Get returned %d, expected 200 after update", val)
	}
}

func TestBitKeyPut_MultipleKeys(t *testing.T) {
	bd := InitBitKey[string](17, 0)
	bd.Put(0x0001, "one")
	bd.Put(0x0002, "two")
	bd.Put(0x0003, "three")

	tests := []struct {
		key      uint64
		expected string
	}{
		{0x0001, "one"},
		{0x0002, "two"},
		{0x0003, "three"},
	}

	for _, tt := range tests {
		val, err := bd.Get(tt.key)
		if err != nil {
			t.Errorf("Get(0x%X) returned error: %v", tt.key, err)
		}
		if val != tt.expected {
			t.Errorf("Get(0x%X) = %q, expected %q", tt.key, val, tt.expected)
		}
	}
}

func TestBitKeyPut_WithCollision(t *testing.T) {
	bd := InitBitKey[int](17, 0)

	for i := 0; i < 10; i++ {
		bd.Put(uint64(i), i*10)
	}

	for i := 0; i < 10; i++ {
		val, err := bd.Get(uint64(i))
		if err != nil {
			t.Errorf("Get(%d) returned error: %v", i, err)
		}
		if val != i*10 {
			t.Errorf("Get(%d) = %d, expected %d", i, val, i*10)
		}
	}
}

func TestBitKeyIsKey_ExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0xABCD, 42)

	if !bd.IsKey(0xABCD) {
		t.Error("IsKey should return true for existing key")
	}
}

func TestBitKeyIsKey_NonExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0xABCD, 42)

	if bd.IsKey(0xDCBA) {
		t.Error("IsKey should return false for non-existing key")
	}
}

func TestBitKeyIsKey_EmptyTable(t *testing.T) {
	bd := InitBitKey[int](17, 0)

	if bd.IsKey(0x1234) {
		t.Error("IsKey should return false on empty table")
	}
}

func TestBitKeyGet_ExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0xFFFF, 999)

	val, err := bd.Get(0xFFFF)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 999 {
		t.Errorf("Get returned %d, expected 999", val)
	}
}

func TestBitKeyGet_NonExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0xFFFF, 999)

	_, err := bd.Get(0x0000)
	if err != ErrKeyNotFound {
		t.Errorf("Get should return ErrKeyNotFound, got: %v", err)
	}
}

func TestBitKeyGet_EmptyTable(t *testing.T) {
	bd := InitBitKey[int](17, 0)

	_, err := bd.Get(0x1234)
	if err != ErrKeyNotFound {
		t.Errorf("Get on empty table should return ErrKeyNotFound, got: %v", err)
	}
}

func TestBitKeyGet_AfterUpdate(t *testing.T) {
	bd := InitBitKey[string](17, 0)
	bd.Put(0x0001, "pending")
	bd.Put(0x0001, "done")

	val, err := bd.Get(0x0001)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != "done" {
		t.Errorf("Get returned %q, expected 'done'", val)
	}
}

func TestBitKeyDelete_ExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0x1234, 42)

	deleted := bd.Delete(0x1234)
	if !deleted {
		t.Error("Delete should return true for existing key")
	}
	if bd.IsKey(0x1234) {
		t.Error("key should not exist after Delete")
	}
}

func TestBitKeyDelete_NonExistingKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0x1234, 42)

	deleted := bd.Delete(0x5678)
	if deleted {
		t.Error("Delete should return false for non-existing key")
	}
}

func TestBitKeyZeroKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	bd.Put(0x0000, 123)

	if !bd.IsKey(0x0000) {
		t.Error("Zero should be valid key")
	}
	val, err := bd.Get(0x0000)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 123 {
		t.Errorf("Get returned %d, expected 123", val)
	}
}

func TestBitKeyMaxKey(t *testing.T) {
	bd := InitBitKey[int](17, 0)
	maxKey := uint64(0xFFFFFFFFFFFFFFFF)
	bd.Put(maxKey, 999)

	if !bd.IsKey(maxKey) {
		t.Error("Max uint64 should be valid key")
	}
	val, err := bd.Get(maxKey)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 999 {
		t.Errorf("Get returned %d, expected 999", val)
	}
}

func TestBitKeyFullTable(t *testing.T) {
	bd := InitBitKey[int](5, 0)

	for i := 0; i < 5; i++ {
		bd.Put(uint64(i), i)
	}

	for i := 0; i < 5; i++ {
		if !bd.IsKey(uint64(i)) {
			t.Errorf("key %d should exist", i)
		}
	}

	bd.Put(0xFF, 999)
	if bd.IsKey(0xFF) {
		t.Error("overflow key should not exist in full table")
	}
}

func TestBitKeyGenericTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	bd := InitBitKey[Person](17, 0)
	bd.Put(0x0001, Person{Name: "Alice", Age: 30})
	bd.Put(0x0002, Person{Name: "Bob", Age: 25})

	alice, err := bd.Get(0x0001)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if alice.Name != "Alice" || alice.Age != 30 {
		t.Errorf("Got %+v, expected Alice/30", alice)
	}
}
