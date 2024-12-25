package calculator

import (
	"testing"
)

func TestIsValidExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   bool
	}{
		{"2 + 2", true},
		{"(1 - 2) * 3", true},
		{"5 / 2.5", true},
		{"2 + a", false},
		{"", false},
		{"5 * (3 + 2)", true},
		{"5 * (3 + 2) - 1", true},
		{"(5 + 3))", false},
	}

	for _, test := range tests {
		result := IsValidExpression(test.expression)
		if result != test.expected {
			t.Errorf("IsValidExpression(%q) = %v; expected %v", test.expression, result, test.expected)
		}
	}
}
