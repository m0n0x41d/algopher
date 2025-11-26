package main

import (
	"errors"
	"os"
	"reflect"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// O(1) – благодаря указателю на хвост
func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
	l.length++
}

// O(1) - ячейка-счетчик в самой структуре
func (l *LinkedList) Count() int {
	return l.length
}

// O(n) – линейный перебор списка до встречи вхождения
func (l *LinkedList) Find(n int) (Node, error) {
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			return *current, nil
		}
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

// По времени: O(n) – линейный перебор всего списку.
// По памяти: O(k), где k – количество найденных элементо
func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			nodes = append(nodes, *current)
		}
	}
	return nodes
}

// O(n) при любом all – линейный перебор списка
// Используется паттерн "два указателя" (prev, current) для однонаправленного списка,
// т.к. нет возможности получить предыдущий узел напрямую.
// Удаление головы вынесено в отдельный цикл, чтобы избежать проверки prev == nil на каждой итерации.
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

// O(k), где k – количество удаляемых элементов, но помним что O(k) ⊆ O(n)
// В случае all:false – O(1)
// В случае all:ture - в worst case (все элементы списка оказываются одного целевого значения) это O(n),
func (l *LinkedList) DeleteFirst(n int, all bool) {
	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		l.length--
		if !all {
			return
		}
	}

}

// O(1) по указателям
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

// O(1) по указателям
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

// O(1) по указателям
// Память освобождается сборщиком мусора Go, когда на узлы не останется ссылок.
// В языках без GC вероятно потребовался бы явный обход и освобождение каждого узла.
func (l *LinkedList) Clean() {
	l.length = 0
	l.tail = nil
	l.head = nil

}
