package pager

import (
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	// It gets integer paramaters limit and offset
	Execute(int, int) ([]entity.SqliteCommand, error)
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewPager(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(limit int, offset int) ([]entity.SqliteCommand, error) {
	result, err := e.repository.Pagination(limit, offset)
	items := []entity.SqliteCommand{}

	if err != nil {
		log.Fatal("Pager:", err)
	}

	for _, value := range result {
		items = append(
			items,
			entity.SqliteCommand(value),
		)
	}

	return items, err
}
