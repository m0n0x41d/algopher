package linked_list

import "errors"

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

func (l *LinkedList) Count() int {
	return l.length
}

// error not nil, if node not found
func (l *LinkedList) Find(n int) (Node, error) {
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			return *current, nil
		}
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			nodes = append(nodes, *current)
		}
	}
	return nodes
}

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

func (l *LinkedList) DeleteFirst(n int, all bool) {
	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		l.length--
		if !all {
			return
		}
	}

}

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

func (l *LinkedList) Clean() {
	l.length = 0
	l.tail = nil
	l.head = nil

}
