package sqlite

import (
	"database/sql"
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository() repositories.ISqliteRepository {
	db, err := OpenDB()

	if err != nil {
		log.Fatal(err)
	}

	return &SQLiteRepository{
		db: db,
	}
}

func (sql *SQLiteRepository) Migrate() error {
	_, err := sql.db.Exec(table)
	return err
}

func (sql *SQLiteRepository) All() ([]repositories.SqliteCmd, error) {
	result, err := sql.db.Query("SELECT * FROM command")
	if err != nil {
		return []repositories.SqliteCmd{}, err
	}
	defer result.Close()

	var data []repositories.SqliteCmd

	for result.Next() {
		var command repositories.SqliteCmd
		if err := result.Scan(
			&command.ID,
			&command.EnTitle,
			&command.Desc,
		); err != nil {
			return []repositories.SqliteCmd{}, err
		}

		data = append(data, command)
	}

	return data, nil
}
