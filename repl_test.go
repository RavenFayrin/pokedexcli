package main

import "testing"

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
			input:    "TESTING  Clean InPuTs",
			expected: []string{"testing", "clean", "inputs"},
		},
		{
			input:    "YAY!",
			expected: []string{"yay!"},
		},
	}

	for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("ERROR: Slice length does not match.\nWant: %d\nGot: %d", len(c.expected), len(actual))
	} 
	
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("ERROR: Words do not match.\nWant: %s\nGot: %s", expectedWord, word)
			}
		}
}
}