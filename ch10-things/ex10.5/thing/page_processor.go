package thing

import (
	"regexp"
	"unicode/utf8"
)

type page struct {
	content string
	number  int
}

type PageProcessor struct {
	wordFreqLimit int
	pages         []page
	wordPages     map[string][]int
}

func NewPageProcessor(pages []page, wordFreqLimit int) *PageProcessor {
	return &PageProcessor{
		pages:         pages,
		wordFreqLimit: wordFreqLimit,
	}
}

func (p *PageProcessor) SplitAndCountWords() {
	re := regexp.MustCompile(`\s+`)
	p.wordPages = map[string][]int{}
	for _, page := range p.pages {
		words := re.Split(page.content, -1)
		for _, w := range words {
			p.wordPages[w] = append(p.wordPages[w], page.number)
		}
	}
}

func (p *PageProcessor) FilterWords() {
	filtered := make(map[string][]int)
	for w, nums := range p.wordPages {
		if len(nums) > p.wordFreqLimit || utf8.RuneCountInString(w) <= 1 {
			continue
		}
		filtered[w] = nums
	}
	p.wordPages = filtered
}

func (p *PageProcessor) RemoveDuplicatedPageNums() {
	for w, nums := range p.wordPages {
		mp := make(map[int]struct{})
		uniqueNums := []int{}
		for _, n := range nums {
			if _, ok := mp[n]; !ok {
				uniqueNums = append(uniqueNums, n)
				mp[n] = struct{}{}
			}
		}
		p.wordPages[w] = uniqueNums
	}
}

func (p *PageProcessor) GetWordPagesMap() map[string][]int {
	return p.wordPages
}
