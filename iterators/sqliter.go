package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// START OMIT
	rows, err := db.Query("SELECT * FROM foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
