package main

import (
	"container/list"
	"errors"
	"os"
)

var _ = os.Args

// === easiest Stack with the tail of the Go slice ===
type Stack[T any] struct {
	container []T
}

// O(1)
func (st *Stack[T]) Size() int {
	return len(st.container)
}

// O(1)
func (st *Stack[T]) Peek() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	result := st.container[len(st.container)-1]
	return result, nil
}

// O(1). Potential issue with preserved pointers.
// On a very large stack with many Pops, this
// implementation will be inefficient in terms of memory usage.
func (st *Stack[T]) Pop() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	result := st.container[len(st.container)-1]
	st.container = st.container[:len(st.container)-1]
	return result, nil
}

// Amortized O(1), reallocate on slice growth
func (st *Stack[T]) Push(itm T) {
	st.container = append(st.container, itm)
}

// Regarding heterogeneous Stack. Well.. In Golang we can do:
// ```
// type Stack struct {
//   container []any
// }
//
// But it completely sacrifices type safety and will force us to type assert on every Pop.
// In my honest opinion this is smelling too bad to use in real scenarios.
//
// If going forward with generic implementation, and still wanting to support any type... Well
// we can write a constructor with the map of stack (slice) per type, which might be a bit costly on memory and
// will introduce a lot of boilerplate code, because chasing type safety we will need to implement Pop & Push for every type:
//   type MultiStack struct {
//     ints    []int
//     strings []string
//     // ...
//  }
//
//  func (m *MultiStack) PushInt(v int)    { ... }
//  func (m *MultiStack) PopInt() int      { ... }
//
//  Thus, in Golang it is better to stick with type-locked stacks. And in case we need to
//  implement a stack of Callable types... We can define a Golang interface and init stack of those.
//

// Task 2 - "Refactor the stack implementation to use the head of the list as the stack's top instead of the tail, using an appropriate data structure that preserves O(1) time complexity."
//
// We can use slices as underlying structure here, because it will lead to inevitable reallocation on every push into zero index.
// For being able precerve O(1) on push we need a singly linked list. Singly, because we do not need prev pointer for such stack implementation.
// For quiqier prototyping I will use here golang stdlib list, which is doubly linked, thus having a slight memory overhead for prev pointers.

type StackOnList[T any] struct {
	container *list.List
}

// O(1)
func (st *StackOnList[T]) Push(itm T) {
	st.container.PushFront(itm)
}

// O(1)
func (st *StackOnList[T]) Pop() (T, error) {
	front := st.container.Front()
	st.container.Remove(front)
	// We need type assetion here thus container/list is hoding `any`
	return front.Value.(T), nil
}

// O(1)
func (st *StackOnList[T]) Size() int {
	return st.container.Len()
}

// O(1)
func (st *StackOnList[T]) Peek() (T, error) {
	front := st.container.Front()
	return front.Value.(T), nil
}

// Overall - slice based implementation will be more efficicient in most of Stack-ish use cases.
