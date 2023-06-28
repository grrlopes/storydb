package entity

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
)

type Command struct {
	Content    string
	Cursor     int
	Ready      bool
	Selected   string
	Viewport   viewport.Model
	Start      *int
	End        *int
	PageTotal  int
	Pagination *paginator.Model
	Count      int
}
