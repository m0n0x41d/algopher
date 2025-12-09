package stack

// First implementation only  for ()
//
//	func IsBalanced(s string) bool {
//		stack := Stack[rune]{}
//
//		for _, char := range s {
//			switch char {
//			case '(':
//				stack.Push(char)
//			case ')':
//				if stack.Size() == 0 {
//					return false
//				}
//				stack.Pop()
//			}
//		}
//
//		return stack.Size() == 0
//	}

func IsBalanced(s string) bool {
	stack := Stack[rune]{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			if stack.Size() == 0 {
				return false
			}
			// err must be unreachable
			item, _ := stack.Pop()
			if item != pairs[char] {
				return false
			}

		}
	}

	return stack.Size() == 0
}
