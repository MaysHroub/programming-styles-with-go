package dbio

import (
	"context"

	"github.com/MaysHroub/programming-styles-with-go/persistent-tables/internal/database"
)

func GetWordsFreq(dbQueries *database.Queries, docID, limit int64) ([]database.GetWordsFreqRow, error) {
	return dbQueries.GetWordsFreq(context.Background(), database.GetWordsFreqParams{
		DocID: docID,
		Limit: limit,
	})
}

func GetWordsFreqPerDoc(dbQueries *database.Queries, limit int64) ([]database.GetWordsFreqPerDocRow, error) {
	return dbQueries.GetWordsFreqPerDoc(context.Background(), limit)
}

func GetWordsCountPerDoc(dbQueries *database.Queries) ([]database.GetWordsCountPerDocRow, error) {
	return dbQueries.GetWordsCountPerDoc(context.Background())
}

func GetCharsCountPerDoc(dbQueries *database.Queries) ([]database.GetCharsCountPerDocRow, error) {
	return dbQueries.GetCharsCountPerDoc(context.Background())
}

func GetLongestWordsPerDoc(dbQueries *database.Queries) ([]database.GetLongestWordsPerDocRow, error) {
	return dbQueries.GetLongestWordsPerDoc(context.Background())
}

func GetCombinedLengthOfTop25WordsPerDoc(dbQueries *database.Queries) ([]database.GetCombinedLengthOfTop25WordsPerDocRow, error) {
	return dbQueries.GetCombinedLengthOfTop25WordsPerDoc(context.Background())
}

func GetAllDocIDs(dbQueries *database.Queries) ([]int64, error) {
	return dbQueries.GetAllDocIDs(context.Background())
}
