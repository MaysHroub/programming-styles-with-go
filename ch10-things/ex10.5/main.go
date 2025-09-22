package main

import "github.com/MaysHroub/programming-styles-with-go/things/ex10.5/thing"

const (
	FREQ_LIMIT_PER_WORD = 100
	INPUT_FILENAME      = "../../input.txt"
	NLINES_PER_PAGE     = 45
)

func main() {
	dm := thing.NewDataManager(INPUT_FILENAME, NLINES_PER_PAGE)
	dm.ExtractLines()
	dm.Normalize()
	pages := dm.SeparateIntoPages()

	p := thing.NewPageProcessor(pages, FREQ_LIMIT_PER_WORD)
	p.SplitAndCountWords()
	p.FilterWords()
	p.RemoveDuplicatedPageNums()

	wm := thing.NewWordIndexManager(p.GetWordPagesMap())
	wm.PrintSorted()
}
