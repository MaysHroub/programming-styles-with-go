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
	for r := 'a'; r <= 'z'; r++ {
		stopWords[string(r)] = struct{}{}
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
			} else {
				if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
					word := strings.ToLower(line[wordStartIdx:i])
					_, exists := stopWords[word]
					if !exists {
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
						if !found {
							wordsFreq = append(wordsFreq, pair{word: word, freq: 1})
						} else {
							// reorder (word with most frequency first)
							for i := pairIdx; i > 0; i-- {
								if wordsFreq[i].freq > wordsFreq[i-1].freq {
									wordsFreq[i], wordsFreq[i-1] =
										wordsFreq[i-1], wordsFreq[i]
								}
							}
						}
					}
					wordStartIdx = -1 //reset
				}
			}
		}
	}

	for i, wf := range wordsFreq {
		fmt.Printf("%v  -  %v\n", wf.word, wf.freq)
		if i == 24 {
			break
		}
	}

}
