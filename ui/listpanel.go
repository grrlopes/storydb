package ui

import "github.com/grrlopes/storydb/entity"

type NewListPanel struct {
	entity.SqliteCommand
}

func (list NewListPanel) Title() string       { return list.EnTitle }
func (list NewListPanel) Description() string { return list.Desc }
func (list NewListPanel) FilterValue() string { return list.EnTitle }
