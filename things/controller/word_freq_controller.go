package controller

import (
	"fmt"

	"github.com/MaysHroub/programming-styles-with-go/things/manager"
)

type WordFrequencyController struct {
	dataStorageMgr manager.DataStorageManager
	stopWordsMgr   manager.StopWordsManager
	wordFreqMgr    manager.WordFreqManager
}

func NewWordFreqController(inputFileName, stopWordsFileName string) WordFrequencyController {
	return WordFrequencyController{
		dataStorageMgr: manager.NewDataStorageManager(inputFileName),
		stopWordsMgr:   manager.NewStopWordManager(stopWordsFileName),
		wordFreqMgr:    manager.NewWordFreqManager(),
	}
}

func (wc WordFrequencyController) Run() {
	words := wc.dataStorageMgr.Words()
	for _, w := range words {
		if wc.stopWordsMgr.IsStopWord(w) {
			continue
		}
		wc.wordFreqMgr.IncrementCount(w)
	}

	sortedPairs := wc.wordFreqMgr.ToSortedPairs()

	for _, p := range sortedPairs[:25] {
		fmt.Printf("%s\n", p.ToString())
	}
}
