package fhistory

import (
	"bufio"
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() int
}

type execute struct {
	frepository repositories.IFileParsedRepository
	srepository repositories.ISqliteRepository
}

func NewFHistory(frepo repositories.IFileParsedRepository, srepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		frepository: frepo,
		srepository: srepo,
	}
}

func (e execute) Execute() int {
	fresult := e.frepository.All()

	scanner := bufio.NewScanner(fresult)
	fcount := 0
	for scanner.Scan() {
		data := scanner.Text()
		fcount += 1
		e.srepository.InsertParsed(data)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fcount
}
