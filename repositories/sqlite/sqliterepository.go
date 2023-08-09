package sqlite

import (
	"database/sql"
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"gorm.io/gorm"
)

type SQLiteRepository struct {
	db       *sql.DB
	database *gorm.DB
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

func NewGormRepostory() repositories.ISqliteRepository {
	db, err := GormOpenDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Commands{})
	if err != nil {
		log.Fatal("not able to migrate", err)
	}

	return &SQLiteRepository{
		database: db,
	}
}

func (sql *SQLiteRepository) Migrate() error {
	_, err := sql.db.Exec(table)
	return err
}

func (sql *SQLiteRepository) All(limit int) ([]entity.Commands, int, error) {
	var commands []entity.Commands

	result := sql.database.Limit(limit).Find(&commands)

	if result.Error != nil {
		return commands, limit, result.Error
	}

	return commands, limit, nil
}

// Pagination implements repositories.ISqliteRepository
func (sql *SQLiteRepository) Pagination(limit int, offset int) ([]entity.Commands, error) {
	var commands []entity.Commands

	result := sql.database.Limit(limit).Offset(offset).Find(&commands)
	if result.Error != nil {
		return commands, result.Error
	}

	return commands, nil
}

func (sql *SQLiteRepository) Count() (int, error) {
	var count int

	err := sql.db.QueryRow("SELECT COUNT(*) FROM commands").Scan(&count)
	if err != nil {
		return count, err
	}

	return count, err
}

func (sql SQLiteRepository) InsertParsed(data string) (int64, error) {
	res, err := sql.db.Exec("INSERT INTO commands(cmd, desc) values(?, ?)", data, "--")
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
	var (
		count    int
		commands []entity.Commands
	)

	result := sql.database.Limit(limit).Offset(skip).Where("cmd LIKE ?", "%"+filter+"%").Find(&commands)
	if result.Error != nil {
		return commands, count, result.Error
	}

	return commands, count, result.Error
}

func (sql *SQLiteRepository) SearchCount(filter string) (int, error) {
	var (
		count   int64
		command entity.Commands
	)

	sql.database.Model(&command).Where("cmd LIKE ?", "%"+filter+"%").Count(&count)

	return int(count), nil
}
