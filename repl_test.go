package main

import "testing"

// func TestCommandHelp(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{}
//
// 	for _, c := range cases {
// 		actual := commandExit(&config{})
//
// 	}
// }

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			// error and continue
			t.Errorf("lengths don't match, Actual length:%d, Got: %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice
			// if they don't match. use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("actual word: %s\ngot: %s", expectedWord, word)
			}
		}

	}
}
