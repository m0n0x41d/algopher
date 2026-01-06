package main

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

// space and time is O(N * M) where N and M are sizes of sets accordingly
func CartesianProduct[T constraints.Ordered](ps1, ps2 PowerSet[T]) [][2]T {
	// preallocate result slice to avoide reallocations on appends
	result := make([][2]T, 0, ps1.Size()*ps2.Size())

	for x := range ps1.storage {
		for y := range ps2.storage {
			result = append(result, [2]T{x, y})
		}

	}
	return result
}

// By Time: O(S * K) where S is size of smallest set, K is number of sets.
// Space: O(R) where R is size of resulting intersection.
// we are itergint over smallest set, trying to be more eficcient, but it will be still slow.
func IntersectMany[T constraints.Ordered](sets ...PowerSet[T]) PowerSet[T] {
	result := Init[T]()

	// The task stated clearly: three or more.
	if len(sets) < 3 {
		return result
	}

	smallestSetIndex := 0
	for i := 1; i < len(sets); i++ {
		if sets[i].count < sets[smallestSetIndex].count {
			smallestSetIndex = i
		}
	}

	for smallestSetElement := range sets[smallestSetIndex].storage {
		foundInAllSets := true
		for i, currentSet := range sets {
			// Skip smallest.
			if i == smallestSetIndex {
				continue
			}

			// can't be in intersect if it's not in all sets.
			if !currentSet.Get(smallestSetElement) {
				foundInAllSets = false
				break
			}
		}
		if foundInAllSets {
			result.Put(smallestSetElement)
		}
	}
	return result
}

// Bag

// Bag is a multi-set thing where each element can appear multiple times.
// I decided to just store counters as valued ¯\_(ツ)_/¯
// and count will be total number of elements (with dublications).
type Bag[T constraints.Ordered] struct {
	storage map[T]int
	count   int
}

func InitBag[T constraints.Ordered]() Bag[T] {
	return Bag[T]{storage: make(map[T]int)}
}

// Time: O(1)
func (b *Bag[T]) Size() int {
	return b.count
}

// Time: O(1)
func (b *Bag[T]) UniqueSize() int {
	return len(b.storage)
}

// Time: O(1)
func (b *Bag[T]) Put(value T) {
	b.storage[value]++
	b.count++
}

// Time: O(1)
func (b *Bag[T]) Get(value T) bool {
	return b.storage[value] > 0
}

// Time: O(1)
func (b *Bag[T]) Count(value T) int {
	return b.storage[value]
}

// Time: O(1)
func (b *Bag[T]) Remove(value T) bool {
	if b.storage[value] == 0 {
		return false
	}
	b.storage[value]--
	if b.storage[value] == 0 {
		delete(b.storage, value)
	}
	b.count--
	return true
}

// Time: O(1)
func (b *Bag[T]) RemoveAll(value T) int {
	count := b.storage[value]
	if count > 0 {
		delete(b.storage, value)
		b.count -= count
	}
	return count
}

// Time: O(N) where N is number of unique elements
// Space: O(N)
func (b *Bag[T]) Frequencies() map[T]int {
	return maps.Clone(b.storage)
}
