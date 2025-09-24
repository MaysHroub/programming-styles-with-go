package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/ex25.2/dbio"
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

	docID, err := dbio.LoadFileIntoDatabase(filepath, db, batchSize)
	if err != nil {
		log.Fatalf("couldn't load file into database: %v", err)
	}
	err = dbio.LoadStopwordsIntoDatabase(stopwordsfile, db)
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
