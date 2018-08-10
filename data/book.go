package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	getBookInfoURL = "http://10.1.10.63:18080/kindergarten/getBookInfo.json?bookId="
	addBookURL     = "http://localhost:8080/book/add"
)

// ImportBook is import book data
func ImportBook(bookMap map[string]string) (string, interface{}) {

	// get book info
	respGB, err := http.Get(getBookInfoURL + bookMap["book_id"])
	if err != nil {
		log.Fatal(err)
	}
	defer respGB.Body.Close()

	bodyGB, err := ioutil.ReadAll(respGB.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bodyGB, &data); err != nil {
		log.Fatal(err)
	}

	book := data["data"].(map[string]interface{})
	book["bookId"] = bookMap["book_id"]
	book["status"] = 0
	book["onlineTime"] = time.Now()
	if bookMap["name"] != "" {
		book["name"] = bookMap["name"]
	}
	if bookMap["author"] != "" {
		book["author"] = bookMap["author"]
	}
	if bookMap["translator"] != "" {
		book["translator"] = bookMap["translator"]
	}
	if bookMap["painter"] != "" {
		book["painter"] = bookMap["painter"]
	}
	if bookMap["introduction"] != "" {
		book["introduction"] = bookMap["introduction"]
	}
	if bookMap["cover_url"] != "" {
		book["coverUrl"] = bookMap["cover_url"]
	}
	if bookMap["page_count"] != "" {
		book["pageCount"] = bookMap["page_count"]
	}
	book["type"] = bookMap["type"]
	book["teachingPlan"] = bookMap["teaching_plan"]
	book["classId"] = bookMap["class_id"]
	if bookMap["min_age"] != "" {
		book["minAge"] = bookMap["min_age"]
	}
	if bookMap["max_age"] != "" {
		book["maxAge"] = bookMap["max_age"]
	}
	book["summary"] = bookMap["summary"]
	book["order"] = bookMap["order"]
	if bookMap["supplier_id"] != "" {
		book["supplierId"] = bookMap["supplier_id"]
	}
	if bookMap["copyright"] != "" {
		book["copyright"] = bookMap["copyright"]
	}
	if bookMap["press"] != "" {
		book["press"] = bookMap["press"]
	}
	if bookMap["proprietor"] != "" {
		book["proprietor"] = bookMap["proprietor"]
	}
	// data["data"].(map[string]interface{})["labelAndRatio"]
	// data["data"].(map[string]interface{})["sceneMapList"]

	// get book downloadUrl
	respBP, err := http.Get("http://cdn.hhdd.com/" + book["downloadUrl"].(string))
	if err != nil {
		log.Fatal(err)
	}
	defer respBP.Body.Close()

	bodyBP, _ := ioutil.ReadAll(respBP.Body)
	var bookPage map[string]interface{}
	if err := json.Unmarshal(bodyBP, &bookPage); err != nil {
		log.Fatal(err)
	}
	if bookMap["sound_url"] != "" {
		book["soundUrl"] = bookMap["sound_url"]
	} else {
		book["soundUrl"] = bookPage["soundUrl"]
	}

	// add book
	nbook, err := json.Marshal(data["data"])
	if err != nil {
		log.Fatal(err)
	}

	respAB, err := http.Post(addBookURL, "application/json", bytes.NewBuffer(nbook))
	if err != nil {
		log.Fatal(err)
	}
	defer respAB.Body.Close()

	bodyAB, _ := ioutil.ReadAll(respAB.Body)
	fmt.Printf("bookId: %s\n\n", bodyAB)

	return string(bodyAB[:]), bookPage["pages"]
}
