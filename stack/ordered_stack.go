package stack

import "cmp"

// Numeric constraint for types that support arithmetic operations
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// === Stack with O(1) Min support ===
// Requires ordered types (int, float, string, etc.)
// Uses auxiliary stack to track minimum history.

type OrderedStack[T cmp.Ordered] struct {
	container []T
	minStack  []T
}

// === Stack with O(1) Min and Avg support ===
// Requires numeric types only (no strings).

type NumericStack[T Numeric] struct {
	container []T
	minStack  []T
	sum       T
}

// O(1)
func (st *OrderedStack[T]) Size() int {
	return len(st.container)
}

// O(1)
func (st *OrderedStack[T]) Peek() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	return st.container[len(st.container)-1], nil
}

// O(1)
func (st *OrderedStack[T]) Pop() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	result := st.container[len(st.container)-1]
	st.container = st.container[:len(st.container)-1]
	st.minStack = st.minStack[:len(st.minStack)-1]
	return result, nil
}

// Amortized O(1)
func (st *OrderedStack[T]) Push(itm T) {
	st.container = append(st.container, itm)
	if len(st.minStack) == 0 || itm < st.minStack[len(st.minStack)-1] {
		st.minStack = append(st.minStack, itm)
	} else {
		st.minStack = append(st.minStack, st.minStack[len(st.minStack)-1])
	}
}

// O(1) - returns current minimum element
func (st *OrderedStack[T]) Min() (T, error) {
	if len(st.minStack) == 0 {
		var zero T
		return zero, nil
	}
	return st.minStack[len(st.minStack)-1], nil
}

// === NumericStack methods ===

// O(1)
func (st *NumericStack[T]) Size() int {
	return len(st.container)
}

// O(1)
func (st *NumericStack[T]) Peek() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	return st.container[len(st.container)-1], nil
}

// O(1)
func (st *NumericStack[T]) Pop() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	result := st.container[len(st.container)-1]
	st.container = st.container[:len(st.container)-1]
	st.minStack = st.minStack[:len(st.minStack)-1]
	st.sum -= result
	return result, nil
}

// Amortized O(1)
func (st *NumericStack[T]) Push(itm T) {
	st.container = append(st.container, itm)
	st.sum += itm
	if len(st.minStack) == 0 || itm < st.minStack[len(st.minStack)-1] {
		st.minStack = append(st.minStack, itm)
	} else {
		st.minStack = append(st.minStack, st.minStack[len(st.minStack)-1])
	}
}

// O(1) - returns current minimum element
func (st *NumericStack[T]) Min() (T, error) {
	if len(st.minStack) == 0 {
		var zero T
		return zero, nil
	}
	return st.minStack[len(st.minStack)-1], nil
}

// O(1) - returns average of all elements
func (st *NumericStack[T]) Avg() float64 {
	if len(st.container) == 0 {
		return 0
	}
	return float64(st.sum) / float64(len(st.container))
}
