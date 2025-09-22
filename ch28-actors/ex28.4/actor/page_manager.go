package actor

import (
	"regexp"
	"unicode/utf8"
)

type page struct {
	content string
	number  int
}

type PageProcessManager struct {
	mailbox          chan Message
	wordPages        map[string][]int
	wic              *WordIndexController
	freqLimitPerWord int
}

func NewPageProcessManager() *PageProcessManager {
	return &PageProcessManager{
		mailbox:   make(chan Message),
		wordPages: make(map[string][]int),
	}
}

func (pm *PageProcessManager) AddToMailbox(message Message) {
	pm.mailbox <- message
}

func (pm *PageProcessManager) Run() {
	for msg := range pm.mailbox {
		funcType := msg[0].(string)
		switch funcType {
		case "init":
			pm.init(msg[1:])
		case "process-page":
			pm.processPage(msg[1:])
		case "clean":
			pm.filterWords()
			pm.removeDuplications()
		case "stop":
			return
		}
	}
}

func (pm *PageProcessManager) removeDuplications() {
	for w, nums := range pm.wordPages {
		mp := make(map[int]struct{})
		uniqueNums := []int{}
		for _, n := range nums {
			if _, ok := mp[n]; !ok {
				uniqueNums = append(uniqueNums, n)
				mp[n] = struct{}{}
			}
		}
		pm.wordPages[w] = uniqueNums
	}
	Send(pm.wic, Message{"top25", pm.wordPages})
}

func (pm *PageProcessManager) filterWords() {
	for w, nums := range pm.wordPages {
		if len(nums) > pm.freqLimitPerWord || utf8.RuneCountInString(w) <= 1 {
			delete(pm.wordPages, w)
		}
	}
}

func (pm *PageProcessManager) processPage(message Message) {
	page := message[0].(page)
	re := regexp.MustCompile(`\s+`)
	words := re.Split(page.content, -1)
	for _, w := range words {
		pm.wordPages[w] = append(pm.wordPages[w], page.number)
	}
}

func (pm *PageProcessManager) init(message Message) {
	pm.freqLimitPerWord = message[0].(int)
	pm.wic = message[1].(*WordIndexController)
}
