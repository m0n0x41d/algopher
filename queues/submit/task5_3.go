package queues

import (
	"errors"
	"testing"
)

func TestQueueSizeAfterEnqueue(t *testing.T) {
	q := Queue[int]{}
	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	q.Enqueue(1)

	if q.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", q.Size())
	}

	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", q.Size())
	}

}

func TestQueueDequeue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 1 {
		t.Errorf("Expected dequeue to return 1, got %d", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 2 {
		t.Errorf("Expected dequeue to return 2, got %d", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	_, err = q.Dequeue()

	if err != ErrQueueEmpty {
		t.Errorf("Expected dequeue to return ErrQueueEmpty, got %v", err)
	}
}

func TestDequeueFromEmptyQueue(t *testing.T) {
	q := Queue[int]{}

	_, err := q.Dequeue()

	if !errors.Is(err, ErrQueueEmpty) {
		t.Errorf("Expected dequeue to return ErrQueueEmpty, got %v", err)
	}
}

func TestSizeAfterDequeue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 1 {
		t.Errorf("Expected size to be 1 after dequeue, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0 after dequeue, got %d", q.Size())
	}
}

func TestQueueReuse(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)

	q.Dequeue()
	q.Dequeue()

	q.Enqueue(3)
	q.Enqueue(4)

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 4 {
		t.Errorf("Expected dequeue to return 4, got %d", dequeued)
	}
}

func TestQueueWithStrings(t *testing.T) {
	q := Queue[string]{}
	q.Enqueue("first")
	q.Enqueue("second")

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != "first" {
		t.Errorf("Expected dequeue to return 'first', got %s", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != "second" {
		t.Errorf("Expected dequeue to return 'second', got %s", dequeued)
	}
}

func TestQueueWithZeroValues(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(0)
	q.Enqueue(0)

	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", q.Size())
	}

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 0 {
		t.Errorf("Expected dequeue to return 0, got %d", dequeued)
	}

	if q.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", q.Size())
	}
}

func TestQueueManyElements(t *testing.T) {
	q := Queue[int]{}

	for i := range 1000 {
		q.Enqueue(i)
	}

	if q.Size() != 1000 {
		t.Errorf("Expected size to be 1000, got %d", q.Size())
	}

	for i := range 1000 {
		dequeued, err := q.Dequeue()
		if err != nil {
			t.Errorf("Expected dequeue to not return an error, got %v", err)
		}
		if dequeued != i {
			t.Errorf("Expected dequeue to return %d, got %d", i, dequeued)
		}
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}
}

func TestReverseQueue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	reversed := ReverseQueue(q)

	dequeued, err := reversed.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 5 {
		t.Errorf("Expected dequeue to return 5, got %d", dequeued)
	}

	dequeued, err = reversed.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 4 {
		t.Errorf("Expected dequeue to return 4, got %d", dequeued)
	}

	dequeued, err = reversed.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	dequeued, err = reversed.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 2 {
		t.Errorf("Expected dequeue to return 2, got %d", dequeued)
	}

	dequeued, err = reversed.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 1 {
		t.Errorf("Expected dequeue to return 1, got %d", dequeued)
	}
}

// === additional tasks tests ===

func TestRotateQueue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	rotated := RotateQueue(q, 2)

	dequeued, err := rotated.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	dequeued, err = rotated.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 4 {
		t.Errorf("Expected dequeue to return 4, got %d", dequeued)
	}

	dequeued, err = rotated.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 5 {
		t.Errorf("Expected dequeue to return 5, got %d", dequeued)
	}

	dequeued, err = rotated.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 1 {
		t.Errorf("Expected dequeue to return 1, got %d", dequeued)
	}

	dequeued, err = rotated.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 2 {
		t.Errorf("Expected dequeue to return 2, got %d", dequeued)
	}
}

// === TwoStacksQueue tests ===

func TestTwoStacksQueueSize(t *testing.T) {
	q := TwoStacksQueue[int]{}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	q.Enqueue(1)

	if q.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", q.Size())
	}

	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", q.Size())
	}
}

func TestTwoStacksQueueDequeue(t *testing.T) {
	q := TwoStacksQueue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 1 {
		t.Errorf("Expected dequeue to return 1, got %d", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 2 {
		t.Errorf("Expected dequeue to return 2, got %d", dequeued)
	}

	dequeued, err = q.Dequeue()
	if err != nil {
		t.Errorf("Expected dequeue to not return an error, got %v", err)
	}
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	_, err = q.Dequeue()

	if !errors.Is(err, ErrQueueEmpty) {
		t.Errorf("Expected dequeue to return ErrQueueEmpty, got %v", err)
	}
}

func TestTwoStacksQueueDequeueFromEmpty(t *testing.T) {
	q := TwoStacksQueue[int]{}

	_, err := q.Dequeue()

	if !errors.Is(err, ErrQueueEmpty) {
		t.Errorf("Expected dequeue to return ErrQueueEmpty, got %v", err)
	}
}

func TestTwoStacksQueueIsEmpty(t *testing.T) {
	q := TwoStacksQueue[int]{}

	if !q.IsEmpty() {
		t.Errorf("Expected IsEmpty to return true")
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Errorf("Expected IsEmpty to return false")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("Expected IsEmpty to return true after dequeue")
	}
}

func TestTwoStacksQueueMixedOperations(t *testing.T) {
	q := TwoStacksQueue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)

	dequeued, _ := q.Dequeue()
	if dequeued != 1 {
		t.Errorf("Expected dequeue to return 1, got %d", dequeued)
	}

	q.Enqueue(3)
	q.Enqueue(4)

	dequeued, _ = q.Dequeue()
	if dequeued != 2 {
		t.Errorf("Expected dequeue to return 2, got %d", dequeued)
	}

	dequeued, _ = q.Dequeue()
	if dequeued != 3 {
		t.Errorf("Expected dequeue to return 3, got %d", dequeued)
	}

	dequeued, _ = q.Dequeue()
	if dequeued != 4 {
		t.Errorf("Expected dequeue to return 4, got %d", dequeued)
	}
}

func TestTwoStacksQueueManyElements(t *testing.T) {
	q := TwoStacksQueue[int]{}

	for i := range 1000 {
		q.Enqueue(i)
	}

	if q.Size() != 1000 {
		t.Errorf("Expected size to be 1000, got %d", q.Size())
	}

	for i := range 1000 {
		dequeued, err := q.Dequeue()
		if err != nil {
			t.Errorf("Expected dequeue to not return an error, got %v", err)
		}
		if dequeued != i {
			t.Errorf("Expected dequeue to return %d, got %d", i, dequeued)
		}
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}
}
