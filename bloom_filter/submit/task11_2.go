package main

// This is counting bloom filter variant that tries to supports deletion by replacing bits with counters ¯\_(ツ)_/¯
// but the thing is that remove on false positives will still corrupt the filter and cause false negatives.
// So... the precondition for using such implementation safely might be informally state like "PLEASE lnly call Remove for elements that you were actually Added!!!"

type CountingBloomFilter struct {
	filter_len int
	counters   []uint8
}

func NewCountingBloomFilter(f_len int) *CountingBloomFilter {
	return &CountingBloomFilter{
		filter_len: f_len,
		counters:   make([]uint8, f_len),
	}
}

func (cbf *CountingBloomFilter) hasher(s string, salt int) int {
	hashSum := 0
	for _, char := range s {
		hashSum = (hashSum*salt + int(char)) % cbf.filter_len
	}
	return hashSum
}

func (cbf *CountingBloomFilter) Hash1(s string) int {
	return cbf.hasher(s, MAGIC_1)
}

func (cbf *CountingBloomFilter) Hash2(s string) int {
	return cbf.hasher(s, MAGIC_2)
}

func (cbf *CountingBloomFilter) Add(s string) {
	pos1 := cbf.Hash1(s)
	pos2 := cbf.Hash2(s)

	if cbf.counters[pos1] < 255 {
		cbf.counters[pos1]++
	}
	if cbf.counters[pos2] < 255 {
		cbf.counters[pos2]++
	}
}

func (cbf *CountingBloomFilter) Remove(s string) {
	pos1 := cbf.Hash1(s)
	pos2 := cbf.Hash2(s)

	if cbf.counters[pos1] > 0 {
		cbf.counters[pos1]--
	}
	if cbf.counters[pos2] > 0 {
		cbf.counters[pos2]--
	}
}

func (cbf *CountingBloomFilter) IsValue(s string) bool {
	pos1 := cbf.Hash1(s)
	pos2 := cbf.Hash2(s)
	return cbf.counters[pos1] > 0 && cbf.counters[pos2] > 0
}

// Merge...?

// The probability of false positives increases somewhat exponentially with the number of elements in the merged filters. This check works best with fairly large filters and is worst when there are either a large number of filters (non-empty) or elements.
func Merge(filters ...*BloomFilter) *BloomFilter {
	if len(filters) == 0 {
		return nil
	}

	result := NewBloomFilter(filters[0].filter_len)
	for _, f := range filters {
		result.bitmask |= f.bitmask
	}
	return result
}

// Regarding task 4 (restoring the original set of values added to the filter):
// TL;DR – impossible, it is an irreversible morphism.
// It feels like a fundamentally unsolvable task in general if we are talking about a "purely Bloom" filter because:
// 1. A pure Bloom filter stores only bits, no elements or parts of elements.
// 2. High collisions (likely) - different strings/elements might and likely will have the same hash values.
// 3. If we are talking especially about our implementation with a 32-bit bitmask –
// it is not possible and not in general - the domain is potentially infinite, and our codomain is just a 32-bit bitmask. So?
//
// If we have some set of "candidates" to check whether they were added to the Bloom filter or not, we might still just iterate over these candidates and check the Bloom filter, keeping false positives in mind, but this is nowhere near "restoring."
