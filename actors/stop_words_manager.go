package main

import (
	"log"
	"os"
	"strings"
)

type StopWordManager struct {
	mailbox   chan Message
	stopwords map[string]struct{}
	wfm       *WordFreqManager
}

func NewStopWordManager() *StopWordManager {
	return &StopWordManager{
		mailbox:   make(chan Message),
		stopwords: make(map[string]struct{}),
	}
}

func (swm *StopWordManager) Run() {
	for message := range swm.mailbox {
		switch message[0].(string) {
		case "init":
			err := swm.init(message[1:])
			if err != nil {
				log.Fatalf("couldn't init stop-word-manager: %v\n", err)
			}
		case "filter":
			swm.filterWord(message[1:])
		case "stop":
			Send(swm.wfm, Message{"stop"})
			return
		default: // forward
			Send(swm.wfm, message)
		}
	}
}

func (swm *StopWordManager) AddToMailbox(message Message) {
	swm.mailbox <- message
}

func (swm *StopWordManager) filterWord(message Message) {
	word := message[0].(string)
	if _, ok := swm.stopwords[word]; !ok {
		Send(swm.wfm, Message{"count-freq", word})
	}
}

func (swm *StopWordManager) init(message Message) error {
	filepath := message[0].(string)
	swm.wfm = message[1].(*WordFreqManager)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	spwords := strings.SplitSeq(string(data), ",")
	for sw := range spwords {
		swm.stopwords[sw] = struct{}{}
	}
	for r := 'a'; r <= 'z'; r++ {
		swm.stopwords[string(r)] = struct{}{}
	}
	return nil
}
