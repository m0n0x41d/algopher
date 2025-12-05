package main

import (
	"fmt"
	"reflect"
)

var _ = reflect.TypeOf("x")

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}
	l.tail = &item
	l.length++
}

func (l *LinkedList) Delete(n int, all bool) {
	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		l.length--
		if !all {
			return
		}
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

func main() {
	// Тест 1: удаляем голову, остаётся один элемент
	list := LinkedList{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})

	fmt.Printf("Before: head=%d, tail=%d, length=%d\n", list.head.value, list.tail.value, list.length)

	list.Delete(1, false)

	fmt.Printf("After delete head: head=%d, tail=%d, length=%d\n", list.head.value, list.tail.value, list.length)
	fmt.Printf("head == tail? %v\n", list.head == list.tail)

	// Тест 2: удаляем хвост, остаётся один элемент
	list2 := LinkedList{}
	list2.AddInTail(Node{value: 1})
	list2.AddInTail(Node{value: 2})

	fmt.Printf("\nBefore: head=%d, tail=%d, length=%d\n", list2.head.value, list2.tail.value, list2.length)

	list2.Delete(2, false)

	fmt.Printf("After delete tail: head=%d, tail=%d, length=%d\n", list2.head.value, list2.tail.value, list2.length)
	fmt.Printf("head == tail? %v\n", list2.head == list2.tail)
}
