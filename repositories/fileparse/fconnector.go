package fileparse

import (
	"log"
	"os"
)

func OpenHist() (*os.File, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find the file:", err, "\n")
		os.Exit(1)
	}

	data, err := os.Open(homedir + "/.bash_history")
	if err != nil {
		log.Fatal("could not load file:", err, "\n")
		os.Exit(1)
	}

	return data, err
}
