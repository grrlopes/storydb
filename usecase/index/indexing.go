package index

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() (bool, error)
}

type execute struct {
	repository repositories.ISqliteRepository
	indexing   repositories.IBleveRepository
}

func NewIndex(repo repositories.ISqliteRepository, index repositories.IBleveRepository) InputBoundary {
	return execute{
		repository: repo,
		indexing:   index,
	}
}

func hashDocument(doc entity.Commands) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d%s%s", doc.ID, doc.Cmd, doc.Desc)))
	return hex.EncodeToString(h.Sum(nil))
}

func (e execute) Execute() (bool, error) {
	result, err := e.repository.AllWithoutLimit()
	if err != nil {
		log.Fatal("findAll:", err)
	}

	var index = e.indexing.Indexes()

	for _, v := range result {
		var indexedDoc entity.Commands
		err = index.Index(fmt.Sprintf("%d", v.ID), v)
		if err != nil {
			return false, err
		}

		if hashDocument(v) != hashDocument(indexedDoc) {
			return false, nil
		}
	}

	return true, nil
}
