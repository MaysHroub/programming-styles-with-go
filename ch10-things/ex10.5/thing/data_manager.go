package thing

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type DataManager struct {
	nlinesPerPage int
	filepath      string
	lines         []string
}

func NewDataManager(filepath string, nlinesPerPage int) *DataManager {
	return &DataManager{
		nlinesPerPage: nlinesPerPage,
		filepath: filepath,
	}
}

func (dm *DataManager) ExtractLines() {
	file, err := os.Open(dm.filepath)
	if err != nil {
		log.Fatalf("couldn't open file: %v\n", err)
	}
	sc := bufio.NewScanner(file)
	dm.lines = make([]string, 0)
	for sc.Scan() {
		dm.lines = append(dm.lines, sc.Text())
	}
}

func (dm *DataManager) Normalize() {
	normalized := make([]string, 0)
	for _, line := range dm.lines {
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
	dm.lines = normalized
}

func (dm *DataManager) SeparateIntoPages() []page {
	pages := []page{}
	p := 1
	for i := 0; i < len(dm.lines); i += dm.nlinesPerPage {

		var content string

		if len(dm.lines) >= i+dm.nlinesPerPage {
			content = strings.Join(dm.lines[i:i+dm.nlinesPerPage], "\n")
		} else {
			content = strings.Join(dm.lines[i:], " ")
		}

		pages = append(pages, page{content: content, number: p})
		p++
	}
	return pages
}
