package fhistory

import (
	"bufio"
	"fmt"
	"log"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute()
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

func (e execute) Execute() {
	fresult := e.frepository.All()
	scanner := bufio.NewScanner(fresult)
	for scanner.Scan() {
		data := scanner.Text()
		number, err := e.srepository.InsertParsed(data)
		fmt.Println(number, err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
