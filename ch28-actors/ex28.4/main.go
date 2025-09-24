package main

import (
	"sync"

	"github.com/MaysHroub/programming-styles-with-go/ch28-actors/ex28.4/actor"
)

func main() {
	inputfilepath := "../../files/input.txt"
	nlinesPerPage := 45
	freqLimitPerWord := 100

	dm := actor.NewDataManager()
	pm := actor.NewPageProcessManager()
	wic := actor.NewWordIndexController()

	actors := []actor.Actor{dm, pm, wic}

	var wg sync.WaitGroup

	for _, ac := range actors {
		wg.Add(1)
		go func(a actor.Actor) {
			defer wg.Done()
			a.Run()
		}(ac)
	}

	actor.Send(dm, actor.Message{"init", inputfilepath, pm})
	actor.Send(pm, actor.Message{"init", freqLimitPerWord, wic})
	actor.Send(wic, actor.Message{"init", dm, nlinesPerPage})

	actor.Send(wic, actor.Message{"start"})

	wg.Wait()
}
