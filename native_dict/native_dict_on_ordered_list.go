package native_dict

// OrderedDictionary based on ordered list with binary search.
//
// Time complexity:
// - Search (Get, IsKey): O(log n) - binary search
// - Insert (Put): O(n) - binary search + shift elements
// - Delete: O(n) - binary search + shift elements
//
// Trade-off vs hash table:
// + Guaranteed O(log n) search (no worst case O(n) on collisions)
// + Keys stored in sorted order (range queries possible)
// + No hash function selection or collision problems
// - Insert/delete O(n) instead of O(1) amortized

type OrderedDictionary[T any] struct {
	count  int
	keys   []string
	values []T
}

// Time: O(n) where n = sz
// Space: O(n)
func InitOrdered[T any](sz int) OrderedDictionary[T] {
	return OrderedDictionary[T]{
		count:  0,
		keys:   make([]string, 0, sz),
		values: make([]T, 0, sz),
	}
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) IsKey(key string) bool {
	_, found := od.binarySearch(key)
	return found
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) Get(key string) (T, error) {
	var result T
	index, found := od.binarySearch(key)
	if !found {
		return result, ErrKeyNotFound
	}
	return od.values[index], nil
}

// Time: O(n) - binary search O(log n) + shift O(n)
// Space: O(1) amortized
func (od *OrderedDictionary[T]) Put(key string, value T) {
	index, found := od.binarySearch(key)
	if found {
		od.values[index] = value
		return
	}

	od.keys = append(od.keys, "")
	od.values = append(od.values, value)

	copy(od.keys[index+1:], od.keys[index:])
	copy(od.values[index+1:], od.values[index:])

	od.keys[index] = key
	od.values[index] = value
	od.count++
}

// Time: O(n) - binary search O(log n) + shift O(n)
// Space: O(1)
func (od *OrderedDictionary[T]) Delete(key string) bool {
	index, found := od.binarySearch(key)
	if !found {
		return false
	}

	od.keys = append(od.keys[:index], od.keys[index+1:]...)
	od.values = append(od.values[:index], od.values[index+1:]...)
	od.count--
	return true
}

// Time: O(1)
// Space: O(1)
func (od *OrderedDictionary[T]) Count() int {
	return od.count
}

// Time: O(log n)
// Space: O(1)
func (od *OrderedDictionary[T]) binarySearch(key string) (index int, found bool) {
	left, right := 0, len(od.keys)-1
	for left <= right {
		mid := left + (right-left)/2
		if od.keys[mid] == key {
			return mid, true
		}
		if od.keys[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left, false
}
