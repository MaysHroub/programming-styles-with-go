package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/MaysHroub/programming-styles-with-go/config"
)

type pair struct {
	word string
	freq int
}

func main() {
	nlines := 200

	data, err := os.ReadFile(config.InputFile)
	if err != nil {
		log.Fatalf("couldn't read file: %v", err)
	}
	chunks := partition(string(data), nlines)

	now := time.Now()

	pairsLists := map_(split, chunks)
	filteredPairs := removeStopWords(config.StopWordsFile, pairsLists)
	wordsFreq := reduce(countWords, filteredPairs)

	fmt.Printf("Execution time: %v\n\n", time.Since(now))

	sort.Slice(wordsFreq, func(i, j int) bool {
		return wordsFreq[i].freq >= wordsFreq[j].freq
	})

	for _, wf := range wordsFreq[:25] {
		fmt.Printf("%s  --  %d\n", wf.word, wf.freq)
	}
}

func reduce(countFunc func([]pair, []pair) []pair, pairsLists [][]pair) []pair {
	n := len(pairsLists)
	for i := 0; i < n-1; i += 2 {
		pairs1 := pairsLists[i]
		pairs2 := pairsLists[i+1]
		result := countFunc(pairs1, pairs2)
		pairsLists = append(pairsLists, result)
		n++
	}
	return pairsLists[n-1]
}

func map_(splitFunc func(string) []pair, dataChunks []string) [][]pair {
	pairsLists := [][]pair{}
	for _, chnk := range dataChunks {
		pairs := splitFunc(chnk)
		pairsLists = append(pairsLists, pairs)
	}
	return pairsLists
}

func partition(data string, nlines int) []string {
	lines := strings.Split(data, "\n")
	chunks := []string{}
	numOfChunks := int(math.Ceil(float64(len(lines)) / float64(nlines)))
	for i := 0; i < numOfChunks-1; i++ {
		start := i * nlines
		end := start + nlines
		chunck := strings.Join(lines[start:end], " ")
		chunks = append(chunks, chunck)
	}
	lastChunk := strings.Join(lines[(numOfChunks-1)*nlines:], " ")
	chunks = append(chunks, lastChunk)
	return chunks
}

func split(text string) []pair {
	normalized := normalize(text)
	re := regexp.MustCompile(`\s+`)
	words := re.Split(normalized, -1)
	pairs := []pair{}

	for _, w := range words {
		pairs = append(pairs, pair{word: w, freq: 1})
	}
	return pairs
}

func countWords(pairs1, pairs2 []pair) []pair {
	mp := map[string]int{}
	for _, p := range pairs1 {
		mp[p.word] += p.freq
	}
	for _, p := range pairs2 {
		mp[p.word] += p.freq
	}
	mergedPairs := []pair{}
	for w, f := range mp {
		mergedPairs = append(mergedPairs, pair{word: w, freq: f})
	}
	return mergedPairs
}

func removeStopWords(stopwordsfilepath string, pairsLists [][]pair) [][]pair {
	stopwords := getStopWords(stopwordsfilepath)
	filteredPairsList := [][]pair{}
	for _, pairs := range pairsLists {
		filteredPairs := []pair{}
		for _, p := range pairs {
			if _, ok := stopwords[p.word]; !ok {
				filteredPairs = append(filteredPairs, p)
			}
		}
		filteredPairsList = append(filteredPairsList, filteredPairs)
	}
	return filteredPairsList
}

func getStopWords(stopwordsfilepath string) map[string]struct{} {
	data, err := os.ReadFile(stopwordsfilepath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	spwords := strings.Split(string(data), ",")
	stopwords := map[string]struct{}{}
	for _, sw := range spwords {
		stopwords[strings.ToLower(sw)] = struct{}{}
	}
	for r := 'a'; r <= 'z'; r++ {
		stopwords[string(r)] = struct{}{}
	}
	return stopwords
}

func normalize(text string) string {
	str := strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return ' '
		}
		return r
	}, text)
	return strings.ToLower(str)
}
