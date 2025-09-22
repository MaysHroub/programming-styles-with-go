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
)

type page struct {
	content string
	number  int
}

func main() {
	filepath := "../../input.txt"
	nlinesPerPage := 45
	freqLimitPerWord := 100
	printSorted(
		removeDuplicatedPageNums(
			filterWords(freqLimitPerWord)(
				splitAndCountWords(
					separateIntoPages(nlinesPerPage)(
						normalize(
							readData(filepath),
						),
					),
				),
			),
		),
	)
}

func readData(filepath string) (lines []string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("couldn't open file: %v\n", err)
	}
	sc := bufio.NewScanner(file)
	lines = make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
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

func separateIntoPages(nlinesPerPage int) func([]string) []page {
	return func(lines []string) []page {
		pages := []page{}
		p := 1
		for i := 0; i < len(lines); i += nlinesPerPage {

			var content string

			if len(lines) >= i+nlinesPerPage {
				content = strings.Join(lines[i:i+nlinesPerPage], "\n")
			} else {
				content = strings.Join(lines[i:], " ")
			}

			pages = append(pages, page{content: content, number: p})
			p++
		}
		return pages
	}
}

func splitAndCountWords(pages []page) map[string][]int {
	re := regexp.MustCompile(`\s+`)
	mp := map[string][]int{}
	for _, page := range pages {
		words := re.Split(page.content, -1)
		for _, w := range words {
			mp[w] = append(mp[w], page.number)
		}
	}
	return mp
}

func filterWords(freqLimitPerWord int) func(map[string][]int) map[string][]int {
	return func(wordPages map[string][]int) map[string][]int {
		filteredMp := make(map[string][]int)
		for w, nums := range wordPages {
			if len(nums) > freqLimitPerWord || utf8.RuneCountInString(w) <= 1 {
				continue
			}
			filteredMp[w] = nums
		}
		return filteredMp
	}
}

func removeDuplicatedPageNums(wordPages map[string][]int) map[string][]int {
	noDuplicateWordPages := make(map[string][]int)
	for w, nums := range wordPages {
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

func printSorted(wordPages map[string][]int) {
	words := []string{}
	for w := range wordPages {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words[:25] {
		fmt.Printf("word: %v\npages: %v\n\n", w, wordPages[w])
	}
}
