package main

import (
	"sync"

	"github.com/MaysHroub/programming-styles-with-go/actors/ex28.4/actor"
)

const (
	INPUT_FILEPATH = "../../input.txt"
	NLINE_PER_PAGE = 45
	FREQ_LIMIT_PER_WORD = 100
)

func main() {
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

	actor.Send(dm, actor.Message{"init", INPUT_FILEPATH, pm})
	actor.Send(pm, actor.Message{"init", FREQ_LIMIT_PER_WORD, wic})
	actor.Send(wic, actor.Message{"init", dm, NLINE_PER_PAGE})

	actor.Send(wic, actor.Message{"start"})

	wg.Wait()
}