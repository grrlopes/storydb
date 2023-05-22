package schema

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() error
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewMigrate(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() error {
	err := e.repository.Migrate()

	if err != nil {
		log.Fatal("migrate:", err)
	}

	return err
}
