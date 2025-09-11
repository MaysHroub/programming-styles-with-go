package manager

import (
	"fmt"
	"sort"
)

type WordFreqManager struct {
	wordsFreq map[string]int
}

type Pair struct {
	word string
	freq int
}

func (p Pair) ToString() string {
	return fmt.Sprintf("%s  -  %d", p.word, p.freq)
}

func NewWordFreqManager() WordFreqManager {
	return WordFreqManager{
		wordsFreq: make(map[string]int),
	}
}

func (w WordFreqManager) IncrementCount(word string) {
	w.wordsFreq[word]++
}

func (w WordFreqManager) ToSortedPairs() []Pair {
	pairs := []Pair{}

	for w, f := range w.wordsFreq {
		pairs = append(pairs, Pair{word: w, freq: f})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq >= pairs[j].freq
	})

	return pairs
}
