package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	var store = "/.local/share/storydb/"

	homedir, err := os.UserHomeDir()
	os.Mkdir(homedir+store, 0755)

	db, err := sql.Open("sqlite3", homedir+store+"sqlite.db")
	if err != nil {
		log.Fatal("connector: ", err)
	}

	return db, err
}
