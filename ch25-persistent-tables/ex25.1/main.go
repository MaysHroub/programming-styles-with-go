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

	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pathToDB := "../sql/schema/testdb.db"
	filepath := "../../files/input.txt"
	stopwordsfile := "../../files/stopwords.txt"
	batchSize := 1000

	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	db.SetMaxOpenConns(1)

	docID, err := loadFileIntoDatabase(filepath, db, batchSize)
	if err != nil {
		log.Fatalf("couldn't load file into database: %v", err)
	}

	err = loadStopwordsIntoDatabase(stopwordsfile, db)
	if err != nil {
		log.Fatalf("couldn't save stopwords in database: %v", err)
	}

	dbQueries := database.New(db)
	limit := 25
	wordsFreq, err := dbQueries.GetWordsFreq(context.Background(), database.GetWordsFreqParams{
		DocID: docID,
		Limit: int64(limit),
	})
	if err != nil {
		log.Fatalf("couldn't retreive words-frequences: %v\n", err)
	}

	for _, wf := range wordsFreq {
		fmt.Printf("%s  -  %d\n", wf.Word, wf.Freq)
	}
}

func loadStopwordsIntoDatabase(stopwordsfile string, db *sql.DB) error {
	data, err := os.ReadFile(stopwordsfile)
	if err != nil {
		return err
	}
	stopwords := strings.Split(string(data), ",")

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := database.New(db).WithTx(tx)

	for _, w := range stopwords {
		err = qtx.AddStopWord(context.Background(), w)
		if err != nil {
			fmt.Printf("couldn't add stopword %s to database", w)
		}
	}
	for r := 'a'; r <= 'z'; r++ {
		err = qtx.AddStopWord(context.Background(), string(r))
		if err != nil {
			fmt.Printf("couldn't add character %s to database", string(r))
		}
	}
	return tx.Commit()
}

func loadFileIntoDatabase(filename string, db *sql.DB, batchSize int) (int64, error) {
	words, err := extractWords(filename)
	if err != nil {
		return -1, err
	}

	queries := database.New(db)

	doc, err := queries.AddDocument(context.Background(), filename)
	if err != nil {
		return -1, err
	}

	for i := 0; i < len(words); i += batchSize {
		tx, err := db.Begin()
		if err != nil {
			return -1, err
		}
		defer tx.Rollback()

		qtx := queries.WithTx(tx)

		for j := 0; j < min(batchSize, len(words)-i); j++ {
			w := strings.TrimSpace(words[j+i])
			_, err := qtx.AddWord(context.Background(), database.AddWordParams{
				Val:   w,
				DocID: doc.ID,
			})
			if err != nil {
				fmt.Printf("couldn't add word %s to database\n", w)
				continue
			}
		}
		err = tx.Commit()
		if err != nil {
			return -1, err
		}
	}

	return doc.ID, nil
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
