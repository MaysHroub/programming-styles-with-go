package main

import "github.com/MaysHroub/programming-styles-with-go/things/ex10.1/thing"

func main() {
	inputFileName := "../input.txt"
	stopWordsFileName := "../stopwords.txt"
	controller := thing.NewWordFreqController(inputFileName, stopWordsFileName)
	controller.Run()
}
