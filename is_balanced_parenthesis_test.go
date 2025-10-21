package main

import "testing"

func TestIsBalancedParenthesis(t *testing.T) {
	entries := map[string]bool{
		"()()":     true,
		"(()))":    false,
		"((())())": true,
		"(()(()":   false,
		")(":       false,
		")()(()":   false,
		"())(()":   false,
		"(":        false,
		"()":       true,
		"(())":     true,
	}
	for entry, value := range entries {
		res := isBalancedParentensis(entry)
		if res != value {
			t.Fatalf("Wrong result for entry: %v , expected: %v, got: %v\n", entry, value, res)
		}
	}
}
