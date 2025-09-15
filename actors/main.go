package main

import "sync"

func main() {

	inputfilepath := "../input.txt"
	stopwordfilepath := "../stopwords.txt"

	dsm := NewDataStorageManager()
	swm := NewStopWordManager()
	wfm := NewWordFreqManager()
	controller := NewWordFreqController()

	actors := []Actor{dsm, swm, wfm, controller}

	var wg sync.WaitGroup

	for _, ac := range actors {
		wg.Add(1)
		go func (a Actor)  {
			defer wg.Done()
			ac.Run()
		}(ac)
	}

	Send(swm, Message{"init", stopwordfilepath, wfm})
	Send(dsm, Message{"init", inputfilepath, swm})
	Send(controller, Message{"run", dsm})

	wg.Wait()
}
