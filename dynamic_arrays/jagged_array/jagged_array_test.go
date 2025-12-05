package jagged_array

import (
	"testing"
)

func TestJaggedArrayInit(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	if ja.RowCount() != 0 {
		t.Errorf("ja.RowCount() != 0, got: %d", ja.RowCount())
	}
}

func TestJaggedArrayInitWithDimensions(t *testing.T) {
	var ja JaggedArray2D[int]
	rows := 3
	cols := 5

	ja.InitWithDimensions(rows, cols)

	if ja.RowCount() != rows {
		t.Errorf("ja.RowCount() != %d, got: %d", rows, ja.RowCount())
	}

	for i := 0; i < rows; i++ {
		row_len, err := ja.RowLength(i)
		if err != nil {
			t.Errorf("Error getting row length: %s", err)
		}
		if row_len != cols {
			t.Errorf("Row %d length != %d, got: %d", i, cols, row_len)
		}
	}
}

func TestJaggedArrayGetSetValidIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	value := 42
	row := 1
	col := 2

	err := ja.Set(value, row, col)
	if err != nil {
		t.Errorf("Error during Set: %s", err)
	}

	got, err := ja.Get(row, col)
	if err != nil {
		t.Errorf("Error during Get: %s", err)
	}

	if got != value {
		t.Errorf("Get returned wrong value: got %d, want %d", got, value)
	}
}

func TestJaggedArrayGetInvalidRowIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	_, err := ja.Get(-1, 0)
	if err == nil {
		t.Error("Expected error for negative row index")
	}

	_, err = ja.Get(10, 0)
	if err == nil {
		t.Error("Expected error for row index > RowCount")
	}
}

func TestJaggedArrayGetInvalidColIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	_, err := ja.Get(0, -1)
	if err == nil {
		t.Error("Expected error for negative col index")
	}

	_, err = ja.Get(0, 10)
	if err == nil {
		t.Error("Expected error for col index > row length")
	}
}

func TestJaggedArraySetInvalidRowIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	err := ja.Set(1, -1, 0)
	if err == nil {
		t.Error("Expected error for negative row index")
	}

	err = ja.Set(1, 10, 0)
	if err == nil {
		t.Error("Expected error for row index > RowCount")
	}
}

func TestJaggedArraySetInvalidColIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	err := ja.Set(1, 0, -1)
	if err == nil {
		t.Error("Expected error for negative col index")
	}

	err = ja.Set(1, 0, 10)
	if err == nil {
		t.Error("Expected error for col index > row length")
	}
}

func TestJaggedArrayAppendRow(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	initial_rows := ja.RowCount()
	if initial_rows != 0 {
		t.Errorf("Initial row count should be 0, got: %d", initial_rows)
	}

	ja.AppendRow()
	ja.AppendRow()
	ja.AppendRow()

	if ja.RowCount() != 3 {
		t.Errorf("After 3 AppendRow, RowCount should be 3, got: %d", ja.RowCount())
	}

	// New rows should be empty
	for i := 0; i < 3; i++ {
		row_len, err := ja.RowLength(i)
		if err != nil {
			t.Errorf("Error getting row length: %s", err)
		}
		if row_len != 0 {
			t.Errorf("New row %d should be empty, got length: %d", i, row_len)
		}
	}
}

func TestJaggedArrayAppendToRow(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	ja.AppendRow()
	row_index := 0

	initial_len, _ := ja.RowLength(row_index)
	if initial_len != 0 {
		t.Errorf("Initial row length should be 0, got: %d", initial_len)
	}

	err := ja.AppendToRow(row_index, 10)
	if err != nil {
		t.Errorf("Error during AppendToRow: %s", err)
	}

	err = ja.AppendToRow(row_index, 20)
	if err != nil {
		t.Errorf("Error during AppendToRow: %s", err)
	}

	row_len, _ := ja.RowLength(row_index)
	if row_len != 2 {
		t.Errorf("Row length should be 2 after 2 appends, got: %d", row_len)
	}

	val, _ := ja.Get(row_index, 0)
	if val != 10 {
		t.Errorf("First element should be 10, got: %d", val)
	}

	val, _ = ja.Get(row_index, 1)
	if val != 20 {
		t.Errorf("Second element should be 20, got: %d", val)
	}
}

func TestJaggedArrayAppendToRowInvalidIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	err := ja.AppendToRow(0, 1) // no rows exist
	if err == nil {
		t.Error("Expected error when appending to non-existent row")
	}

	err = ja.AppendToRow(-1, 1)
	if err == nil {
		t.Error("Expected error for negative row index")
	}
}

// Key test: independent row growth!
func TestJaggedArrayIndependentRowGrowth(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 2) // 3 rows with 2 elements each

	// Verify initial state
	for i := 0; i < 3; i++ {
		len, _ := ja.RowLength(i)
		if len != 2 {
			t.Errorf("Initial row %d length should be 2, got: %d", i, len)
		}
	}

	// Grow only first row
	ja.AppendToRow(0, 100)
	ja.AppendToRow(0, 101)
	ja.AppendToRow(0, 102)

	// Grow only third row
	ja.AppendToRow(2, 200)

	// Verify rows grew independently
	len0, _ := ja.RowLength(0)
	if len0 != 5 { // 2 initial + 3 appended
		t.Errorf("Row 0 should have length 5, got: %d", len0)
	}

	len1, _ := ja.RowLength(1)
	if len1 != 2 {
		t.Errorf("Row 1 should still have length 2, got: %d", len1)
	}

	len2, _ := ja.RowLength(2)
	if len2 != 3 { // 2 initial + 1 appended
		t.Errorf("Row 2 should have length 3, got: %d", len2)
	}

	val, _ := ja.Get(0, 4)
	if val != 102 {
		t.Errorf("Row 0, col 4 should be 102, got: %d", val)
	}

	val, _ = ja.Get(2, 2)
	if val != 200 {
		t.Errorf("Row 2, col 2 should be 200, got: %d", val)
	}
}

func TestJaggedArrayFillAndRead(t *testing.T) {
	var ja JaggedArray2D[int]
	rows := 10
	cols := 10

	ja.InitWithDimensions(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			value := i*100 + j
			err := ja.Set(value, i, j)
			if err != nil {
				t.Errorf("Error setting [%d,%d]: %s", i, j, err)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			expected := i*100 + j
			got, err := ja.Get(i, j)
			if err != nil {
				t.Errorf("Error getting [%d,%d]: %s", i, j, err)
			}
			if got != expected {
				t.Errorf("At [%d,%d]: expected %d, got %d", i, j, expected, got)
			}
		}
	}
}

func TestJaggedArrayRowLengthInvalidIndex(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.InitWithDimensions(3, 3)

	_, err := ja.RowLength(-1)
	if err == nil {
		t.Error("Expected error for negative row index")
	}

	_, err = ja.RowLength(10)
	if err == nil {
		t.Error("Expected error for row index >= RowCount")
	}
}

func TestJaggedArraySparseUsage(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	ja.AppendRow()
	for i := 0; i < 100; i++ {
		ja.AppendToRow(0, i)
	}

	ja.AppendRow()
	ja.AppendToRow(1, 1)
	ja.AppendToRow(1, 2)

	ja.AppendRow()
	ja.AppendToRow(2, 10)

	len0, _ := ja.RowLength(0)
	len1, _ := ja.RowLength(1)
	len2, _ := ja.RowLength(2)

	if len0 != 100 {
		t.Errorf("Row 0 length should be 100, got: %d", len0)
	}
	if len1 != 2 {
		t.Errorf("Row 1 length should be 2, got: %d", len1)
	}
	if len2 != 1 {
		t.Errorf("Row 2 length should be 1, got: %d", len2)
	}
}

func TestJaggedArrayManyRowsGrowth(t *testing.T) {
	var ja JaggedArray2D[int]
	ja.Init()

	for i := 0; i < 100; i++ {
		ja.AppendRow()
	}

	if ja.RowCount() != 100 {
		t.Errorf("Should have 100 rows, got: %d", ja.RowCount())
	}

	ja.AppendToRow(0, 1)
	ja.AppendToRow(50, 2)
	ja.AppendToRow(50, 3)
	ja.AppendToRow(99, 4)

	len50, _ := ja.RowLength(50)
	if len50 != 2 {
		t.Errorf("Row 50 should have 2 elements, got: %d", len50)
	}

	val, _ := ja.Get(50, 1)
	if val != 3 {
		t.Errorf("Row 50, col 1 should be 3, got: %d", val)
	}
}
