package cache

import "container/list"

// Why I wrote an LFU Cache here with Doubly Linked Lists instead of using hash table with linear probing (as Labaratory task were suggesting)?
//
// the naive implementation is in same module, see native_cache.go
//
// Well, yes - original task suggested implementing NativeCache based on a hash table with
// open addressing (quadratic probing) and an eviction scheme that finds
// the element with minimum hits by scanning the entire hits[] array.
//
// However, I assumed several problems with this approach:
//
// 1. Eviction is O(n) - algorithm must scan entire hits[] array to find minimum...
// 2. And after eviction, the new key might hash to a completely different slot,
//    requiring potentially O(n) probing to reach the freed slot
// 3. Deleting from such open-addressed hash tables is also problematic:
//    Consider: Put("A") -> slot[2], Put("B") -> slot[3] (collision), Put("C") -> slot[4] (collision)
//    all three keys hash to 2, but B and C were placed via probing!
//    Now if delete "B": slots become [_, _, A, nil, C, _, _]
//    Find("C") starts at hash("C")=2, steps to slot[3], sees nil and STOPS.
//    We lost "C" even though it exists! The nil breaks the probe chain.
//    Solutions I found are not pleasant:
//    - Special "Tombstones" valued. We can mark deleted slots as DELETED (or something like that, some magic constant) instead of nil,
//    so probing continues.But tombstones can accumulate, which might lead to degrading performance (need to rehash map periodically)
//    - Rehash on delete: recalc all positions of all elements in collision chain. This is too expensive I guess.
//    - Just don't support deletion. Not an option for a cache (no, really, it is a bad interface)
// 4. on top of that, such implementation gives worst-case O(n) for Put when cache is full
//
// This implementation below utilizes the standard LFU Cache design:
// - HashMap for O(1) key lookup (I might reuse my own implementation, which may be a bit updated, but I decided to use Go's map and linked list)
// - Frequency buckets (map[freq] -> doubly linked list) for O(1) eviction (how cool is that?!)
// - Each bucket maintains LRU order thanks to linked list, so comparison ties are broken by recency
//
// So... if we are comparing complexity this LFUCache gives us O(1) for all operations!
// But for sure there is also a trade-off, and this approach uses much more memory (linked list nodes and pointers, two maps).
// And even so, for most of mediocre business applications such tradeoff is justified by great performance for user experience.

type cacheEntry struct {
	key   string
	value any
	freq  int
}

type LFUCache struct {
	capacity int
	minFreq  int
	cache    map[string]*list.Element
	freqMap  map[int]*list.List
}

// Time: O(1)
// Space: O(1) initialization, O(capacity) when full
func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		minFreq:  0,
		cache:    make(map[string]*list.Element),
		freqMap:  make(map[int]*list.List),
	}
}

// Time: O(1) - map lookup + incrementFreq (both O(1))
// Space: O(1)
func (c *LFUCache) Get(key string) (any, bool) {
	item, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	c.incrementFreq(item)
	return item.Value.(*cacheEntry).value, true
}

// Time: O(1) - list removal and insertion are O(1) for doubly linked list
// Space: O(1)
func (c *LFUCache) incrementFreq(listNode *list.Element) {
	currentElement := listNode.Value.(*cacheEntry)
	prevFreq := currentElement.freq

	c.freqMap[prevFreq].Remove(listNode)
	if c.freqMap[prevFreq].Len() == 0 && c.minFreq == prevFreq {
		c.minFreq++
	}

	currentElement.freq++
	if c.freqMap[currentElement.freq] == nil {
		c.freqMap[currentElement.freq] = list.New()
	}
	newElem := c.freqMap[currentElement.freq].PushFront(currentElement)
	c.cache[currentElement.key] = newElem
}

// Time: O(1) - map lookup, possible evict O(1), list insertion O(1)
// Space: O(1) per call, O(capacity) total for stored entries
func (c *LFUCache) Put(key string, value any) {
	if c.capacity <= 0 {
		return
	}

	if item, ok := c.cache[key]; ok {
		item.Value.(*cacheEntry).value = value
		c.incrementFreq(item)
		return
	}

	if len(c.cache) >= c.capacity {
		c.evict()
	}

	entry := &cacheEntry{key: key, value: value, freq: 1}
	if c.freqMap[1] == nil {
		c.freqMap[1] = list.New()
	}
	elem := c.freqMap[1].PushFront(entry)
	c.cache[key] = elem
	c.minFreq = 1
}

// Time: O(1) - minFreq gives direct access to eviction candidate bucket,
//
//	Back() and Remove() are O(1) for doubly linked list
//
// Space: O(1)
func (c *LFUCache) evict() {
	bucket := c.freqMap[c.minFreq]
	elem := bucket.Back()
	e := elem.Value.(*cacheEntry)

	bucket.Remove(elem)
	delete(c.cache, e.key)
}
