package repositories

import "github.com/grrlopes/storydb/entity"

type SqliteCmd entity.Commands

type ISqliteRepository interface {
	Migrate() error
	All(int) ([]entity.Commands, int, error)
	Pagination(int, int) ([]entity.Commands, error)
	Count() (int, error)
	InsertParsed(string) (int64, error)
	Search(string, int, int) ([]entity.Commands, int, error)
	SearchCount(string) (int, error)
}
