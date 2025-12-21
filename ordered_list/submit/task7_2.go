package main

import (
	"os"
	"slices"

	"golang.org/x/exp/constraints"
)

// Note for a reviewer: The rest of the * tasks are implemented as methods in main implementation for the server.

type IndexedOrderedList[T constraints.Ordered] struct {
	items      []T
	_ascending bool
}

// O(1)
func (l *IndexedOrderedList[T]) Count() int {
	return len(l.items)
}

// O(n) - binary search O(log n) + insert O(n)
func (l *IndexedOrderedList[T]) Add(item T) {
	idx := l.findInsertIndex(item)
	l.items = slices.Insert(l.items, idx, item)
}

// O(log n)
func (l *IndexedOrderedList[T]) Find(n T) (T, error) {
	idx, found := l.binarySearch(n)
	if !found {
		var zero T
		return zero, os.ErrNotExist
	}
	return l.items[idx], nil
}

// O(log n) - the main reason for this implementation
func (l *IndexedOrderedList[T]) FindIndex(n T) (int, error) {
	idx, found := l.binarySearch(n)
	if !found {
		return -1, os.ErrNotExist
	}
	return idx, nil
}

// O(n) - binary search O(log n) + delete O(n)
func (l *IndexedOrderedList[T]) Delete(n T) {
	idx, found := l.binarySearch(n)
	if !found {
		return
	}
	l.items = slices.Delete(l.items, idx, idx+1)
}

// O(1)
func (l *IndexedOrderedList[T]) Clear(asc bool) {
	l.items = nil
	l._ascending = asc
}

// O(n)
func (l *IndexedOrderedList[T]) Dedup() {
	if len(l.items) <= 1 {
		return
	}

	seen := make(map[T]bool)
	result := make([]T, 0, len(l.items))

	for _, item := range l.items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	l.items = result
}

// O(n + m)
func (l *IndexedOrderedList[T]) IsSublist(l2 *IndexedOrderedList[T]) bool {
	if l2.Count() == 0 {
		return true
	}
	if l.Count() < l2.Count() {
		return false
	}

	// Find start of potential sublist
	startIdx, found := l.binarySearch(l2.items[0])
	if !found {
		return false
	}

	// Check if remaining elements match
	for i := 0; i < len(l2.items); i++ {
		if startIdx+i >= len(l.items) || l.items[startIdx+i] != l2.items[i] {
			return false
		}
	}

	return true
}

// O(n)
func (l *IndexedOrderedList[T]) TopFrequent() T {
	if len(l.items) == 0 {
		var zero T
		return zero
	}

	meetCounter := make(map[T]int)
	for _, item := range l.items {
		meetCounter[item]++
	}

	var max int
	var maxValue T
	for key, value := range meetCounter {
		if value > max {
			max = value
			maxValue = key
		}
	}
	return maxValue
}

// O(log n) - binary search for element
func (l *IndexedOrderedList[T]) binarySearch(n T) (int, bool) {
	if len(l.items) == 0 {
		return 0, false
	}

	if l._ascending {
		return slices.BinarySearch(l.items, n)
	}

	// For descending: reverse comparison
	low, high := 0, len(l.items)
	for low < high {
		mid := (low + high) / 2
		if l.items[mid] > n {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if low < len(l.items) && l.items[low] == n {
		return low, true
	}
	return low, false
}

// O(log n) - find index where item should be inserted
func (l *IndexedOrderedList[T]) findInsertIndex(item T) int {
	if len(l.items) == 0 {
		return 0
	}

	if l._ascending {
		idx, _ := slices.BinarySearch(l.items, item)
		return idx
	}

	// For descending: find first element <= item
	low, high := 0, len(l.items)
	for low < high {
		mid := (low + high) / 2
		if l.items[mid] > item {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// O(1)
func (l *IndexedOrderedList[T]) Get(idx int) (T, error) {
	if idx < 0 || idx >= len(l.items) {
		var zero T
		return zero, os.ErrNotExist
	}
	return l.items[idx], nil
}
