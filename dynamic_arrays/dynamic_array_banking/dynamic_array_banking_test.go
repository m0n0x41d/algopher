package dynamic_array_banking

import (
	"testing"
)

func TestDynamicArrayInit(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()
	if da.count != 0 {
		t.Errorf("da.count != 0")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", DEFAULT_MIN_BUFFER_SIZE)
	}

	if len(da.array) != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("len(da.array) != DEFAULT_MIN_BUFFER_SIZE, got: %d, want: %d", len(da.array), DEFAULT_MIN_BUFFER_SIZE)
	}
}

func TestDynamicArrayInserItemNoReallocate(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()
	item := 1
	index := 0
	err := da.Insert(item, index)
	if err != nil {
		t.Errorf("Error happened during da.Insert: %s", err)
	}

	if da.count != 1 {
		t.Errorf("da.count != 1")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", da.capacity)
	}

	item_from_da, err := da.GetItem(0)
	if err != nil {
		t.Errorf("Error happened during da.GetItem: %s", err)
	}
	if item_from_da != item {
		t.Errorf("item_from_da != item, got: %d, want: %d", item_from_da, item)
	}
}

func TestDynamicArrayAppendItemNoRealloc(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()
	item := 1
	da.Append(item)

	if da.count != 1 {
		t.Errorf("da.count != 1")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", da.capacity)
	}

	another_item := 2

	da.Append(another_item)

	if da.count != 2 {
		t.Errorf("da.count != 2")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", da.capacity)
	}

	item_1, err := da.GetItem(0)
	if err != nil {
		t.Errorf("Error happened during da.GetItem: %s", err)
	}
	if item_1 != item {
		t.Errorf("item_1 != item, got: %d, want: %d", item_1, item)
	}

	item_2, err := da.GetItem(1)
	if err != nil {
		t.Errorf("Error happened during da.GetItem: %s", err)
	}
	if item_2 != another_item {
		t.Errorf("item_2 != another_item, got: %d, want: %d", item_2, another_item)
	}
}

func TestDynamicArrayRemoveValidIndex(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()
	item := 1
	index := 0
	err := da.Insert(item, index)
	if err != nil {
		t.Errorf("Error happened during da.Insert: %s", err)
	}

	if da.count != 1 {
		t.Errorf("da.count != 1")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", da.capacity)
	}

	remove_err := da.Remove(index)
	if remove_err != nil {
		t.Errorf("Error happened during da.Remove: %s", remove_err)
	}

	if da.count != 0 {
		t.Errorf("da.count != 0")
	}
	if da.capacity != DEFAULT_MIN_BUFFER_SIZE {
		t.Errorf("da.capacity != DEFAULT_MIN_BUFFER_SIZE, got: %d", da.capacity)
	}
}

func TestDynamicArrayReallocationUpOnConsequensiveInsert(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	for i := 1; i <= 17; i++ {
		insert_err := da.Insert(i, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	if da.count != 17 {
		t.Errorf("Incorrect counter in the array")
	}

	if da.capacity != 32 {
		t.Errorf("This implementation of DynamicArray must grow up at twice of previous size. New capasity got: %d, want: 32", da.capacity)
	}
}

func TestDynamicArrayInsertInTheMiddleNoResize(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	for i := 1; i <= 5; i++ {
		insert_err := da.Insert(i, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	da.Insert(108, 3)

	got, err := da.GetItem(4)
	if err != nil {
		t.Errorf("Unexpected error happened on GetItem: %s", err)
	}

	if got != 4 {
		t.Errorf("Insert is working incorrectly on insert in the middle tail is not moving forward. got: %d, want: 4", got)
	}
}

func TestDynamicArrayInsertInTheMiddleWithResize(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	for i := 1; i <= 6; i++ {
		insert_err := da.Insert(1, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	for i := 1; i <= 6; i++ {
		insert_err := da.Insert(2, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	for i := 1; i <= 6; i++ {
		insert_err := da.Insert(3, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	if da.capacity != 32 {
		t.Errorf("Capacity must be 32, got: %d", da.capacity)
	}
}

func TestDymanicArrayRemoveMustShrink(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	for i := 1; i <= 20; i++ {
		insert_err := da.Insert(1, i-1)
		if insert_err != nil {
			t.Errorf("Failed to inser in the array: %s", insert_err)
		}
	}

	for i := 1; i <= 5; i++ {
		da.Remove(0)
	}

	if da.capacity != 21 {
		t.Errorf("Array capacity must shink down by 1.5 of current capacity when usage is less then 50percents. got: %d, want: 21", da.capacity)
	}

	for i := 1; i <= 5; i++ {
		da.Remove(0)
	}

	if da.capacity != 16 {
		t.Errorf("Array capacity must shink down by 1.5 of current capacity when usage is less then 50percents, and minimum size must be 16! got: %d, want: 16", da.capacity)
	}
}

func TestDynamicArrayInsertInvalidIndex(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	err := da.Insert(1, -1) // негативный индекс
	if err == nil {
		t.Error("Expected error for negative index")
	}

	err = da.Insert(1, 100) // индекс > count
	if err == nil {
		t.Error("Expected error for index > count")
	}
}

func TestDynamicArrayRemoveInvalidIndex(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()
	da.Append(1)

	err := da.Remove(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	err = da.Remove(100)
	if err == nil {
		t.Error("Expected error for index >= count")
	}
}

func TestDynamicArrayGetItemInvalidIndex(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	_, err := da.GetItem(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	_, err = da.GetItem(0) // пустой массив
	if err == nil {
		t.Error("Expected error for empty array")
	}
}

// ========== Banking Method Specific Tests ==========

// TestBankingCreditNeverNegative verifies that credit balance never goes negative
func TestBankingCreditNeverNegative(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	if da.GetCredit() < 0 {
		t.Errorf("Credit should never be negative, got: %d", da.GetCredit())
	}

	// Insert many elements triggering multiple reallocations
	for i := 0; i < 100; i++ {
		da.Append(i)
		if da.GetCredit() < 0 {
			t.Errorf("Credit went negative after insert %d: credit=%d", i, da.GetCredit())
		}
	}
}

// TestBankingSufficientCreditBeforeReallocation proves that credit is always sufficient
// when reallocation is needed
func TestBankingSufficientCreditBeforeReallocation(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	// Track credit before each potential reallocation
	for i := 0; i < 100; i++ {
		prevCapacity := da.capacity
		creditBefore := da.GetCredit()

		da.Append(i)

		// If reallocation happened
		if da.capacity > prevCapacity {
			// Credit before must have been >= cost of copying
			costOfCopy := i // we had i elements to copy
			if creditBefore < costOfCopy {
				t.Errorf("Insufficient credit before reallocation at i=%d: credit=%d, cost=%d",
					i, creditBefore, costOfCopy)
			}
		}
	}
}

// TestBankingInvariantMaintained tests the core banking method invariant:
// credit >= (elements that would need to be copied in worst case)
func TestBankingInvariantMaintained(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	for i := 1; i <= 64; i++ {
		da.Append(i)

		// At any point, credit should be sufficient for next reallocation
		// Each insert adds 3 credits and costs 1 (the insert itself)
		// So we accumulate 2 credits per insert until reallocation
		credit := da.GetCredit()

		// Invariant: credit should be non-negative
		if credit < 0 {
			t.Errorf("Invariant violated at count=%d: credit=%d", da.count, credit)
		}

		t.Logf("count=%d, capacity=%d, credit=%d", da.count, da.capacity, credit)
	}
}

// TestBankingWorstCaseSequence tests the worst-case sequence that triggers
// reallocation on every power of 2
func TestBankingWorstCaseSequence(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	// Insert exactly at reallocation boundaries
	reallocationPoints := []int{16, 32, 64, 128}

	for _, target := range reallocationPoints {
		for da.count < target {
			da.Append(da.count)
		}

		credit := da.GetCredit()
		t.Logf("After %d inserts: capacity=%d, credit=%d", target, da.capacity, credit)

		// Credit must still be non-negative
		if credit < 0 {
			t.Errorf("Credit went negative after reaching capacity %d: credit=%d", da.capacity, credit)
		}
	}
}

// TestBankingCreditAccumulation verifies that credit accumulates correctly:
// Each insert adds 3, each reallocation subtracts count
func TestBankingCreditAccumulation(t *testing.T) {
	var da DynArrayBanking[int]
	da.Init()

	expectedCredit := 0

	for i := 0; i < 50; i++ {
		prevCount := da.count
		prevCapacity := da.capacity

		da.Append(i)
		expectedCredit += 3 // each insert adds 3

		// If reallocation happened
		if da.capacity > prevCapacity {
			expectedCredit -= prevCount // subtract cost of copying
			t.Logf("Reallocation at count=%d: capacity %d->%d, credit after=%d",
				prevCount, prevCapacity, da.capacity, da.GetCredit())
		}

		actualCredit := da.GetCredit()
		if actualCredit != expectedCredit {
			t.Errorf("Credit mismatch at i=%d: expected=%d, actual=%d",
				i, expectedCredit, actualCredit)
		}
	}
}

// TestBankingMathematicalProof provides mathematical proof that banking method works
func TestBankingMathematicalProof(t *testing.T) {
	t.Log("=== Proof of Banking Method ===")
	t.Log("Claim: With cost 3 per insert, credit is always sufficient for reallocation")
	t.Log("")
	t.Log("Proof:")
	t.Log("- Each insert adds 3 credits")
	t.Log("- Reallocation from capacity C to 2C requires copying C elements (cost C)")
	t.Log("- Between reallocations, we do C inserts (filling the buffer)")
	t.Log("- These C inserts accumulate: 3*C = 3C credits")
	t.Log("- Cost of reallocation: C credits")
	t.Log("- Surplus after reallocation: 3C - C = 2C credits")
	t.Log("- Therefore, credit >= 0 always (actually credit >= 2C)")
	t.Log("")

	var da DynArrayBanking[int]
	da.Init()

	minCredit := 0

	for i := 0; i < 256; i++ {
		da.Append(i)
		credit := da.GetCredit()

		if credit < minCredit {
			minCredit = credit
		}
	}

	t.Logf("Result: minimum credit observed = %d", minCredit)

	if minCredit < 0 {
		t.Errorf("Banking method failed: credit went negative (min=%d)", minCredit)
	} else {
		t.Log("Banking method invariant verified :)")
	}
}
