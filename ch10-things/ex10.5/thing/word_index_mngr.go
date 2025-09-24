package thing

import (
	"fmt"
	"sort"
)

type WordIndexManager struct {
	wordPages map[string][]int
}

func NewWordIndexManager(wordPages map[string][]int) *WordIndexManager {
	return &WordIndexManager{
		wordPages: wordPages,
	}
}

func (wm *WordIndexManager) PrintSorted() {
	words := []string{}
	for w := range wm.wordPages {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words[:min(25, len(words))] {
		fmt.Printf("word: %v\npages: %v\n\n", w, wm.wordPages[w])
	}
}
