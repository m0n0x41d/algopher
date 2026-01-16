package cache

import (
	"testing"
)

func TestNewLFUCache(t *testing.T) {
	c := NewLFUCache(10)
	if c.capacity != 10 {
		t.Errorf("capacity is %d, expected 10", c.capacity)
	}
	if c.minFreq != 0 {
		t.Errorf("minFreq is %d, expected 0", c.minFreq)
	}
	if c.cache == nil {
		t.Error("cache map should be initialized")
	}
	if c.freqMap == nil {
		t.Error("freqMap should be initialized")
	}
}

func TestPut_NewKey(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("key1", 100)

	val, ok := c.Get("key1")
	if !ok {
		t.Error("key1 should exist after Put")
	}
	if val != 100 {
		t.Errorf("Get returned %v, expected 100", val)
	}
}

func TestPut_UpdateExistingKey(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("key1", 100)
	c.Put("key1", 200)

	val, ok := c.Get("key1")
	if !ok {
		t.Error("key1 should exist")
	}
	if val != 200 {
		t.Errorf("Get returned %v, expected 200 after update", val)
	}
}

func TestPut_MultipleKeys(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	tests := []struct {
		key      string
		expected int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	for _, tt := range tests {
		val, ok := c.Get(tt.key)
		if !ok {
			t.Errorf("key %q should exist", tt.key)
		}
		if val != tt.expected {
			t.Errorf("Get(%q) = %v, expected %d", tt.key, val, tt.expected)
		}
	}
}

func TestPut_ZeroCapacity(t *testing.T) {
	c := NewLFUCache(0)
	c.Put("key", 100)

	_, ok := c.Get("key")
	if ok {
		t.Error("zero capacity cache should not store anything")
	}
}

func TestGet_NonExistingKey(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("exists", 42)

	_, ok := c.Get("missing")
	if ok {
		t.Error("Get should return false for non-existing key")
	}
}

func TestGet_EmptyCache(t *testing.T) {
	c := NewLFUCache(10)

	_, ok := c.Get("any")
	if ok {
		t.Error("Get should return false on empty cache")
	}
}

func TestGet_IncrementsFrequency(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("key", 100)

	c.Get("key")
	c.Get("key")
	c.Get("key")

	elem := c.cache["key"]
	entry := elem.Value.(*cacheEntry)
	if entry.freq != 4 {
		t.Errorf("freq is %d, expected 4 (1 from Put + 3 from Get)", entry.freq)
	}
}

func TestEvict_LeastFrequentlyUsed(t *testing.T) {
	c := NewLFUCache(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	c.Get("a")
	c.Get("a")
	c.Get("b")

	c.Put("d", 4)

	_, ok := c.Get("c")
	if ok {
		t.Error("c should be evicted (least frequently used)")
	}

	if _, ok := c.Get("a"); !ok {
		t.Error("a should still exist")
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("b should still exist")
	}
	if _, ok := c.Get("d"); !ok {
		t.Error("d should exist")
	}
}

func TestEvict_LRUAmongSameFrequency(t *testing.T) {
	c := NewLFUCache(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	c.Put("d", 4)

	_, ok := c.Get("a")
	if ok {
		t.Error("a should be evicted (oldest among same frequency)")
	}

	if _, ok := c.Get("b"); !ok {
		t.Error("b should still exist")
	}
	if _, ok := c.Get("c"); !ok {
		t.Error("c should still exist")
	}
	if _, ok := c.Get("d"); !ok {
		t.Error("d should exist")
	}
}

func TestEvict_MultipleEvictions(t *testing.T) {
	c := NewLFUCache(2)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	c.Put("d", 4)

	if _, ok := c.Get("a"); ok {
		t.Error("a should be evicted")
	}
	if _, ok := c.Get("b"); ok {
		t.Error("b should be evicted")
	}
	if _, ok := c.Get("c"); !ok {
		t.Error("c should exist")
	}
	if _, ok := c.Get("d"); !ok {
		t.Error("d should exist")
	}
}

func TestMinFreq_UpdatesOnEviction(t *testing.T) {
	c := NewLFUCache(2)

	c.Put("a", 1)
	c.Get("a")
	c.Get("a")

	c.Put("b", 2)
	c.Get("b")

	c.Put("c", 3)

	if c.minFreq != 1 {
		t.Errorf("minFreq is %d, expected 1 after adding new element", c.minFreq)
	}
}

func TestMinFreq_UpdatesWhenBucketEmpty(t *testing.T) {
	c := NewLFUCache(2)

	c.Put("a", 1)
	c.Put("b", 2)

	c.Get("a")
	c.Get("b")

	if c.minFreq != 2 {
		t.Errorf("minFreq is %d, expected 2 after all elements incremented", c.minFreq)
	}
}

func TestCapacityOne(t *testing.T) {
	c := NewLFUCache(1)

	c.Put("a", 1)
	c.Put("b", 2)

	if _, ok := c.Get("a"); ok {
		t.Error("a should be evicted")
	}
	val, ok := c.Get("b")
	if !ok {
		t.Error("b should exist")
	}
	if val != 2 {
		t.Errorf("Get(b) = %v, expected 2", val)
	}
}

func TestUpdateExistingKey_DoesNotEvict(t *testing.T) {
	c := NewLFUCache(2)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 10)

	if _, ok := c.Get("a"); !ok {
		t.Error("a should exist after update")
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("b should exist (update should not evict)")
	}
}

func TestGenericValues(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	c := NewLFUCache(10)
	c.Put("alice", Person{Name: "Alice", Age: 30})
	c.Put("bob", Person{Name: "Bob", Age: 25})

	val, ok := c.Get("alice")
	if !ok {
		t.Error("alice should exist")
	}
	alice := val.(Person)
	if alice.Name != "Alice" || alice.Age != 30 {
		t.Errorf("Got %+v, expected Alice/30", alice)
	}
}

func TestEmptyStringKey(t *testing.T) {
	c := NewLFUCache(10)
	c.Put("", 123)

	val, ok := c.Get("")
	if !ok {
		t.Error("empty string should be valid key")
	}
	if val != 123 {
		t.Errorf("Get returned %v, expected 123", val)
	}
}

func TestLargeCache(t *testing.T) {
	c := NewLFUCache(1000)

	for i := 0; i < 1000; i++ {
		c.Put(string(rune(i)), i)
	}

	for i := 0; i < 1000; i++ {
		val, ok := c.Get(string(rune(i)))
		if !ok {
			t.Errorf("key %d should exist", i)
		}
		if val != i {
			t.Errorf("Get(%d) = %v, expected %d", i, val, i)
		}
	}
}

func TestFrequencyBucketsCorrectness(t *testing.T) {
	c := NewLFUCache(5)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	c.Get("a")
	c.Get("a")
	c.Get("b")

	entryA := c.cache["a"].Value.(*cacheEntry)
	entryB := c.cache["b"].Value.(*cacheEntry)
	entryC := c.cache["c"].Value.(*cacheEntry)

	if entryA.freq != 3 {
		t.Errorf("a.freq is %d, expected 3", entryA.freq)
	}
	if entryB.freq != 2 {
		t.Errorf("b.freq is %d, expected 2", entryB.freq)
	}
	if entryC.freq != 1 {
		t.Errorf("c.freq is %d, expected 1", entryC.freq)
	}
}
