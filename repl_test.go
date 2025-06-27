package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "    ",
			expected: []string{""},
		},
		{
			input: "Normal",
			expected: []string{"normal"},
		},
	// add more cases here
	}

	failed, passed := 0, 0
	for _, c := range cases {
		previousFailed := failed
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The size of actual %v does not match expected %v", actual, c.expected)
			failed++
			continue
		}
		// if they don't match, use t.Errorf to print an error messag
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %s but got %s", expectedWord, word)
				failed++
				break
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
		if previousFailed == failed {
			passed++
		}
	}
	fmt.Printf("Tests failed: %v\nTests Passed: %v\n", failed, passed)
}