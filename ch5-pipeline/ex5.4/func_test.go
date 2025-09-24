package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestNormalizeLines(t *testing.T) {
	input := []string{
		"Programming STYles",
		"programming, styles!?",
		"programming12 styles345",
		"programming$%^*@():!styles",
		"",
		"@#$%^&*",
		"123456789",
		"programming12%^&*styles345",
		"Café naïve résumé",
		"programming\nstyles\t",
		"    ",
	}
	expected := []string{
		"programming styles",
		"programming  styles  ",
		"programming   styles   ",
		"programming         styles",
		"       ",
		"         ",
		"programming      styles   ",
		"café naïve résumé",
		"programming styles ",
		"    ",
	}

	output := normalize(input)

	if !slices.Equal(output, expected) {
		t.Errorf("got = %q, want = %q", output, expected)
	}
}

func TestNormalizeLines_IdempotentProperty(t *testing.T) {
	input := []string{
		"Programming STYles",
		"programming, styles!?",
		"programming12 styles345",
		"programming$%^*@():!styles",
		"",
		"@#$%^&*",
		"123456789",
		"programming12%^&*styles345",
		"Café naïve résumé",
		"programming\nstyles\t",
		"    ",
	}

	output1 := normalize(input)
	output2 := normalize(input)

	if !slices.Equal(output1, output2) {
		t.Errorf("not idempotent: output1 = %q, output2 = %q", output1, output2)
	}
}

func TestSeparateIntoPages(t *testing.T) {
	testCases := []struct {
		name          string
		nlinesPerPage int
		inputLines    []string
		expected      []page
	}{
		{
			name:          "basic pagination with exact pages",
			nlinesPerPage: 2,
			inputLines:    []string{"line1", "line2", "line3", "line4"},
			expected: []page{
				page{content: "line1\nline2", number: 1},
				page{content: "line3\nline4", number: 2},
			},
		}, {
			name:          "uneven pages",
			nlinesPerPage: 3,
			inputLines:    []string{"line1", "line2", "line3", "line4"},
			expected: []page{
				page{content: "line1\nline2\nline3", number: 1},
				page{content: "line4", number: 2},
			},
		}, {
			name:          "single line per page",
			nlinesPerPage: 1,
			inputLines:    []string{"line1", "line2", "line3", "line4"},
			expected: []page{
				page{content: "line1", number: 1},
				page{content: "line2", number: 2},
				page{content: "line3", number: 3},
				page{content: "line4", number: 4},
			},
		}, {
			name:          "lines less than page size",
			nlinesPerPage: 5,
			inputLines:    []string{"line1", "line2", "line3", "line4"},
			expected: []page{
				page{content: "line1\nline2\nline3\nline4", number: 1},
			},
		}, {
			name:          "empty input lines",
			nlinesPerPage: 2,
			inputLines:    []string{},
			expected:      []page{},
		}, {
			name:          "lines but with zero page size",
			nlinesPerPage: 0,
			inputLines:    []string{},
			expected:      []page{},
		}, {
			name:          "lines but with negative page size",
			nlinesPerPage: -2,
			inputLines:    []string{},
			expected:      []page{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			paginator := separateIntoPages(tc.nlinesPerPage)
			output := paginator(tc.inputLines)

			if !slices.Equal(output, tc.expected) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}

func TestRecordPageNumbersForWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    []page
		expected map[string][]int
	}{
		{
			name: "one page with unique words",
			input: []page{
				page{content: "word1 word2\nword3 word4", number: 1},
			},
			expected: map[string][]int{
				"word2": {1},
				"word4": {1},
				"word1": {1},
				"word3": {1},
			},
		}, {
			name: "one page with repeated words",
			input: []page{
				page{content: "word1 word2\nword1 word2\nword1", number: 1},
			},
			expected: map[string][]int{
				"word1": {1, 1, 1},
				"word2": {1, 1},
			},
		}, {
			name: "many pages with unique words",
			input: []page{
				page{content: "word1 word2\nword3 word4", number: 1},
				page{content: "word5 word6", number: 2},
				page{content: "word7", number: 3},
			},
			expected: map[string][]int{
				"word1": {1},
				"word2": {1},
				"word3": {1},
				"word4": {1},
				"word5": {2},
				"word6": {2},
				"word7": {3},
			},
		}, {
			name: "many pages with repeated words",
			input: []page{
				page{content: "word1 word2", number: 1},
				page{content: "word1 word2", number: 2},
				page{content: "word1 word3", number: 3},
			},
			expected: map[string][]int{
				"word1": {1, 2, 3},
				"word2": {1, 2},
				"word3": {3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := recordPageNumbersForWords(tc.input)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}

func TestFilterWords(t *testing.T) {
	testCases := []struct {
		name          string
		wordFreqLimit int
		input         map[string][]int
		expected      map[string][]int
	}{
		{
			name:          "words with frequency bellow limit and not singled-char",
			wordFreqLimit: 10,
			input: map[string][]int{
				"word1": {1, 2},
				"word2": {1, 2},
				"word3": {3, 3},
			},
			expected: map[string][]int{
				"word1": {1, 2},
				"word2": {1, 2},
				"word3": {3, 3},
			},
		}, {
			name:          "words with frequency above limit and not singled-char",
			wordFreqLimit: 3,
			input: map[string][]int{
				"word1": {1, 2, 3, 4, 5},
				"word2": {1, 2, 4},
				"word3": {3, 3},
				"word4": {4, 4, 4, 4},
			},
			expected: map[string][]int{
				"word2": {1, 2, 4},
				"word3": {3, 3},
			},
		}, {
			name:          "words with frequency bellow limit and singled-char",
			wordFreqLimit: 10,
			input: map[string][]int{
				"word1": {1, 2},
				"c":     {1, 2},
				"b":     {3, 3},
			},
			expected: map[string][]int{
				"word1": {1, 2},
			},
		}, {
			name:          "empty map",
			wordFreqLimit: 10,
			input:         map[string][]int{},
			expected:      map[string][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			filter := filterWords(tc.wordFreqLimit)
			output := filter(tc.input)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}

func TestRemoveDuplicatedPageNums(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string][]int
		expected map[string][]int
	}{
		{
			name: "words with no duplicated pages",
			input: map[string][]int{
				"word1": {1},
				"word2": {1, 2},
				"word3": {1, 3, 4, 5},
			},
			expected: map[string][]int{
				"word1": {1},
				"word2": {1, 2},
				"word3": {1, 3, 4, 5},
			},
		}, {
			name: "words with duplicated pages",
			input: map[string][]int{
				"word1": {1, 1},
				"word2": {1, 2, 2, 2},
				"word3": {4, 4, 4, 5},
				"word4": {3},
			},
			expected: map[string][]int{
				"word1": {1},
				"word2": {1, 2},
				"word3": {4, 5},
				"word4": {3},
			},
		}, {
			name:     "empty map",
			input:    map[string][]int{},
			expected: map[string][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := removeDuplicatedPageNumbers(tc.input)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("got = %q, want = %q", output, tc.expected)
			}
		})
	}
}
