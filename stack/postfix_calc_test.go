package stack

import "testing"

func TestPostfixCalculator(t *testing.T) {
	tests := []struct {
		name       string
		expression Expression
		expected   string
	}{
		// Basic operations
		{"simple add", "1 2 +", "3"},
		{"simple subtract", "5 3 -", "2"},
		{"simple multiply", "3 4 *", "12"},
		{"simple divide", "8 2 /", "4"},

		// From task description
		{"task example 1", "1 2 + 3 *", "9"},
		{"task example 2", "8 2 + 5 * 9 + =", "59"},

		// Multiple operations
		{"chained add", "1 2 + 3 +", "6"},
		{"mixed ops", "2 3 + 4 *", "20"},

		// Single number
		{"single number", "42", "42"},
		{"single with equals", "42 =", "42"},

		// Errors
		{"not enough operands", "+ 1 2", "error: not enough operands"},
		{"empty expression", "", "error: empty expression"},
		{"empty stack on equals", "=", "error: empty stack"},
		{"division by zero", "5 0 /", "error: division by zero"},
		{"invalid token", "1 foo +", "error: invalid token foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostfixCalculator(tt.expression)
			if got.String() != tt.expected {
				t.Errorf("PostfixCalculator(%q) = %q, want %q", tt.expression, got.String(), tt.expected)
			}
		})
	}
}
