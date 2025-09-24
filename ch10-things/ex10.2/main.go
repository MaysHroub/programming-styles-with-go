package main

import (
	"fmt"

	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.2/thing"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	dsm := thing.NewDataStorageManager(config.InputFile)
	swm := thing.NewStopWordManager(config.StopWordsFile)
	wfm := thing.NewWordFreqManager()
	controller := thing.NewWordFreqController(dsm, swm, wfm)
	controller.Run()

	fmt.Println()

	informers := []thing.Informer{&dsm, &swm, &wfm, &controller}
	for _, inf := range informers {
		fmt.Println(inf.Info())
	}
}
