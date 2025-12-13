package queues

import "github.com/m0n0x41d/algopher/stack"

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
