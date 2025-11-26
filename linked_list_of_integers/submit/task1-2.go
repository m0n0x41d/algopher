package main


// Это наивное решение дополнительноый задачи, потому что в виду упомянутой слабости
// интерфейса связанного списка приходится использовать реализацию
// Эту ситуацию можно было бы улучшить добавив в Node геттер текущего значения и следующего узла, если он есть
// Либо можно было бы имплементировать интерфейс-итератор для самого списка, с Next и HasNext методами.
//
// Сложность сумматора выглядит так что в любом случае будет O(n) и по времени и по памяти, зависит от длинны входных списков.

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
