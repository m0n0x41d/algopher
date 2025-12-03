package doubly_linked_list_of_ints

import (
	"errors"
	"os"
	"reflect"
)

var _, _ = os.Args, reflect.TypeOf(0)

type SortStrategy string

const (
	MergeSort SortStrategy = "merge"
)

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head   *Node
	tail   *Node
	length int
}

// O(1) - tail pointer allows constant time append
func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
	l.length++
}

// O(1) - length counter stored in struct
func (l *LinkedList2) Count() int {
	return l.length
}

// O(n) - linear search until match found
func (l *LinkedList2) Find(n int) (Node, error) {
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			return *current, nil
		}
	}
	return Node{value: -1, next: nil}, errors.New("Node not found")
}

// Time: O(n) - full list traversal
// Space: O(k) where k is number of found elements
func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	for current := l.head; current != nil; current = current.next {
		if current.value == n {
			nodes = append(nodes, *current)
		}
	}
	return nodes
}

// O(n) for any value of all - linear list traversal
func (l *LinkedList2) Delete(n int, all bool) {
	for l.head != nil && l.head.value == n {
		if l.head.next != nil {
			l.head.next.prev = nil
		}
		l.head = l.head.next
		l.length--
		if l.head == nil {
			l.tail = nil
		}
		if !all {
			return
		}
	}

	if l.head == nil {
		l.tail = nil
		return
	}

	currentNode := l.head.next

	for currentNode != nil {

		if currentNode.value == n {
			currentNode.prev.next = currentNode.next
			if currentNode.next != nil {
				currentNode.next.prev = currentNode.prev
			}
			l.length--
			if currentNode == l.tail {
				l.tail = currentNode.prev
			}
			// Continue only if `all` is True.
			if !all {
				return
			}
			currentNode = currentNode.next
		} else {
			currentNode = currentNode.next
		}
	}
}

// O(1) - pointer operations only
func (l *LinkedList2) Insert(after *Node, add Node) {
	if l.tail == after {
		add.prev = l.tail
		l.tail.next = &add
		l.tail = &add
		l.length++
		return
	}

	add.next = after.next
	add.prev = after
	after.next.prev = &add
	after.next = &add
	l.length++

}

// O(1) - pointer operations only
func (l *LinkedList2) InsertFirst(first Node) {
	if l.head == nil {
		l.head = &first
		l.tail = &first
		l.length++
	} else {
		first.next = l.head
		l.head.prev = &first
		l.head = &first
		l.length++
	}
}

// O(1) - pointer operations only
// Memory is freed by Go garbage collector when no references to nodes remain.
// In languages without GC explicit traversal and deallocation of each node would be required.
func (l *LinkedList2) Clean() {
	l.length = 0
	l.tail = nil
	l.head = nil

}

// O(n) - linear traversal
func (l *LinkedList2) Reverse() {
	if l.Count() <= 1 {
		return
	}

	currentNode := l.head

	for {
		nextNode := currentNode.next
		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next

		if nextNode == nil {
			break
		}

		currentNode = nextNode

	}

	l.tail, l.head = l.head, l.tail

}

// Floyds Cycle Detection – O(n) by time and O(1) by memory, due to
// pointers usage.
func (l *LinkedList2) isLooped() bool {
	if l.head == nil {
		return false
	}
	slow := l.head
	fast := l.head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false

}

// MergeSort delegated O(n log n) — other sort strategies
// are not implemented.
func (l *LinkedList2) Sort(strategy SortStrategy) {
	if strategy == "" {
		strategy = "merge"
	}

	if l.head == nil || l.head.next == nil {
		return
	}

	switch strategy {
	case MergeSort:
		result := l.mergeSort()
		l.head = result.head
		l.tail = result.tail
		l.length = result.length
	default:
		result := l.mergeSort()
		l.head = result.head
		l.tail = result.tail
		l.length = result.length
	}
}

// O(n log n) by time, and O(log n) by memory dues to stack recursion.
func (l *LinkedList2) mergeSort() *LinkedList2 {
	if l.head == nil || l.head.next == nil {
		return l
	}

	middleNode := l.findMiddleNode()
	rigthList := l.splitAt(middleNode)
	leftList := l

	left := leftList.mergeSort()
	right := rigthList.mergeSort()

	return l.Merge(left, right)

}

// First of all – this method is exposing another issue with the LinkedList2 interface:
// LinkedList2 needs a method for popping out a node, like a head from the stack.
// This is needed not only for separating the interface from the implementation
// but at least for correctly updating the list length counter.
//
// O(n log n + m log m) – where (n log n) and (m log m) are complexities of
// merge sorts used under the hood for both of the passed lists.
// There is also an additional O(n + m) on top of that – sorted lists merging cycle,
// but it is "absorbed" by a greater complexity expression.
func (l *LinkedList2) Merge(left, right *LinkedList2) *LinkedList2 {
	left.Sort(MergeSort)
	right.Sort(MergeSort)
	mergedList := LinkedList2{}

	for left.head != nil && right.head != nil {
		if left.head.value <= right.head.value {
			mergedList.AddInTail(*left.head)
			left.head = left.head.next
		} else {
			mergedList.AddInTail(*right.head)
			right.head = right.head.next
		}
	}

	for left.head != nil {
		mergedList.AddInTail(*left.head)
		left.head = left.head.next
	}

	for right.head != nil {
		mergedList.AddInTail(*right.head)
		right.head = right.head.next
	}

	return &mergedList
}

// O(n) - linear traversal
func (l *LinkedList2) findMiddleNode() *Node {
	if l.head == nil || l.head.next == nil {
		return nil
	}
	slow := l.head
	fast := l.head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// O(n) - linear traversal for length recalculation
func (l *LinkedList2) splitAt(node *Node) *LinkedList2 {
	rightList := LinkedList2{}
	rightList.head = node.next
	rightList.tail = l.tail

	node.next.prev = nil
	node.next = nil

	l.tail = node

	originalLength := l.Count()
	leftLength := 0
	for current := l.head; current != nil; current = current.next {
		leftLength++
	}

	l.length = leftLength
	rightList.length = originalLength - leftLength

	return &rightList

}

// Additional reflexion:
// As I mentioned in the comment on the Merge method, the linked list interface design continues to reveal
// its shortcomings - it requires the use of structure fields to implement "external methods."
// Even though these methods are external exclusively to the user, it seems they should be "external"
// primarily to the implementation itself.
// 1. About violating encapsulation
// The Merge method directly manipulates left.head and right.head,
// bypassing the public interface. This violates encapsulation and makes
// the code brittle - changing the internal structure of LinkedList2
// will break Merge, even though it is formally an "external" method.
//
// 2. On the need for the Pop/Shift operation
// The absence of the PopHead() method (remove and return the head) forces us to
// write `left.head = left.head.next` directly, which:
// 1) Doesn't update the length
// 2) Doesn't clear the prev of the new head
// 3) Doesn't handle the case when the list becomes empty
// A proper interface should provide atomic operations.
//
// 3. On the receiver of the Merge method
// Merge is called as l.Merge(left, right), but the receiver `l`
// is not used at all—the method creates a new list.
// This should be either a function (not a method) or a constructor:
// func MergeLists(left, right *LinkedList2) *LinkedList2
//
// 4. On the mutation of input parameters
// Merge mutates the input lists (left.head, right.head change).
// After calling Merge, the original lists are in an invalid
// state — head is shifted, but length is not updated.
// Ideally, everything should be rewritten in the spirit of functional immutability
// and return new list instances, albeit at the cost of memory – Go's GC is good,
// and the additional guarantees provided by immutability are invaluable
// Ideally, everything should be rewritten in the spirit of functional immutability
// and return new list instances, albeit at the cost of memory – Go's GC is good,
// and the additional guarantees provided by immutability are invaluable.
//
// 5. General Conclusion
// Conclusion: A linked list as a data structure is conceptually simple,
// but designing a clean API for it is non-trivial (as for any non-trivial type).
// Each new operation reveals the flaws of the underlying interface.
// It would certainly be more correct to start by defining the full set of operations
// via interfaces, following the Abstract Data Types approach. But we know that's the context of a separate course.
