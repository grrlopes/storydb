package repositories

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/grrlopes/storydb/entity"
)

type IBleveRepository interface {
	Indexes() bleve.Index
	Search(string) []entity.Commands
}
