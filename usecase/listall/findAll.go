package listall

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(int) ([]repositories.SqliteCmd, int, error)
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewListAll(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(limit int) ([]repositories.SqliteCmd, int, error) {
	result, count, err := e.repository.All(limit)
	if err != nil {
		log.Fatal("findAll:", err)
	}

	return result, count, err
}
