package main

import (
	"fmt"
	"os"
)

// Go manages pointers and memory allocation under the hood,
// so this Dynamic Array implementation is clearly "implementation for the sake of implementation" –
// the language provides slices as dynamic arrays out of the box.
// Nevertheless, implementing data structures like this is valuable even in high-level languages,
// as it demonstrates core concepts: amortized O(1) operations, multiplicative growth strategies,
// and the trade-offs between memory usage and performance.

var _ = os.Args

const (
	DEFAULT_MIN_BUFFER_SIZE = 16
	SHRINK_AT_PERCENT       = 50
)
const SHRINK_FACTOR float32 = 1.5

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
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
func (da *DynArray[T]) MakeArray(sz int) {
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
func (da *DynArray[T]) Insert(itm T, index int) error {
	if index < 0 || index > da.count {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	if da.count == da.capacity {
		da.growUp()
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
//   - Best case: O(1) - remowin last item without shrinking
//   - Worst case: O(n) - removing first item with shrinking
//   - Amortized: O(1) - again, thaks to multiplicative strategy
//
// by space:
//   - O(1) - regular shift
//   - O(n) - while shrinking
func (da *DynArray[T]) Remove(index int) error {
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
func (da *DynArray[T]) Append(itm T) {
	da.Insert(itm, da.count)
}

// O(1) - pointer access.
func (da *DynArray[T]) GetItem(index int) (T, error) {
	if index < 0 || index >= da.count {
		var zero T
		return zero, fmt.Errorf("index out of bounts: %d", index)
	}

	return da.array[index], nil
}

func (da *DynArray[T]) Count() int {
	return da.count
}

// Complexity is O(n)
func (da *DynArray[T]) growUp() {
	new_capacity := da.capacity * 2
	new_array := make([]T, new_capacity)
	copy(new_array, da.array)
	da.array = new_array
	da.capacity = new_capacity
}

// Complexity is O(n)
func (da *DynArray[T]) shrink() {
	new_capacity := int(float32(da.capacity) / SHRINK_FACTOR)
	if new_capacity < 16 {
		new_capacity = 16
	}
	new_array := make([]T, new_capacity)
	copy(new_array, da.array)
	da.array = new_array
	da.capacity = new_capacity
}
