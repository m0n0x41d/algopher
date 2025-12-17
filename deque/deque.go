package deque

import (
	"errors"
	"os"
	// "fmt"
)

var _ = os.Args

var ErrEmptyDeque = errors.New("Deque is empty")

// A naive implementation using a slice as a container for a deque will result
// in AddFront always being O(n) due to copying all elements to the new slice.
// AddTail will be amortized O(1). The whole purpose of a deque, when
// this structure is indeed needed, is to have aligned complexity for both
// ends' operations of the deque. Thus, a good idea is to use a doubly linked
// list; this will align both ends' add and remove operations as O(1), and
// the slight overhead on pointers is acceptable in this case.
//
// This implementation is kept with the slice because the laboratory test server awaits no
// changes in signatures. Check deque_linked.go for the doubly linked list implementation.

type Deque[T any] struct {
	container []T
}

func (d *Deque[T]) Size() int {
	return len(d.container)
}

func (d *Deque[T]) AddFront(itm T) {
	d.container = append([]T{itm}, d.container...)
}

func (d *Deque[T]) AddTail(itm T) {
	d.container = append(d.container, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyDeque
	}

	result := d.container[0]
	d.container = d.container[1:]

	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	if d.Size() <= 0 {
		var zero T
		return zero, ErrEmptyDeque
	}

	result := d.container[d.Size()-1]
	d.container = d.container[:d.Size()-1]

	return result, nil
}
