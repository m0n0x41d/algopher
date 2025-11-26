package main

import (
	"errors"
	"os"
	"reflect"
)

var _, _ = os.Args, reflect.TypeOf(0)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// O(1) - tail pointer allows constant time append
func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
	l.length++
}

// O(1) - length counter stored in struct
func (l *LinkedList) Count() int {
	return l.length
}

// O(n) - linear search until match found
func (l *LinkedList) Find(n int) (Node, error) {
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			return *current, nil
		}
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

// Time: O(n) - full list traversal
// Space: O(k) where k is number of found elements
func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			nodes = append(nodes, *current)
		}
	}
	return nodes
}

// O(n) for any value of all - linear list traversal
// Uses two-pointer pattern (prev, current) for singly linked list
// since there is no way to get previous node directly.
// Head deletion extracted to separate loop to avoid prev == nil check on each iteration.
func (l *LinkedList) Delete(n int, all bool) {
	for l.head != nil && l.head.value == n {
		l.DeleteFirst(n, all)
	}

	if l.head == nil {
		l.tail = nil
		return
	}

	prevNode := l.head
	currentNode := l.head.next

	for currentNode != nil {
		if currentNode.value == n {
			prevNode.next = currentNode.next
			l.length--
			if currentNode == l.tail {
				l.tail = prevNode
			}
			if !all {
				return
			}
			currentNode = prevNode.next
		} else {
			prevNode = currentNode
			currentNode = currentNode.next
		}
	}
}

// O(k) where k is number of deleted elements, but O(k) is subset of O(n)
// When all:false - O(1)
// When all:true - worst case (all elements have target value) is O(n)
func (l *LinkedList) DeleteFirst(n int, all bool) {
	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		l.length--
		if !all {
			return
		}
	}

}

// O(1) - pointer operations only
func (l *LinkedList) Insert(after *Node, add Node) {
	if l.tail == after {
		l.tail.next = &add
		l.tail = &add
		l.length++
		return
	}

	add.next = after.next
	after.next = &add
	l.length++

}

// O(1) - pointer operations only
func (l *LinkedList) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
		l.length++
		return
	}
	first.next = l.head
	l.head = &first
	l.length++

}

// O(1) - pointer operations only
// Memory is freed by Go garbage collector when no references to nodes remain.
// In languages without GC explicit traversal and deallocation of each node would be required.
func (l *LinkedList) Clean() {
	l.length = 0
	l.tail = nil
	l.head = nil

}


