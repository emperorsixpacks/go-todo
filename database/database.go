package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(db_path string) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error initializing db", err)
	}
}
