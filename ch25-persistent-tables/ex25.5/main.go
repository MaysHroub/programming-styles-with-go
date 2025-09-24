package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/MaysHroub/programming-styles-with-go/ch25-persistent-tables/internal/database"
	"github.com/MaysHroub/programming-styles-with-go/config"

	_ "github.com/mattn/go-sqlite3"
)

type page struct {
	content string
	number  int
}

func main() {
	dbDriver := "sqlite3"
	nlinesPerPage := 45
	batchSize := 1000
	limit := 25

	db, err := sql.Open(dbDriver, config.PathToDB)
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}

	lines := readData(config.InputFile)
	normalizedLines := normalize(lines)
	pages := separateIntoPages(normalizedLines, nlinesPerPage)

	docID, err := saveDocument(db, config.InputFile)
	if err != nil {
		log.Fatalf("couldn't save document with id %d: %v", docID, err)
	}
	err = savePages(db, docID, pages, batchSize)
	if err != nil {
		log.Fatalf("couldn't save words and page numbers: %v", err)
	}

	queries := database.New(db)
	wordPages, err := queries.GetWordPagesPairs(context.Background(), database.GetWordPagesPairsParams{
		DocID: docID,
		Limit: int64(limit),
	})
	if err != nil {
		log.Fatalf("couldn't query page number of each word: %v", err)
	}

	mp := make(map[string][]int)

	for _, wp := range wordPages {
		mp[wp.Word] = append(mp[wp.Word], int(wp.PageNumber))
	}

	fmt.Printf("Doc ID: %d\n", docID)
	printSorted(mp)
}

func printSorted(wordPages map[string][]int) {
	words := []string{}
	for w := range wordPages {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words[:25] {
		fmt.Printf("word: %v\npages: %v\n\n", w, wordPages[w])
	}
}

func saveDocument(db *sql.DB, docName string) (int64, error) {
	queries := database.New(db)
	doc, err := queries.AddDocument(context.Background(), docName)
	if err != nil {
		return -1, err
	}
	return doc.ID, nil
}

func savePages(db *sql.DB, docID int64, pages []page, batchSize int) error {
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := database.New(db).WithTx(tx)
	re := regexp.MustCompile(`\s+`)
	i := 0
	for _, page := range pages {
		words := re.Split(page.content, -1)
		for _, w := range words {
			wr, err := qtx.AddWord(ctx, database.AddWordParams{
				DocID: docID,
				Val:   w,
			})
			if err != nil {
				fmt.Printf("couldn't save word %s of doc id %d: %v\n", w, docID, err)
				continue
			}
			qtx.AddPage(ctx, database.AddPageParams{
				WordID: wr.ID,
				Number: int64(page.number),
			})
			i++
			if i == batchSize {
				i = 0
				err := tx.Commit()
				if err != nil {
					return err
				}
				tx, err = db.Begin()
				if err != nil {
					return err
				}
				qtx = qtx.WithTx(tx)
			}
		}
	}
	return tx.Commit()
}

func readData(filepath string) (lines []string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("couldn't open file: %v\n", err)
	}
	sc := bufio.NewScanner(file)
	lines = make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
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

func separateIntoPages(lines []string, nlinesPerPage int) []page {
	pages := []page{}
	p := 1
	for i := 0; i < len(lines); i += nlinesPerPage {

		var content string

		if len(lines) >= i+nlinesPerPage {
			content = strings.Join(lines[i:i+nlinesPerPage], "\n")
		} else {
			content = strings.Join(lines[i:], " ")
		}

		pages = append(pages, page{content: content, number: p})
		p++
	}
	return pages
}
