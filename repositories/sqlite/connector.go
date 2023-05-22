package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal("connector: ", err)
	}

	return db, err
}
