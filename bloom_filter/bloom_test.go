package bloomfilter

import (
	"fmt"
	"testing"
)

func generateTestStrings() []string {
	base := "0123456789"
	result := make([]string, 10)
	for i := range 10 {
		result[i] = base[i:] + base[:i]
	}
	return result
}

func TestNewBloomFilter(t *testing.T) {
	bf := NewBloomFilter(32)
	if bf.filter_len != 32 {
		t.Errorf("filter_len is %d, expected 32", bf.filter_len)
	}
	if bf.bitmask != 0 {
		t.Errorf("bitmask is %d, expected 0", bf.bitmask)
	}
}

func TestHash1_ReturnsValidIndex(t *testing.T) {
	bf := NewBloomFilter(32)
	for _, s := range generateTestStrings() {
		idx := bf.Hash1(s)
		if idx < 0 || idx >= 32 {
			t.Errorf("Hash1(%q) returned %d, expected 0..31", s, idx)
		}
	}
}

func TestHash2_ReturnsValidIndex(t *testing.T) {
	bf := NewBloomFilter(32)
	for _, s := range generateTestStrings() {
		idx := bf.Hash2(s)
		if idx < 0 || idx >= 32 {
			t.Errorf("Hash2(%q) returned %d, expected 0..31", s, idx)
		}
	}
}

func TestHash1_Consistency(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "0123456789"
	h1 := bf.Hash1(s)
	h2 := bf.Hash1(s)
	if h1 != h2 {
		t.Errorf("Hash1 not consistent: %d != %d", h1, h2)
	}
}

func TestHash2_Consistency(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "0123456789"
	h1 := bf.Hash2(s)
	h2 := bf.Hash2(s)
	if h1 != h2 {
		t.Errorf("Hash2 not consistent: %d != %d", h1, h2)
	}
}

func TestHash1_Hash2_Different(t *testing.T) {
	bf := NewBloomFilter(32)
	sameCount := 0
	for _, s := range generateTestStrings() {
		h1 := bf.Hash1(s)
		h2 := bf.Hash2(s)
		if h1 == h2 {
			sameCount++
		}
	}
	if sameCount == 10 {
		t.Error("Hash1 and Hash2 return same values for all test strings")
	}
}

func TestAdd_SetsBits(t *testing.T) {
	bf := NewBloomFilter(32)
	if bf.bitmask != 0 {
		t.Errorf("bitmask should be 0 before Add")
	}
	bf.Add("0123456789")
	if bf.bitmask == 0 {
		t.Errorf("bitmask should not be 0 after Add")
	}
}

func TestAdd_SetsCorrectBits(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "0123456789"
	pos1 := bf.Hash1(s)
	pos2 := bf.Hash2(s)

	bf.Add(s)

	bit1 := (bf.bitmask >> uint(pos1)) & 1
	bit2 := (bf.bitmask >> uint(pos2)) & 1

	if bit1 != 1 {
		t.Errorf("bit at position %d should be 1", pos1)
	}
	if bit2 != 1 {
		t.Errorf("bit at position %d should be 1", pos2)
	}
}

func TestIsValue_EmptyFilter(t *testing.T) {
	bf := NewBloomFilter(32)
	for _, s := range generateTestStrings() {
		if bf.IsValue(s) {
			t.Errorf("IsValue(%q) returned true on empty filter", s)
		}
	}
}

func TestIsValue_AfterAdd(t *testing.T) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		bf.Add(s)
	}

	for _, s := range testStrings {
		if !bf.IsValue(s) {
			t.Errorf("IsValue(%q) returned false after Add", s)
		}
	}
}

func TestIsValue_NoFalseNegatives(t *testing.T) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		bf.Add(s)
		if !bf.IsValue(s) {
			t.Errorf("false negative for %q", s)
		}
	}
}

func TestIsValue_SingleElement(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "0123456789"
	bf.Add(s)

	if !bf.IsValue(s) {
		t.Error("IsValue should return true for added element")
	}
}

func TestIsValue_NonExistingElement(t *testing.T) {
	bf := NewBloomFilter(32)
	bf.Add("0123456789")

	notAdded := "completely_different_string_not_in_filter"
	result := bf.IsValue(notAdded)

	t.Logf("IsValue(%q) = %v (may be false positive)", notAdded, result)
}

func TestBloomFilter_AllTestStrings(t *testing.T) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	t.Log("Test strings:")
	for i, s := range testStrings {
		t.Logf("  [%d] %q", i, s)
	}

	for _, s := range testStrings {
		bf.Add(s)
	}

	t.Logf("Bitmask after adding all: %032b", bf.bitmask)

	allFound := true
	for _, s := range testStrings {
		if !bf.IsValue(s) {
			t.Errorf("IsValue(%q) = false, expected true", s)
			allFound = false
		}
	}

	if allFound {
		t.Log("All test strings found in filter")
	}
}

func TestBloomFilter_HashDistribution(t *testing.T) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	t.Log("Hash distribution:")
	for _, s := range testStrings {
		h1 := bf.Hash1(s)
		h2 := bf.Hash2(s)
		t.Logf("  %q: Hash1=%d, Hash2=%d", s, h1, h2)
	}
}

func TestAdd_Idempotent(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "0123456789"

	bf.Add(s)
	mask1 := bf.bitmask

	bf.Add(s)
	mask2 := bf.bitmask

	if mask1 != mask2 {
		t.Errorf("Add should be idempotent: %032b != %032b", mask1, mask2)
	}
}

func TestIsValue_PartialBitsSet(t *testing.T) {
	bf := NewBloomFilter(32)
	s := "test_string"
	pos1 := bf.Hash1(s)

	bf.bitmask |= 1 << uint(pos1)

	if bf.IsValue(s) {
		pos2 := bf.Hash2(s)
		if pos1 == pos2 {
			t.Log("Hash1 and Hash2 returned same position, IsValue is true")
		} else {
			t.Error("IsValue should return false when only one bit is set")
		}
	}
}

func TestBloomFilter_FalsePositiveExample(t *testing.T) {
	bf := NewBloomFilter(32)

	bf.Add("abc")
	bf.Add("def")

	candidates := []string{"ghi", "jkl", "mno", "pqr", "stu", "vwx", "xyz"}
	falsePositives := []string{}

	for _, s := range candidates {
		if bf.IsValue(s) {
			falsePositives = append(falsePositives, s)
		}
	}

	if len(falsePositives) > 0 {
		t.Logf("False positives found: %v", falsePositives)
	} else {
		t.Log("No false positives in test set")
	}
}

func TestBloomFilter_BitManipulation(t *testing.T) {
	bf := NewBloomFilter(32)

	bf.bitmask |= 1 << 0
	if (bf.bitmask>>0)&1 != 1 {
		t.Error("bit 0 should be set")
	}

	bf.bitmask |= 1 << 31
	if (bf.bitmask>>31)&1 != 1 {
		t.Error("bit 31 should be set")
	}

	if bf.bitmask != (1 | (1 << 31)) {
		t.Errorf("unexpected bitmask: %032b", bf.bitmask)
	}
}

func TestHash_EmptyString(t *testing.T) {
	bf := NewBloomFilter(32)

	h1 := bf.Hash1("")
	h2 := bf.Hash2("")

	if h1 != 0 {
		t.Errorf("Hash1('') = %d, expected 0", h1)
	}
	if h2 != 0 {
		t.Errorf("Hash2('') = %d, expected 0", h2)
	}
}

func TestAdd_EmptyString(t *testing.T) {
	bf := NewBloomFilter(32)
	bf.Add("")

	if (bf.bitmask>>0)&1 != 1 {
		t.Error("bit 0 should be set after adding empty string")
	}
}

func TestIsValue_EmptyString(t *testing.T) {
	bf := NewBloomFilter(32)
	bf.Add("")

	if !bf.IsValue("") {
		t.Error("IsValue('') should return true after Add('')")
	}
}

func BenchmarkAdd(b *testing.B) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			bf.Add(s)
		}
	}
}

func BenchmarkIsValue(b *testing.B) {
	bf := NewBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		bf.Add(s)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			bf.IsValue(s)
		}
	}
}

func BenchmarkHash1(b *testing.B) {
	bf := NewBloomFilter(32)
	s := "0123456789"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Hash1(s)
	}
}

func BenchmarkHash2(b *testing.B) {
	bf := NewBloomFilter(32)
	s := "0123456789"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.Hash2(s)
	}
}

// merge filters
func TestMerge_EmptyInput(t *testing.T) {
	result := Merge()
	if result != nil {
		t.Error("Merge() with no arguments should return nil")
	}
}

func TestMerge_SingleFilter(t *testing.T) {
	bf := NewBloomFilter(32)
	bf.Add("hello")

	merged := Merge(bf)

	if merged.bitmask != bf.bitmask {
		t.Errorf("merged bitmask %032b != original %032b", merged.bitmask, bf.bitmask)
	}
	if !merged.IsValue("hello") {
		t.Error("merged filter should contain 'hello'")
	}
}

func TestMerge_TwoFilters(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf1.Add("cat")

	bf2 := NewBloomFilter(32)
	bf2.Add("dog")

	merged := Merge(bf1, bf2)

	if !merged.IsValue("cat") {
		t.Error("merged filter should contain 'cat'")
	}
	if !merged.IsValue("dog") {
		t.Error("merged filter should contain 'dog'")
	}
}

func TestMerge_ThreeFilters(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf1.Add("one")

	bf2 := NewBloomFilter(32)
	bf2.Add("two")

	bf3 := NewBloomFilter(32)
	bf3.Add("three")

	merged := Merge(bf1, bf2, bf3)

	for _, s := range []string{"one", "two", "three"} {
		if !merged.IsValue(s) {
			t.Errorf("merged filter should contain %q", s)
		}
	}
}

func TestMerge_BitwiseOR(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf1.bitmask = 0b00001111

	bf2 := NewBloomFilter(32)
	bf2.bitmask = 0b11110000

	merged := Merge(bf1, bf2)

	expected := uint32(0b11111111)
	if merged.bitmask != expected {
		t.Errorf("merged bitmask %032b != expected %032b", merged.bitmask, expected)
	}
}

func TestMerge_OverlappingBits(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf1.bitmask = 0b00111100

	bf2 := NewBloomFilter(32)
	bf2.bitmask = 0b00011110

	merged := Merge(bf1, bf2)

	expected := uint32(0b00111110)
	if merged.bitmask != expected {
		t.Errorf("merged bitmask %032b != expected %032b", merged.bitmask, expected)
	}
}

func TestMerge_PreservesFilterLen(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf2 := NewBloomFilter(32)

	merged := Merge(bf1, bf2)

	if merged.filter_len != 32 {
		t.Errorf("merged filter_len %d != 32", merged.filter_len)
	}
}

func TestMerge_EmptyFilters(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf2 := NewBloomFilter(32)

	merged := Merge(bf1, bf2)

	if merged.bitmask != 0 {
		t.Errorf("merging empty filters should give bitmask 0, got %032b", merged.bitmask)
	}
}

func TestMerge_DoesNotModifyOriginals(t *testing.T) {
	bf1 := NewBloomFilter(32)
	bf1.Add("cat")
	original1 := bf1.bitmask

	bf2 := NewBloomFilter(32)
	bf2.Add("dog")
	original2 := bf2.bitmask

	Merge(bf1, bf2)

	if bf1.bitmask != original1 {
		t.Error("Merge modified bf1")
	}
	if bf2.bitmask != original2 {
		t.Error("Merge modified bf2")
	}
}

func TestMerge_NoFalseNegatives(t *testing.T) {
	filters := make([]*BloomFilter, 5)
	allStrings := []string{}

	for i := range 5 {
		filters[i] = NewBloomFilter(32)
		s := fmt.Sprintf("string_%d", i)
		filters[i].Add(s)
		allStrings = append(allStrings, s)
	}

	merged := Merge(filters...)

	for _, s := range allStrings {
		if !merged.IsValue(s) {
			t.Errorf("false negative for %q after merge", s)
		}
	}
}

func TestMerge_IncreasedFalsePositives(t *testing.T) {
	bf1 := NewBloomFilter(32)
	for i := range 5 {
		bf1.Add(fmt.Sprintf("set1_%d", i))
	}
	bits1 := countBits(bf1.bitmask)

	bf2 := NewBloomFilter(32)
	for i := range 5 {
		bf2.Add(fmt.Sprintf("set2_%d", i))
	}
	bits2 := countBits(bf2.bitmask)

	merged := Merge(bf1, bf2)
	bitsMerged := countBits(merged.bitmask)

	t.Logf("bf1 bits: %d, bf2 bits: %d, merged bits: %d", bits1, bits2, bitsMerged)

	if bitsMerged < bits1 || bitsMerged < bits2 {
		t.Error("merged filter should have at least as many bits as each original")
	}
}

func countBits(n uint32) int {
	count := 0
	for n != 0 {
		count += int(n & 1)
		n >>= 1
	}
	return count
}

func TestMerge_DistributedScenario(t *testing.T) {
	nodeA := NewBloomFilter(32)
	nodeA.Add("user_1")
	nodeA.Add("user_2")

	nodeB := NewBloomFilter(32)
	nodeB.Add("user_3")
	nodeB.Add("user_4")

	nodeC := NewBloomFilter(32)
	nodeC.Add("user_5")

	global := Merge(nodeA, nodeB, nodeC)

	for i := 1; i <= 5; i++ {
		user := fmt.Sprintf("user_%d", i)
		if !global.IsValue(user) {
			t.Errorf("global filter should contain %q", user)
		}
	}

	t.Logf("Global filter bitmask: %032b", global.bitmask)
	t.Logf("Bits set: %d/32", countBits(global.bitmask))
}

// deletion bloom

func TestCountingBloomFilter_New(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	if cbf.filter_len != 32 {
		t.Errorf("filter_len is %d, expected 32", cbf.filter_len)
	}
	if len(cbf.counters) != 32 {
		t.Errorf("counters length is %d, expected 32", len(cbf.counters))
	}
	for i, c := range cbf.counters {
		if c != 0 {
			t.Errorf("counter[%d] is %d, expected 0", i, c)
		}
	}
}

func TestCountingBloomFilter_Hash1(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	bf := NewBloomFilter(32)

	for _, s := range generateTestStrings() {
		if cbf.Hash1(s) != bf.Hash1(s) {
			t.Errorf("Hash1 mismatch for %q", s)
		}
	}
}

func TestCountingBloomFilter_Hash2(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	bf := NewBloomFilter(32)

	for _, s := range generateTestStrings() {
		if cbf.Hash2(s) != bf.Hash2(s) {
			t.Errorf("Hash2 mismatch for %q", s)
		}
	}
}

func TestCountingBloomFilter_Add(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "hello"
	pos1 := cbf.Hash1(s)
	pos2 := cbf.Hash2(s)

	cbf.Add(s)

	if cbf.counters[pos1] != 1 {
		t.Errorf("counter[%d] is %d, expected 1", pos1, cbf.counters[pos1])
	}
	if cbf.counters[pos2] != 1 {
		t.Errorf("counter[%d] is %d, expected 1", pos2, cbf.counters[pos2])
	}
}

func TestCountingBloomFilter_Add_Increments(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "hello"
	pos1 := cbf.Hash1(s)

	cbf.Add(s)
	cbf.Add(s)
	cbf.Add(s)

	if cbf.counters[pos1] != 3 {
		t.Errorf("counter[%d] is %d, expected 3", pos1, cbf.counters[pos1])
	}
}

func TestCountingBloomFilter_IsValue_Empty(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	for _, s := range generateTestStrings() {
		if cbf.IsValue(s) {
			t.Errorf("IsValue(%q) returned true on empty filter", s)
		}
	}
}

func TestCountingBloomFilter_IsValue_AfterAdd(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		cbf.Add(s)
	}

	for _, s := range testStrings {
		if !cbf.IsValue(s) {
			t.Errorf("IsValue(%q) returned false after Add", s)
		}
	}
}

func TestCountingBloomFilter_Remove_Basic(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "hello"

	cbf.Add(s)
	if !cbf.IsValue(s) {
		t.Error("IsValue should return true after Add")
	}

	cbf.Remove(s)
	if cbf.IsValue(s) {
		t.Error("IsValue should return false after Remove")
	}
}

func TestCountingBloomFilter_Remove_Decrements(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "hello"
	pos1 := cbf.Hash1(s)

	cbf.Add(s)
	cbf.Add(s)
	cbf.Add(s)
	cbf.Remove(s)

	if cbf.counters[pos1] != 2 {
		t.Errorf("counter[%d] is %d, expected 2", pos1, cbf.counters[pos1])
	}
}

func TestCountingBloomFilter_Remove_NoUnderflow(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "hello"
	pos1 := cbf.Hash1(s)
	pos2 := cbf.Hash2(s)

	cbf.Remove(s)
	cbf.Remove(s)
	cbf.Remove(s)

	if cbf.counters[pos1] != 0 {
		t.Errorf("counter[%d] is %d, expected 0 (no underflow)", pos1, cbf.counters[pos1])
	}
	if cbf.counters[pos2] != 0 {
		t.Errorf("counter[%d] is %d, expected 0 (no underflow)", pos2, cbf.counters[pos2])
	}
}

func TestCountingBloomFilter_Remove_PreservesOthers(t *testing.T) {
	cbf := NewCountingBloomFilter(32)

	cbf.Add("cat")
	cbf.Add("dog")

	cbf.Remove("cat")

	if cbf.IsValue("cat") {
		t.Error("'cat' should not be found after Remove")
	}
	if !cbf.IsValue("dog") {
		t.Error("'dog' should still be found after removing 'cat'")
	}
}

func TestCountingBloomFilter_CollisionHandling(t *testing.T) {
	cbf := NewCountingBloomFilter(32)

	cbf.Add("cat")
	cbf.Add("dog")

	pos1Cat := cbf.Hash1("cat")
	pos1Dog := cbf.Hash1("dog")

	if pos1Cat == pos1Dog {
		t.Logf("'cat' and 'dog' collide at position %d", pos1Cat)
		if cbf.counters[pos1Cat] != 2 {
			t.Errorf("counter should be 2 after collision, got %d", cbf.counters[pos1Cat])
		}

		cbf.Remove("cat")
		if cbf.counters[pos1Cat] != 1 {
			t.Errorf("counter should be 1 after removing 'cat', got %d", cbf.counters[pos1Cat])
		}
		if !cbf.IsValue("dog") {
			t.Error("'dog' should still be found after removing 'cat' (collision case)")
		}
	}
}

func TestCountingBloomFilter_Add_NoOverflow(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	s := "overflow_test"
	pos1 := cbf.Hash1(s)

	for range 300 {
		cbf.Add(s)
	}

	if cbf.counters[pos1] != 255 {
		t.Errorf("counter should cap at 255, got %d", cbf.counters[pos1])
	}
}

func TestCountingBloomFilter_FalsePositiveRemoval(t *testing.T) {
	cbf := NewCountingBloomFilter(32)

	cbf.Add("real_element")

	fakeElement := "fake_element"
	if cbf.IsValue(fakeElement) {
		t.Log("False positive detected for 'fake_element'")

		cbf.Remove(fakeElement)

		if !cbf.IsValue("real_element") {
			t.Log("FALSE NEGATIVE caused by removing false positive — this is expected behavior")
		}
	} else {
		t.Log("No false positive for 'fake_element' in this case")
	}
}

func TestCountingBloomFilter_MultipleAddRemove(t *testing.T) {
	cbf := NewCountingBloomFilter(32)

	cbf.Add("item")
	cbf.Add("item")
	cbf.Add("item")

	cbf.Remove("item")
	if !cbf.IsValue("item") {
		t.Error("item should still exist after 1 remove (added 3 times)")
	}

	cbf.Remove("item")
	if !cbf.IsValue("item") {
		t.Error("item should still exist after 2 removes (added 3 times)")
	}

	cbf.Remove("item")
	if cbf.IsValue("item") {
		t.Error("item should not exist after 3 removes (added 3 times)")
	}
}

func TestCountingBloomFilter_AllTestStrings(t *testing.T) {
	cbf := NewCountingBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		cbf.Add(s)
	}

	for _, s := range testStrings {
		if !cbf.IsValue(s) {
			t.Errorf("IsValue(%q) = false after Add", s)
		}
	}

	for i, s := range testStrings {
		cbf.Remove(s)

		if cbf.IsValue(s) {
			t.Logf("IsValue(%q) still true after Remove (likely collision)", s)
		}

		for j := i + 1; j < len(testStrings); j++ {
			if !cbf.IsValue(testStrings[j]) {
				t.Errorf("IsValue(%q) = false, but it wasn't removed yet", testStrings[j])
			}
		}
	}
}

func BenchmarkCountingBloomFilter_Add(b *testing.B) {
	cbf := NewCountingBloomFilter(32)
	testStrings := generateTestStrings()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			cbf.Add(s)
		}
	}
}

func BenchmarkCountingBloomFilter_Remove(b *testing.B) {
	cbf := NewCountingBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		for range 1000 {
			cbf.Add(s)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			cbf.Remove(s)
		}
	}
}

func BenchmarkCountingBloomFilter_IsValue(b *testing.B) {
	cbf := NewCountingBloomFilter(32)
	testStrings := generateTestStrings()

	for _, s := range testStrings {
		cbf.Add(s)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			cbf.IsValue(s)
		}
	}
}
