package lfucache

import (
	"testing"
)

func TestInitNativeCache(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	if nc.size != 17 {
		t.Errorf("size is %d, expected 17", nc.size)
	}
	if nc.step != 3 {
		t.Errorf("step is %d, expected 3", nc.step)
	}
	if len(nc.slots) != 17 {
		t.Errorf("slots len is %d, expected 17", len(nc.slots))
	}
	if len(nc.values) != 17 {
		t.Errorf("values len is %d, expected 17", len(nc.values))
	}
	if len(nc.hits) != 17 {
		t.Errorf("hits len is %d, expected 17", len(nc.hits))
	}
}

func TestNativeCache_HashFun(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	idx := nc.HashFun("test")
	if idx < 0 || idx >= nc.size {
		t.Errorf("HashFun returned %d, expected 0..%d", idx, nc.size-1)
	}
}

func TestNativeCache_HashFun_Consistent(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	key := "consistent_key"
	hash1 := nc.HashFun(key)
	hash2 := nc.HashFun(key)
	if hash1 != hash2 {
		t.Errorf("same key should hash to same value: %d != %d", hash1, hash2)
	}
}

func TestNativeCache_Put_NewKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("key1", 100)

	val, err := nc.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 100 {
		t.Errorf("Get returned %d, expected 100", val)
	}
}

func TestNativeCache_Put_UpdateExistingKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("key1", 100)
	nc.Put("key1", 200)

	val, err := nc.Get("key1")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 200 {
		t.Errorf("Get returned %d, expected 200 after update", val)
	}
}

func TestNativeCache_Put_MultipleKeys(t *testing.T) {
	nc := InitNativeCache[string](17, 3)
	nc.Put("name", "Alice")
	nc.Put("city", "Berlin")
	nc.Put("lang", "Go")

	tests := []struct {
		key      string
		expected string
	}{
		{"name", "Alice"},
		{"city", "Berlin"},
		{"lang", "Go"},
	}

	for _, tt := range tests {
		val, err := nc.Get(tt.key)
		if err != nil {
			t.Errorf("Get(%q) returned error: %v", tt.key, err)
		}
		if val != tt.expected {
			t.Errorf("Get(%q) = %q, expected %q", tt.key, val, tt.expected)
		}
	}
}

func TestNativeCache_Get_NonExistingKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("exists", 42)

	_, err := nc.Get("missing")
	if err != ErrKeyNotFound {
		t.Errorf("Get should return ErrKeyNotFound, got: %v", err)
	}
}

func TestNativeCache_Get_EmptyCache(t *testing.T) {
	nc := InitNativeCache[int](17, 3)

	_, err := nc.Get("any")
	if err != ErrKeyNotFound {
		t.Errorf("Get on empty cache should return ErrKeyNotFound, got: %v", err)
	}
}

func TestNativeCache_IsKey_ExistingKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("exists", 42)

	if !nc.IsKey("exists") {
		t.Error("IsKey should return true for existing key")
	}
}

func TestNativeCache_IsKey_NonExistingKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("exists", 42)

	if nc.IsKey("not_exists") {
		t.Error("IsKey should return false for non-existing key")
	}
}

func TestNativeCache_IsKey_EmptyCache(t *testing.T) {
	nc := InitNativeCache[int](17, 3)

	if nc.IsKey("any") {
		t.Error("IsKey should return false on empty cache")
	}
}

func TestNativeCache_HitsCounter_IncrementsOnGet(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("key", 100)

	nc.Get("key")
	nc.Get("key")
	nc.Get("key")

	hits := nc.GetHits("key")
	if hits != 3 {
		t.Errorf("hits is %d, expected 3 after 3 Gets", hits)
	}
}

func TestNativeCache_HitsCounter_TracksAccessesCorrectly(t *testing.T) {
	nc := InitNativeCache[int](17, 3)

	nc.Put("a", 1)
	nc.Put("b", 2)
	nc.Put("c", 3)

	for i := 0; i < 10; i++ {
		nc.Get("a")
	}
	for i := 0; i < 5; i++ {
		nc.Get("b")
	}
	for i := 0; i < 2; i++ {
		nc.Get("c")
	}

	hitsA := nc.GetHits("a")
	hitsB := nc.GetHits("b")
	hitsC := nc.GetHits("c")

	if hitsA != 10 {
		t.Errorf("a.hits is %d, expected 10", hitsA)
	}
	if hitsB != 5 {
		t.Errorf("b.hits is %d, expected 5", hitsB)
	}
	if hitsC != 2 {
		t.Errorf("c.hits is %d, expected 2", hitsC)
	}
}

func TestNativeCache_FullCache_EvictionOnOverflow(t *testing.T) {
	nc := InitNativeCache[int](5, 1)

	nc.Put("k1", 1)
	nc.Put("k2", 2)
	nc.Put("k3", 3)
	nc.Put("k4", 4)
	nc.Put("k5", 5)

	nc.Get("k2")
	nc.Get("k3")
	nc.Get("k4")
	nc.Get("k5")

	nc.Put("k6", 6)

	if nc.IsKey("k1") {
		t.Error("k1 should be evicted (0 hits, minimum)")
	}

	for _, key := range []string{"k2", "k3", "k4", "k5", "k6"} {
		if !nc.IsKey(key) {
			t.Errorf("%s should still exist", key)
		}
	}
}

func TestNativeCache_FullCache_EvictsByMinHits(t *testing.T) {
	nc := InitNativeCache[int](5, 1)

	nc.Put("rare", 1)
	nc.Put("common", 2)
	nc.Put("popular", 3)
	nc.Put("viral", 4)
	nc.Put("legendary", 5)

	nc.Get("common")
	nc.Get("popular")
	nc.Get("popular")
	nc.Get("viral")
	nc.Get("viral")
	nc.Get("viral")
	nc.Get("legendary")
	nc.Get("legendary")
	nc.Get("legendary")
	nc.Get("legendary")

	nc.Put("newcomer", 6)

	if nc.IsKey("rare") {
		t.Error("rare should be evicted (0 hits, minimum)")
	}

	if !nc.IsKey("common") {
		t.Error("common should exist (1 hit)")
	}
	if !nc.IsKey("popular") {
		t.Error("popular should exist (2 hits)")
	}
	if !nc.IsKey("viral") {
		t.Error("viral should exist (3 hits)")
	}
	if !nc.IsKey("legendary") {
		t.Error("legendary should exist (4 hits)")
	}
	if !nc.IsKey("newcomer") {
		t.Error("newcomer should exist")
	}
}

func TestNativeCache_FullCache_EvictsMinHitsElement(t *testing.T) {
	nc := InitNativeCache[int](3, 1)

	nc.Put("hot", 1)
	nc.Put("warm", 2)
	nc.Put("cold", 3)

	for i := 0; i < 100; i++ {
		nc.Get("hot")
	}
	for i := 0; i < 50; i++ {
		nc.Get("warm")
	}

	nc.Put("new", 4)

	if nc.IsKey("cold") {
		t.Error("cold should be evicted (0 hits, minimum)")
	}

	hotHits := nc.GetHits("hot")
	warmHits := nc.GetHits("warm")

	if hotHits != 100 {
		t.Errorf("hot.hits is %d, expected 100", hotHits)
	}
	if warmHits != 50 {
		t.Errorf("warm.hits is %d, expected 50", warmHits)
	}
}

func TestNativeCache_FullCache_HitsResetOnEviction(t *testing.T) {
	nc := InitNativeCache[int](2, 1)

	nc.Put("old", 1)
	nc.Get("old")
	nc.Get("old")
	nc.Get("old")

	nc.Put("newer", 2)

	nc.Put("newest", 3)

	if nc.IsKey("newer") {
		t.Error("newer should be evicted (0 hits)")
	}

	newestHits := nc.GetHits("newest")
	if newestHits != 0 {
		t.Errorf("newest.hits is %d, expected 0 (fresh entry)", newestHits)
	}
}

func TestNativeCache_Collision_PutAndGet(t *testing.T) {
	nc := InitNativeCache[int](5, 1)

	nc.Put("a", 1)
	nc.Put("f", 2)
	nc.Put("k", 3)

	val1, err := nc.Get("a")
	if err != nil || val1 != 1 {
		t.Errorf("Get(a) = %d, err=%v, expected 1", val1, err)
	}

	val2, err := nc.Get("f")
	if err != nil || val2 != 2 {
		t.Errorf("Get(f) = %d, err=%v, expected 2", val2, err)
	}

	val3, err := nc.Get("k")
	if err != nil || val3 != 3 {
		t.Errorf("Get(k) = %d, err=%v, expected 3", val3, err)
	}
}

func TestNativeCache_ManyCollisions_EvictsCorrectly(t *testing.T) {
	nc := InitNativeCache[int](5, 1)

	nc.Put("a", 1)
	nc.Put("f", 2)
	nc.Put("k", 3)
	nc.Put("p", 4)
	nc.Put("u", 5)

	nc.Get("a")
	nc.Get("a")
	nc.Get("f")
	nc.Get("k")
	nc.Get("p")

	nc.Put("z", 6)

	if nc.IsKey("u") {
		t.Error("u should be evicted (0 hits, minimum)")
	}

	if !nc.IsKey("a") {
		t.Error("a should exist (2 hits)")
	}
	if !nc.IsKey("f") {
		t.Error("f should exist (1 hit)")
	}
	if !nc.IsKey("k") {
		t.Error("k should exist (1 hit)")
	}
	if !nc.IsKey("p") {
		t.Error("p should exist (1 hit)")
	}
}

func TestNativeCache_EmptyStringKey(t *testing.T) {
	nc := InitNativeCache[int](17, 3)
	nc.Put("", 123)

	val, err := nc.Get("")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if val != 123 {
		t.Errorf("Get returned %d, expected 123", val)
	}
}

func TestNativeCache_GenericTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	nc := InitNativeCache[Person](17, 3)
	nc.Put("alice", Person{Name: "Alice", Age: 30})
	nc.Put("bob", Person{Name: "Bob", Age: 25})

	alice, err := nc.Get("alice")
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	}
	if alice.Name != "Alice" || alice.Age != 30 {
		t.Errorf("Got %+v, expected Alice/30", alice)
	}
}
