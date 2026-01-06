package powersets

import (
	"testing"
)

// === Init tests ===

func TestInit(t *testing.T) {
	ps := Init[string]()
	if ps.storage == nil {
		t.Error("storage should be initialized")
	}
	if ps.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", ps.Size())
	}
}

// === Size tests ===

func TestSize_Empty(t *testing.T) {
	ps := Init[string]()
	if ps.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", ps.Size())
	}
}

func TestSize_AfterPut(t *testing.T) {
	ps := Init[string]()
	ps.Put("a")
	ps.Put("b")
	ps.Put("c")
	if ps.Size() != 3 {
		t.Errorf("Size() is %d, expected 3", ps.Size())
	}
}

func TestSize_AfterDuplicatePut(t *testing.T) {
	ps := Init[string]()
	ps.Put("a")
	ps.Put("a")
	ps.Put("a")
	if ps.Size() != 1 {
		t.Errorf("Size() is %d, expected 1 (duplicates should not increase size)", ps.Size())
	}
}

// === Put tests ===

func TestPut_NewElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("hello")

	if !ps.Get("hello") {
		t.Error("element should exist after Put")
	}
}

func TestPut_DuplicateElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("hello")
	ps.Put("hello")

	if ps.Size() != 1 {
		t.Errorf("Size() is %d, expected 1 (duplicate Put should not add element)", ps.Size())
	}
}

func TestPut_MultipleElements(t *testing.T) {
	ps := Init[int]()
	ps.Put(1)
	ps.Put(2)
	ps.Put(3)

	for _, v := range []int{1, 2, 3} {
		if !ps.Get(v) {
			t.Errorf("element %d should exist", v)
		}
	}
}

// === Get tests ===

func TestGet_ExistingElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("exists")

	if !ps.Get("exists") {
		t.Error("Get should return true for existing element")
	}
}

func TestGet_NonExistingElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("exists")

	if ps.Get("not_exists") {
		t.Error("Get should return false for non-existing element")
	}
}

func TestGet_EmptySet(t *testing.T) {
	ps := Init[string]()

	if ps.Get("any") {
		t.Error("Get should return false on empty set")
	}
}

// === Remove tests ===

func TestRemove_ExistingElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("element")

	removed := ps.Remove("element")
	if !removed {
		t.Error("Remove should return true for existing element")
	}
	if ps.Get("element") {
		t.Error("element should not exist after Remove")
	}
	if ps.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 after Remove", ps.Size())
	}
}

func TestRemove_NonExistingElement(t *testing.T) {
	ps := Init[string]()
	ps.Put("element")

	removed := ps.Remove("missing")
	if removed {
		t.Error("Remove should return false for non-existing element")
	}
	if ps.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", ps.Size())
	}
}

func TestRemove_EmptySet(t *testing.T) {
	ps := Init[string]()

	removed := ps.Remove("any")
	if removed {
		t.Error("Remove should return false on empty set")
	}
}

// === Union tests ===

func TestUnion_BothNonEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("b")
	ps2.Put("c")

	result := ps1.Union(ps2)

	if result.Size() != 3 {
		t.Errorf("Size() is %d, expected 3", result.Size())
	}
	for _, v := range []string{"a", "b", "c"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in union", v)
		}
	}
}

func TestUnion_FirstEmpty(t *testing.T) {
	ps1 := Init[string]()

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Union(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in union", v)
		}
	}
}

func TestUnion_SecondEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()

	result := ps1.Union(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in union", v)
		}
	}
}

func TestUnion_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	result := ps1.Union(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestUnion_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Union(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
}

func TestUnion_DisjointSets(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)
	ps1.Put(2)

	ps2 := Init[int]()
	ps2.Put(3)
	ps2.Put(4)

	result := ps1.Union(ps2)

	if result.Size() != 4 {
		t.Errorf("Size() is %d, expected 4", result.Size())
	}
	for _, v := range []int{1, 2, 3, 4} {
		if !result.Get(v) {
			t.Errorf("element %d should exist in union", v)
		}
	}
}

// === Intersection tests ===

func TestIntersection_WithOverlap(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("b")
	ps2.Put("c")
	ps2.Put("d")

	result := ps1.Intersection(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"b", "c"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in intersection", v)
		}
	}
	if result.Get("a") || result.Get("d") {
		t.Error("elements 'a' and 'd' should not exist in intersection")
	}
}

func TestIntersection_DisjointSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("c")
	ps2.Put("d")

	result := ps1.Intersection(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 (disjoint sets)", result.Size())
	}
}

func TestIntersection_FirstEmpty(t *testing.T) {
	ps1 := Init[string]()

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Intersection(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestIntersection_SecondEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()

	result := ps1.Intersection(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestIntersection_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	result := ps1.Intersection(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestIntersection_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Intersection(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in intersection", v)
		}
	}
}

func TestIntersection_SubsetRelation(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)
	ps1.Put(2)
	ps1.Put(3)

	ps2 := Init[int]()
	ps2.Put(2)

	result := ps1.Intersection(ps2)

	if result.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", result.Size())
	}
	if !result.Get(2) {
		t.Error("element 2 should exist in intersection")
	}
}

// === Equals tests ===

func TestEquals_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")
	ps2.Put("c")

	if !ps1.Equals(ps2) {
		t.Error("identical sets should be equal")
	}
}

func TestEquals_DifferentElements(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("c")

	if ps1.Equals(ps2) {
		t.Error("sets with different elements should not be equal")
	}
}

func TestEquals_DifferentSizes(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	if ps1.Equals(ps2) {
		t.Error("sets with different sizes should not be equal")
	}
}

func TestEquals_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	if !ps1.Equals(ps2) {
		t.Error("empty sets should be equal")
	}
}

func TestEquals_OneEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")

	ps2 := Init[string]()

	if ps1.Equals(ps2) {
		t.Error("empty and non-empty sets should not be equal")
	}
}

func TestEquals_Symmetric(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)
	ps1.Put(2)

	ps2 := Init[int]()
	ps2.Put(1)
	ps2.Put(2)

	if ps1.Equals(ps2) != ps2.Equals(ps1) {
		t.Error("equality should be symmetric")
	}
}

// === IsSubset tests ===

func TestIsSubset_AllElementsInCurrent(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	if !ps1.IsSubset(ps2) {
		t.Error("set2 should be a subset of ps1 (all elements of set2 are in ps1)")
	}
}

func TestIsSubset_CurrentIsSubsetOfParam(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")
	ps2.Put("c")

	if ps1.IsSubset(ps2) {
		t.Error("set2 is not a subset of ps1 (set2 is larger)")
	}
}

func TestIsSubset_NotAllElementsInCurrent(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("d")

	if ps1.IsSubset(ps2) {
		t.Error("set2 should not be a subset (element 'd' is not in ps1)")
	}
}

func TestIsSubset_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	if !ps1.IsSubset(ps2) {
		t.Error("identical sets: set2 should be a subset of ps1")
	}
}

func TestIsSubset_EmptyParam(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()

	if !ps1.IsSubset(ps2) {
		t.Error("empty set should be a subset of any set")
	}
}

func TestIsSubset_EmptyCurrent(t *testing.T) {
	ps1 := Init[string]()

	ps2 := Init[string]()
	ps2.Put("a")

	if ps1.IsSubset(ps2) {
		t.Error("non-empty set cannot be a subset of empty set")
	}
}

func TestIsSubset_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	if !ps1.IsSubset(ps2) {
		t.Error("empty set should be a subset of empty set")
	}
}

// === Difference tests ===

func TestDifference_WithOverlap(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("b")
	ps2.Put("c")
	ps2.Put("d")

	result := ps1.Difference(ps2)

	if result.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", result.Size())
	}
	if !result.Get("a") {
		t.Error("element 'a' should exist in difference")
	}
	if result.Get("b") || result.Get("c") || result.Get("d") {
		t.Error("elements 'b', 'c', 'd' should not exist in difference")
	}
}

func TestDifference_DisjointSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("c")
	ps2.Put("d")

	result := ps1.Difference(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in difference", v)
		}
	}
}

func TestDifference_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Difference(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 (identical sets)", result.Size())
	}
}

func TestDifference_FirstEmpty(t *testing.T) {
	ps1 := Init[string]()

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := ps1.Difference(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestDifference_SecondEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()

	result := ps1.Difference(ps2)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in difference", v)
		}
	}
}

func TestDifference_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	result := ps1.Difference(ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

func TestDifference_SubsetRelation(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)
	ps1.Put(2)
	ps1.Put(3)

	ps2 := Init[int]()
	ps2.Put(1)
	ps2.Put(2)

	result := ps1.Difference(ps2)

	if result.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", result.Size())
	}
	if !result.Get(3) {
		t.Error("element 3 should exist in difference")
	}
}

// === CartesianProduct tests ===

func TestCartesianProduct_BothNonEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("1")
	ps2.Put("2")

	result := CartesianProduct(ps1, ps2)

	if len(result) != 4 {
		t.Errorf("len is %d, expected 4", len(result))
	}

	expected := map[[2]string]bool{
		{"a", "1"}: true,
		{"a", "2"}: true,
		{"b", "1"}: true,
		{"b", "2"}: true,
	}

	for _, pair := range result {
		if !expected[pair] {
			t.Errorf("unexpected pair %v", pair)
		}
		delete(expected, pair)
	}

	if len(expected) != 0 {
		t.Errorf("missing pairs: %v", expected)
	}
}

func TestCartesianProduct_FirstEmpty(t *testing.T) {
	ps1 := Init[string]()

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	result := CartesianProduct(ps1, ps2)

	if len(result) != 0 {
		t.Errorf("len is %d, expected 0", len(result))
	}
}

func TestCartesianProduct_SecondEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()

	result := CartesianProduct(ps1, ps2)

	if len(result) != 0 {
		t.Errorf("len is %d, expected 0", len(result))
	}
}

func TestCartesianProduct_BothEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()

	result := CartesianProduct(ps1, ps2)

	if len(result) != 0 {
		t.Errorf("len is %d, expected 0", len(result))
	}
}

func TestCartesianProduct_SingleElements(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)

	ps2 := Init[int]()
	ps2.Put(2)

	result := CartesianProduct(ps1, ps2)

	if len(result) != 1 {
		t.Errorf("len is %d, expected 1", len(result))
	}
	if result[0] != [2]int{1, 2} {
		t.Errorf("got %v, expected [1, 2]", result[0])
	}
}

func TestCartesianProduct_SameSet(t *testing.T) {
	ps := Init[string]()
	ps.Put("a")
	ps.Put("b")

	result := CartesianProduct(ps, ps)

	if len(result) != 4 {
		t.Errorf("len is %d, expected 4", len(result))
	}

	expected := map[[2]string]bool{
		{"a", "a"}: true,
		{"a", "b"}: true,
		{"b", "a"}: true,
		{"b", "b"}: true,
	}

	for _, pair := range result {
		if !expected[pair] {
			t.Errorf("unexpected pair %v", pair)
		}
		delete(expected, pair)
	}
}

// === IntersectMany tests ===

func TestIntersectMany_ThreeSetsWithOverlap(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")
	ps1.Put("c")

	ps2 := Init[string]()
	ps2.Put("b")
	ps2.Put("c")
	ps2.Put("d")

	ps3 := Init[string]()
	ps3.Put("c")
	ps3.Put("d")
	ps3.Put("e")

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", result.Size())
	}
	if !result.Get("c") {
		t.Error("element 'c' should exist in intersection")
	}
}

func TestIntersectMany_FourSets(t *testing.T) {
	ps1 := Init[int]()
	ps1.Put(1)
	ps1.Put(2)
	ps1.Put(3)

	ps2 := Init[int]()
	ps2.Put(2)
	ps2.Put(3)
	ps2.Put(4)

	ps3 := Init[int]()
	ps3.Put(2)
	ps3.Put(3)
	ps3.Put(5)

	ps4 := Init[int]()
	ps4.Put(2)
	ps4.Put(3)
	ps4.Put(6)

	result := IntersectMany(ps1, ps2, ps3, ps4)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []int{2, 3} {
		if !result.Get(v) {
			t.Errorf("element %d should exist in intersection", v)
		}
	}
}

func TestIntersectMany_DisjointSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")

	ps2 := Init[string]()
	ps2.Put("b")

	ps3 := Init[string]()
	ps3.Put("c")

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 (disjoint sets)", result.Size())
	}
}

func TestIntersectMany_OneEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	ps3 := Init[string]()

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 (one empty set)", result.Size())
	}
}

func TestIntersectMany_AllEmpty(t *testing.T) {
	ps1 := Init[string]()
	ps2 := Init[string]()
	ps3 := Init[string]()

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}

// === Bag tests ===

func TestBag_Init(t *testing.T) {
	b := InitBag[string]()
	if b.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", b.Size())
	}
	if b.UniqueSize() != 0 {
		t.Errorf("UniqueSize() is %d, expected 0", b.UniqueSize())
	}
}

func TestBag_Put(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")
	b.Put("a")
	b.Put("a")
	b.Put("b")

	if b.Size() != 4 {
		t.Errorf("Size() is %d, expected 4", b.Size())
	}
	if b.UniqueSize() != 2 {
		t.Errorf("UniqueSize() is %d, expected 2", b.UniqueSize())
	}
}

func TestBag_Get(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")

	if !b.Get("a") {
		t.Error("Get should return true for existing element")
	}
	if b.Get("b") {
		t.Error("Get should return false for non-existing element")
	}
}

func TestBag_Count(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")
	b.Put("a")
	b.Put("a")
	b.Put("b")

	if b.Count("a") != 3 {
		t.Errorf("Count('a') is %d, expected 3", b.Count("a"))
	}
	if b.Count("b") != 1 {
		t.Errorf("Count('b') is %d, expected 1", b.Count("b"))
	}
	if b.Count("c") != 0 {
		t.Errorf("Count('c') is %d, expected 0", b.Count("c"))
	}
}

func TestBag_Remove(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")
	b.Put("a")
	b.Put("a")

	removed := b.Remove("a")
	if !removed {
		t.Error("Remove should return true")
	}
	if b.Count("a") != 2 {
		t.Errorf("Count('a') is %d, expected 2", b.Count("a"))
	}
	if b.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", b.Size())
	}
}

func TestBag_Remove_LastInstance(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")

	b.Remove("a")

	if b.Get("a") {
		t.Error("element should not exist after removing last instance")
	}
	if b.UniqueSize() != 0 {
		t.Errorf("UniqueSize() is %d, expected 0", b.UniqueSize())
	}
}

func TestBag_Remove_NonExisting(t *testing.T) {
	b := InitBag[string]()

	removed := b.Remove("a")
	if removed {
		t.Error("Remove should return false for non-existing element")
	}
}

func TestBag_RemoveAll(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")
	b.Put("a")
	b.Put("a")
	b.Put("b")

	count := b.RemoveAll("a")

	if count != 3 {
		t.Errorf("RemoveAll returned %d, expected 3", count)
	}
	if b.Get("a") {
		t.Error("element 'a' should not exist after RemoveAll")
	}
	if b.Size() != 1 {
		t.Errorf("Size() is %d, expected 1", b.Size())
	}
	if !b.Get("b") {
		t.Error("element 'b' should still exist")
	}
}

func TestBag_RemoveAll_NonExisting(t *testing.T) {
	b := InitBag[string]()

	count := b.RemoveAll("a")

	if count != 0 {
		t.Errorf("RemoveAll returned %d, expected 0", count)
	}
}

func TestBag_Frequencies(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")
	b.Put("a")
	b.Put("b")
	b.Put("c")
	b.Put("c")
	b.Put("c")

	freq := b.Frequencies()

	expected := map[string]int{
		"a": 2,
		"b": 1,
		"c": 3,
	}

	if len(freq) != len(expected) {
		t.Errorf("len(freq) is %d, expected %d", len(freq), len(expected))
	}

	for k, v := range expected {
		if freq[k] != v {
			t.Errorf("freq[%q] is %d, expected %d", k, freq[k], v)
		}
	}
}

func TestBag_Frequencies_IsCopy(t *testing.T) {
	b := InitBag[string]()
	b.Put("a")

	freq := b.Frequencies()
	freq["a"] = 999

	if b.Count("a") != 1 {
		t.Error("modifying Frequencies result should not affect bag")
	}
}

func TestIntersectMany_IdenticalSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")
	ps1.Put("b")

	ps2 := Init[string]()
	ps2.Put("a")
	ps2.Put("b")

	ps3 := Init[string]()
	ps3.Put("a")
	ps3.Put("b")

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 2 {
		t.Errorf("Size() is %d, expected 2", result.Size())
	}
	for _, v := range []string{"a", "b"} {
		if !result.Get(v) {
			t.Errorf("element %q should exist in intersection", v)
		}
	}
}

func TestIntersectMany_LessThanThreeSets(t *testing.T) {
	ps1 := Init[string]()
	ps1.Put("a")

	ps2 := Init[string]()
	ps2.Put("a")

	result := IntersectMany(ps1, ps2)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0 (less than 3 sets)", result.Size())
	}
}

func TestIntersectMany_SmallestSetOptimization(t *testing.T) {
	// smallest set has element not in others
	ps1 := Init[int]()
	ps1.Put(999)

	ps2 := Init[int]()
	for i := range 100 {
		ps2.Put(i)
	}

	ps3 := Init[int]()
	for i := range 100 {
		ps3.Put(i)
	}

	result := IntersectMany(ps1, ps2, ps3)

	if result.Size() != 0 {
		t.Errorf("Size() is %d, expected 0", result.Size())
	}
}
