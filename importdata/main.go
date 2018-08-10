package main

import (
	"bytes"
	"encoding/csv"
	"importdata/data"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./csv/k_book.csv")
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.NewReader(bytes.NewBuffer(content)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bpmap := data.ReadBookPageCSV()

	for i := 1; i < len(records); i++ {
		bmap := make(map[string]string)
		for j := 0; j < len(records[i]); j++ {
			bmap[records[0][j]] = records[i][j]
		}
		bookID, bookPages := data.ImportBook(bmap)
		data.ImportBookPage(bookID, bpmap[bmap["book_id"]], bookPages)
	}
}
