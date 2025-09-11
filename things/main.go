package main

import "github.com/MaysHroub/programming-styles-with-go/things/controller"

func main() {
	inputFileName := "../input.txt"
	stopWordsFileName := "../stopwords.txt"
	controller := controller.NewWordFreqController(inputFileName, stopWordsFileName)
	controller.Run()
}
