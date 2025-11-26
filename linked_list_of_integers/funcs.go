package linked_list

func NaiveSumLists(list1 *LinkedList, list2 *LinkedList) *LinkedList {
	sumList := LinkedList{}
	if list1.length != list2.length {
		return &sumList
	}

	list1Сurrent := list1.head
	list2Сurrent := list2.head
	for list1Сurrent != nil && list2Сurrent != nil {
		sumList.AddInTail(Node{value: list1Сurrent.value + list2Сurrent.value})
		list1Сurrent = list1Сurrent.next
		list2Сurrent = list2Сurrent.next
	}

	return &sumList

}
