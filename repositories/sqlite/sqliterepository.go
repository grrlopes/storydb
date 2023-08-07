package sqlite

import (
	"database/sql"
	"log"

	"github.com/grrlopes/storydb/entity"
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

func (sql *SQLiteRepository) All(limit int) ([]repositories.SqliteCmd, int, error) {
	var count int
	result, err := sql.db.Query("SELECT * FROM command limit ?", limit)

	if err != nil {
		return []repositories.SqliteCmd{}, count, err
	}
	defer result.Close()

	err = sql.db.QueryRow("SELECT COUNT(*) FROM command").Scan(&count)

	var data []repositories.SqliteCmd

	for result.Next() {
		var command repositories.SqliteCmd
		if err := result.Scan(
			&command.ID,
			&command.EnTitle,
			&command.Desc,
		); err != nil {
			return []repositories.SqliteCmd{}, count, err
		}

		data = append(data, command)
	}

	return data, count, nil
}

// Pagination implements repositories.ISqliteRepository
func (sql *SQLiteRepository) Pagination(limit int, offset int) ([]repositories.SqliteCmd, error) {
	result, err := sql.db.Query("SELECT * FROM command limit ? offset ?", limit, offset)

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

func (sql *SQLiteRepository) Count() (int, error) {
	var count int

	err := sql.db.QueryRow("SELECT COUNT(*) FROM command").Scan(&count)
	if err != nil {
		return count, err
	}

	return count, err
}

func (sql SQLiteRepository) InsertParsed(data string) (int64, error) {
	res, err := sql.db.Exec("INSERT INTO command(title, description) values(?, ?)", data, "--")
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return id, err
	}

	return id, err
}

func (sql *SQLiteRepository) Search(filter string, limit int, skip int) ([]entity.Commands, int, error) {
	var count int

	stmt, err := sql.db.Prepare("SELECT * FROM command WHERE title LIKE ? limit ? offset ?")
	if err != nil {
		return []entity.Commands{}, count, err
	}

	result, err := stmt.Query("%"+filter+"%", limit, skip)
	if err != nil {
		return []entity.Commands{}, count, err
	}

	defer result.Close()

	err = sql.db.QueryRow("SELECT COUNT(*) FROM command WHERE Title LIKE ?", "%"+filter+"%").Scan(&count)

	var data []entity.Commands

	for result.Next() {
		var command entity.Commands
		if err := result.Scan(
			&command.ID,
			&command.EnTitle,
			&command.Desc,
		); err != nil {
			return []entity.Commands{}, count, err
		}

		data = append(data, command)
	}

	return data, count, nil
}

func (sql *SQLiteRepository) SearchCount(filter string) (int, error) {
	var count int

	sql.db.QueryRow("SELECT COUNT(*) FROM command WHERE Title LIKE ?", "%"+filter+"%").Scan(&count)

	return count, nil
}
