package fileparse

import (
	"log"
	"os"
)

func OpenHist() (*os.File, error) {
	data, err := os.Open("/tmp/history")
	if err != nil {
		log.Fatal("could not load file:", err)
		os.Exit(1)
	}

	return data, err
}
