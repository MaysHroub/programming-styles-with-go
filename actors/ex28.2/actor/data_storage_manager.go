package actor

import (
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type DataStorageManager struct {
	mailbox   chan Message
	words     []string
	stopwords map[string]struct{}
}

func NewDataStorageManager() *DataStorageManager {
	return &DataStorageManager{
		mailbox:   make(chan Message),
		stopwords: make(map[string]struct{}),
	}
}

func (dsm *DataStorageManager) Run() {
	for message := range dsm.mailbox {
		switch message[0].(string) {
		case "init-data":
			err := dsm.initData(message[1:])
			if err != nil {
				log.Fatalf("couldn't process file content: %v\n", err)
			}
		case "init-stopwords":
			err := dsm.initStopWords(message[1:])
			if err != nil {
				log.Fatalf("couldn't process file content: %v\n", err)
			}
		case "process-words":
			dsm.processWords(message[1:])
		case "stop":
			return
		}
	}
}

func (dsm *DataStorageManager) AddToMailbox(message Message) {
	dsm.mailbox <- message
}

func (dsm *DataStorageManager) processWords(message Message) {
	wfm := message[0].(*WordFreqManager)
	wfc := message[1].(*WordFreqController)
	for _, w := range dsm.words {
		if _, ok := dsm.stopwords[w]; !ok {
			Send(wfm, Message{"count", w})
		}
	}
	Send(wfm, Message{"top25", wfc})
}

func (dsm *DataStorageManager) initData(message Message) error {
	filepath := message[0].(string)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	normalizedText := normalizeText(data)
	dsm.words = splitText(normalizedText)
	return nil
}

func (dsm *DataStorageManager) initStopWords(message Message) error {
	filepath := message[0].(string)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	spwords := strings.SplitSeq(string(data), ",")
	for sw := range spwords {
		dsm.stopwords[sw] = struct{}{}
	}
	for r := 'a'; r <= 'z'; r++ {
		dsm.stopwords[string(r)] = struct{}{}
	}
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
