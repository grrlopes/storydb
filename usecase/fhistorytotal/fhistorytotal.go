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
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find the file:", err, "\n")
		os.Exit(1)
	}

	file, err := os.Open(homedir + "/.bash_history")
	if err != nil {
		log.Println("could not load file:", err)
	}

	defer file.Close()

	fcount := 0
	prevByte := make([]byte, 1)
	b := make([]byte, 1)

	for {
		_, err := file.Read(b)

		if err != nil {
			log.Println("could not open file:", err)
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
