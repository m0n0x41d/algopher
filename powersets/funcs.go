package powersets

import "golang.org/x/exp/constraints"

// space and time is O(N * M) where N and M are sizes of sets accordingly
func CartesianProduct[T constraints.Ordered](ps1, ps2 PowerSet[T]) [][2]T {
	// preallocate result slice to avoide reallocations on appends
	result := make([][2]T, 0, ps1.Size()*ps2.Size())

	for x := range ps1.storage {
		for y := range ps2.storage {
			result = append(result, [2]T{x, y})
		}

	}
	return result
}

// By Time: O(S * K) where S is size of smallest set, K is number of sets.
// Space: O(R) where R is size of resulting intersection.
// we are itergint over smallest set, trying to be more eficcient, but it will be still slow.
func IntersectMany[T constraints.Ordered](sets ...PowerSet[T]) PowerSet[T] {
	result := Init[T]()

	// The task stated clearly: three or more.
	if len(sets) < 3 {
		return result
	}

	smallestSetIndex := 0
	for i := 1; i < len(sets); i++ {
		if sets[i].count < sets[smallestSetIndex].count {
			smallestSetIndex = i
		}
	}

	for smallestSetElement := range sets[smallestSetIndex].storage {
		foundInAllSets := true
		for i, currentSet := range sets {
			// Skip smallest.
			if i == smallestSetIndex {
				continue
			}

			// can't be in intersect if it's not in all sets.
			if !currentSet.Get(smallestSetElement) {
				foundInAllSets = false
				break
			}
		}
		if foundInAllSets {
			result.Put(smallestSetElement)
		}
	}
	return result
}
