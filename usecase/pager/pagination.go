package pager

import (
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(int, int) ([]repositories.SqliteCmd, error)
}

type NewListPanel struct {
	entity.SqliteCommand
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewPager(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(limit int, offset int) ([]repositories.SqliteCmd, error) {
	result, err := e.repository.Pagination(limit, offset)

	if err != nil {
		log.Fatal("Pager:", err)
	}

	return result, err
}
