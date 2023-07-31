package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/fileparse"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/usecase/count"
	"github.com/grrlopes/storydb/usecase/fhistory"
	"github.com/grrlopes/storydb/usecase/fhistorytotal"
	"github.com/grrlopes/storydb/usecase/finder"
	"github.com/grrlopes/storydb/usecase/pager"
)

var (
	repository          repositories.ISqliteRepository     = sqlite.NewSQLiteRepository()
	frepository         repositories.IFileParsedRepository = fileparse.NewFparsedRepository()
	usecasePager        pager.InputBoundary                = pager.NewPager(repository)
	usecaseCount        count.InputBoundary                = count.NewCount(repository)
	usecaseHistory      fhistory.InputBoundary             = fhistory.NewFHistory(frepository, repository)
	usecaseHistoryTotal fhistorytotal.InputBoundary        = fhistorytotal.NewFHistoryTotal(frepository, repository)
	usecaseFinder       finder.InputBoundary               = finder.NewFinder(repository)
)

type ModelHome struct {
	home entity.Command
}

func NewHome(m *entity.Command) *ModelHome {
	count := usecaseCount.Execute()
	ftotal := usecaseHistoryTotal.Execute()
	p := paginator.New()
	p.PerPage = 18
	p.SetTotalPages(count)
	pro := progress.New(progress.WithDefaultGradient())
	txt := textinput.New()
	txt.Placeholder = "type"
	txt.CharLimit = 156
	txt.Width = 20
	txt.Prompt = "Finder: "

	home := ModelHome{
		home: entity.Command{
			Content:          m.Content,
			Ready:            false,
			Viewport:         viewport.Model{},
			PageTotal:        m.PageTotal,
			Pagination:       &p,
			Count:            &count,
			ActiveSyncScreen: false,
			StatusSyncScreen: false,
			ProgressSync:     pro,
			Ftotal:           ftotal,
			Finder:           txt,
		},
	}
	return &home
}

func (m ModelHome) HeaderView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) FooterView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) Update(msg tea.Msg) (*ModelHome, tea.Cmd) {
	if m.home.ActiveSyncScreen {
		m.home.Ready = true
		synced, cmd := syncUpdate(msg, m)
		return synced, cmd
	}

	// if m.home.ActiveFinderScreen {
	// 	m.home.Ready = true
	// 	finder, cmd := finderUpdate(msg, m)
	// 	return finder, cmd
	// }

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.home.Finder.Focused() {
			if msg.String() == "ctrl+c" {
				m.home.Finder.Blur()
			}
			m.home.Store, *m.home.Count = finderCmd(m.home.Finder.Value(), 18, m.home.Start)
			start, end := m.updatepagination()
			m.home.Start = start
			m.home.End = end
			m.home.Finder, cmd = m.home.Finder.Update(msg)
		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				return &m, tea.Quit
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
			case "ctrl+g":
				m.home.Cursor = m.home.PageTotal - 1
			case "s":
				m.home.StatusSyncScreen = false
				m.home.ActiveSyncScreen = true
				m.home.Viewport.SetContent(syncView(&m))
				return &m, cmd
			case "ctrl+u":
				m.home.Cursor = 0
			case "enter":
				return &m, tea.Quit
			case "/":
				m.home.Finder.Focus()
			}
		}
	case tea.WindowSizeMsg:
		m.home.Content = "window"
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height - 6
		m.home.Viewport.SetContent(m.GetDataView())
		m.home.Ready = true
	}

	*m.home.Pagination, cmd = m.home.Pagination.Update(msg)
	if !m.home.Finder.Focused() {
		start, end := m.updatepagination()
		m.home.Start = start
		m.home.End = end
	}

	m.home.Viewport.SetContent(m.GetDataView())
	return &m, cmd
}

func (m ModelHome) View() string {
	view := lipgloss.NewStyle()
	content := lipgloss.NewStyle()
	if !m.home.Ready {
		return "\n  Loading..."
	}

	return view.Render(m.HeaderView()) + "\n" + m.home.Finder.View() +
		content.Render(m.home.Viewport.View()) + "\n" +
		m.FooterView() + "\n" +
		m.paginationView()
}

func (m *ModelHome) GetSelected() string {
	return m.home.Selected
}

func (m *ModelHome) updatepagination() (int, int) {
	start, end := m.home.Pagination.GetSliceBounds(*m.home.Count)
	return start, end
}

func (m *ModelHome) GetDataView() string {
	var (
		pagey   = m.home.PageTotal - 1
		selecty = m.home.Content
	)

	if !m.home.Finder.Focused() {
		m.home.Store, _ = usecasePager.Execute(18, m.home.Start)
	}
	m.home.PageTotal = len(m.home.Store)
	m.home.Pagination.SetTotalPages(*m.home.Count)
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

func (m *ModelHome) paginationView() string {
	var b strings.Builder
	b.WriteString("  " + m.home.Pagination.View())
	b.WriteString("\n\n  h/l ←/→ page • q: quit\n")
	return b.String()
}
