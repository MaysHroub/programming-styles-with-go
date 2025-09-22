package dbio

import (
	"context"
	"database/sql"

	"github.com/MaysHroub/programming-styles-with-go/persistent-tables/internal/database"
)

func GetWordsFreq(db *sql.DB, docID, limit int64) ([]database.GetWordsFreqRow, error) {
	dbQueries := database.New(db)
	return dbQueries.GetWordsFreq(context.Background(), database.GetWordsFreqParams{
		DocID: docID,
		Limit: limit,
	})
}
