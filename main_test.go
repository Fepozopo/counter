package main_test

import (
	"testing"

	counter "github.com/Fepozopo/counter"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "5 words",
			input:    "one two three four five",
			expected: 5,
		},
		{
			name:     "empty input",
			input:    "",
			expected: 0,
		},
		{
			name:     "single space",
			input:    " ",
			expected: 0,
		},
		{
			name:     "new lines",
			input:    "one two three\nfour five",
			expected: 5,
		},
		{
			name:     "multiple spaces",
			input:    "This is a sentence.  This is another",
			expected: 7,
		},
		{
			name:     "Prefixed multiple spaces",
			input:    "  Hello",
			expected: 1,
		},
		{
			name:     "Suffixed multiple spaces",
			input:    "Hello     ",
			expected: 1,
		},
		{
			name:     "Tab character in code",
			input:    "Hello\tWord\n",
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountWords([]byte(tc.input))
			if result != tc.expected {
				t.Logf("expected: %d | got: %d", tc.expected, result)
				t.Fail()
			}
		})
	}
}
