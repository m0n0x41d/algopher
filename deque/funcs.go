package deque

func IsPalindrome(s string) bool {
	deq := Deque[rune]{}

	for _, r := range s {
		deq.AddFront(r)
	}

	for deq.Size() > 1 {
		front, _ := deq.RemoveFront()
		tail, _ := deq.RemoveTail()
		if front != tail {
			return false
		}
	}

	return true
}
