package bleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type BleveRepository struct {
	indexing bleve.Index
}

func NewBleveRepostory() repositories.IBleveRepository {
	index := CreateIndexDB()

	return &BleveRepository{
		indexing: index,
	}
}

func (b *BleveRepository) Search(filter string) []entity.Commands {
	var (
		command []entity.Commands
	)

	return command
}

func (b BleveRepository) Indexes() bleve.Index {
	return b.indexing
}
