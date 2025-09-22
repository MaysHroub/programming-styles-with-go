package thing

import (
	"os"
	"strings"
)

type StopWordsManager struct {
	stopWords map[string]struct{}
}

func NewStopWordManager(filepath string) StopWordsManager {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return StopWordsManager{}
	}

	stopWordsAsSlice := strings.Split(string(data), ",")

	stopWordsMap := make(map[string]struct{})

	for _, sp := range stopWordsAsSlice {
		stopWordsMap[strings.ToLower(sp)] = struct{}{}
	}
	for r := 'a'; r <= 'z'; r++ {
		stopWordsMap[string(r)] = struct{}{}
	}

	return StopWordsManager{stopWords: stopWordsMap}
}

func (s *StopWordsManager) IsStopWord(word string) bool {
	_, exists := s.stopWords[strings.ToLower(word)]
	return exists
}

func (s *StopWordsManager) Info() string {
	return "Name: StopWordsManager\nJob: Rertieve stop words from specified file and checks if a word is a stop word or not\n"
}
