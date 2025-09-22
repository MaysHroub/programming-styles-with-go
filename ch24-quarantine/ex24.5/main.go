package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode/utf8"

	"sort"
	"strings"
	"unicode"
)

const (
	FREQ_LIMIT_PER_WORD = 100
	NLINES_PER_PAGE     = 45
)

type page struct {
	content string
	number  int
}

func main() {
	qn := NewQuarantine()
	qn.bind(getInputFromConsole).
		bind(extractLines).
		bind(normalize).
		bind(separateIntoPages).
		bind(splitAndCountWords).
		bind(filterWords).
		bind(removeDuplicatedPageNums).
		bind(printSorted).
		execute()
}

func getInputFromConsole(any) (functionReturningConsoleInput any) {
	return func() any {
		var input string
		fmt.Scanf("%s", &input)
		return input
	}
}

func extractLines(filepath any) (functionReturningExtractedLines any) {
	return func() any {
		file, err := os.Open(filepath.(string))
		if err != nil {
			log.Fatalf("couldn't open file: %v\n", err)
		}
		sc := bufio.NewScanner(file)
		lines := make([]string, 0)
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		return lines
	}
}

func normalize(lines any) (normalizedLines any) {
	normalized := make([]string, 0)
	for _, line := range lines.([]string) {
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

func separateIntoPages(lines any) (pages any) {
	pages_ := []page{}
	lines_ := lines.([]string)
	p := 1
	for i := 0; i < len(lines_); i += NLINES_PER_PAGE {

		var content string

		if len(lines_) >= i+NLINES_PER_PAGE {
			content = strings.Join(lines_[i:i+NLINES_PER_PAGE], "\n")
		} else {
			content = strings.Join(lines_[i:], " ")
		}

		pages_ = append(pages_, page{content: content, number: p})
		p++
	}
	return pages_
}

func splitAndCountWords(pages any) (wordPagesMap any) {
	re := regexp.MustCompile(`\s+`)
	mp := map[string][]int{}
	for _, page := range pages.([]page) {
		words := re.Split(page.content, -1)
		for _, w := range words {
			mp[w] = append(mp[w], page.number)
		}
	}
	return mp
}

func filterWords(wordPagesMap any) (filteredWordPagesMap any) {
	filteredMp := make(map[string][]int)
	for w, nums := range wordPagesMap.(map[string][]int) {
		if len(nums) > FREQ_LIMIT_PER_WORD || utf8.RuneCountInString(w) <= 1 {
			continue
		}
		filteredMp[w] = nums
	}
	return filteredMp
}

func removeDuplicatedPageNums(wordPages any) (wordPagesWithNoDups any) {
	noDuplicateWordPages := make(map[string][]int)
	for w, nums := range wordPages.(map[string][]int) {
		mp := make(map[int]struct{})
		uniqueNums := []int{}
		for _, n := range nums {
			if _, ok := mp[n]; !ok {
				uniqueNums = append(uniqueNums, n)
				mp[n] = struct{}{}
			}
		}
		noDuplicateWordPages[w] = uniqueNums
	}
	return noDuplicateWordPages
}

func printSorted(wordPages any) (functionPrintingToStdout any) {
	return func() any {
		words := []string{}
		wordPages_ := wordPages.(map[string][]int)
		for w := range wordPages_ {
			words = append(words, w)
		}
		sort.Strings(words)
		for _, w := range words[:25] {
			fmt.Printf("word: %v\npages: %v\n\n", w, wordPages_[w])
		}
		return nil
	}
}
