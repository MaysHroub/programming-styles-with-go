package main

import "github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.1/thing"

func main() {
	filepath := "../../input.txt"
	stopWordsFilepath := "../../stopwords.txt"
	controller := thing.NewWordFreqController(filepath, stopWordsFilepath)
	controller.Run()
}
