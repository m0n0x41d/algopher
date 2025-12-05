package main

import (
	"fmt"
	"os"
)

// Tests for both of additional tasks are located in upper directories of same module in separate modules accordingly.

var _ = os.Args

const (
	DEFAULT_MIN_BUFFER_SIZE_BANKING = 16
)
const SHRINK_FACTOR_BANKING float32 = 1.5

type DynArrayBanking[T any] struct {
	count    int
	capacity int
	array    []T
	credit   int
}

func (da *DynArrayBanking[T]) Init() {
	da.count = 0
	da.MakeArray(DEFAULT_MIN_BUFFER_SIZE)
}

// Complexity must be linear O(1) for empty arrays, and O(n) while copying
//
// additionally:
// It should be a 'private' method returning new
// instance of the array of the size, so it will be
// functionally used in other private OR public methods
// for example - in growUp.
func (da *DynArrayBanking[T]) MakeArray(sz int) {
	arr := make([]T, sz)
	copy(arr, da.array)
	da.capacity = sz
	da.array = arr
}

// Complexity depends:
// by time:
//   - Best case: O(1) - inserting ad the end with no reallocation and shifting
//   - Worst case: O(n) - inserting in beginning with reallocation
//   - Overall amortized: O(1) - while using multiplicative growth strategy
//
// by space: O(1) except for reallocation it is O(n)
func (da *DynArrayBanking[T]) Insert(itm T, index int) error {
	if index < 0 || index > da.count {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	da.credit += 3

	if da.count == da.capacity {
		cost := da.count // cost of copying existing elements
		if da.credit >= cost {
			da.credit -= cost
			da.growUp()
		} else {
			panic(fmt.Sprintf("Banking method invariant violated: credit=%d, cost=%d", da.credit, cost))
		}
	}

	if index <= da.count {
		copy(da.array[index+1:], da.array[index:da.count])
	}

	da.array[index] = itm
	da.count++
	return nil
}

// Complexity:
// by time:
//   - Best case: O(1) - removing last item without shrinking
//   - Worst case: O(n) - removing first item with shrinking
//   - Amortized: O(1) - again, thaks to multiplicative strategy
//
// by space:
//   - O(1) - regular shift
//   - O(n) - while shrinking
func (da *DynArrayBanking[T]) Remove(index int) error {
	if index < 0 || index >= da.count {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	copy(da.array[index:], da.array[index+1:da.count])
	da.count--
	da.array[da.count] = *new(T)

	if da.count*100 < da.capacity*SHRINK_AT_PERCENT {
		da.shrink()
	}

	return nil
}

// Complexity:
// by time: Amortized O(1) - constant time on average due to multiplicative growth.
// Some operations may be O(n) during reallocation, but amortized cost is O(1).
//
// by space: O(1) except during reallocation O(n)
//
// Additionally:
// Not doing with erorr anything? ¯\_(ツ)_/¯
// Because appen should be always successfull, until os or hardware has enought memory)
func (da *DynArrayBanking[T]) Append(itm T) {
	da.Insert(itm, da.count)
}

// O(1) - pointer access.
func (da *DynArrayBanking[T]) GetItem(index int) (T, error) {
	if index < 0 || index >= da.count {
		var zero T
		return zero, fmt.Errorf("index out of bounds: %d", index)
	}

	return da.array[index], nil
}

// Complexity is O(n)
func (da *DynArrayBanking[T]) growUp() {
	new_capacity := da.capacity * 2
	new_array := make([]T, new_capacity)
	copy(new_array, da.array)
	da.array = new_array
	da.capacity = new_capacity
}

// for testing banking method invariant
func (da *DynArrayBanking[T]) GetCredit() int {
	return da.credit
}

// Complexity is O(n)
func (da *DynArrayBanking[T]) shrink() {
	new_capacity := int(float32(da.capacity) / SHRINK_FACTOR_BANKING)
	if new_capacity < 16 {
		new_capacity = 16
	}
	new_array := make([]T, new_capacity)
	copy(new_array, da.array)
	da.array = new_array
	da.capacity = new_capacity
}

// ===== JAGGED ARRAY ======
// package jagged_array

// Jagged Array (array of arrays) implementation for multidimensional dynamic array.
// Each dimension can grow independently, which addresses the requirement of Labaratory task.
// "easch dimension can expand on demand"
//
// But the thinkg is that Go doesn't support myArr[i,j,k] syntax from the labaratory assignment.
// Instead, I decided to use Get(i,j) and Set(value, i,j) methods.
//
// Trade-offs:
// + Each row can grow independently (true dynamic behavior per dimension)
// + Memory efficient for sparse data (rows can have different lengths)
// + Adding/removing rows is O(1) amortized
// - More pointer indirection = worse cache locality
// - Limited to specific dimensionality (2D, 3D need separate implementations)
//
// But I think that this realization is still better then emulation of jagged array with the only one array under the hood,
// because in this case too many reallocation will happen.

// JaggedArray2D represents a 2D array where each row can have independent length
type JaggedArray2D[T any] struct {
	rows DynArray[*DynArray[T]]
}

func (ja *JaggedArray2D[T]) Init() {
	ja.rows.Init()
}

// InitWithDimensions creates jagged array with specified initial dimensions
// rows x cols, but each row can grow independently later
func (ja *JaggedArray2D[T]) InitWithDimensions(rows, cols int) {
	ja.rows.Init()

	for i := 0; i < rows; i++ {
		row := &DynArray[T]{}
		row.Init()

		// Pre-fill with zero values
		var zero T
		for j := 0; j < cols; j++ {
			row.Append(zero)
		}

		ja.rows.Append(row)
	}
}

// Complexity: O(1)
func (ja *JaggedArray2D[T]) Get(row, col int) (T, error) {
	var zero T

	if row < 0 || row >= ja.rows.Count() {
		return zero, fmt.Errorf("row index out of bounds: %d", row)
	}

	rowArray, _ := ja.rows.GetItem(row)

	if col < 0 || col >= rowArray.Count() {
		return zero, fmt.Errorf("col index out of bounds: %d at row %d", col, row)
	}

	return rowArray.GetItem(col)
}

// Complexity: O(1) if no reallocation needed
func (ja *JaggedArray2D[T]) Set(value T, row, col int) error {
	if row < 0 || row >= ja.rows.Count() {
		return fmt.Errorf("row index out of bounds: %d", row)
	}

	rowArray, _ := ja.rows.GetItem(row)

	if col < 0 || col >= rowArray.Count() {
		return fmt.Errorf("col index out of bounds: %d at row %d", col, row)
	}

	// Direct assignment by reconstructing the element
	rowArray.array[col] = value
	return nil
}

// Complexity: Amortized O(1)
func (ja *JaggedArray2D[T]) AppendRow() {
	row := &DynArray[T]{}
	row.Init()
	ja.rows.Append(row)
}

// Complexity: Amortized O(1)
func (ja *JaggedArray2D[T]) AppendToRow(rowIndex int, value T) error {
	if rowIndex < 0 || rowIndex >= ja.rows.Count() {
		return fmt.Errorf("row index out of bounds: %d", rowIndex)
	}

	rowArray, _ := ja.rows.GetItem(rowIndex)
	rowArray.Append(value)
	return nil
}

// RowCount returns number of rows
func (ja *JaggedArray2D[T]) RowCount() int {
	return ja.rows.Count()
}

// RowLength returns length of specific row
func (ja *JaggedArray2D[T]) RowLength(rowIndex int) (int, error) {
	if rowIndex < 0 || rowIndex >= ja.rows.Count() {
		return 0, fmt.Errorf("row index out of bounds: %d", rowIndex)
	}

	rowArray, _ := ja.rows.GetItem(rowIndex)
	return rowArray.Count(), nil
}
