package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/MaysHroub/programming-styles-with-go/config"
)

// declare our global variables
type pair struct {
	word string
	freq int
}

var (
	wordsFreq = make([]pair, 0)
	stopWords = make(map[string]struct{})
)

func main() {
	// retrieve all stop words (with single letters)
	stopWordsFileContent, err := os.ReadFile(config.StopWordsFile)
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return
	}
	stopWordsAsSlice := strings.Split(string(stopWordsFileContent), ",")
	for _, spword := range stopWordsAsSlice {
		stopWords[spword] = struct{}{}
	}
	for r := 'a'; r <= 'z'; r++ {
		stopWords[string(r)] = struct{}{}
	}

	// open the file and read it line by line
	inputFile, err := os.Open(config.InputFile)
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		line += "\n" // so it doesn't ignore the last word
		wordStartIdx := -1
		for i, c := range line {

			// find the start of the word
			if wordStartIdx == -1 {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					wordStartIdx = i
				}
				continue
			}

			// continue looping if it's not the end of the word
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				continue
			}

			// retrieve the word
			word := strings.ToLower(line[wordStartIdx:i])

			// if it's a stop word, ignore it and reset the start index
			_, exists := stopWords[word]
			if exists {
				wordStartIdx = -1
				continue
			}

			// look for the word and update its frequency if it's found
			pairIdx := 0
			found := false
			for i, wf := range wordsFreq {
				if word == wf.word {
					wordsFreq[i].freq++
					found = true
					break
				}
				pairIdx++
			}

			// if it's not found, just append it and reset the start index
			if !found {
				wordsFreq = append(wordsFreq, pair{word: word, freq: 1})
				wordStartIdx = -1
				continue
			}

			// reorder (word with most frequency first)
			for i := pairIdx; i > 0; i-- {
				if wordsFreq[i].freq > wordsFreq[i-1].freq {
					wordsFreq[i], wordsFreq[i-1] =
						wordsFreq[i-1], wordsFreq[i]
				}
			}

			wordStartIdx = -1 //reset
		}
	}

	for i, wf := range wordsFreq {
		fmt.Printf("%v  -  %v\n", wf.word, wf.freq)
		if i == 24 {
			break
		}
	}

}
