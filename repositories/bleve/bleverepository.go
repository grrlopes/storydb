package bleve

import (
	"log"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type BleveRepository struct {
	indexdb bleve.Index
}

func NewBleveRepostory() repositories.IBleveRepository {
	db, err := BleveOpenDB()
	if err != nil {
		log.Fatal(err)
	}

	return &BleveRepository{
		indexdb: db,
	}
}

func (b *BleveRepository) Search(filter string) []interface{} {
	query := bleve.NewMatchPhraseQuery(filter)

	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Size = 1000
	results, err := b.indexdb.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	ids := make([]interface{}, 0, len(results.Hits))
	for _, hit := range results.Hits {
		ids = append(ids, hit.ID)
	}

	defer b.indexdb.Close()

	return ids
}

func (b BleveRepository) Indexes(data []entity.Commands) {
	for id, doc := range data {
		err := b.indexdb.Index(strconv.Itoa(id), doc)
		if err != nil {
			log.Fatalf("Error indexing document %s: %v", strconv.Itoa(id), err)
		}
	}
	defer b.indexdb.Close()
}
