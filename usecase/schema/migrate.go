package schema

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute()
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewMigrate(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() {
	err := e.repository.Migrate()

	if err != nil {
		log.Fatal("migrate:", err)
	}
}
