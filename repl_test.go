package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello, test!",
			expected: []string{"hello,", "test!"},
		},
		{
			input:    "multiple   whitespaces between	words",
			expected: []string{"multiple", "whitespaces", "between", "words"},
		},
		{
			input:    "VariOUs CaSES oF ChArAcTeRs",
			expected: []string{"various", "cases", "of", "characters"},
		},
		{
			input:    "     preceeding whitespace",
			expected: []string{"preceeding", "whitespace"},
		},
		{
			input:    "following whitespace    ",
			expected: []string{"following", "whitespace"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("%v should result in length %v, resulted in %v", c.input, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%v does not match expected: %v", word, expectedWord)
			}
		}
	}
}
