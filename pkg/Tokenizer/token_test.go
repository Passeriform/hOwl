package tokenizer

import (
	"strings"
	"testing"
)

func TestParentheses(t *testing.T) {
	input := NewScanner(strings.NewReader("("))

	if !isParentheses(input) {
		t.Errorf("expected true but got %v", isParentheses(input))
	}
}
