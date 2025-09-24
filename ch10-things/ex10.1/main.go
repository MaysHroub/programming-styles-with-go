package main

import (
	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.1/thing"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	controller := thing.NewWordFreqController(config.InputFile, config.StopWordsFile)
	controller.Run()
}
