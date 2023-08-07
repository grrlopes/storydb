package ui

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
)

func finderCmd(filter string, limit int, offset int) ([]entity.Commands, int) {
	data, total, _ := usecaseFinder.Execute(filter, limit, offset)
	return data, total
}

func finderCount(filter string) int {
	count := usecaseFinderCount.Execute(filter)
	return count
}

func finderUpdate(msg tea.Msg, m ModelHome) (*ModelHome, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			if m.home.ActiveFinderScreen {
				m.home.ActiveFinderScreen = false
				m.home.Finder.Reset()
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
	m.home.Finder, cmd = m.home.Finder.Update(msg)
	m.home.FinderFilter = m.home.Finder.Value()
	m.home.Viewport.SetContent(finderView(&m))
	return &m, cmd
}

func finderView(m *ModelHome) string {
	return view.Render(
		m.home.Finder.View(),
		finderDataView(m, m.home.FinderFilter),
	)
}

func finderDataView(m *ModelHome, filter string) string {
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
			m.home.Selected = v.EnTitle
			v.EnTitle = SelecRow.Render(v.EnTitle)
		}

		if pagey == i && selecty == "window" {
			v.EnTitle = SelecRow.Render(v.EnTitle)
		}

		if len(v.EnTitle) > maxLen {
			title := ShrinkWordMiddle(v.EnTitle, maxLen)
			v.EnTitle = title
		}

		result = append(result, fmt.Sprintf("\n%s", v.EnTitle))
	}

	rowData := strings.Trim(fmt.Sprintf("%s", result), "[]")

	return rowData
}
