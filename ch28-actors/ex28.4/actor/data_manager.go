package actor

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type DataManager struct {
	mailbox chan Message
	lines   []string
	pm      *PageProcessManager
}

func NewDataManager() *DataManager {
	return &DataManager{
		mailbox: make(chan Message),
	}
}

func (dm *DataManager) AddToMailbox(message Message) {
	dm.mailbox <- message
}

func (dm *DataManager) Run() {
	for msg := range dm.mailbox {
		funcType := msg[0].(string)
		switch funcType {
		case "init":
			dm.init(msg[1:])
		case "process-lines":
			dm.processLinesToPages(msg[1:])
		case "stop":
			Send(dm.pm, Message{"stop"})
			return
		}
	}
}

func (dm *DataManager) processLinesToPages(message Message) {
	nlinesPerPage := message[0].(int)
	p := 1
	for i := 0; i < len(dm.lines); i += nlinesPerPage {
		var content string
		if len(dm.lines) >= i+nlinesPerPage {
			content = strings.Join(dm.lines[i:i+nlinesPerPage], "\n")
		} else {
			content = strings.Join(dm.lines[i:], " ")
		}
		page := page{content: content, number: p}
		Send(dm.pm, Message{"process-page", page})
		p++
	}
	Send(dm.pm, Message{"clean"})
}

func (dm *DataManager) init(message Message) {
	filepath := message[0].(string)
	dm.pm = message[1].(*PageProcessManager)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("couldn't open file: %v\n", err)
	}
	sc := bufio.NewScanner(file)
	dm.lines = make([]string, 0)
	for sc.Scan() {
		dm.lines = append(dm.lines, sc.Text())
	}
	dm.lines = normalize(dm.lines)
}

func normalize(lines []string) []string {
	normalized := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		newline := strings.Map(func(r rune) rune {
			if !unicode.IsLetter(r) {
				return ' '
			}
			return unicode.ToLower(r)
		}, line)
		normalized = append(normalized, newline)
	}
	return normalized
}
