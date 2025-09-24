package main

import (
	"sync"

	"github.com/MaysHroub/programming-styles-with-go/ch28-actors/ex28.2/actor"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	dsm := actor.NewDataStorageManager()
	wfm := actor.NewWordFreqManager()
	wfc := actor.NewWordFreqController()

	actors := []actor.Actor{dsm, wfm, wfc}

	var wg sync.WaitGroup

	for _, ac := range actors {
		wg.Add(1)
		go func(a actor.Actor) {
			defer wg.Done()
			a.Run()
		}(ac)
	}

	actor.Send(dsm, actor.Message{"init-data", config.InputFile})
	actor.Send(dsm, actor.Message{"init-stopwords", config.StopWordsFile})
	actor.Send(wfc, actor.Message{"run", dsm, wfm})

	wg.Wait()
}
