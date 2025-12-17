package main

import (
	"cmp"
	"container/list"
	"errors"
)

func IsPalindrome(s string) bool {
	deq := Deque[rune]{}

	for _, r := range s {
		deq.AddFront(r)
	}

	for deq.Size() > 1 {
		front, _ := deq.RemoveFront()
		tail, _ := deq.RemoveTail()
		if front != tail {
			return false
		}
	}

	return true
}

// linked list for base deque. Nice to have O(1) at both ends.

var ErrEmptyLinkedDeque = errors.New("LinkedDeque is empty")

type LinkedDeque[T any] struct {
	container *list.List
}

func (d *LinkedDeque[T]) init() {
	if d.container == nil {
		d.container = list.New()
	}
}

func (d *LinkedDeque[T]) Size() int {
	d.init()
	return d.container.Len()
}

func (d *LinkedDeque[T]) AddFront(itm T) {
	d.init()
	d.container.PushFront(itm)
}

func (d *LinkedDeque[T]) AddTail(itm T) {
	d.init()
	d.container.PushBack(itm)
}

func (d *LinkedDeque[T]) RemoveFront() (T, error) {
	d.init()
	if d.container.Len() == 0 {
		var zero T
		return zero, ErrEmptyLinkedDeque
	}

	elem := d.container.Front()
	d.container.Remove(elem)
	return elem.Value.(T), nil
}

func (d *LinkedDeque[T]) RemoveTail() (T, error) {
	d.init()
	if d.container.Len() == 0 {
		var zero T
		return zero, ErrEmptyLinkedDeque
	}

	elem := d.container.Back()
	d.container.Remove(elem)
	return elem.Value.(T), nil
}

// MinDeque provides O(1) minimum retrieval using two stacks with min-tracking.
//
// Complexity:
//   - AddFront: O(1)
//   - AddTail:  O(1)
//   - RemoveFront: O(1) amortized (may trigger O(n) rebalance when front stack exhausted)
//   - RemoveTail:  O(1) amortized (may trigger O(n) rebalance when tail stack exhausted)
//   - Min:      O(1) - this was the main goal of the Laboratory Task
//   - PeekFront/PeekTail: O(1)
//
// Why two stacks?
// A single min-stack only tracks minimum from one direction (bottom to top).
// When we pop from the opposite end, we lose min information — the min was
// computed relative to the other end. With two stacks (front and tail), each
// tracks its own minimum independently. The global minimum is simply
// min(front.min, tail.min). When one stack empties, we rebalance by moving
// half the elements from the other stack, rebuilding min information for
// the new direction.
//
// I decided to make this Min method in a separate implementation to avoid overloading the main deque.go for the reviewer and because we need Peek methods and all the related functionality.
//
// Also, this implementation is unsafe for the last additional task, which is stated as:
// "Implement a deque using a dynamic array. The methods for adding and removing elements from both ends of the deque should run in amortized time O(1)."
//
// Slices in Go are dynamic arrays under the hood, so I believe it counts and does not require reusing my own educational implementation of a dynamic array from this repository.

var ErrEmptyMinDeque = errors.New("MinDeque is empty")

type minStackEntry[T cmp.Ordered] struct {
	value T
	min   T
}

type MinDeque[T cmp.Ordered] struct {
	front []minStackEntry[T] // stack growing toward front
	tail  []minStackEntry[T] // stack growing toward tail
}

func (d *MinDeque[T]) Size() int {
	return len(d.front) + len(d.tail)
}

func (d *MinDeque[T]) PeekFront() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyMinDeque
	}
	if len(d.front) > 0 {
		return d.front[len(d.front)-1].value, nil
	}
	return d.tail[0].value, nil
}

func (d *MinDeque[T]) PeekTail() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyMinDeque
	}
	if len(d.tail) > 0 {
		return d.tail[len(d.tail)-1].value, nil
	}
	return d.front[0].value, nil
}

func (d *MinDeque[T]) AddFront(itm T) {
	minVal := itm
	if len(d.front) > 0 && d.front[len(d.front)-1].min < minVal {
		minVal = d.front[len(d.front)-1].min
	}
	d.front = append(d.front, minStackEntry[T]{value: itm, min: minVal})
}

func (d *MinDeque[T]) AddTail(itm T) {
	minVal := itm
	if len(d.tail) > 0 && d.tail[len(d.tail)-1].min < minVal {
		minVal = d.tail[len(d.tail)-1].min
	}
	d.tail = append(d.tail, minStackEntry[T]{value: itm, min: minVal})
}

func (d *MinDeque[T]) RemoveFront() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyMinDeque
	}

	if len(d.front) > 0 {
		result := d.front[len(d.front)-1].value
		d.front = d.front[:len(d.front)-1]
		return result, nil
	}

	// front is empty, need to redistribute from tail
	d.rebalanceToFront()
	result := d.front[len(d.front)-1].value
	d.front = d.front[:len(d.front)-1]
	return result, nil
}

func (d *MinDeque[T]) RemoveTail() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyMinDeque
	}

	if len(d.tail) > 0 {
		result := d.tail[len(d.tail)-1].value
		d.tail = d.tail[:len(d.tail)-1]
		return result, nil
	}

	// tail is empty, need to redistribute from front
	d.rebalanceToTail()
	result := d.tail[len(d.tail)-1].value
	d.tail = d.tail[:len(d.tail)-1]
	return result, nil
}

func (d *MinDeque[T]) rebalanceToFront() {
	// Move half of tail elements to front
	mid := len(d.tail) / 2
	if mid == 0 {
		mid = 1
	}

	// Elements [0, mid) go to front (in reverse order)
	for i := mid - 1; i >= 0; i-- {
		d.AddFront(d.tail[i].value)
	}

	// Rebuild tail with remaining elements [mid, len)
	remaining := make([]T, len(d.tail)-mid)
	for i := mid; i < len(d.tail); i++ {
		remaining[i-mid] = d.tail[i].value
	}

	d.tail = nil
	for _, v := range remaining {
		d.AddTail(v)
	}
}

func (d *MinDeque[T]) rebalanceToTail() {
	// Move half of front elements to tail
	mid := len(d.front) / 2
	if mid == 0 {
		mid = 1
	}

	// Elements [0, mid) go to tail (in reverse order)
	for i := mid - 1; i >= 0; i-- {
		d.AddTail(d.front[i].value)
	}

	// Rebuild front with remaining elements [mid, len)
	remaining := make([]T, len(d.front)-mid)
	for i := mid; i < len(d.front); i++ {
		remaining[i-mid] = d.front[i].value
	}

	d.front = nil
	for _, v := range remaining {
		d.AddFront(v)
	}
}

func (d *MinDeque[T]) Min() (T, error) {
	if d.Size() == 0 {
		var zero T
		return zero, ErrEmptyMinDeque
	}

	if len(d.front) == 0 {
		return d.tail[len(d.tail)-1].min, nil
	}
	if len(d.tail) == 0 {
		return d.front[len(d.front)-1].min, nil
	}

	frontMin := d.front[len(d.front)-1].min
	tailMin := d.tail[len(d.tail)-1].min
	if frontMin < tailMin {
		return frontMin, nil
	}
	return tailMin, nil
}
