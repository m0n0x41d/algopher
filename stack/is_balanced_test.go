package stack

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Balanced cases - parentheses only
		{"empty string", "", true},
		{"single pair", "()", true},
		{"nested", "(())", true},
		{"sequential", "()()", true},
		{"complex nested", "(()((())()))", true},
		{"deep nesting", "(((())))", true},
		{"mixed", "(()())()", true},

		// Unbalanced cases - parentheses only
		{"only open", "(", false},
		{"only close", ")", false},
		{"extra close", "())", false},
		{"extra open", "(()", false},
		{"wrong order", ")(", false},
		{"close then open", ")()(", false},
		{"unbalanced from task", "())(", false},
		{"double wrong", "))((", false},
		{"almost balanced", "((())", false},
		{"close in middle", "(()()(())", false},

		// Balanced cases - multiple bracket types
		{"curly pair", "{}", true},
		{"square pair", "[]", true},
		{"all types simple", "(){}[]", true},
		{"nested different types", "{[()]}", true},
		{"complex multi-type", "({[]})[{}]", true},
		{"deeply nested multi", "{([{()}])}", true},
		{"sequential multi", "(){}[](){}[]", true},
		{"mixed nesting", "[{()}()]", true},

		// Unbalanced cases - multiple bracket types
		{"mismatched pair 1", "(]", false},
		{"mismatched pair 2", "[}", false},
		{"mismatched pair 3", "{)", false},
		{"open curly only", "{", false},
		{"close square only", "]", false},
		{"wrong close type", "([)]", false},
		{"interleaved wrong", "[(])", false},
		{"almost right multi", "{[()}", false},
		{"nested mismatch", "{[(])}", false},
		{"extra close curly", "{}}", false},
		{"extra open square", "[[]", false},
		{"all opens", "([{", false},
		{"all closes", ")]}", false},
		{"reversed closes", "([{}]))", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBalanced(tt.input)
			if got != tt.expected {
				t.Errorf("IsBalanced(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
