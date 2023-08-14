package fileparse

import (
	"log"
	"os"

	"github.com/grrlopes/storydb/helper"
)

func OpenHist() (*os.File, error) {
	env := os.Getenv("HISTFILE")
	if env == "" {
		log.Fatalf("%s %s", "could not read HISTFILE\n", helper.ErrEnvHISTFailed)
	}

	data, err := os.Open(env)
	if err != nil {
		log.Fatalln("could not load file:", err)
		os.Exit(1)
	}

	return data, err
}
