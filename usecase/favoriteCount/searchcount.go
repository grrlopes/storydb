package favoritecount

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

func NewFavoriteCount(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(filter string) int {
	count, err := e.repository.SearchFavoriteCount(filter)
	if err != nil {
		log.Fatal("Favorite Search count:", err)
	}

	return count
}
