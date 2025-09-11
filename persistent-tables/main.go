package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/MaysHroub/programming-styles-with-go/persistent-tables/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pathToDB := "./testdb"
	filename := "../test.txt"
	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	dbQueries := database.New(db)
	loadFileIntoDatabase(filename, dbQueries)

	wordsFreq, err := dbQueries.GetWordsFreq(context.Background())
	if err != nil {
		log.Fatalf("couldn't retreive words-frequences: %v\n", err)
	}

	for _, wf := range wordsFreq {
		fmt.Printf("%s  -  %d\n", wf.Word, wf.Freq)
	}
}

func loadFileIntoDatabase(filename string, db *database.Queries) error {
	words, err := extractWords(filename)
	if err != nil {
		return err
	}

	doc, err := db.AddDocument(context.Background(), filename)
	if err != nil {
		return err
	}

	for _, w := range words {
		w = strings.TrimSpace(w)
		wr, err := db.AddWord(context.Background(), database.AddWordParams{
			Val:   w,
			DocID: doc.ID,
		})
		if err != nil {
			fmt.Printf("couldn't add word %s to database\n", w)
			continue
		}
		for _, r := range w {
			_, err := db.AddChar(context.Background(), database.AddCharParams{
				Val:    string(r),
				WordID: wr.ID,
			})
			if err != nil {
				fmt.Printf("couldn't add character %s of word %s to database\n", string(r), w)
				continue
			}
		}
	}
	fmt.Println("document is loaded!")
	return nil
}

func extractWords(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	normalized := normalizeText(string(data))
	re := regexp.MustCompile(`\s+`)
	return re.Split(normalized, -1), nil
}

func normalizeText(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return ' '
	}, text)
}
