package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite"
)

func init() {
	db, err := sql.Open("sqlite3", "blog.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	PrintDb(db)
	initMigration()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/fisk", fisk)
	http.ListenAndServe(":8080", mux)
}
