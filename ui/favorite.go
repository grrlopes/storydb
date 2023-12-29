package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
)

type (
	favoriteMsg      []entity.Commands
	favoriteCountMsg int
	favoritePagMsg   struct{}
)

func FavoriteCmd(filter textinput.Model, limit int, offset int) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = func() tea.Msg {
		data, _, _ := usecaseFinder.Execute(filter.Value(), limit, offset)
		return favoriteMsg(data)
	}

	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func favoriteUpdate(msg tea.Msg, m ModelHome) (*ModelHome, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			if m.home.ActiveFavScreen {
				m.home.ActiveFavScreen = false
				m.home.Viewport.SetContent(m.GetDataView())
				return &m, nil
			}
		}
		switch msg.String() {
		case "up", "k":
			if m.home.Cursor > 0 {
				m.home.Content = "arrow"
				m.home.Cursor--
			}
		case "down", "j":
			if m.home.Cursor < m.home.PageTotal-1 {
				m.home.Content = "arrow"
				m.home.Cursor++
			}
		}
	}
	m.home.Viewport.SetContent(finderView(&m))
	return &m, cmd
}

func favoriteView(m *ModelHome) string {
	return view.Render(
		m.home.Favorite.View(),
		favoriteDataView(m, m.home.FavoriteFilter),
	)
}

func favoriteDataView(m *ModelHome, filter string) string {
	m.home.Start, _ = strconv.Atoi(filter)
	*m.home.Count = 2
	var (
		pagey   = m.home.PageTotal - 1
		selecty = m.home.Content
	)

	m.home.Store, _, _ = usecaseFinder.Execute(filter, 18, 1)
	m.home.PageTotal = len(m.home.Store)
	var (
		result []string
		maxLen = m.home.Viewport.Width
	)

	for i, v := range m.home.Store {
		if m.home.Cursor == i && selecty == "arrow" {
			m.home.Selected = v.Cmd
			v.Cmd = SelecRow.Render(v.Cmd)
		}

		if pagey == i && selecty == "window" {
			v.Cmd = SelecRow.Render(v.Cmd)
		}

		if len(v.Cmd) > maxLen {
			title := ShrinkWordMiddle(v.Cmd, maxLen)
			v.Cmd = title
		}

		result = append(result, fmt.Sprintf("\n%s", v.Cmd))
	}

	rowData := strings.Trim(fmt.Sprintf("%s", result), "[]")

	return rowData
}
