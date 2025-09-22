package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.3/thing"
)

func main() {
	filepath := "../../input.txt"
	stopwordsFilepath := "../../stopwords.txt"

	inputDataReader := thing.NewDataReader(filepath)
	stopwordsDataReader := thing.NewDataReader(stopwordsFilepath)

	inputData, err := inputDataReader.Read()
	if err != nil {
		log.Fatalf("couldn't read data from file %s: %v\n", filepath, err)
	}
	stopwordsData, err := stopwordsDataReader.Read()
	if err != nil {
		log.Fatalf("couldn't read data from file %s: %v\n", stopwordsFilepath, err)
	}

	wordProcessor := thing.NewDataProcessor(inputData)
	stopwordProcessor := thing.NewDataProcessor(stopwordsData)

	wordProcessor.NormalizeData()
	wordRegx := regexp.MustCompile(`\s+`)
	words := wordProcessor.ConvertToSlice(wordRegx)

	stopwordRegx := regexp.MustCompile(`\s*,\s*`)
	stopwords := stopwordProcessor.ConvertToMap(stopwordRegx)
	for r := 'a'; r <= 'z'; r++ {
		stopwords[string(r)] = struct{}{}
	}

	wordTokenProcessor := thing.NewTokenProcessor(words)
	wordTokenProcessor.Clean(stopwords)
	wordTokenProcessor.CountFrequencies()
	pairs := wordTokenProcessor.ToSortedPairs()

	for _, p := range pairs[:min(25, len(pairs))] {
		fmt.Printf("%s  -  %d\n", p.Token, p.Freq)
	}
}
