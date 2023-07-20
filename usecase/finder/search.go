package finder

import (
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(string, int, int) ([]entity.SqliteCommand, int, error)
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewFinder(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(filter string, limit int, skip int) ([]entity.SqliteCommand, int, error) {
	result, count, err := e.repository.Search(filter, limit, skip)
	if err != nil {
		log.Fatal("Search:", err)
	}

	return result, count, err
}
