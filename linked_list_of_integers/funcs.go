package linked_list

// This is a naive solution for the additional task because due to the mentioned weakness
// of the linked list interface we have to access implementation details directly.
// This situation could be improved by adding getters for current value and next node to Node.
// Alternatively, an iterator interface could be implemented for the list itself with Next and HasNext methods.
//
// Complexity: O(n) for both time and space, depends on the length of input lists.

func NaiveSumLists(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	sumList := LinkedList{}
	if list1.length != list2.length {
		return &sumList
	}

	list1Current := list1.head
	list2Current := list2.head
	for list1Current != nil && list2Current != nil {
		sumList.AddInTail(Node{value: list1Current.value + list2Current.value})
		list1Current = list1Current.next
		list2Current = list2Current.next
	}

	return &sumList
}
