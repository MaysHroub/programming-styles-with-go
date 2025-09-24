package actor

import (
	"fmt"
)

type WordFreqController struct {
	mailbox chan Message
	dsm     *DataStorageManager
}

func NewWordFreqController() *WordFreqController {
	return &WordFreqController{
		mailbox: make(chan Message),
	}
}

func (wfc *WordFreqController) Run() {
	for message := range wfc.mailbox {
		switch message[0].(string) {
		case "display-top25":
			wfc.displayPairs(message[1:])
		case "run":
			wfc.startExecuting(message[1:])
		case "stop":
			Send(wfc.dsm, Message{"stop"})
			return
		}
	}
}

func (wfc *WordFreqController) AddToMailbox(message Message) {
	wfc.mailbox <- message
}

func (wfc *WordFreqController) displayPairs(message Message) {
	pairs := message[0].([]pair)
	for _, pair := range pairs[:min(25, len(pairs))] {
		fmt.Printf("%s  --  %d\n", pair.word, pair.freq)
	}
	go func() {
		Send(wfc, Message{"stop"})
	}()
}

func (wfc *WordFreqController) startExecuting(message Message) {
	wfc.dsm = message[0].(*DataStorageManager)
	Send(wfc.dsm, Message{"process-words", wfc})
}
