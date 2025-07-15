package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Rows struct {
	title     string
	url       string
	abstract  string
	body_text string
	body_html string
}

func main() {
	fmt.Println("Starting processing with WikiBooks ...")

	db, err := sql.Open("sqlite3", "wikibooks.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select title, url, abstract, body_text, body_html from en")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var mRow = Rows{}
		rows.Scan(&mRow.title)
		rows.Scan(&mRow.url)
		rows.Scan(&mRow.abstract)
		rows.Scan(&mRow.body_text)
		rows.Scan(&mRow.body_html)
		fmt.Printf("Value: %s\n", mRow.title)
	}

}
