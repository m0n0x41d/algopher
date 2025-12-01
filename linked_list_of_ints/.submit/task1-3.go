package main

import "testing"

func TestCountEmpty(t *testing.T) {
	list := LinkedList{}

	got := list.Count()
	want := 0

	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}
}

func TestCountThreeElements(t *testing.T) {
	list := LinkedList{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	got := list.Count()
	want := 3

	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}
}

func TestFindNodeOk(t *testing.T) {
	list := LinkedList{}
	nodeToFind := Node{value: 777}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(nodeToFind)
	list.AddInTail(Node{value: 3})

	got, err := list.Find(777)
	if err != nil {
		t.Errorf("Failed to find node.")
	}

	if got.value != nodeToFind.value {
		t.Errorf("Find() = %v, want %v", got.value, nodeToFind.value)
	}
}

func TestFindNodeError(t *testing.T) {
	list := LinkedList{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	_, err := list.Find(10)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}
}

func TestFindAllOk(t *testing.T) {
	list := LinkedList{}
	nodeToFind := Node{value: 777}
	nodeToFind2 := Node{value: 777}

	list.AddInTail(nodeToFind)
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(nodeToFind2)

	got := list.FindAll(777)

	if len(got) != 2 {
		t.Errorf("FindAll() = %v, want %v", len(got), 2)
	}
}

func TestDeleteHeadNodeOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}

	list.AddInTail(nodeToDelete)
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	list.Delete(777, false)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 3
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}

	got_head_value := list.head.value
	want := 1
	if got_head_value != want {
		t.Errorf("Head value = %v, want %v", got_head_value, want)
	}
}

func TestDeleteTailNodeOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(nodeToDelete)

	list.Delete(777, false)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 3
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}

	got_tail_value := list.tail.value
	want := 3
	if got_tail_value != want {
		t.Errorf("Tail value = %v, want %v", got_tail_value, want)
	}
}

func TestDeleteFromTheMiddleOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(nodeToDelete)
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	list.Delete(777, false)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 4
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}
}

func TestDeleteAllFromTheHeadOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}
	nodeToDelete2 := Node{value: 777}

	list.AddInTail(nodeToDelete)
	list.AddInTail(nodeToDelete2)
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	list.Delete(777, true)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 3
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}
}

func TestDeleteAllFromTheTailOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}
	nodeToDelete2 := Node{value: 777}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(nodeToDelete)
	list.AddInTail(nodeToDelete2)

	list.Delete(777, true)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 3
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}
}

func TestDeleteAllFromTheMiddleOk(t *testing.T) {
	list := LinkedList{}
	nodeToDelete := Node{value: 777}
	nodeToDelete2 := Node{value: 777}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(nodeToDelete)
	list.AddInTail(nodeToDelete2)
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	list.Delete(777, true)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 4
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}
}

func TestCleanOk(t *testing.T) {
	list := LinkedList{}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	list.Clean()

	length_after_deletion := list.Count()
	want_length_after_deletion := 0
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}

	var want_head, want_tail *Node = nil, nil
	if list.head != want_head {
		t.Errorf("Head = %v, want %v", list.head, want_head)
	}
	if list.tail != want_tail {
		t.Errorf("Tail = %v, want %v", list.tail, want_tail)
	}
}

func TestInsertAfterHeadOk(t *testing.T) {
	list := LinkedList{}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	nodeToInsert := Node{value: 777}
	list.Insert(list.head, nodeToInsert)

	got := list.head.value
	want := 1
	if got != want {
		t.Errorf("Head value = %v, want %v", got, want)
	}

	got = list.head.next.value
	want = 777
	if got != want {
		t.Errorf("Head next value = %v, want %v", got, want)
	}

	got = list.head.next.next.value
	want = 2
	if got != want {
		t.Errorf("Head next next value = %v, want %v", got, want)
	}

	got_length := list.Count()
	want_length := 4
	if got_length != want_length {
		t.Errorf("Length = %v, want %v", got_length, want_length)
	}
}

func TestInsertAfterTailOk(t *testing.T) {
	list := LinkedList{}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	nodeToInsert := Node{value: 777}
	list.Insert(list.tail, nodeToInsert)

	got := list.tail.value
	want := 777
	if got != want {
		t.Errorf("Tail value = %v, want %v", got, want)
	}

	if list.tail.next != nil {
		t.Errorf("Tail next value = %v, want %v", got, want)
	}

	got_length := list.Count()
	want_length := 4
	if got_length != want_length {
		t.Errorf("Length = %v, want %v", got_length, want_length)
	}
}

func TestInsertFirstOk(t *testing.T) {
	list := LinkedList{}

	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	nodeToInsert := Node{value: 777}
	list.InsertFirst(nodeToInsert)

	got := list.head.value
	want := 777
	if got != want {
		t.Errorf("Head value = %v, want %v", got, want)
	}

	got = list.head.next.value
	want = 1
	if got != want {
		t.Errorf("Head next value = %v, want %v", got, want)
	}

	got = list.head.next.next.value
	want = 2
	if got != want {
		t.Errorf("Head next next value = %v, want %v", got, want)
	}

	got_length := list.Count()
	want_length := 4
	if got_length != want_length {
		t.Errorf("Length = %v, want %v", got_length, want_length)
	}
}

func TestInsertFirstInEmptyOk(t *testing.T) {
	list := LinkedList{}

	nodeToInsert := Node{value: 777}
	list.InsertFirst(nodeToInsert)

	got := list.head.value
	want := 777
	if got != want {
		t.Errorf("Head value = %v, want %v", got, want)
	}

	got = list.tail.value
	want = 777
	if got != want {
		t.Errorf("Head next value = %v, want %v", got, want)
	}

	got_length := list.Count()
	want_length := 1
	if got_length != want_length {
		t.Errorf("Length = %v, want %v", got_length, want_length)
	}
}

func TestDeleteFromEmptyList(t *testing.T) {
	list := LinkedList{}

	list.Delete(777, false)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 0
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}
}

func TestDeleteOneNodeFromOneNodeList(t *testing.T) {
	list := LinkedList{}
	list.AddInTail(Node{value: 777})

	list.Delete(777, false)

	_, err := list.Find(777)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}

	length_after_deletion := list.Count()
	want_length_after_deletion := 0
	if length_after_deletion != want_length_after_deletion {
		t.Errorf("Length after deletion = %v, want %v", length_after_deletion, want_length_after_deletion)
	}

	if list.tail != nil {
		t.Errorf("Tail = %v, want %v", list.tail, nil)
	}
	if list.head != nil {
		t.Errorf("Head = %v, want %v", list.head, nil)
	}
}

func TestFindAllFromEmptyList(t *testing.T) {
	list := LinkedList{}

	got := list.FindAll(777)

	if len(got) != 0 {
		t.Errorf("FindAll() = %v, want %v", len(got), 0)
	}
}

func TestDeleteHeadFromTwoNodesList(t *testing.T) {
	list := LinkedList{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})

	list.Delete(1, false)

	got := list.Count()
	want := 1
	if got != want {
		t.Errorf("Length after deletion = %v, want %v", got, want)
	}

	got = list.head.value
	want = 2
	if got != want {
		t.Errorf("Head value = %v, want %v", got, want)
	}

	got = list.tail.value
	want = 2
	if got != want {
		t.Errorf("Tail value = %v, want %v", got, want)
	}

	if list.head != list.tail {
		t.Errorf("head != tail, but should be same node")
	}
	if list.head.next != nil {
		t.Errorf("head.next = %v, want nil", list.head.next)
	}
}

func TestDeleteTailFromTwoNodesList(t *testing.T) {
	list := LinkedList{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})

	list.Delete(2, false)

	got := list.Count()
	want := 1
	if got != want {
		t.Errorf("Length after deletion = %v, want %v", got, want)
	}

	got = list.head.value
	want = 1
	if got != want {
		t.Errorf("Head value = %v, want %v", got, want)
	}

	got = list.tail.value
	want = 1
	if got != want {
		t.Errorf("Tail value = %v, want %v", got, want)
	}

	if list.head != list.tail {
		t.Errorf("head != tail, but should be same node")
	}
}

// Tests for NaiveSumLists
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
