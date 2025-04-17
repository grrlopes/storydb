package index

import (
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute()
}

type execute struct {
	repository repositories.IBleveRepository
	database   repositories.ISqliteRepository
}

func NewBleveIndex(repo repositories.IBleveRepository, sqlrepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
		database:   sqlrepo,
	}
}

func (e execute) Execute() {
	data, _ := e.database.AllWithoutLimit()
	e.repository.Indexes(data)
}
