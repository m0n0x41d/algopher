package ordlist

import "golang.org/x/exp/constraints"

func MergeOrdLists[T constraints.Ordered](l1 *OrderedList[T], l2 *OrderedList[T]) *OrderedList[T] {
	if l1._ascending != l2._ascending {
		// There was no condition in the task to support merging of differently ordered lists.
		panic("cannot merge lists with different ordering")
	}

	mergedList := OrderedList[T]{_ascending: l1._ascending}

	l1CurrentNode := l1.head
	l2CurrentNode := l2.head
	for l1CurrentNode != nil && l2CurrentNode != nil {
		comp := mergedList.Compare(l1CurrentNode.value, l2CurrentNode.value)

		takeFromL1 := (l1._ascending && comp <= 0) || (!l1._ascending && comp >= 0)

		if takeFromL1 {
			mergedList.addInTail(l1CurrentNode.value)
			l1CurrentNode = l1CurrentNode.next
		} else {
			mergedList.addInTail(l2CurrentNode.value)
			l2CurrentNode = l2CurrentNode.next
		}
	}

	if l1CurrentNode != nil {
		for l1CurrentNode != nil {
			mergedList.addInTail(l1CurrentNode.value)
			l1CurrentNode = l1CurrentNode.next
		}
	} else {
		for l2CurrentNode != nil {
			mergedList.addInTail(l2CurrentNode.value)
			l2CurrentNode = l2CurrentNode.next
		}
	}

	return &mergedList
}
