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
