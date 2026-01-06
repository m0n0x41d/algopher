package main

import (
	//      "fmt"
	"os"
	"strconv"

	"constraints"
)

var _ = os.Stdout
var _ = strconv.Itoa

// An empty struct as a value is memory efficient in Golang because it leads to
// zero allocation.
type PowerSet[T constraints.Ordered] struct {
	storage map[T]struct{}
	count   int
}

func Init[T constraints.Ordered]() PowerSet[T] {
	return PowerSet[T]{storage: make(map[T]struct{})}
}

func (ps *PowerSet[T]) Size() int {
	return ps.count
}

func (ps *PowerSet[T]) Put(value T) {
	isExist := ps.Get(value)
	if !isExist {
		ps.storage[value] = struct{}{}
		ps.count++
	}
}

func (ps *PowerSet[T]) Get(value T) bool {
	_, ok := ps.storage[value]
	return ok
}

func (ps *PowerSet[T]) Remove(value T) bool {
	isExist := ps.Get(value)
	if isExist {
		delete(ps.storage, value)
		ps.count--
	}
	return isExist
}

// Time: O(min(N, M) linear, iterating over smallest
// Space: O(K), where K is amount of all similar elements
func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) PowerSet[T] {
	bigger, smaller := ps, &set2
	if ps.count < set2.count {
		bigger, smaller = &set2, ps
	}

	var result = Init[T]()

	for i := range smaller.storage {
		if bigger.Get(i) {
			result.Put(i)
		}
	}

	return result
}

// Time: O(N + M) where N and M are sizes of sets
// Space: O(max(N,M)), in worse case O(N + M) too, if sets are completely different
func (ps *PowerSet[T]) Union(set2 PowerSet[T]) PowerSet[T] {
	bigger, smaller := ps, &set2
	if ps.count < set2.count {
		bigger, smaller = &set2, ps
	}

	result := Init[T]()
	for k := range bigger.storage {
		result.storage[k] = struct{}{}
	}
	result.count = bigger.count

	for i := range smaller.storage {
		result.Put(i)
	}

	return result
}

// Time O(N) where N is size of ps set
// Space O(N) where N is size diffirent elements
func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) PowerSet[T] {
	result := Init[T]()

	for i := range ps.storage {
		if !set2.Get(i) {
			result.Put(i)
		}
	}
	return result
}

// Time O(N) where N is size of set2
// Space is ontouched.
func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) bool {
	if ps.count >= set2.count {
		for i := range set2.storage {
			if !ps.Get(i) {
				return false
			}
		}
		return true
	}
	return false
}

// Time: O(N) where N is size of set
// Space: No additional space used.
func (ps *PowerSet[T]) Equals(set2 PowerSet[T]) bool {
	if ps.count != set2.count {
		return false
	}

	for i := range ps.storage {
		if !set2.Get(i) {
			return false
		}
	}

	return true
}
