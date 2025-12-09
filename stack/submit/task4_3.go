package main

import (
	"cmp"
	"errors"
	"strconv"
	"strings"
)

// Additional tests for these tasts are locaed in upper dir.

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

// Ordered stack with min and avg

// Numeric constraint for types that support arithmetic operations
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// === Stack with O(1) Min support ===
// Requires ordered types (int, float, string, etc.)
// Uses auxiliary stack to track minimum history.

type OrderedStack[T cmp.Ordered] struct {
	container []T
	minStack  []T
}

// === Stack with O(1) Min and Avg support ===
// Requires numeric types only (no strings).

type NumericStack[T Numeric] struct {
	container []T
	minStack  []T
	sum       T
}

// O(1)
func (st *OrderedStack[T]) Size() int {
	return len(st.container)
}

// O(1)
func (st *OrderedStack[T]) Peek() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	return st.container[len(st.container)-1], nil
}

// O(1)
func (st *OrderedStack[T]) Pop() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	result := st.container[len(st.container)-1]
	st.container = st.container[:len(st.container)-1]
	st.minStack = st.minStack[:len(st.minStack)-1]
	return result, nil
}

// Amortized O(1)
func (st *OrderedStack[T]) Push(itm T) {
	st.container = append(st.container, itm)
	if len(st.minStack) == 0 || itm < st.minStack[len(st.minStack)-1] {
		st.minStack = append(st.minStack, itm)
	} else {
		st.minStack = append(st.minStack, st.minStack[len(st.minStack)-1])
	}
}

// O(1) - returns current minimum element
func (st *OrderedStack[T]) Min() (T, error) {
	if len(st.minStack) == 0 {
		var zero T
		return zero, nil
	}
	return st.minStack[len(st.minStack)-1], nil
}

// === NumericStack methods ===

// O(1)
func (st *NumericStack[T]) Size() int {
	return len(st.container)
}

// O(1)
func (st *NumericStack[T]) Peek() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	return st.container[len(st.container)-1], nil
}

// O(1)
func (st *NumericStack[T]) Pop() (T, error) {
	if len(st.container) == 0 {
		var zero T
		return zero, nil
	}
	result := st.container[len(st.container)-1]
	st.container = st.container[:len(st.container)-1]
	st.minStack = st.minStack[:len(st.minStack)-1]
	st.sum -= result
	return result, nil
}

// Amortized O(1)
func (st *NumericStack[T]) Push(itm T) {
	st.container = append(st.container, itm)
	st.sum += itm
	if len(st.minStack) == 0 || itm < st.minStack[len(st.minStack)-1] {
		st.minStack = append(st.minStack, itm)
	} else {
		st.minStack = append(st.minStack, st.minStack[len(st.minStack)-1])
	}
}

// O(1) - returns current minimum element
func (st *NumericStack[T]) Min() (T, error) {
	if len(st.minStack) == 0 {
		var zero T
		return zero, nil
	}
	return st.minStack[len(st.minStack)-1], nil
}

// O(1) - returns average of all elements
func (st *NumericStack[T]) Avg() float64 {
	if len(st.container) == 0 {
		return 0
	}
	return float64(st.sum) / float64(len(st.container))
}

// Postfix Calc

type Expression string

type CalculationResult struct {
	Value int
	Err   error
}

func (r CalculationResult) String() string {
	if r.Err != nil {
		return "error: " + r.Err.Error()
	}
	return strconv.Itoa(r.Value)
}

func (r CalculationResult) IsError() bool {
	return r.Err != nil
}

// PostfixCalculator evaluates a postfix expression using two stacks.
// Supported operations: +, *, =
// Example: "8 2 + 5 * 9 + =" returns CalculationResult{Value: 59}
//
// Algorithm:
// - Numbers are pushed to the operand stack (S2)
// - Operators pop two operands, compute result, push back to S2
// - "=" returns the top of S2 as result
func PostfixCalculator(expression Expression) CalculationResult {
	operands := Stack[int]{}

	tokens := strings.Fields(string(expression))

	for _, token := range tokens {
		switch token {
		case "+":
			if operands.Size() < 2 {
				return CalculationResult{Err: errors.New("not enough operands")}
			}
			b, _ := operands.Pop()
			a, _ := operands.Pop()
			operands.Push(a + b)
		case "-":
			if operands.Size() < 2 {
				return CalculationResult{Err: errors.New("not enough operands")}
			}
			b, _ := operands.Pop()
			a, _ := operands.Pop()
			operands.Push(a - b)
		case "*":
			if operands.Size() < 2 {
				return CalculationResult{Err: errors.New("not enough operands")}
			}
			b, _ := operands.Pop()
			a, _ := operands.Pop()
			operands.Push(a * b)
		case "/":
			if operands.Size() < 2 {
				return CalculationResult{Err: errors.New("not enough operands")}
			}
			b, _ := operands.Pop()
			a, _ := operands.Pop()
			if b == 0 {
				return CalculationResult{Err: errors.New("division by zero")}
			}
			operands.Push(a / b)
		case "=":
			if operands.Size() == 0 {
				return CalculationResult{Err: errors.New("empty stack")}
			}
			result, _ := operands.Pop()
			return CalculationResult{Value: result}
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				return CalculationResult{Err: errors.New("invalid token " + token)}
			}
			operands.Push(num)
		}
	}

	if operands.Size() > 0 {
		result, _ := operands.Pop()
		return CalculationResult{Value: result}
	}

	return CalculationResult{
		Err: errors.New("empty expression"),
	}
}
