package main

import (
	"github.com/MaysHroub/programming-styles-with-go/ch10-things/ex10.5/thing"
	"github.com/MaysHroub/programming-styles-with-go/config"
)

func main() {
	nlinesPerPage := 45
	freqLimitPerWord := 100

	dm := thing.NewDataManager(config.InputFile, nlinesPerPage)
	dm.ExtractLines()
	dm.Normalize()
	pages := dm.SeparateIntoPages()

	p := thing.NewPageProcessor(pages, freqLimitPerWord)
	p.SplitAndCountWords()
	p.FilterWords()
	p.RemoveDuplicatedPageNums()

	wm := thing.NewWordIndexManager(p.GetWordPagesMap())
	wm.PrintSorted()
}
