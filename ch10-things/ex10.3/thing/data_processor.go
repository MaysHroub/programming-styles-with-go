package thing

import (
	"regexp"
	"strings"
	"unicode"
)

type DataProcessor struct {
	data string
}

func NewDataProcessor(data string) DataProcessor {
	return DataProcessor{
		data: data,
	}
}

func (d *DataProcessor) NormalizeData() {
	d.data = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, d.data)
}

func (d DataProcessor) ConvertToSlice(re *regexp.Regexp) []string {
	return re.Split(d.data, -1)
}

func (d DataProcessor) ConvertToMap(re *regexp.Regexp) map[string]struct{} {
	slice := d.ConvertToSlice(re)
	mp := make(map[string]struct{})
	for _, sp := range slice {
		mp[strings.TrimSpace(sp)] = struct{}{}
	}
	return mp
}
