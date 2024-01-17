package stack

import "testing"

func TestBracketMatch(t *testing.T) {
	str := "([{}])"
	if !BracketMatch(str) {
		t.Error("BracketMatch error")
	}
}
