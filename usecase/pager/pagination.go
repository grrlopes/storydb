package pager

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
)

type InputBoundary interface {
	Execute(int, int) ([]list.Item, error)
}

type execute struct {
	repository repositories.ISqliteRepository
}

func NewPager(repo repositories.ISqliteRepository) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(limit int, offset int) ([]list.Item, error) {
	result, err := e.repository.Pagination(limit, offset)
	items := []list.Item{}

	if err != nil {
		log.Fatal("Pager:", err)
	}

	for _, value := range result {
		items = append(
			items,
			entity.NewListPanel{
				SqliteCommand: entity.SqliteCommand(value),
			},
		)
	}

	return items, err
}
