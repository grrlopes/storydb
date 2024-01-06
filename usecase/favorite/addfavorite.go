package favorite

import (
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(uint) entity.Warning
}

type execute struct {
	favrepository repositories.ISqliteRepository
}

func NewFavorite(favrepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		favrepository: favrepo,
	}
}

func (e execute) Execute(id uint) entity.Warning {
	result := e.favrepository.AddFavorite(id)
	if result > 0 {
		return entity.Warning{
			Active:  true,
			Message: "That cmd has been added to Favorite!!!",
			Color:   "#006633",
		}
	}
	return entity.Warning{
		Active:  true,
		Message: "That cmd is already added to favorite!..",
		Color:   "#990000",
	}
}
