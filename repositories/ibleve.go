package repositories

import (
	"github.com/grrlopes/storydb/entity"
)

type IBleveRepository interface {
	Indexes([]entity.Commands)
	Search(string) []interface{}
}
