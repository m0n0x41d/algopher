package queues

import (
	"errors"
	"os"
)

var _ = os.Args

var (
	ErrQueueFull  = errors.New("queue is full")
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue[T any] struct {
	container []T
}

func (q *Queue[T]) Size() int {
	return len(q.container)
}

// Pure O(1) by time, but underlying array never shrinks.
// For production: we must implement some sort of shrink mechanism
// for example when size < capacity/4) to avoid memory waste on unused pointers.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.Size() == 0 {
		var zero T
		return zero, ErrQueueEmpty
	}

	out := q.container[0]
	q.container = q.container[1:]
	return out, nil
}

// O(1) amortized. Worst case O(n) when reallocation happens,
// but reallocs are rare (capacity doubles each time for underlying array in Go slices)
func (q *Queue[T]) Enqueue(itm T) {
	q.container = append(q.container, itm)
}

func (q *Queue[T]) Split(splitIndex int) (Queue[T], Queue[T], error) {
	firstQueue := Queue[T]{}
	firstQueue.container = q.container[:splitIndex]
	secondQueue := Queue[T]{}
	secondQueue.container = q.container[splitIndex:]

	return firstQueue, secondQueue, nil

}
