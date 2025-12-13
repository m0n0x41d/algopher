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
