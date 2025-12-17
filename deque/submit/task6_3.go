package deque

import (
	"testing"
)

func TestDequeAddFrontAndSize(t *testing.T) {

	d := Deque[int]{}

	d.AddFront(1)
	d.AddFront(2)
	d.AddFront(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	pop1, err := d.RemoveFront()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop1 != 3 {
		t.Errorf("Expected 3 to be removed from the front of the deque, got %d", pop1)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2, got %d", d.Size())
	}

	pop2, err := d.RemoveFront()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop2 != 2 {
		t.Errorf("Expected 2 to be removed from the front of the deque, got %d", pop2)
	}

	if d.Size() != 1 {
		t.Errorf("Expected size 1, got %d", d.Size())
	}

	pop3, err := d.RemoveFront()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop3 != 1 {
		t.Errorf("Expected 1 to be removed from the front of the deque, got %d", pop3)
	}

	if d.Size() != 0 {
		t.Errorf("Expected size 0, got %d", d.Size())
	}

	_, err = d.RemoveFront()
	if err != ErrEmptyDeque {
		t.Errorf("Expected error %v, got %v", ErrEmptyDeque, err)
	}
}

func TestDequeAddTailAndSize(t *testing.T) {

	d := Deque[int]{}

	d.AddTail(1)
	d.AddTail(2)
	d.AddTail(3)

	if d.Size() != 3 {
		t.Errorf("Expected size 3, got %d", d.Size())
	}

	pop1, err := d.RemoveTail()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop1 != 3 {
		t.Errorf("Expected 3 to be removed from the tail of the deque, got %d", pop1)
	}

	if d.Size() != 2 {
		t.Errorf("Expected size 2, got %d", d.Size())
	}

	pop2, err := d.RemoveTail()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop2 != 2 {
		t.Errorf("Expected 2 to be removed from the tail of the deque, got %d", pop2)
	}

	if d.Size() != 1 {
		t.Errorf("Expected size 1, got %d", d.Size())
	}

	pop3, err := d.RemoveTail()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if pop3 != 1 {
		t.Errorf("Expected 1 to be removed from the tail of the deque, got %d", pop3)
	}

	if d.Size() != 0 {
		t.Errorf("Expected size 0, got %d", d.Size())
	}

	_, err = d.RemoveTail()
	if err != ErrEmptyDeque {
		t.Errorf("Expected error %v, got %v", ErrEmptyDeque, err)
	}
}

func TestIsPalindrome(t *testing.T) {
	palindrome := "racecar"
	not_palindrome := "youmuststudycs"

	x := IsPalindrome(palindrome)
	y := IsPalindrome(not_palindrome)

	if x != true {
		t.Errorf("Expected IsPalindrome to be true, got %v", x)
	}

	if y != false {
		t.Errorf("Expected IsPalindrome to be false, got %v", y)
	}
}
