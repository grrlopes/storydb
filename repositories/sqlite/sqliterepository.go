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
	err := sql.database.Exec(table).Error
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
	var (
		count   int64
		command entity.Commands
	)

	sql.database.Model(&command).Count(&count)

	return int(count), nil
}

func (sql SQLiteRepository) InsertParsed(data string) {
	command := entity.Commands{Cmd: data, Desc: "---"}
	sql.database.Create(&command)
}

func (sql *SQLiteRepository) Search(filter string, limit int, skip int) ([]entity.Commands, int, error) {
	var (
		count    int
		commands []entity.Commands
		result   *gorm.DB
	)

	if filter == `""*` {
		result = sql.database.Limit(limit).Offset(skip).Where("cmd LIKE ?", "%"+""+"%").Find(&commands)
	} else {
		result = sql.database.Raw(
			"SELECT cmd FROM commands_fts WHERE cmd MATCH ? ORDER BY rank LIMIT ? OFFSET ?",
			filter, limit, skip,
		).Find(&commands)
	}
	if result.Error != nil {
		return commands, count, result.Error
	}

	return commands, count, result.Error
}

func (sql *SQLiteRepository) SearchCount(filter string) (int, error) {
	var (
		count       int64
		countResult int
		command     []entity.Commands
		err         error
	)

	if filter == `""*` {
		err = sql.database.Model(&command).Where("cmd LIKE ?", "%"+""+"%").Count(&count).Error
		countResult = int(count)
	} else {
		err = sql.database.Raw("SELECT cmd FROM commands_fts WHERE cmd MATCH ?", filter).Find(&command).Error
		countResult = len(command)
	}
	return countResult, err
}
