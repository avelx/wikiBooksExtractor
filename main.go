package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	fmt.Println("Starting processing with WikiBooks ...")

	db, err := sql.Open("sqlite3", "wikibooks.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from wikibooks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var value int
		rows.Scan(&value)
		fmt.Printf("value: %d\n", value)
	}

}
