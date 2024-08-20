package bleve

import (
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
)

func CreateIndexDB() bleve.Index {
	var store = "/.local/share/bleve/"
	homedir, err := os.UserHomeDir()
	os.Mkdir(homedir+store, 0755)

	// Define the mapping
	mapping := bleve.NewIndexMapping()

	// Create the index
	index, err := bleve.New("storydb.bleve", mapping)
	if err != nil {
		log.Fatal(err)
	}

	return index
}
