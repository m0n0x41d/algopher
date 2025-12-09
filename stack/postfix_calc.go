package stack

import (
	"errors"
	"strconv"
	"strings"
)

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
