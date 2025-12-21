package main

// The constraints exp package is a requirement for the Laboratory test server.
// Perhaps an ordered list can be implemented with the cmp package in Go 1.21+

import (
	"constraints"
	"os"
)

var _ = os.Args

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	length     int
	_ascending bool
}

// O(1)
func (l *OrderedList[T]) Count() int {
	return l.length
}

// O(N)
func (l *OrderedList[T]) Add(item T) {
	newNode := &Node[T]{value: item}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	currentNode := l.head
	for currentNode != nil {
		comp := l.Compare(currentNode.value, item)
		shouldInsert := (l._ascending && comp >= 0) || (!l._ascending && comp <= 0)

		if shouldInsert {
			newNode.next = currentNode
			newNode.prev = currentNode.prev

			if currentNode.prev != nil {
				currentNode.prev.next = newNode
			} else {
				l.head = newNode
			}
			currentNode.prev = newNode

			l.length++
			return
		}
		currentNode = currentNode.next
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.length++

}

// O(n)
func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	if l.Count() <= 0 {
		return Node[T]{}, os.ErrNotExist
	}
	currentNode := l.head

	for currentNode != nil {
		comp := l.Compare(currentNode.value, n)

		if comp == 0 {
			return *currentNode, nil
		}

		// Early return: there's no point in searching further in a sorted list
		if (l._ascending && comp > 0) || (!l._ascending && comp < 0) {
			return Node[T]{}, os.ErrNotExist
		}

		currentNode = currentNode.next
	}

	return Node[T]{}, os.ErrNotExist
}

// O(n)
func (l *OrderedList[T]) Delete(n T) {
	if l.Count() <= 0 {
		return
	}
	currentNode := l.head

	for {
		if currentNode.value == n {
			if currentNode.prev == nil {
				l.head = currentNode.next
			} else {
				currentNode.prev.next = currentNode.next
			}
			if currentNode.next == nil {
				l.tail = currentNode.prev
			} else {
				currentNode.next.prev = currentNode.prev
			}
			l.length--
			return
		}
		currentNode = currentNode.next
		if currentNode == nil {
			return
		}

	}

}

// O(1)
func (l *OrderedList[T]) Clear(asc bool) {
	l.head = nil
	l.tail = nil
	l.length = 0
	l._ascending = asc
}

// O(1)
func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
}

// O(n)
func (l *OrderedList[T]) Dedup() {
	seenMap := make(map[T]bool)

	currentNode := l.head
	for currentNode != nil {
		if seenMap[currentNode.value] {
			currentNode.prev.next = currentNode.next
			if currentNode.next != nil {
				currentNode.next.prev = currentNode.prev
			} else {
				l.tail = currentNode.prev
			}
			l.length--
		} else {
			seenMap[currentNode.value] = true
		}

		currentNode = currentNode.next
	}

}

// O(n + m). O(n) if sublist length <= list length
func (l *OrderedList[T]) IsSublist(l2 *OrderedList[T]) bool {
	if l2.Count() <= 0 {
		return true
	}

	if l.Count() < l2.Count() {
		return false
	}

	lCurrentNode := l.head
	for lCurrentNode != nil && lCurrentNode.value != l2.head.value {
		lCurrentNode = lCurrentNode.next
	}

	if lCurrentNode == nil {
		return false
	}

	l2CurrentNode := l2.head
	for l2CurrentNode != nil {
		if lCurrentNode == nil || lCurrentNode.value != l2CurrentNode.value {
			return false
		}
		lCurrentNode = lCurrentNode.next
		l2CurrentNode = l2CurrentNode.next

	}

	return true

}

// O(n)
func (l *OrderedList[T]) TopFrequent() T {
	meetCounter := make(map[T]int)

	currentNode := l.head
	for currentNode != nil {
		meetCounter[currentNode.value]++
		currentNode = currentNode.next
	}

	var max int
	var maxValue T
	for key, value := range meetCounter {
		if value > max {
			max = value
			maxValue = key
		}
	}
	return maxValue
}

// O(1) - tail pointer allows constant time append
func (l *OrderedList[T]) addInTail(item T) {
	newNode := &Node[T]{value: item}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.length++
}
