package findercount

import (
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(string) int
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewFinderCount(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(filter string) int {
	count, err := e.repository.SearchCount(filter)
	if err != nil {
		log.Fatal("Search count:", err)
	}

	return count
}
