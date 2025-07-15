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

	//, abstract, body_text, body_html
	rows, err := db.Query("select title, url from en")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var url string
		rows.Scan(&title, &url)
		fmt.Printf("Title: %s - Url: %s\n", title, url)
	}

}
