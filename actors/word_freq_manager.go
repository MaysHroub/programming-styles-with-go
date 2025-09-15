package main

import (
	"sort"
)

type pair struct {
	word string
	freq int
}

type WordFreqManager struct {
	mailbox   chan Message
	wordfreqs map[string]int
}

func NewWordFreqManager() *WordFreqManager {
	return &WordFreqManager{
		mailbox:   make(chan Message),
		wordfreqs: make(map[string]int),
	}
}

func (wfm *WordFreqManager) Run() {
	for message := range wfm.mailbox {
		switch message[0].(string) {
		case "count-freq":
			wfm.incrementCount(message[1:])
		case "top25":
			wfm.generateTop25Words(message[1:])
		case "stop":
			return
		}
	}
}

func (wfm *WordFreqManager) AddToMailbox(message Message) {
	wfm.mailbox <- message
}

func (wfm *WordFreqManager) incrementCount(message Message) {
	word := message[0].(string)
	wfm.wordfreqs[word]++
}

func (wfm *WordFreqManager) generateTop25Words(message Message) {
	pairs := []pair{}
	wfc := message[0].(*WordFreqController)
	for w, f := range wfm.wordfreqs {
		pairs = append(pairs, pair{word: w, freq: f})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq >= pairs[j].freq
	})
	Send(wfc, Message{"display-top25", pairs})
}
