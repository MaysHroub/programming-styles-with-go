package main

import (
	"sync"

	"github.com/MaysHroub/programming-styles-with-go/ch28-actors/ex28.1/actor"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	dsm := actor.NewDataStorageManager()
	swm := actor.NewStopWordManager()
	wfm := actor.NewWordFreqManager()
	controller := actor.NewWordFreqController()

	actors := []actor.Actor{dsm, swm, wfm, controller}

	var wg sync.WaitGroup

	for _, ac := range actors {
		wg.Add(1)
		go func(a actor.Actor) {
			defer wg.Done()
			a.Run()
		}(ac)
	}

	actor.Send(swm, actor.Message{"init", config.StopWordsFile, wfm})
	actor.Send(dsm, actor.Message{"init", config.InputFile, swm})
	actor.Send(controller, actor.Message{"run", dsm})

	wg.Wait()
}
