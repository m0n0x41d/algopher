package main

import "testing"

func TestAddOrderAndCountAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(6)
	list.Add(10)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 3 {
		t.Errorf("Expected head.next to be 3, got %d", list.head.next.value)
	}

	if list.tail.value != 10 {
		t.Errorf("Expected tail to be 10, got %d", list.tail.value)
	}

	if list.tail.prev.value != 6 {
		t.Errorf("Expected tail.prev to be 6, got %d", list.tail.prev.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestClearNonEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Clear(true)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items in list, got %d", list.Count())
	}

	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("Expected tail to be nil, got %v", list.tail)
	}
}

func TestClearEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	list.Clear(false)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items in list, got %d", list.Count())
	}

	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("Expected tail to be nil, got %v", list.tail)
	}

	if list._ascending != false {
		t.Errorf("Expected _ascending to be false, got %v", list._ascending)
	}
}

func TestClearChangesOrder(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	list.Clear(false)

	list.Add(1)
	list.Add(3)

	if list.head.value != 3 {
		t.Errorf("Expected head to be 3, got %d", list.head.value)
	}

	if list.tail.value != 1 {
		t.Errorf("Expected tail to be 1, got %d", list.tail.value)
	}
}

func TestDeleteHead(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(1)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items in list, got %d", list.Count())
	}

	if list.head.value != 3 {
		t.Errorf("Expected head to be 3, got %d", list.head.value)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestDeleteTail(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(5)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items in list, got %d", list.Count())
	}

	if list.tail.value != 3 {
		t.Errorf("Expected tail to be 3, got %d", list.tail.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}
}

func TestDeleteMiddle(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(3)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 5 {
		t.Errorf("Expected head.next to be 5, got %d", list.head.next.value)
	}

	if list.tail.prev.value != 1 {
		t.Errorf("Expected tail.prev to be 1, got %d", list.tail.prev.value)
	}
}

func TestDeleteOnlyElement(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)

	list.Delete(5)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items in list, got %d", list.Count())
	}

	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("Expected tail to be nil, got %v", list.tail)
	}
}

func TestDeleteNonExistent(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(999)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestDeleteFromEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	list.Delete(5)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items in list, got %d", list.Count())
	}
}

func TestDeleteDuplicates(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	list.Delete(5)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestAddEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)

	if list.Count() != 1 {
		t.Errorf("Expected 1 item in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}

	if list.head != list.tail {
		t.Errorf("Expected head and tail to be the same node")
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}

	if list.head.next != nil {
		t.Errorf("Expected head.next to be nil, got %v", list.head.next)
	}
}

func TestAddDuplicatesAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.head.next.value != 5 {
		t.Errorf("Expected head.next to be 5, got %d", list.head.next.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestAddDuplicatesDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.head.next.value != 5 {
		t.Errorf("Expected head.next to be 5, got %d", list.head.next.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestAddInsertAtHeadAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(10)
	list.Add(20)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.head.next.value != 10 {
		t.Errorf("Expected head.next to be 10, got %d", list.head.next.value)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestAddInsertAtHeadDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(5)
	list.Add(10)
	list.Add(20)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 20 {
		t.Errorf("Expected head to be 20, got %d", list.head.value)
	}

	if list.head.next.value != 10 {
		t.Errorf("Expected head.next to be 10, got %d", list.head.next.value)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestAddInsertAtTailAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(10)
	list.Add(20)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.tail.value != 20 {
		t.Errorf("Expected tail to be 20, got %d", list.tail.value)
	}

	if list.tail.prev.value != 10 {
		t.Errorf("Expected tail.prev to be 10, got %d", list.tail.prev.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}
}

func TestAddInsertAtTailDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(20)
	list.Add(10)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}

	if list.tail.prev.value != 10 {
		t.Errorf("Expected tail.prev to be 10, got %d", list.tail.prev.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}
}

func TestAddInsertInMiddleAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(10)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 5 {
		t.Errorf("Expected head.next to be 5, got %d", list.head.next.value)
	}

	if list.tail.value != 10 {
		t.Errorf("Expected tail to be 10, got %d", list.tail.value)
	}

	middle := list.head.next
	if middle.prev.value != 1 {
		t.Errorf("Expected middle.prev to be 1, got %d", middle.prev.value)
	}

	if middle.next.value != 10 {
		t.Errorf("Expected middle.next to be 10, got %d", middle.next.value)
	}
}

func TestAddInsertInMiddleDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(10)
	list.Add(1)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 10 {
		t.Errorf("Expected head to be 10, got %d", list.head.value)
	}

	if list.head.next.value != 5 {
		t.Errorf("Expected head.next to be 5, got %d", list.head.next.value)
	}

	if list.tail.value != 1 {
		t.Errorf("Expected tail to be 1, got %d", list.tail.value)
	}

	middle := list.head.next
	if middle.prev.value != 10 {
		t.Errorf("Expected middle.prev to be 10, got %d", middle.prev.value)
	}

	if middle.next.value != 1 {
		t.Errorf("Expected middle.next to be 1, got %d", middle.next.value)
	}
}

func TestAddOrderAndCountDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(3)
	list.Add(6)
	list.Add(10)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items in list, got %d", list.Count())
	}

	if list.head.value != 10 {
		t.Errorf("Expected head to be 10, got %d", list.head.value)
	}

	if list.head.next.value != 6 {
		t.Errorf("Expected head.next to be 6, got %d", list.head.next.value)
	}

	if list.tail.value != 1 {
		t.Errorf("Expected tail to be 1, got %d", list.tail.value)
	}

	if list.tail.prev.value != 3 {
		t.Errorf("Expected tail.prev to be 3, got %d", list.tail.prev.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestFindExistingAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	node, err := list.Find(3)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.value != 3 {
		t.Errorf("Expected value to be 3, got %d", node.value)
	}
}

func TestFindExistingDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	node, err := list.Find(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.value != 5 {
		t.Errorf("Expected value to be 5, got %d", node.value)
	}
}

func TestDedupNoDuplicates(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestDedupAllSame(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	list.Dedup()

	if list.Count() != 1 {
		t.Errorf("Expected 1 item in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}

	if list.head != list.tail {
		t.Errorf("Expected head and tail to be the same node")
	}
}

func TestDedupMiddleDuplicates(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 3 {
		t.Errorf("Expected head.next to be 3, got %d", list.head.next.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestDedupHeadDuplicates(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 3 {
		t.Errorf("Expected head.next to be 3, got %d", list.head.next.value)
	}

	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
}

func TestDedupTailDuplicates(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}

	if list.tail.prev.value != 3 {
		t.Errorf("Expected tail.prev to be 3, got %d", list.tail.prev.value)
	}

	if list.tail.next != nil {
		t.Errorf("Expected tail.next to be nil, got %v", list.tail.next)
	}
}

func TestDedupEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	list.Dedup()

	if list.Count() != 0 {
		t.Errorf("Expected 0 items in list, got %d", list.Count())
	}
}

func TestDedupSingleElement(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)

	list.Dedup()

	if list.Count() != 1 {
		t.Errorf("Expected 1 item in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}
}

func TestDedupMultipleGroupsAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.value)
	}

	if list.head.next.value != 3 {
		t.Errorf("Expected head.next to be 3, got %d", list.head.next.value)
	}

	if list.tail.value != 5 {
		t.Errorf("Expected tail to be 5, got %d", list.tail.value)
	}
}

func TestDedupMultipleGroupsDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items in list, got %d", list.Count())
	}

	if list.head.value != 5 {
		t.Errorf("Expected head to be 5, got %d", list.head.value)
	}

	if list.head.next.value != 3 {
		t.Errorf("Expected head.next to be 3, got %d", list.head.next.value)
	}

	if list.tail.value != 1 {
		t.Errorf("Expected tail to be 1, got %d", list.tail.value)
	}
}

func TestDedupPreservesLinks(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.head.next.prev != list.head {
		t.Errorf("Expected head.next.prev to be head")
	}

	if list.tail.prev.next != list.tail {
		t.Errorf("Expected tail.prev.next to be tail")
	}
}

func TestIsSublistFullMatch(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(1)
	sublist.Add(3)
	sublist.Add(5)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist")
	}
}

func TestIsSublistAtStart(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(1)
	sublist.Add(3)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist")
	}
}

func TestIsSublistAtEnd(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(5)
	sublist.Add(7)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist")
	}
}

func TestIsSublistInMiddle(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(5)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist")
	}
}

func TestIsSublistSingleElement(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(3)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist")
	}
}

func TestIsSublistEmptySublist(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = OrderedList[int]{_ascending: true}

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected empty sublist to be a sublist")
	}
}

func TestIsSublistEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(1)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be a sublist of empty list")
	}
}

func TestIsSublistBothEmpty(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	var sublist = OrderedList[int]{_ascending: true}

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected empty sublist to be a sublist of empty list")
	}
}

func TestIsSublistNotFound(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(2)
	sublist.Add(4)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be a sublist")
	}
}

func TestIsSublistPartialMatch(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(7)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be a sublist (partial match)")
	}
}

func TestIsSublistLongerThanList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(1)
	sublist.Add(3)
	sublist.Add(5)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be a sublist (longer)")
	}
}

func TestIsSublistDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = OrderedList[int]{_ascending: false}
	sublist.Add(3)
	sublist.Add(5)

	// list: 7 -> 5 -> 3 -> 1
	// sublist: 5 -> 3
	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be a sublist (descending)")
	}
}

func TestIsSublistNotContiguous(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = OrderedList[int]{_ascending: true}
	sublist.Add(1)
	sublist.Add(5)

	// sublist elements exist but not contiguous
	if list.IsSublist(&sublist) {
		t.Errorf("Expected non-contiguous elements to not be a sublist")
	}
}

func TestTopFrequentSingleMost(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestTopFrequentAllSame(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	result := list.TopFrequent()

	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestTopFrequentAllUnique(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	// Все уникальные — любой валиден, просто проверяем что один из них
	if result != 1 && result != 3 && result != 5 {
		t.Errorf("Expected one of 1, 3, 5, got %d", result)
	}
}

func TestTopFrequentSingleElement(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(7)

	result := list.TopFrequent()

	if result != 7 {
		t.Errorf("Expected 7, got %d", result)
	}
}

func TestTopFrequentEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	result := list.TopFrequent()

	// Пустой список — вернёт zero value
	if result != 0 {
		t.Errorf("Expected 0 (zero value), got %d", result)
	}
}

func TestTopFrequentTwoMostFrequent(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	// 1 и 3 оба по 2 раза — любой валиден
	if result != 1 && result != 3 {
		t.Errorf("Expected 1 or 3, got %d", result)
	}
}

func TestTopFrequentDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestTopFrequentAtHead(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestTopFrequentAtTail(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(5)
	list.Add(5)

	result := list.TopFrequent()

	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestMergeBasicAsc(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)
	l1.Add(3)
	l1.Add(5)

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(2)
	l2.Add(4)
	l2.Add(6)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 6 {
		t.Errorf("Expected 6 items, got %d", result.Count())
	}

	if result.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", result.head.value)
	}

	if result.tail.value != 6 {
		t.Errorf("Expected tail to be 6, got %d", result.tail.value)
	}

	// Проверяем порядок: 1 -> 2 -> 3 -> 4 -> 5 -> 6
	expected := []int{1, 2, 3, 4, 5, 6}
	current := result.head
	for i, exp := range expected {
		if current.value != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, current.value)
		}
		current = current.next
	}
}

func TestMergeBasicDesc(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: false}
	l1.Add(1)
	l1.Add(3)
	l1.Add(5)

	var l2 = OrderedList[int]{_ascending: false}
	l2.Add(2)
	l2.Add(4)
	l2.Add(6)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 6 {
		t.Errorf("Expected 6 items, got %d", result.Count())
	}

	if result.head.value != 6 {
		t.Errorf("Expected head to be 6, got %d", result.head.value)
	}

	if result.tail.value != 1 {
		t.Errorf("Expected tail to be 1, got %d", result.tail.value)
	}

	// Проверяем порядок: 6 -> 5 -> 4 -> 3 -> 2 -> 1
	expected := []int{6, 5, 4, 3, 2, 1}
	current := result.head
	for i, exp := range expected {
		if current.value != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, current.value)
		}
		current = current.next
	}
}

func TestMergeFirstEmpty(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(1)
	l2.Add(2)
	l2.Add(3)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", result.Count())
	}

	if result.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", result.head.value)
	}

	if result.tail.value != 3 {
		t.Errorf("Expected tail to be 3, got %d", result.tail.value)
	}
}

func TestMergeSecondEmpty(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)

	var l2 = OrderedList[int]{_ascending: true}

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", result.Count())
	}

	if result.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", result.head.value)
	}

	if result.tail.value != 3 {
		t.Errorf("Expected tail to be 3, got %d", result.tail.value)
	}
}

func TestMergeBothEmpty(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	var l2 = OrderedList[int]{_ascending: true}

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", result.Count())
	}

	if result.head != nil {
		t.Errorf("Expected head to be nil")
	}

	if result.tail != nil {
		t.Errorf("Expected tail to be nil")
	}
}

func TestMergeWithDuplicates(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)
	l1.Add(3)
	l1.Add(5)

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(1)
	l2.Add(3)
	l2.Add(5)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 6 {
		t.Errorf("Expected 6 items, got %d", result.Count())
	}

	// Проверяем порядок: 1 -> 1 -> 3 -> 3 -> 5 -> 5
	expected := []int{1, 1, 3, 3, 5, 5}
	current := result.head
	for i, exp := range expected {
		if current.value != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, current.value)
		}
		current = current.next
	}
}

func TestMergeDifferentLengths(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)
	l1.Add(5)

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(2)
	l2.Add(3)
	l2.Add(4)
	l2.Add(6)
	l2.Add(7)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 7 {
		t.Errorf("Expected 7 items, got %d", result.Count())
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	current := result.head
	for i, exp := range expected {
		if current.value != exp {
			t.Errorf("Expected %d at position %d, got %d", exp, i, current.value)
		}
		current = current.next
	}
}

func TestMergePreservesLinks(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)
	l1.Add(3)

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(2)
	l2.Add(4)

	result := MergeOrdLists(&l1, &l2)

	// Проверяем prev/next связи
	if result.head.prev != nil {
		t.Errorf("Expected head.prev to be nil")
	}

	if result.tail.next != nil {
		t.Errorf("Expected tail.next to be nil")
	}

	current := result.head
	for current.next != nil {
		if current.next.prev != current {
			t.Errorf("Broken prev link")
		}
		current = current.next
	}
}

func TestMergeSingleElements(t *testing.T) {
	var l1 = OrderedList[int]{_ascending: true}
	l1.Add(1)

	var l2 = OrderedList[int]{_ascending: true}
	l2.Add(2)

	result := MergeOrdLists(&l1, &l2)

	if result.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", result.Count())
	}

	if result.head.value != 1 {
		t.Errorf("Expected head to be 1, got %d", result.head.value)
	}

	if result.tail.value != 2 {
		t.Errorf("Expected tail to be 2, got %d", result.tail.value)
	}
}

func TestFindHead(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	node, err := list.Find(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.value != 1 {
		t.Errorf("Expected value to be 1, got %d", node.value)
	}
}

func TestFindTail(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	node, err := list.Find(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.value != 5 {
		t.Errorf("Expected value to be 5, got %d", node.value)
	}
}

func TestFindNonExistentAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	_, err := list.Find(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindNonExistentDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	_, err := list.Find(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindEmptyList(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}

	_, err := list.Find(5)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindEarlyTerminationAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(5)
	list.Add(10)
	list.Add(20)

	_, err := list.Find(3)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindEarlyTerminationDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(1)
	list.Add(5)
	list.Add(10)
	list.Add(20)

	_, err := list.Find(7)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindBeyondTailAsc(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	_, err := list.Find(10)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindBeyondTailDesc(t *testing.T) {
	var list = OrderedList[int]{_ascending: false}
	list.Add(10)
	list.Add(5)
	list.Add(3)

	_, err := list.Find(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFindDuplicate(t *testing.T) {
	var list = OrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	node, err := list.Find(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if node.value != 5 {
		t.Errorf("Expected value to be 5, got %d", node.value)
	}
}

// Indexed ordered list tests

package ordlist

import "testing"

func TestIndexedAddAndCountAsc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(1)
	list.Add(3)
	list.Add(7)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items, got %d", list.Count())
	}

	// Проверяем порядок: 1, 3, 5, 7
	expected := []int{1, 3, 5, 7}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedAddAndCountDesc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: false}
	list.Add(5)
	list.Add(1)
	list.Add(3)
	list.Add(7)

	if list.Count() != 4 {
		t.Errorf("Expected 4 items, got %d", list.Count())
	}

	// Проверяем порядок: 7, 5, 3, 1
	expected := []int{7, 5, 3, 1}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedAddDuplicates(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)
	list.Add(5)
	list.Add(5)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}
}

func TestIndexedFindExisting(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	val, err := list.Find(3)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestIndexedFindNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	_, err := list.Find(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedFindEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	_, err := list.Find(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedFindIndexAsc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)

	idx, err := list.FindIndex(30)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexDesc(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: false}
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)

	// Порядок: 50, 40, 30, 20, 10
	idx, err := list.FindIndex(30)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexHead(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 0 {
		t.Errorf("Expected index 0, got %d", idx)
	}
}

func TestIndexedFindIndexTail(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(5)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
}

func TestIndexedFindIndexNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	idx, err := list.FindIndex(4)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if idx != -1 {
		t.Errorf("Expected index -1, got %d", idx)
	}
}

func TestIndexedFindIndexEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	_, err := list.FindIndex(1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedDeleteExisting(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(3)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	_, err := list.Find(3)
	if err == nil {
		t.Errorf("Expected element to be deleted")
	}
}

func TestIndexedDeleteHead(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(1)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	val, _ := list.Get(0)
	if val != 3 {
		t.Errorf("Expected first element to be 3, got %d", val)
	}
}

func TestIndexedDeleteTail(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(5)

	if list.Count() != 2 {
		t.Errorf("Expected 2 items, got %d", list.Count())
	}

	val, _ := list.Get(1)
	if val != 3 {
		t.Errorf("Expected last element to be 3, got %d", val)
	}
}

func TestIndexedDeleteNonExistent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Delete(999)

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}
}

func TestIndexedDeleteEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	list.Delete(1)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}
}

func TestIndexedClear(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	list.Clear(false)

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}

	if list._ascending != false {
		t.Errorf("Expected ascending to be false")
	}
}

func TestIndexedDedup(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	list.Dedup()

	if list.Count() != 3 {
		t.Errorf("Expected 3 items, got %d", list.Count())
	}

	expected := []int{1, 3, 5}
	for i, exp := range expected {
		val, _ := list.Get(i)
		if val != exp {
			t.Errorf("Expected %d at index %d, got %d", exp, i, val)
		}
	}
}

func TestIndexedDedupEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	list.Dedup()

	if list.Count() != 0 {
		t.Errorf("Expected 0 items, got %d", list.Count())
	}
}

func TestIndexedDedupSingle(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(5)

	list.Dedup()

	if list.Count() != 1 {
		t.Errorf("Expected 1 item, got %d", list.Count())
	}
}

func TestIndexedIsSublistTrue(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)

	var sublist = IndexedOrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(5)

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to be found")
	}
}

func TestIndexedIsSublistFalse(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	var sublist = IndexedOrderedList[int]{_ascending: true}
	sublist.Add(3)
	sublist.Add(7)

	if list.IsSublist(&sublist) {
		t.Errorf("Expected sublist to not be found")
	}
}

func TestIndexedIsSublistEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	var sublist = IndexedOrderedList[int]{_ascending: true}

	if !list.IsSublist(&sublist) {
		t.Errorf("Expected empty sublist to be found")
	}
}

func TestIndexedTopFrequent(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(3)
	list.Add(3)
	list.Add(5)

	result := list.TopFrequent()

	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestIndexedTopFrequentEmpty(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}

	result := list.TopFrequent()

	if result != 0 {
		t.Errorf("Expected 0 (zero value), got %d", result)
	}
}

func TestIndexedGet(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)
	list.Add(5)

	val, err := list.Get(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestIndexedGetOutOfBounds(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)
	list.Add(3)

	_, err := list.Get(10)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestIndexedGetNegativeIndex(t *testing.T) {
	var list = IndexedOrderedList[int]{_ascending: true}
	list.Add(1)

	_, err := list.Get(-1)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
