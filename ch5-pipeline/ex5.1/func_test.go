package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "text with mixed cases",
			input:    "Programming STYles",
			expected: "programming styles",
		}, {
			name:     "text with punctuation",
			input:    "programming, styles!?",
			expected: "programming  styles  ",
		}, {
			name:     "text with numbers",
			input:    "programming12 styles345",
			expected: "programming12 styles345",
		}, {
			name:     "text with special characters",
			input:    "programming$%^*@():!styles",
			expected: "programming         styles",
		}, {
			name:     "empty text",
			input:    "",
			expected: "",
		}, {
			name:     "text with only special characters",
			input:    "@#$%^&*",
			expected: "       ",
		}, {
			name:     "text with only numbers",
			input:    "123456789",
			expected: "123456789",
		}, {
			name:     "text with numbers and special characters",
			input:    "programming12%^&*styles345",
			expected: "programming12    styles345",
		}, {
			name:     "text with unicode characters",
			input:    "Café naïve résumé",
			expected: "café naïve résumé",
		}, {
			name:     "text with newlines and tabs",
			input:    "programming\nstyles\t",
			expected: "programming styles ",
		}, {
			name:     "text with whitespaces only",
			input:    "    ",
			expected: "    ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(tc.input)
			if got != tc.expected {
				t.Errorf("normalize(%q) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestNormalizeProperties(t *testing.T) {
	inputs := []string{
		"Programming Styles",
		"test123",
		"test@#$%^&",
		"#$%^&,.!?",
		" ",
	}
	t.Run("idempotent property", func(t *testing.T) {
		for _, inp := range inputs {
			output1 := normalize(inp)
			output2 := normalize(inp)
			if output1 != output2 {
				t.Errorf("normalize is not idempotent: normalize(%q) = %q, but normalize(%q) = %q", inp, output1, inp, output2)
			}
		}
	})
	t.Run("length property", func(t *testing.T) {
		for _, inp := range inputs {
			output := normalize(inp)
			if len(output) != len(inp) {
				t.Errorf("normalize changed input's length: len(input)=%d, len(output)=%d", len(inp), len(output))
			}
		}
	})
}

func TestConvertToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "text with single space between words",
			input:    "programming styles in go",
			expected: []string{"programming", "styles", "in", "go"},
		}, {
			name:     "text with multiple spaces between words",
			input:    "programming      styles   in               go",
			expected: []string{"programming", "styles", "in", "go"},
		}, {
			name:     "text with new lines and tabs",
			input:    "programming\nstyles\tin\r\ngo",
			expected: []string{"programming", "styles", "in", "go"},
		}, {
			name:     "text with several whitespaces",
			input:    "programming\n   styles  \tin \r\ngo",
			expected: []string{"programming", "styles", "in", "go"},
		}, {
			name:     "text with single word surrounded by whitespaces",
			input:    " one  \t",
			expected: []string{"one"},
		}, {
			name:     "text with whitespaces only",
			input:    "   \t \n\n  \t   ",
			expected: []string{},
		}, {
			name:     "empty text",
			input:    "",
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := convertToSlice(tc.input)
			if !slices.Equal(got, tc.expected) {
				t.Errorf("got = %q, want = %q", got, tc.expected)
			}
		})
	}
}

func TestConvertToSliceIdempotentProperty(t *testing.T) {
	inputs := []string{
		"programming styles",
		"programming      styles   ",
		"\nprogramming \nstyles\t\tin go\n",
	}
	for _, inp := range inputs {
		output1 := convertToSlice(inp)
		output2 := convertToSlice(inp)
		if !slices.Equal(output1, output2) {
			t.Errorf("ConvertToSlice is not idempotent: output1 = %q, output2 = %q", output1, output2)
		}
	}
}

func TestCountFrequencies(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:  "slice with several words",
			input: []string{"hello", "world", "world"},
			expected: map[string]int{
				"hello": 1,
				"world": 2,
			},
		}, {
			name:  "slice with one word",
			input: []string{"hello"},
			expected: map[string]int{
				"hello": 1,
			},
		}, {
			name:  "slice with unique words",
			input: []string{"hello", "world", "go", "bye"},
			expected: map[string]int{
				"hello": 1,
				"world": 1,
				"go":    1,
				"bye":   1,
			},
		}, {
			name:  "slice with repeated word",
			input: []string{"hello", "hello", "hello", "hello"},
			expected: map[string]int{
				"hello": 4,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := countFrequencies(tc.input)
			if !reflect.DeepEqual(tc.expected, output) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}

func TestCountFrequenciesIdempotent(t *testing.T) {
	inputs := [][]string{
		{"hello", "world", "world"},
		{"hello"},
		{"hello", "world", "go", "bye"},
		{"hello", "hello", "hello", "hello"},
	}

	for _, inp := range inputs {
		output1 := countFrequencies(inp)
		output2 := countFrequencies(inp)

		if !reflect.DeepEqual(output1, output2) {
			t.Errorf("countFrequencies is not idempotent: output1 = %q, output2 = %q", output1, output2)
		}
	}
}

func TestSortedPairs(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]int
		expected []pair
	}{
		{
			name: "words with different frequencies",
			input: map[string]int{
				"hello": 1,
				"bye":   4,
				"hi":    25,
				"go":    9,
			},
			expected: []pair{
				{word: "hi", freq: 25},
				{word: "go", freq: 9},
				{word: "bye", freq: 4},
				{word: "hello", freq: 1},
			},
		}, {
			name: "single word",
			input: map[string]int{
				"hello": 2,
			},
			expected: []pair{
				{word: "hello", freq: 2},
			},
		}, {
			name:     "empty map",
			input:    map[string]int{},
			expected: []pair{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := sortedPairs(tc.input)
			if !slices.Equal(tc.expected, output) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}
