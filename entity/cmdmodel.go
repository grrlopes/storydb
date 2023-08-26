package entity

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/grrlopes/storydb/helper"
)

type CmdModel struct {
	Content            string
	Cursor             int
	Ready              bool
	Selected           string
	RowChosen          string
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
	Finder             textinput.Model
	FinderFilter       string
	Store              []Commands
	HomeKeys           helper.Keymap
	FinderKeys         helper.Keymap
	Help               help.Model
	Spinner            spinner.Model
}
