package listall

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() ([]repositories.SqliteCmd, error)
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewListAll(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() ([]repositories.SqliteCmd, error) {
	result, err := e.repository.All()

	if err != nil {
		log.Fatal("migrate:", err)
	}

	return result, err
}
