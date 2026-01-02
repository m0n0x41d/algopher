package native_dict

import (
	//      "fmt"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var _ = strconv.Atoi
var _ = os.Args

const MAGIC_NUMBER = 42

var ErrKeyNotFound = errors.New("key not found")

type NativeDictionary[T any] struct {
	size   int
	step   int
	salt   uint // random salt for HashDoS protection
	slots  []string
	filled []bool
	values []T
}

// Time: O(n) where n = sz
// Space: O(n)
func Init[T any](sz int) NativeDictionary[T] {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nd := NativeDictionary[T]{
		size:   sz,
		step:   3,
		salt:   uint(r.Uint64()),
		slots:  nil,
		filled: nil,
		values: nil,
	}
	nd.slots = make([]string, sz)
	nd.filled = make([]bool, sz)
	nd.values = make([]T, sz)
	return nd
}

// Time: O(k) where k = len(value)
// Space: O(1)
func (nd *NativeDictionary[T]) HashFun(value string) int {
	hash := nd.salt
	for _, x := range []byte(value) {
		hash = hash*MAGIC_NUMBER + uint(x)
	}
	return int(hash % uint(nd.size))
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (nd *NativeDictionary[T]) IsKey(key string) bool {
	idx := nd.findSlot(key)
	return idx != -1
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (nd *NativeDictionary[T]) Get(key string) (T, error) {
	var result T
	idx := nd.findSlot(key)
	if idx == -1 {
		return result, ErrKeyNotFound
	}
	return nd.values[idx], nil
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (nd *NativeDictionary[T]) Put(key string, value T) {

	existingIdx := nd.findSlot(key)
	if existingIdx != -1 {
		nd.values[existingIdx] = value
		return
	}

	idx := nd.seekSlot(key)
	if idx != -1 {
		nd.slots[idx] = key
		nd.filled[idx] = true
		nd.values[idx] = value
	}
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (nd *NativeDictionary[T]) findSlot(key string) int {
	start := nd.HashFun(key)
	idx := start
	for {
		if !nd.filled[idx] {
			return -1
		}
		if nd.slots[idx] == key {
			return idx
		}
		idx = (idx + nd.step) % nd.size
		if idx == start {
			return -1
		}
	}
}

// Time: O(1) average, O(n) worst case
// Space: O(1)
func (nd *NativeDictionary[T]) seekSlot(key string) int {
	start := nd.HashFun(key)
	idx := start
	for {
		if !nd.filled[idx] {
			return idx
		}
		idx = (idx + nd.step) % nd.size
		if idx == start {
			return -1
		}
	}
}
