package data

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	bookPageCSV    = "./csv/k_book_page.csv"
	addBookPageURL = "http://localhost:8080/bookPage/add"
)

// ReadBookPageCSV read file to csv, return map
func ReadBookPageCSV() map[string][]map[string]string {
	content, err := ioutil.ReadFile(bookPageCSV)
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.NewReader(bytes.NewBuffer(content)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bpmap := make(map[string][]map[string]string)
	for i := 1; i < len(records); i++ {
		bprmap := make(map[string]string)
		for j := 0; j < len(records[i]); j++ {
			bprmap[records[0][j]] = records[i][j]
		}
		bpmap[bprmap["book_id"]] = append(bpmap[bprmap["book_id"]], bprmap)
	}
	return bpmap
}

// ImportBookPage import bookPage data
func ImportBookPage(id string, bookPages []map[string]string, oBookPages interface{}) {
	obpmap := make(map[string]interface{})
	for _, v := range oBookPages.([]interface{}) {
		obpmap[fmt.Sprint(v.(map[string]interface{})["pageIndex"])] = v
	}

	for _, v := range bookPages {
		obp := obpmap[v["page_number"]].(map[string]interface{})

		bpmap := make(map[string]interface{})
		bpmap["bookId"] = id
		if v["background_url"] != "" {
			bpmap["backgroundUrl"] = v["background_url"]
		} else {
			bpmap["backgroundUrl"] = obp["imgUrl"]
		}
		if v["sound_url"] != "" {
			bpmap["soundUrl"] = v["sound_url"]
		} else {
			delete(obp, "text")
			delete(obp, "imgUrl")
			delete(obp, "pageSoundUrl")
			delete(obp, "pageIndex")
			tmp, err := json.Marshal(obp)
			if err != nil {
				log.Fatal(err)
			}
			bpmap["soundUrl"] = bytes.NewBuffer(tmp).String()
		}
		bpmap["pageNumber"] = v["page_number"]

		bpjson, err := json.Marshal(bpmap)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(addBookPageURL, "application/json", bytes.NewBuffer(bpjson))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	}
}
