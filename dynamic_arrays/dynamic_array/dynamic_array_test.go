package dynamic_array

import (
	"testing"
)

func TestDynamicArrayInit(t *testing.T) {
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
	var da DynArray[int]
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
