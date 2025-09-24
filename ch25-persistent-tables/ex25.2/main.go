package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/ex25.2/dbio"
	"github.com/MaysHroub/programming-styles-with-go/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	batchSize := 1000

	db, err := sql.Open("sqlite3", config.PathToDB)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}

	docID, err := dbio.LoadFileIntoDatabase(config.InputFile, db, batchSize)
	if err != nil {
		log.Fatalf("couldn't load file into database: %v", err)
	}
	err = dbio.LoadStopwordsIntoDatabase(config.StopWordsFile, db)
	if err != nil {
		log.Fatalf("couldn't save stopwords in database: %v", err)
	}

	limit := 25
	wordsFreq, err := dbio.GetWordsFreq(db, docID, int64(limit))
	if err != nil {
		log.Fatalf("couldn't retreive words-frequences: %v\n", err)
	}

	for _, wf := range wordsFreq {
		fmt.Printf("%s  -  %d\n", wf.Word, wf.Freq)
	}
}
