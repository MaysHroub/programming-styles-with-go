package actor

import (
	"fmt"
	"sort"
)

type WordIndexController struct {
	mailbox       chan Message
	dm            *DataManager
	nlinesPerPage int
}

func NewWordIndexController() *WordIndexController {
	return &WordIndexController{
		mailbox: make(chan Message),
	}
}

func (w *WordIndexController) AddToMailbox(message Message) {
	w.mailbox <- message
}

func (w *WordIndexController) Run() {
	for msg := range w.mailbox {
		funcType := msg[0].(string)
		switch funcType {
		case "init":
			w.init(msg[1:])
		case "start":
			w.start()
		case "top25":
			w.printTop25(msg[1:])
		case "stop":
			Send(w.dm, Message{"stop"})
			return
		}
	}
}

func (w *WordIndexController) init(message Message) {
	w.dm = message[0].(*DataManager)
	w.nlinesPerPage = message[1].(int)
}

func (w *WordIndexController) start() {
	Send(w.dm, Message{"process-lines", w.nlinesPerPage})
}

func (w *WordIndexController) printTop25(message Message) {
	wordPages := message[0].(map[string][]int)
	words := []string{}
	for w := range wordPages {
		words = append(words, w)
	}
	sort.Strings(words)

	for _, w := range words[:min(25, len(words))] {
		fmt.Printf("word: %v\npages: %v\n\n", w, wordPages[w])
	}

	go func() {
		Send(w, Message{"stop"})
	}()
}
