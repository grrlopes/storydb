package repositories

import "github.com/grrlopes/storydb/entity"

type SqliteCmd entity.SqliteCommand

type ISqliteRepository interface {
	Migrate() error
	All(int) ([]SqliteCmd, int, error)
	Pagination(int, int) ([]SqliteCmd, error)
	Count() (int, error)
	InsertParsed(string) (int64, error)
	Search(string, int, int) ([]entity.SqliteCommand, int, error)
}
