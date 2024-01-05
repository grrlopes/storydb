package favorite

import (
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(uint) string
}

type execute struct {
	favrepository repositories.ISqliteRepository
}

func NewFavorite(favrepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		favrepository: favrepo,
	}
}

func (e execute) Execute(id uint) string {
	result := e.favrepository.AddFavorite(id)
	if result > 0 {
		return "That cmd was added to Favorite!!!"
	}
	return "That favorite already exists!!!"
}
