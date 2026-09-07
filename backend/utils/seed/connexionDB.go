package seed

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() *sql.DB {
	db, err := sql.Open("sqlite3", "../backend/database/social_network.db")
	if err != nil {
		log.Fatal(err)
	}

	return db

}

var DB = CreateDB()
