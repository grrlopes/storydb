package favorite

import (
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/helper"
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
			Message: helper.OKAddFavorite,
			Color:   "#006633",
		}
	}
	return entity.Warning{
		Active:  true,
		Message: helper.ErrAddFavorite.Error(),
		Color:   "#990000",
	}
}
