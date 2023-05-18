package repositories

import "github.com/grrlopes/storydb/entity"

type SqliteCmd entity.SqliteCommand

type ISqliteRepository interface {
	Migrate() error
	All() ([]SqliteCmd, error)
}
