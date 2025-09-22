package thing

import (
	"sort"
)

type Pair struct {
	Token string
	Freq  int
}

type TokenProcessor struct {
	tokens []string
	freqs  map[string]int
}

func NewTokenProcessor(tokens []string) *TokenProcessor {
	return &TokenProcessor{
		tokens: tokens,
	}
}

func (t *TokenProcessor) Clean(tokensToRemove map[string]struct{}) {
	filtered := []string{}
	for _, tk := range t.tokens {
		if _, ok := tokensToRemove[tk]; !ok {
			filtered = append(filtered, tk)
		}
	}
	t.tokens = filtered
}

func (t *TokenProcessor) CountFrequencies() {
	t.freqs = make(map[string]int)
	for _, tk := range t.tokens {
		t.freqs[tk]++
	}
}

func (t *TokenProcessor) ToSortedPairs() []Pair {
	pairs := []Pair{}
	for tk, f := range t.freqs {
		pairs = append(pairs, Pair{Token: tk, Freq: f})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Freq >= pairs[j].Freq
	})
	return pairs
}
