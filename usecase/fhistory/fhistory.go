package fhistory

import (
	"bufio"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/repositories"
)

type SyncMsg int

type InputBoundary interface {
	Execute() tea.Cmd
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

func (e execute) Execute() tea.Cmd {
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

	cmd := func() tea.Msg {
		return SyncMsg(fcount)
	}
	return cmd
}
