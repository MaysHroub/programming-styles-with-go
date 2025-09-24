package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/MaysHroub/programming-styles-with-go/config"
)

const (
	FREQ_LIMIT_PER_WORD = 100
)

type page struct {
	content string
	number  int
}

func main() {
	nlinesPerPage := 45

	file, err := os.Open(config.InputFile)
	if err != nil {
		log.Fatalf("couldn't open file: %v\n", err)
	}
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	normalizedLines := normalize(lines)
	pages := separateIntoPages(normalizedLines, nlinesPerPage)
	wordPagesList := map_(splitAndCount, pages)
	reducedWordPagesMp := reduce(merge, wordPagesList)
	removeDuplicatedPageNums(&reducedWordPagesMp)
	sortedWords := toSortedWords(reducedWordPagesMp)
	print_(reducedWordPagesMp, sortedWords)
}

func print_(wordPages map[string][]int, words []string) {
	for _, w := range words[:25] {
		fmt.Printf("word: %v\npages: %v\n\n", w, wordPages[w])
	}
}

func toSortedWords(wordPages map[string][]int) []string {
	words := []string{}
	for w := range wordPages {
		words = append(words, w)
	}
	sort.Strings(words)
	return words
}

func removeDuplicatedPageNums(wordPages *map[string][]int) {
	for w, nums := range *wordPages {
		mp := make(map[int]struct{})
		uniqueNums := []int{}
		for _, n := range nums {
			if _, ok := mp[n]; !ok {
				uniqueNums = append(uniqueNums, n)
				mp[n] = struct{}{}
			}
		}
		(*wordPages)[w] = uniqueNums
	}
}

func reduce(
	mergeFunc func(map[string][]int, map[string][]int) map[string][]int,
	wordPagesList []map[string][]int,
) map[string][]int {
	n := len(wordPagesList)
	for i := 0; i < n-1; i += 2 {
		mergedMp := mergeFunc(wordPagesList[i], wordPagesList[i+1])
		wordPagesList = append(wordPagesList, mergedMp)
		n++
	}
	return wordPagesList[n-1]
}

func map_(
	splitFunc func(page, *regexp.Regexp) map[string][]int,
	pages []page,
) []map[string][]int {
	re := regexp.MustCompile(`\s+`)
	wordPages := []map[string][]int{}
	for _, page := range pages {
		mp := splitFunc(page, re)
		wordPages = append(wordPages, mp)
	}
	return wordPages
}

func merge(mp1, mp2 map[string][]int) map[string][]int {
	// merge results into mp1
	for w, nums := range mp2 {
		mp1[w] = append(mp1[w], nums...)
	}
	filterWords(&mp1)
	return mp1
}

func filterWords(mp *map[string][]int) {
	for w, nums := range *mp {
		if len(nums) > FREQ_LIMIT_PER_WORD || utf8.RuneCountInString(w) <= 1 {
			delete(*mp, w)
		}
	}
}

func splitAndCount(page page, re *regexp.Regexp) map[string][]int {
	words := re.Split(page.content, -1)

	mp := map[string][]int{}

	for _, w := range words {
		mp[w] = append(mp[w], page.number)
	}
	return mp
}

func separateIntoPages(lines []string, nlinesInPage int) []page {
	pages := []page{}
	p := 1
	for i := 0; i < len(lines); i += nlinesInPage {

		var content string

		if len(lines) >= i+nlinesInPage {
			content = strings.Join(lines[i:i+nlinesInPage], "\n")
		} else {
			content = strings.Join(lines[i:], " ")
		}

		pages = append(pages, page{content: content, number: p})
		p++
	}
	return pages
}

func normalize(lines []string) []string {
	normalized := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		newline := strings.Map(func(r rune) rune {
			if !unicode.IsLetter(r) {
				return ' '
			}
			return unicode.ToLower(r)
		}, line)
		normalized = append(normalized, newline)
	}
	return normalized
}
