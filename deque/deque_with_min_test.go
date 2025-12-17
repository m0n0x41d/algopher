package deque

import (
	"testing"
)

func TestMinDequeBasicOperations(t *testing.T) {
	d := MinDeque[int]{}

	d.AddTail(3)
	d.AddTail(1)
	d.AddTail(2)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	min, err := d.Min()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}
}

func TestMinDequeMinUpdatesOnRemove(t *testing.T) {
	d := MinDeque[int]{}

	d.AddTail(2)
	d.AddTail(1)
	d.AddTail(3)

	// Min should be 1
	min, _ := d.Min()
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}

	// Remove 2 from front
	d.RemoveFront()
	min, _ = d.Min()
	if min != 1 {
		t.Errorf("Expected min 1 after removing 2, got %d", min)
	}

	// Remove 1 from front, min should now be 3
	d.RemoveFront()
	min, _ = d.Min()
	if min != 3 {
		t.Errorf("Expected min 3 after removing 1, got %d", min)
	}
}

func TestMinDequeAddFront(t *testing.T) {
	d := MinDeque[int]{}

	d.AddFront(3)
	d.AddFront(5)
	d.AddFront(1)

	min, _ := d.Min()
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}

	// Deque should be [1, 5, 3]
	front, _ := d.PeekFront()
	if front != 1 {
		t.Errorf("Expected front 1, got %d", front)
	}

	tail, _ := d.PeekTail()
	if tail != 3 {
		t.Errorf("Expected tail 3, got %d", tail)
	}
}

func TestMinDequeRemoveTail(t *testing.T) {
	d := MinDeque[int]{}

	d.AddTail(3)
	d.AddTail(1)
	d.AddTail(2)

	// Remove 2 from tail
	removed, _ := d.RemoveTail()
	if removed != 2 {
		t.Errorf("Expected removed 2, got %d", removed)
	}

	min, _ := d.Min()
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}

	// Remove 1 from tail, min should be 3
	d.RemoveTail()
	min, _ = d.Min()
	if min != 3 {
		t.Errorf("Expected min 3, got %d", min)
	}
}

func TestMinDequeEmptyErrors(t *testing.T) {
	d := MinDeque[int]{}

	_, err := d.Min()
	if err != ErrEmptyMinDeque {
		t.Errorf("Expected ErrEmptyMinDeque, got %v", err)
	}

	_, err = d.RemoveFront()
	if err != ErrEmptyMinDeque {
		t.Errorf("Expected ErrEmptyMinDeque, got %v", err)
	}

	_, err = d.RemoveTail()
	if err != ErrEmptyMinDeque {
		t.Errorf("Expected ErrEmptyMinDeque, got %v", err)
	}

	_, err = d.PeekFront()
	if err != ErrEmptyMinDeque {
		t.Errorf("Expected ErrEmptyMinDeque, got %v", err)
	}

	_, err = d.PeekTail()
	if err != ErrEmptyMinDeque {
		t.Errorf("Expected ErrEmptyMinDeque, got %v", err)
	}
}

func TestMinDequeMixedOperations(t *testing.T) {
	d := MinDeque[int]{}

	d.AddTail(5)
	d.AddFront(3)
	d.AddTail(7)
	d.AddFront(1)

	// Deque: [1, 3, 5, 7]
	min, _ := d.Min()
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}

	d.RemoveFront() // remove 1
	min, _ = d.Min()
	if min != 3 {
		t.Errorf("Expected min 3, got %d", min)
	}

	d.RemoveTail() // remove 7
	min, _ = d.Min()
	if min != 3 {
		t.Errorf("Expected min 3, got %d", min)
	}

	d.AddTail(2)
	min, _ = d.Min()
	if min != 2 {
		t.Errorf("Expected min 2, got %d", min)
	}
}
