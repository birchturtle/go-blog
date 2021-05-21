package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

func initMigration() {
	fmt.Println("init run")
}

func PrintDb(conn *sql.DB) {
	rows, err := conn.Query("SELECT * FROM test", nil)
	if err != nil {
		panic("rows broken")
	}
	for rows.Next() {
		var buffer string
		if err = rows.Scan(&buffer); err != nil {
			panic("Scan broken")
		}
		fmt.Println("fra db: %s", buffer)
	}
}
