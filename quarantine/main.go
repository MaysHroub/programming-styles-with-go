package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	word string
	freq int
}

func main() {
	qn := NewQuarantine()
	qn.bind(getInputFromConsole).
		bind(extractWords).
		bind(removeStopWords).
		bind(countFrequencies).
		bind(toSortedPairs).
		bind(printTop25).
		execute()
}

// i/o infected
func getInputFromConsole(any) (functionReturningConsoleInput any) {
	return func() any {
		var input string
		fmt.Scanf("%s", &input)
		return input
	}
}

// i/o infected
func extractWords(filename any) (functionReturningExtractedWords any) {
	return func() any {
		filename_ := filename.(string)
		data, err := os.ReadFile(filename_)
		if err != nil {
			return nil
		}
		normalizedText := strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return unicode.ToLower(r)
			}
			return ' '
		}, string(data))

		re := regexp.MustCompile(`\s+`)
		words := re.Split(normalizedText, -1)
		return words
	}
}

// also i/o infected
func removeStopWords(words any) (functionReturningFilteredWords any) {
	return func() any {
		words_ := words.([]string)
		data, err := os.ReadFile("../stopwords.txt")
		if err != nil {
			return nil
		}
		stopwordsAsSlice := strings.Split(string(data), ",")
		stopwords := make(map[string]struct{})
		for _, sp := range stopwordsAsSlice {
			stopwords[strings.TrimSpace(sp)] = struct{}{}
		}
		for r := 'a'; r <= 'z'; r++ {
			stopwords[string(r)] = struct{}{}
		}
		filteredWords := []string{}
		for _, w := range words_ {
			if _, ok := stopwords[w]; !ok {
				filteredWords = append(filteredWords, w)
			}
		}
		return filteredWords
	}
}

// pure function
func countFrequencies(words any) (wordsFreqMap any) {
	words_ := words.([]string)
	wordsFreq := make(map[string]int)
	for _, w := range words_ {
		wordsFreq[w]++
	}
	return wordsFreq
}

// pure function
func toSortedPairs(wordsFreqMap any) (sortedPairs any) {
	wordsFreqMap_ := wordsFreqMap.(map[string]int)
	pairs := []pair{}
	for w, f := range wordsFreqMap_ {
		pairs = append(pairs, pair{word: w, freq: f})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq >= pairs[j].freq
	})
	return pairs
}

// pure function
func getTop25(pairs any) (top25AsString any) {
	pairs_ := pairs.([]pair)
	top25 := ""
	for _, p := range pairs_ {
		top25 += fmt.Sprintf("%s  -  %d\n", p.word, p.freq)
	}
	return top25
}

// as part of exercise 24.3
func printTop25(pairs any) (functionPrintingToStdout any) {
	return func() any {
		pairs_ := pairs.([]pair)
		for _, p := range pairs_ {
			fmt.Printf("%s  -  %d\n", p.word, p.freq)
		}
		return nil
	}
}
