package favorite

import (
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(string) int64
}

type execute struct {
	favrepository repositories.ISqliteRepository
}

func NewFavorite(favrepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		favrepository: favrepo,
	}
}

func (e execute) Execute(data string) int64 {
	result := e.favrepository.AddFavorite(data)
	return result
}
