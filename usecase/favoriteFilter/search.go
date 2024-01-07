package favoriteFilter

import (
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"gorm.io/gorm"
)

type InputBoundary interface {
	Execute(string, int, int) ([]entity.Commands, int, error)
}

type execute struct {
	favrepository repositories.ISqliteRepository
}

func NewFavoriteFilter(favrepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		favrepository: favrepo,
	}
}

func (e execute) Execute(filter string, limit int, skip int) ([]entity.Commands, int, error) {
	var cmds []entity.Commands

	result, count, err := e.favrepository.SearchFavorite(filter, limit, skip)
	if err != nil {
		log.Fatal("Search:", err)
	}

	for _, v := range result {
		cmds = append(cmds, entity.Commands{
			Model:    gorm.Model{ID: uint(v.CommandsID)},
			Cmd:      v.CommandsCMD,
			Desc:     v.CommandsDESC,
			Favorite: entity.Favorite{},
		})
	}

	return cmds, count, err
}
