package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.3/thing"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	inputDataReader := thing.NewDataReader(config.InputFile)
	stopwordsDataReader := thing.NewDataReader(config.StopWordsFile)

	inputData, err := inputDataReader.Read()
	if err != nil {
		log.Fatalf("couldn't read data from file %s: %v\n", config.InputFile, err)
	}
	stopwordsData, err := stopwordsDataReader.Read()
	if err != nil {
		log.Fatalf("couldn't read data from file %s: %v\n", config.StopWordsFile, err)
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
