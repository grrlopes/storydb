package count

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() int
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewCount(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() int {
	result, err := e.repository.Count()

	if err != nil {
		log.Fatal("PagerTotal:", err)
	}

	return result
}
