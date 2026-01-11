package bloomfilter

import (
	"os"
)

var _ = os.Stdout

const HEX_31 = 0x1F

// Magic numbers are laboratory server requirements
// Either way, there will be no good option to test students' code anyhow :D
const MAGIC_1 = 17
const MAGIC_2 = 223

type BloomFilter struct {
	filter_len int
	bitmask    uint32
}

func NewBloomFilter(f_len int) *BloomFilter {
	return &BloomFilter{
		filter_len: f_len,
		bitmask:    0,
	}
}

// All operations are O(n), where n = len(input string).
// And with fixed-length keys, it is even O(1).

func (bf *BloomFilter) hasher(s string, veryRandomNumber int) int {
	hashSum := 0
	for _, char := range s {
		symbolCode := int(char)
		hashSum = (hashSum*veryRandomNumber + symbolCode) % bf.filter_len
	}
	return hashSum
}

// HashFunc with magic number 17
func (bf *BloomFilter) Hash1(s string) int {
	return bf.hasher(s, MAGIC_1)
}

// HashFunc with magic number 223
func (bf *BloomFilter) Hash2(s string) int {
	return bf.hasher(s, MAGIC_2)
}

func (bf *BloomFilter) Add(s string) {
	position1 := bf.Hash1(s)
	position2 := bf.Hash2(s)
	bf.bitmask |= 1 << uint(position1)
	bf.bitmask |= 1 << uint(position2)
}

func (bf *BloomFilter) IsValue(s string) bool {
	pos1 := bf.Hash1(s)
	pos2 := bf.Hash2(s)
	return (bf.bitmask>>uint(pos1))&1 == 1 && (bf.bitmask>>uint(pos2))&1 == 1
}
