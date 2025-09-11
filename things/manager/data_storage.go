package manager

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

type DataStorageManager struct {
	data string
}

func NewDataStorageManager(filename string) DataStorageManager {
	data, err := os.ReadFile(filename)
	if err != nil {
		return DataStorageManager{}
	}

	normalizedText := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, string(data))

	return DataStorageManager{data: strings.ToLower(normalizedText)}
}

func (d DataStorageManager) Words() []string {
	re := regexp.MustCompile(`\s+`)
	return re.Split(d.data, -1)
}
