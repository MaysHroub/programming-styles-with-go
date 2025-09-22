package thing

import (
	"fmt"
)

type WordFrequencyController struct {
	dataStorageMgr DataStorageManager
	stopWordsMgr   StopWordsManager
	wordFreqMgr    WordFreqManager
}

func NewWordFreqController(
	dataStorageMgr DataStorageManager, 
	stopWordsMgr StopWordsManager, 
	wordFreqMgr WordFreqManager,
) WordFrequencyController {
	return WordFrequencyController{
		dataStorageMgr: dataStorageMgr,
		stopWordsMgr:   stopWordsMgr,
		wordFreqMgr:    wordFreqMgr,
	}
}

func (wc *WordFrequencyController) Run() {
	words := wc.dataStorageMgr.Words()
	for _, w := range words {
		if wc.stopWordsMgr.IsStopWord(w) {
			continue
		}
		wc.wordFreqMgr.IncrementCount(w)
	}

	sortedPairs := wc.wordFreqMgr.ToSortedPairs()

	for _, p := range sortedPairs[:min(25, len(sortedPairs))] {
		fmt.Printf("%s\n", p.ToString())
	}
}

func (wc *WordFrequencyController) Info() string {
	return "Name: WordFrequencyController\nJob: Run the program by grouping other manager structs and executing their methods\n"
}
