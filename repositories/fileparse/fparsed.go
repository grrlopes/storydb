package fileparse

import (
	"log"
	"os"

	"github.com/grrlopes/storydb/repositories"
)

type fparsed struct {
	repositories.IFileParsedRepository
	content *os.File
}

func NewFparsedRepository() repositories.IFileParsedRepository {
	data, err := OpenHist()
	if err != nil {
		log.Fatal(err, data)
	}

	return fparsed{
		content: data,
	}
}

func (p fparsed) All() *os.File {
	return p.content
}
