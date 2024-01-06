package entity

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/grrlopes/storydb/helper"
)

type Warning struct{
  Active bool
  Message string
  Color string
}

type CmdModel struct {
	Content            string
	Cursor             int
	Ready              bool
	Selected           Commands
	RowChosen          string
	Viewport           viewport.Model
	Start              int
	End                int
	PageTotal          int
	Pagination         *paginator.Model
	Count              *int
	Fcount             int
	ActiveSyncScreen   bool
	ActiveFinderScreen bool
	ActiveFavScreen    bool
	StatusSyncScreen   bool
	FavoriteScreen     bool
	Finder             textinput.Model
	FinderFilter       string
	Favorite           textinput.Model
	FavoriteFilter     string
	Store              []Commands
	HomeKeys           helper.Keymap
	FinderKeys         helper.Keymap
	SyncKeys           helper.Keymap
	FavoriteKeys       helper.Keymap
	Help               help.Model
	Spinner            spinner.Model
	Warning            Warning
}
