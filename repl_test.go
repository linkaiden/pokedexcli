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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  you  passed  this test",
			expected: []string{"you", "passed", "this", "test"},
		},
		{
			input:    "something  about   pokemon  I guess...",
			expected: []string{"something", "about", "pokemon", "i", "guess..."},
		},
		{
			input:    "a FEW tests  they said...",
			expected: []string{"a", "few", "tests", "they", "said..."},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Want len: %v, Got len: %v", len(actual), len(c.expected))
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Want: %v, Got: %v", expectedWord, word)
				t.FailNow()
			}
		}
	}
}
