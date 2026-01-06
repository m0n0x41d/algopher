package powersets

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

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
