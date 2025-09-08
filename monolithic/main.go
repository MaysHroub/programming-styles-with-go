package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// declare our global variables
type pair struct {
	word string
	freq int
}

var wordsFreq = make([]pair, 0)
var stopWords = make(map[string]struct{})

func main() {
	// retrieve all stop words
	stopWordsFileContent, err := os.ReadFile("../stopwords.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return
	}
	stopWordsAsSlice := strings.Split(string(stopWordsFileContent), ",")
	for _, spword := range stopWordsAsSlice {
		stopWords[spword] = struct{}{}
	}

	// open the file
	inputFile, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		line += "\n"
		wordStartIdx := -1
		for i, c := range line {
			// find the start of the word
			if wordStartIdx == -1 {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					wordStartIdx = i
				}
				continue
			}

			// if it's not the end of the word, just keep looping
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				continue
			}

			// if we reach the end of the word, we process it
			word := strings.ToLower(line[wordStartIdx:i])

			// if it's just one letter, ignore it
			if len(word) == 1 {
				wordStartIdx = -1 // reset
				continue
			}

			// if it's a stop word, ignore it
			_, exists := stopWords[word]
			if exists {
				wordStartIdx = -1 // reset
				continue
			}

			// if the word doesn't exist, add it to the map
			// otherwise, increment its frequency then reorder
			pairIdx := 0
			for ; pairIdx < len(wordsFreq); pairIdx++ {
				if word == wordsFreq[pairIdx].word {
					wordsFreq[pairIdx].freq++
					break
				}
			}

			// the word doesn't exist
			if pairIdx == len(wordsFreq) {
				wordsFreq = append(wordsFreq, pair{word: word, freq: 1})
				wordStartIdx = -1 // reset
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

	for _, wf := range wordsFreq {
		fmt.Printf("%v  -  %v\n", wf.word, wf.freq)
	}

}
