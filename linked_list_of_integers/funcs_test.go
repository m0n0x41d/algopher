package linked_list

import "testing"

func TestNaiveSumListsDifferentLenth(t *testing.T) {
	list1 := LinkedList{}
	list2 := LinkedList{}
	list1.AddInTail(Node{value: 1})
	list1.AddInTail(Node{value: 2})
	list2.AddInTail(Node{value: 3})
	sumList := NaiveSumLists(&list1, &list2)

	want_length := 0
	if sumList.Count() != want_length {
		t.Errorf("want length %d, got %d", want_length, sumList.Count())
	}
}

func TestNaiveSumListsSameLength(t *testing.T) {
	list1 := LinkedList{}
	list2 := LinkedList{}
	list1.AddInTail(Node{value: 1})
	list1.AddInTail(Node{value: 2})
	list2.AddInTail(Node{value: 3})
	list2.AddInTail(Node{value: 4})
	sumList := NaiveSumLists(&list1, &list2)

	want_length := 2
	if sumList.Count() != want_length {
		t.Errorf("want length %d, got %d", want_length, sumList.Count())
	}

	want_sum_head := 4
	want_sum_tail := 6
	if sumList.head.value != want_sum_head {
		t.Errorf("want sum head %d, got %d", want_sum_head, sumList.head.value)
	}
	if sumList.tail.value != want_sum_tail {
		t.Errorf("want sum tail %d, got %d", want_sum_tail, sumList.tail.value)
	}
}

func TestNaiveSumListsEmpty(t *testing.T) {
	list1 := LinkedList{}
	list2 := LinkedList{}
	sumList := NaiveSumLists(&list1, &list2)

	want_length := 0
	if sumList.Count() != want_length {
		t.Errorf("want length %d, got %d", want_length, sumList.Count())
	}
}
