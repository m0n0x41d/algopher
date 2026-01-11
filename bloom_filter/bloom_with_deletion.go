package bloomfilter

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
