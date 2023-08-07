package entity

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type CmdModel struct {
	Content            string
	Cursor             int
	Ready              bool
	Selected           string
	Viewport           viewport.Model
	Start              int
	End                int
	PageTotal          int
	Pagination         *paginator.Model
	Count              *int
	Fcount             int
	Ftotal             int
	ActiveSyncScreen   bool
	ActiveFinderScreen bool
	StatusSyncScreen   bool
	ProgressSync       progress.Model
	Finder             textinput.Model
	FinderFilter       string
	Store            []Commands
}
