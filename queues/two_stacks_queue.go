package queues

import (
	"github.com/m0n0x41d/algopher/stack"
)

type TwoStacksQueue[T any] struct {
	inStack  stack.Stack[T]
	outStack stack.Stack[T]
}

// O(1)
func (q *TwoStacksQueue[T]) Size() int {
	return q.inStack.Size() + q.outStack.Size()
}

// O(1)
func (q *TwoStacksQueue[T]) Enqueue(itm T) {
	q.inStack.Push(itm)
}

// Amortized O(1), worst case O(n) only when outStack is empty,
// but we are moving items between stack not too often
func (q *TwoStacksQueue[T]) Dequeue() (T, error) {
	if q.Size() == 0 {
		var zero T
		return zero, ErrQueueEmpty
	}

	if q.outStack.Size() <= 0 {
		for q.inStack.Size() > 0 {
			itm, err := q.inStack.Pop()
			if err != nil {
				return *new(T), err
			}
			q.outStack.Push(itm)
		}
	}
	out, err := q.outStack.Pop()
	if err != nil {
		return out, err
	}
	return out, nil
}

func (q *TwoStacksQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Nice thing, love the efficency.
