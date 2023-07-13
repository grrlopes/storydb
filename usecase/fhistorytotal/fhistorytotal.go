package fhistorytotal

import (
	"log"
	"os"

	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute() int
}

type execute struct {
	frepository repositories.IFileParsedRepository
	srepository repositories.ISqliteRepository
}

func NewFHistoryTotal(frepo repositories.IFileParsedRepository, srepo repositories.ISqliteRepository) InputBoundary {
	return execute{
		frepository: frepo,
		srepository: srepo,
	}
}

func (e execute) Execute() int {
	file, err := os.Open("~/.bash_history")
	if err != nil {
		log.Println("Error to open file:", err)
	}

  defer file.Close()

	fcount := 0
	prevByte := make([]byte, 1)
	b := make([]byte, 1)

	for {
		_, err := file.Read(b)

		if err != nil {
			log.Println("Error to open file:", err)
		}

		if b[0] == '\n' && prevByte[0] != '\r' {
			fcount++
		}

		if err == nil {
			prevByte[0] = b[0]
		} else {
			break
		}
	}

	return fcount
}
