package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting processing with WikiBooks ...")

	db, err := sql.Open("sqlite3", "wikibooks.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
