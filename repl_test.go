package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		// Basic string
		{
			input: "This is simple text",
			expected: []string{"this", "is", "simple", "text"},
		},
		// No valid separator
		{
			input: "These,are;not/the-separators,you/are,looking;for",
			expected: []string{"these,are;not/the-separators,you/are,looking;for"},
		},
		// Trim Whitespace
		{
			input: "       I NEED A TRIMMING     ",
			expected: []string{"i", "need", "a", "trimming"},
		},
		// Double Whitespace
		{
			input: "Why  would  you  do  this? why?  !  !!",
			expected: []string{"why", "would", "you", "do", "this?", "why?", "!", "!!"},
		},
	}
	// Run testcases
	for _, tc := range cases {
		// Get Actual Value
		actual := cleanInput(tc.input)
		// Check Lengths Match
		if len(actual) != len(tc.expected) {
			t.Errorf("[TestCleanInput] Mismatched Length. Expected: %d; Actual: %d", len(tc.expected), len(actual))
		}
		// Check Each Element Matches
		for i := range actual {
			if actual[i] != tc.expected[i] {
				t.Errorf("[TestCleanInput] Mismatched Slices. Expected: %v; Actual: %v", tc.expected, actual)
			}
		}

	}
}