package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/ex25.4/dbio"
	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pathToDB := "../sql/schema/testdb.db"
	filepath := "../../files/input.txt"
	stopwordsfile := "../../files/stopwords.txt"
	batchSize := 1000
	limit := 25

	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}

	_, err = dbio.LoadFileIntoDatabase(filepath, db, batchSize)
	if err != nil {
		log.Fatalf("couldn't load file into database: %v", err)
	}
	err = dbio.LoadStopwordsIntoDatabase(stopwordsfile, db)
	if err != nil {
		log.Fatalf("couldn't save stopwords in database: %v", err)
	}

	dbQueries := database.New(db)

	wordsCount, err := dbio.GetWordsCountPerDoc(dbQueries)
	if err != nil {
		log.Fatalf("couldn't retrieve words count: %v", err)
	}
	charsCount, err := dbio.GetCharsCountPerDoc(dbQueries)
	if err != nil {
		log.Fatalf("couldn't retrieve chars count: %v", err)
	}
	longestWords, err := dbio.GetLongestWordsPerDoc(dbQueries)
	if err != nil {
		log.Fatalf("couldn't retrieve longest word(s): %v", err)
	}
	combinedWordLength, err := dbio.GetCombinedLengthOfTop25WordsPerDoc(dbQueries)
	if err != nil {
		log.Fatalf("couldn't retrieve combined words length: %v", err)
	}
	docIDs, err := dbio.GetAllDocIDs(dbQueries)
	if err != nil {
		log.Fatalf("couldn't retrieve doc IDs: %v", err)
	}

	fmt.Println("Top 25 - most frequent words per doc")
	for _, did := range docIDs {
		fmt.Printf("DocID: %d\n", did)
		wordsFreq, err := dbio.GetWordsFreq(dbQueries, did, int64(limit))
		if err != nil {
			log.Fatalf("couldn't retrieve words frequences for doc id %d: %v", did, err)
		}
		for _, wf := range wordsFreq {
			fmt.Printf("Word: %s  -  Freq: %d\n", wf.Word, wf.Freq)
		}
		fmt.Println()
	}

	fmt.Println("-------------------------------\nWords count per document:")
	for _, wc := range wordsCount {
		fmt.Printf("DocID: %d  -  %d\n", wc.DocID, wc.WordsCount)
	}

	fmt.Println("-------------------------------\nChars count per document:")
	for _, cc := range charsCount {
		fmt.Printf("DocID: %d  -  %v\n", cc.DocID, cc.CharsCount.Float64)
	}

	fmt.Println("-------------------------------\nLongest word(s) per document:")
	for _, w := range longestWords {
		fmt.Printf("DocID: %d  -  %s\n", w.DocID, w.LongestWord)
	}

	fmt.Println("-------------------------------\nCombined length of top 25 words per document:")
	for _, cl := range combinedWordLength {
		fmt.Printf("DocID: %d  -  %v\n", cl.DocID, cl.CombinedLength.Float64)
	}
}
