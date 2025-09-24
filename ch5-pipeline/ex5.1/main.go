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
	// composing all functions together
	filepath := "../../files/input.txt"
	printAll(sortedPairs(countFrequencies(removeStopWords(convertToSlice(normalize(readData(filepath)))))))
}

func readData(filepath string) (fileContent string) {
	data, err := os.ReadFile(filepath)
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
	re := regexp.MustCompile(`\S+`)
	return re.FindAllString(text, -1)
}

func removeStopWords(words []string) []string {
	data, err := os.ReadFile("../../files/stopwords.txt")
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
	for _, w := range words {
		if _, ok := stopwords[w]; !ok {
			filteredWords = append(filteredWords, w)
		}
	}
	return filteredWords
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
	if len(wordsFreq) > 0 {
		fmt.Printf("%s  -  %d\n", wordsFreq[0].word, wordsFreq[0].freq)
		printAll(wordsFreq[1:])
	}
}
