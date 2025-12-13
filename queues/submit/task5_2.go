package main


// Funcs


// O(n) by time and space
// This is first realization using additional Split method from the Queue
func RotateQueue[T any](queue Queue[T], steps int) Queue[T] {
	result := Queue[T]{}

	first, second, _ := queue.Split(steps)

	for second.Size() > 0 {
		deq, _ := second.Dequeue()
		result.Enqueue(deq)
	}

	for first.Size() > 0 {
		deq, _ := first.Dequeue()
		result.Enqueue(deq)
	}

	return result
}

// Same O(n) by time and space due to slice
// reallocations on each Dequeue+Enqueue.
// Not a big difference with the implementation above, just a shorter code
func RotateQueueNoSplit[T any](queue Queue[T], steps int) Queue[T] {
	for range steps {
		item, _ := queue.Dequeue()
		queue.Enqueue(item)
	}
	return queue
}

// O(n) by time and space
func ReverseQueue[T any](queue Queue[T]) Queue[T] {
	stackForHelp := stack.Stack[T]{}

	resutQueue := Queue[T]{}

	for queue.Size() > 0 {
		item, _ := queue.Dequeue()
		stackForHelp.Push(item)
	}

	for stackForHelp.Size() > 0 {
		item, _ := stackForHelp.Pop()
		resutQueue.Enqueue(item)
	}

	return resutQueue
}


// two stack queue
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

// Cycled queue with fixed size

package queues

// Circular buffer queue with fixed capacity.
// Uses head/tail pointers to avoid data shifts — all operations O(1).
type CycledQueue[T any] struct {
	buffer []T
	head   int // index of first element
	tail   int // index of next free slot
	size   int
	cap    int
}

func NewCycledQueue[T any](capacity int) *CycledQueue[T] {
	return &CycledQueue[T]{
		buffer: make([]T, capacity),
		cap:    capacity,
	}
}

// O(1)
func (q *CycledQueue[T]) Size() int {
	return q.size
}

// O(1)
func (q *CycledQueue[T]) Capacity() int {
	return q.cap
}

// O(1)
func (q *CycledQueue[T]) IsFull() bool {
	return q.size == q.cap
}

// O(1)
func (q *CycledQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// O(1). Returns ErrQueueFull if queue is at capacity.
func (q *CycledQueue[T]) Enqueue(itm T) error {
	if q.IsFull() {
		return ErrQueueFull
	}

	q.buffer[q.tail] = itm
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return nil
}

// O(1). Returns ErrQueueEmpty if queue is empty.
func (q *CycledQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrQueueEmpty
	}

	out := q.buffer[q.head]

	// Clear reference to allow GC (important for pointer types)
	var zero T
	q.buffer[q.head] = zero

	q.head = (q.head + 1) % q.cap
	q.size--
	return out, nil
}

// O(1). Peek at head without removing.
func (q *CycledQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrQueueEmpty
	}
	return q.buffer[q.head], nil
}

// Reflection
//
// Both 2-stack queue and cycled queue are very nice implementations of the structure.
//
// Cycled queue is best because there is no amortization — O(1) asymptotic complexity
// is guaranteed due to the fixed allocated size. The only downside is that we need
// to predict max capacity accurately. CycledQueue is nice for real-time systems.
//
// 2-stack queue is also quite good, even though O(1) is amortized and from time to time
// we have to move elements into the second stack. The good thing is it's dynamic and
// we don't need to know max capacity upfront. It's still fast, just not as
// 'blazingly fast' as the cycled queue.
