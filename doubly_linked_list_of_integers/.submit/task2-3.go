package main

import "testing"

func TestCountEmpty(t *testing.T) {
	list := LinkedList2{}

	got := list.Count()
	want := 0

	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}
}

func TestCountThreeElements(t *testing.T) {
	list := LinkedList2{}
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
	list := LinkedList2{}
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
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	_, err := list.Find(10)
	if err == nil {
		t.Errorf("Should have failed to find node.")
	}
}

func TestFindAllOk(t *testing.T) {
	list := LinkedList2{}
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
	list := LinkedList2{}
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
	list := LinkedList2{}
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
	list := LinkedList2{}
	nodeBefore := Node{value: 2}
	nodeToDelete := Node{value: 777}
	nodeAfter := Node{value: 3}

	list.AddInTail(Node{value: 1})
	list.AddInTail(nodeBefore)
	list.AddInTail(nodeToDelete)
	list.AddInTail(nodeAfter)
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

	if list.tail.prev.value != nodeAfter.value {
		t.Errorf("nodeAfter.value = %v, want %v", nodeAfter.value, list.tail.prev.value)
	}
	if list.head.next.value != nodeBefore.value {
		t.Errorf("nodeBefore.value = %v, want %v", nodeBefore.value, list.head.next.value)
	}
}

func TestDeleteAllFromTheHeadOk(t *testing.T) {
	list := LinkedList2{}
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
	list := LinkedList2{}
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
	list := LinkedList2{}
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
	list := LinkedList2{}

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
	list := LinkedList2{}

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

	if list.head.next.prev == nil {
		t.Errorf("Head next prev = %v, want %v", list.head.next.prev, list.head)
	}

	if list.head.next.next.prev.value != 777 {
		t.Errorf("Head next next prev value = %v, want %v", list.head.next.next.prev.value, 777)
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
	list := LinkedList2{}

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
	list := LinkedList2{}

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
	list := LinkedList2{}

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
	list := LinkedList2{}

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
	list := LinkedList2{}
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
	list := LinkedList2{}

	got := list.FindAll(777)

	if len(got) != 0 {
		t.Errorf("FindAll() = %v, want %v", len(got), 0)
	}
}

func TestDeleteHeadFromTwoNodesList(t *testing.T) {
	list := LinkedList2{}
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
	list := LinkedList2{}
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

// Additional tasks tests

func TestReverseOneNodeList(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})

	list.Reverse()

	got := list.Count()
	want := 1
	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}

	if list.head != list.tail {
		t.Errorf("head != tail, but should be same node")
	}

	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}

	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}

	if list.head.next != nil {
		t.Errorf("head.next = %v, want nil", list.head.next)
	}
}

func TestReverseTwoNodesList(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})

	list.Reverse()
	got := list.Count()
	want := 2
	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}

	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}

	if list.head.next == nil {
		t.Errorf("head.next = %v, want not nil", list.head.next)
	}

	if list.head.value != 2 {
		t.Errorf("head.value = %v, want 2", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("tail.value = %v, want 1", list.tail.value)
	}

	if list.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", list.tail.next)
	}

	if list.tail.prev == nil {
		t.Errorf("tail.prev = %v, want not nil", list.tail.prev)
	}

	if list.head.next != list.tail {
		t.Errorf("head.next != tail, but should be same node")
	}

	if list.tail.prev != list.head {
		t.Errorf("tail.prev != head, but should be same node")
	}
}

func TestReverseListManyNodes(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})
	list.AddInTail(Node{value: 5})

	list.Reverse()
	got := list.Count()
	want := 5
	if got != want {
		t.Errorf("Count() = %v, want %v", got, want)
	}

	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}

	if list.head.next == nil {
		t.Errorf("head.next = %v, want not nil", list.head.next)
	}

	if list.head.value != 5 {
		t.Errorf("head.value = %v, want 5", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("tail.value = %v, want 1", list.tail.value)
	}

	if list.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", list.tail.next)
	}

	if list.tail.prev == nil {
		t.Errorf("tail.prev = %v, want not nil", list.tail.prev)
	}

	if list.head.next.value != 4 {
		t.Errorf("head.next.value = %v, want 4", list.head.next.value)
	}

	if list.head.next.next.value != 3 {
		t.Errorf("head.next.next.value = %v, want 3", list.head.next.next.value)
	}

	if list.head.next.next.next.value != 2 {
		t.Errorf("head.next.next.next.value = %v, want 2", list.head.next.next.next.value)
	}

	if list.tail.prev.value != 2 {
		t.Errorf("tail.prev.value = %v, want 2", list.tail.prev.value)
	}

}

func TestIsLoopedEmptyList(t *testing.T) {
	list := LinkedList2{}

	got := list.isLooped()
	want := false
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestIsLoopedNoLoop(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	got := list.isLooped()
	want := false
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestIsLoopedWithLoop(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	// create loop: tail.next points back to head
	list.tail.next = list.head

	got := list.isLooped()
	want := true
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestIsLoopedWithLoopToMiddle(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	// create loop: tail.next points to second node
	list.tail.next = list.head.next

	got := list.isLooped()
	want := true
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestIsLoopedSingleNode(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})

	got := list.isLooped()
	want := false
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestIsLoopedSingleNodeWithLoop(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})

	// create loop: node points to itself
	list.head.next = list.head

	got := list.isLooped()
	want := true
	if got != want {
		t.Errorf("isLooped() = %v, want %v", got, want)
	}
}

func TestFindMiddleNodeEmptyList(t *testing.T) {
	list := LinkedList2{}
	got := list.findMiddleNode()
	if got != nil {
		t.Errorf("findMiddleNode() = %v, want %v", got, nil)
	}
}

func TestFindMiddleNodeEvenLengthList(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	got := list.findMiddleNode()
	want := list.head.next
	if got != want {
		t.Errorf("findMiddleNode() = %v, want %v", got, want)
	}
}

func TestFindMiddleNodeOddLengthList(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	got := list.findMiddleNode()
	want := list.head.next
	if got != want {
		t.Errorf("findMiddleNode() = %v, want %v", got, want)
	}
}

func TestSplitAtMiddle(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})

	middle := list.findMiddleNode()
	rightList := list.splitAt(middle)

	// check left list
	if list.Count() != 2 {
		t.Errorf("left list Count() = %v, want 2", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("left list head.value = %v, want 1", list.head.value)
	}
	if list.tail.value != 2 {
		t.Errorf("left list tail.value = %v, want 2", list.tail.value)
	}
	if list.tail.next != nil {
		t.Errorf("left list tail.next = %v, want nil", list.tail.next)
	}

	// check right list
	if rightList.Count() != 2 {
		t.Errorf("right list Count() = %v, want 2", rightList.Count())
	}
	if rightList.head.value != 3 {
		t.Errorf("right list head.value = %v, want 3", rightList.head.value)
	}
	if rightList.tail.value != 4 {
		t.Errorf("right list tail.value = %v, want 4", rightList.tail.value)
	}
	if rightList.head.prev != nil {
		t.Errorf("right list head.prev = %v, want nil", rightList.head.prev)
	}
}

func TestSplitAtMiddleOddLength(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 4})
	list.AddInTail(Node{value: 5})

	middle := list.findMiddleNode()
	rightList := list.splitAt(middle)

	// check left list (1 -> 2 -> 3)
	if list.Count() != 3 {
		t.Errorf("left list Count() = %v, want 3", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("left list head.value = %v, want 1", list.head.value)
	}
	if list.tail.value != 3 {
		t.Errorf("left list tail.value = %v, want 3", list.tail.value)
	}

	// check right list (4 -> 5)
	if rightList.Count() != 2 {
		t.Errorf("right list Count() = %v, want 2", rightList.Count())
	}
	if rightList.head.value != 4 {
		t.Errorf("right list head.value = %v, want 4", rightList.head.value)
	}
	if rightList.tail.value != 5 {
		t.Errorf("right list tail.value = %v, want 5", rightList.tail.value)
	}
}

func TestSplitAtTwoElements(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})

	middle := list.findMiddleNode()
	rightList := list.splitAt(middle)

	// check left list
	if list.Count() != 1 {
		t.Errorf("left list Count() = %v, want 1", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("left list head.value = %v, want 1", list.head.value)
	}
	if list.head != list.tail {
		t.Errorf("left list head != tail, but should be same node")
	}

	// check right list
	if rightList.Count() != 1 {
		t.Errorf("right list Count() = %v, want 1", rightList.Count())
	}
	if rightList.head.value != 2 {
		t.Errorf("right list head.value = %v, want 2", rightList.head.value)
	}
	if rightList.head != rightList.tail {
		t.Errorf("right list head != tail, but should be same node")
	}
}

func TestSortEmptyList(t *testing.T) {
	list := LinkedList2{}

	list.Sort(MergeSort)

	if list.Count() != 0 {
		t.Errorf("Count() = %v, want 0", list.Count())
	}
	if list.head != nil {
		t.Errorf("head = %v, want nil", list.head)
	}
	if list.tail != nil {
		t.Errorf("tail = %v, want nil", list.tail)
	}
}

func TestSortSingleElement(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})

	list.Sort(MergeSort)

	if list.Count() != 1 {
		t.Errorf("Count() = %v, want 1", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}
	if list.head != list.tail {
		t.Errorf("head != tail, but should be same node")
	}
}

func TestSortTwoElements(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 1})

	list.Sort(MergeSort)

	if list.Count() != 2 {
		t.Errorf("Count() = %v, want 2", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}
	if list.tail.value != 2 {
		t.Errorf("tail.value = %v, want 2", list.tail.value)
	}
	if list.head.next != list.tail {
		t.Errorf("head.next != tail")
	}
	if list.tail.prev != list.head {
		t.Errorf("tail.prev != head")
	}
	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}
	if list.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", list.tail.next)
	}
}

func TestSortAlreadySorted(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 3})

	list.Sort(MergeSort)

	if list.Count() != 3 {
		t.Errorf("Count() = %v, want 3", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}
	if list.head.next.value != 2 {
		t.Errorf("head.next.value = %v, want 2", list.head.next.value)
	}
	if list.tail.value != 3 {
		t.Errorf("tail.value = %v, want 3", list.tail.value)
	}
}

func TestSortReversed(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 1})

	list.Sort(MergeSort)

	if list.Count() != 3 {
		t.Errorf("Count() = %v, want 3", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}
	if list.head.next.value != 2 {
		t.Errorf("head.next.value = %v, want 2", list.head.next.value)
	}
	if list.tail.value != 3 {
		t.Errorf("tail.value = %v, want 3", list.tail.value)
	}
	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}
	if list.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", list.tail.next)
	}
}

func TestSortWithDuplicates(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 3})
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 1})

	list.Sort(MergeSort)

	if list.Count() != 4 {
		t.Errorf("Count() = %v, want 4", list.Count())
	}
	if list.head.value != 1 {
		t.Errorf("head.value = %v, want 1", list.head.value)
	}
	if list.head.next.value != 1 {
		t.Errorf("head.next.value = %v, want 1", list.head.next.value)
	}
	if list.head.next.next.value != 2 {
		t.Errorf("head.next.next.value = %v, want 2", list.head.next.next.value)
	}
	if list.tail.value != 3 {
		t.Errorf("tail.value = %v, want 3", list.tail.value)
	}
}

func TestSortManyElements(t *testing.T) {
	list := LinkedList2{}
	list.AddInTail(Node{value: 5})
	list.AddInTail(Node{value: 2})
	list.AddInTail(Node{value: 8})
	list.AddInTail(Node{value: 1})
	list.AddInTail(Node{value: 9})
	list.AddInTail(Node{value: 3})

	list.Sort(MergeSort)

	if list.Count() != 6 {
		t.Errorf("Count() = %v, want 6", list.Count())
	}

	// check sorted order by traversing forward
	expected := []int{1, 2, 3, 5, 8, 9}
	current := list.head
	for i, want := range expected {
		if current == nil {
			t.Errorf("list ended early at index %d", i)
			break
		}
		if current.value != want {
			t.Errorf("element at index %d = %v, want %v", i, current.value, want)
		}
		current = current.next
	}

	// check prev links by traversing backward
	current = list.tail
	for i := len(expected) - 1; i >= 0; i-- {
		if current == nil {
			t.Errorf("list ended early at index %d (backward)", i)
			break
		}
		if current.value != expected[i] {
			t.Errorf("element at index %d (backward) = %v, want %v", i, current.value, expected[i])
		}
		current = current.prev
	}

	// check boundary conditions
	if list.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", list.head.prev)
	}
	if list.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", list.tail.next)
	}
}

func TestMergeTwoEmptyLists(t *testing.T) {
	left := LinkedList2{}
	right := LinkedList2{}
	list := LinkedList2{}

	result := list.Merge(&left, &right)

	if result.Count() != 0 {
		t.Errorf("Count() = %v, want 0", result.Count())
	}
	if result.head != nil {
		t.Errorf("head = %v, want nil", result.head)
	}
	if result.tail != nil {
		t.Errorf("tail = %v, want nil", result.tail)
	}
}

func TestMergeLeftEmpty(t *testing.T) {
	left := LinkedList2{}
	right := LinkedList2{}
	right.AddInTail(Node{value: 1})
	right.AddInTail(Node{value: 2})

	list := LinkedList2{}
	result := list.Merge(&left, &right)

	if result.Count() != 2 {
		t.Errorf("Count() = %v, want 2", result.Count())
	}
	if result.head.value != 1 {
		t.Errorf("head.value = %v, want 1", result.head.value)
	}
	if result.tail.value != 2 {
		t.Errorf("tail.value = %v, want 2", result.tail.value)
	}
}

func TestMergeRightEmpty(t *testing.T) {
	left := LinkedList2{}
	left.AddInTail(Node{value: 3})
	left.AddInTail(Node{value: 4})
	right := LinkedList2{}

	list := LinkedList2{}
	result := list.Merge(&left, &right)

	if result.Count() != 2 {
		t.Errorf("Count() = %v, want 2", result.Count())
	}
	if result.head.value != 3 {
		t.Errorf("head.value = %v, want 3", result.head.value)
	}
	if result.tail.value != 4 {
		t.Errorf("tail.value = %v, want 4", result.tail.value)
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	left := LinkedList2{}
	left.AddInTail(Node{value: 1})
	left.AddInTail(Node{value: 3})
	left.AddInTail(Node{value: 5})

	right := LinkedList2{}
	right.AddInTail(Node{value: 2})
	right.AddInTail(Node{value: 4})
	right.AddInTail(Node{value: 6})

	list := LinkedList2{}
	result := list.Merge(&left, &right)

	if result.Count() != 6 {
		t.Errorf("Count() = %v, want 6", result.Count())
	}

	expected := []int{1, 2, 3, 4, 5, 6}
	current := result.head
	for i, want := range expected {
		if current == nil {
			t.Errorf("list ended early at index %d", i)
			break
		}
		if current.value != want {
			t.Errorf("element at index %d = %v, want %v", i, current.value, want)
		}
		current = current.next
	}

	if result.head.prev != nil {
		t.Errorf("head.prev = %v, want nil", result.head.prev)
	}
	if result.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", result.tail.next)
	}
}

func TestMergeTwoUnsortedLists(t *testing.T) {
	left := LinkedList2{}
	left.AddInTail(Node{value: 5})
	left.AddInTail(Node{value: 1})
	left.AddInTail(Node{value: 3})

	right := LinkedList2{}
	right.AddInTail(Node{value: 6})
	right.AddInTail(Node{value: 2})
	right.AddInTail(Node{value: 4})

	list := LinkedList2{}
	result := list.Merge(&left, &right)

	if result.Count() != 6 {
		t.Errorf("Count() = %v, want 6", result.Count())
	}

	// result should be sorted even though inputs were unsorted
	expected := []int{1, 2, 3, 4, 5, 6}
	current := result.head
	for i, want := range expected {
		if current == nil {
			t.Errorf("list ended early at index %d", i)
			break
		}
		if current.value != want {
			t.Errorf("element at index %d = %v, want %v", i, current.value, want)
		}
		current = current.next
	}
}

func TestMergeWithDuplicates(t *testing.T) {
	left := LinkedList2{}
	left.AddInTail(Node{value: 1})
	left.AddInTail(Node{value: 2})
	left.AddInTail(Node{value: 2})

	right := LinkedList2{}
	right.AddInTail(Node{value: 2})
	right.AddInTail(Node{value: 3})

	list := LinkedList2{}
	result := list.Merge(&left, &right)

	if result.Count() != 5 {
		t.Errorf("Count() = %v, want 5", result.Count())
	}

	expected := []int{1, 2, 2, 2, 3}
	current := result.head
	for i, want := range expected {
		if current == nil {
			t.Errorf("list ended early at index %d", i)
			break
		}
		if current.value != want {
			t.Errorf("element at index %d = %v, want %v", i, current.value, want)
		}
		current = current.next
	}
}
