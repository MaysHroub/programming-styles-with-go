package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type DataStorageManager struct {
	mailbox chan Message
	words   []string
	swm     StopWordManager
}

func NewDataStorageManager() DataStorageManager {
	return DataStorageManager{
		mailbox: make(chan Message),
	}
}

func (dsm DataStorageManager) Run() {
	defer close(dsm.mailbox)
	for message := range dsm.mailbox {
		switch message[0].(string) {
		case "init":
			err := dsm.init(message[1:])
			if err != nil {
				log.Fatalf("couldn't process file content: %v\n", err)
			}
		case "process-words":
			dsm.processWords(message[1:])
		case "stop":
			fmt.Println("stop from dsm")
			// Send(dsm.swm, Message{"stop"})
			return
		}
	}
}

func (dsm DataStorageManager) AddToMailbox(message Message) {
	dsm.mailbox <- message
}

func (dsm DataStorageManager) processWords(message Message) {
	wfc := message[0].(WordFreqController)
	for _, w := range dsm.words {
		Send(dsm.swm, Message{"filter", w})
	}
	Send(dsm.swm, Message{"top25", wfc})
	go func() {
		Send(dsm, Message{"stop"})
	}()
}

func (dsm *DataStorageManager) init(message Message) error {
	filepath := message[0].(string)
	dsm.swm = message[1].(StopWordManager)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	normalizedText := normalizeText(data)
	dsm.words = splitText(normalizedText)
	return nil
}

func normalizeText(data []byte) string {
	str := strings.Map(func(r rune) rune {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			return ' '
		}
		return r
	}, string(data))
	return strings.ToLower(str)
}

func splitText(text string) []string {
	re := regexp.MustCompile(`\s+`)
	return re.Split(text, -1)
}
