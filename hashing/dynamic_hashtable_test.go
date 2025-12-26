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
