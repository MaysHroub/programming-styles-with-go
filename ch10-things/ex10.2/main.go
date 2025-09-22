package main

import (
	"fmt"

	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.2/thing"
)

func main() {
	filepath := "../../input.txt"
	stopWordsFilepath := "../../stopwords.txt"
	dsm := thing.NewDataStorageManager(filepath)
	swm := thing.NewStopWordManager(stopWordsFilepath)
	wfm := thing.NewWordFreqManager()
	controller := thing.NewWordFreqController(dsm, swm, wfm)
	controller.Run()

	fmt.Println()

	informers := []thing.Informer{&dsm, &swm, &wfm, &controller}
	for _, inf := range informers {
		fmt.Println(inf.Info())
	}
}
