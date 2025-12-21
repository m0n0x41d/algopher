package ordlist

import "testing"

func TestIndexedAddAndCountAsc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(1)
	list.Add(3)
	list.Add(7)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items, got %d", list.Count())
	}

	// Проверяем порядок: 1, 3, 5, 7
	expected := []int{1, 3, 5, 7}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedAddAndCountDesc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: false}
	list.Add(5)
	list.Add(1)
	list.Add(3)
	list.Add(7)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items, got %d", list.Count())
	}

	// Проверяем порядок: 7, 5, 3, 1
	expected := []int{7, 5, 3, 1}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedAddDuplicates(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}
}

func TestIndexedFindExisting(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	val, err := list.Find(3)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestIndexedFindNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	_, err := list.Find(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedFindEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	_, err := list.Find(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedFindIndexAsc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)

	idx, err := list.FindIndex(30)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexDesc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: false}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)

	// Порядок: 50, 40, 30, 20, 10
	idx, err := list.FindIndex(30)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexHead(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 0 {
		t.Errorf("Expected index 0, got %d", idx)
	}
}

func TestIndexedFindIndexTail(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if idx != -1 {
		t.Errorf("Expected index -1, got %d", idx)
	}
}

func TestIndexedFindIndexEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	_, err := list.FindIndex(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedDeleteExisting(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(3)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	_, err := list.Find(3)
	if err == nil {
		t.Errorf("Expected element to be deleted")
	}
}

func TestIndexedDeleteHead(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(1)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	val, _ := list.Get(0)
	if val != 3 {
		t.Errorf("Expected first element to be 3, got %d", val)
	}
}

func TestIndexedDeleteTail(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(5)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	val, _ := list.Get(1)
	if val != 3 {
		t.Errorf("Expected last element to be 3, got %d", val)
	}
}

func TestIndexedDeleteNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(999)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}
}

func TestIndexedDeleteEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	list.Delete(1)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}
}

func TestIndexedClear(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Clear(false)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}

	if list._ascending != false {
		t.Errorf("Expected ascending to be false")
	}
}

func TestIndexedDedup(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}

	expected := []int{1, 3, 5}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedDedupEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	list.Dedup()

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}
}

func TestIndexedDedupSingle(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)

	list.Dedup()

	if list.Count() != 1 {
		t.Errorf("Expected 1 item, got %d", list.Count())
	}
}

func TestIndexedIsSublistTrue(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = IndexedOrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(5)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be found")
	}
}

func TestIndexedIsSublistFalse(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = IndexedOrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(7)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be found")
	}
}

func TestIndexedIsSublistEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	var sublist = IndexedOrderedList[int]{_ascending: true}

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected empty sublist to be found")
	}
}

func TestIndexedTopFrequent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestIndexedTopFrequentEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	result := list.TopFrequent()

	if result != 0 {
		t.Errorf("Expected 0 (zero value), got %d", result)
	}
}

func TestIndexedGet(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	val, err := list.Get(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestIndexedGetOutOfBounds(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	_, err := list.Get(10)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedGetNegativeIndex(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)

	_, err := list.Get(-1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
