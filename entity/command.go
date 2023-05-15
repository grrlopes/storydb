package entity

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
)

type Command struct {
	Content  list.Model
	Ready    bool
	Viewport viewport.Model
}
