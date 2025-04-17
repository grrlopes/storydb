package repositories

import "github.com/grrlopes/storydb/entity"

type SqliteCmd entity.Commands

type ISqliteRepository interface {
	Migrate() error
	All(int) ([]entity.Commands, int, error)
	AllWithoutLimit() ([]entity.Commands, error)
	Pagination(int, int) ([]entity.Commands, error)
	Count() (int, error)
	InsertParsed(string)
	Search(string, int, int) ([]entity.Commands, int, error)
	SearchCount(string) (int, error)
	IndexSearch(filter []interface{}) ([]entity.Commands, int, error)
}
