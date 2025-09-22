package thing

import (
	"fmt"
)

type WordFrequencyController struct {
	dataStorageMgr DataStorageManager
	stopWordsMgr   StopWordsManager
	wordFreqMgr    WordFreqManager
}

func NewWordFreqController(inputFileName, stopWordsFileName string) WordFrequencyController {
	return WordFrequencyController{
		dataStorageMgr: NewDataStorageManager(inputFileName),
		stopWordsMgr:   NewStopWordManager(stopWordsFileName),
		wordFreqMgr:    NewWordFreqManager(),
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
