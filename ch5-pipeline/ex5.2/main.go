package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/MaysHroub/programming-styles-with-go/config"
)

type pair struct {
	word string
	freq int
}

func main() {
	printAll(sortedPairs(countFrequencies(removeStopWordsGivenFileName(config.StopWordsFile)(convertToSlice(normalize(readData(config.InputFile)))))))
}

func readData(filename string) (fileContent string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func normalize(text string) (normalizedText string) {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, text)
}

func convertToSlice(text string) []string {
	re := regexp.MustCompile(`\s+`)
	return re.Split(text, -1)
}

func countFrequencies(words []string) map[string]int {
	wordsFreq := make(map[string]int)
	for _, w := range words {
		wordsFreq[w]++
	}
	return wordsFreq
}

func sortedPairs(wordsFreq map[string]int) []pair {
	pairs := []pair{}
	for w, f := range wordsFreq {
		pairs = append(pairs, pair{word: w, freq: f})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq >= pairs[j].freq
	})
	return pairs
}

func printAll(wordsFreq []pair) {
	for _, wf := range wordsFreq[:min(25, len(wordsFreq))] {
		fmt.Printf("%s  -  %d\n", wf.word, wf.freq)
	}
}

func removeStopWordsGivenFileName(filename string) func([]string) []string {
	return func(words []string) []string {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil
		}
		stopwordsAsSlice := strings.Split(string(data), ",")
		stopwords := make(map[string]struct{})
		for _, sp := range stopwordsAsSlice {
			stopwords[strings.ToLower(strings.TrimSpace(sp))] = struct{}{}
		}
		for r := 'a'; r <= 'z'; r++ {
			stopwords[string(r)] = struct{}{}
		}
		filteredWords := []string{}
		for _, w := range words {
			if w == "" {
				continue
			}
			if _, ok := stopwords[w]; !ok {
				filteredWords = append(filteredWords, w)
			}
		}
		return filteredWords
	}
}
