package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/usecase/count"
	"github.com/grrlopes/storydb/usecase/pager"
)

var (
	repository   repositories.ISqliteRepository = sqlite.NewSQLiteRepository()
	usecasePager pager.InputBoundary            = pager.NewPager(repository)
	usecaseCount count.InputBoundary            = count.NewCount(repository)
)

type ModelHome struct {
	home entity.Command
}

func NewHome(m entity.Command) *ModelHome {
	count := usecaseCount.Execute()
	p := paginator.New()
	p.PerPage = 10
	p.SetTotalPages(count)

	home := ModelHome{
		home: entity.Command{
			Content:    m.Content,
			Ready:      false,
			Selected:   "",
			Viewport:   viewport.Model{},
			PageTotal:  m.PageTotal,
			Pagination: &p,
			Count:      count,
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
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return &m, tea.Quit
		case "up", "k":
			if m.home.Cursor > 0 {
				m.home.Cursor--
			}
		case "down", "j":
			if m.home.Cursor < m.home.PageTotal-1 {
				m.home.Cursor++
			}
		case "ctrl+g":
			m.home.Cursor = m.home.PageTotal - 1
		case "ctrl+u":
			m.home.Cursor = 0
		case "l", "right":
		case "enter":
			fmt.Print(m.home.Cursor + 1)
		}
	case tea.WindowSizeMsg:
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height - 6
		m.home.Viewport.SetContent(m.GetDataView())
		m.home.Ready = true
	}
	*m.home.Pagination, cmd = m.home.Pagination.Update(msg)
	m.home.Viewport.SetContent(m.GetDataView())
	return &m, cmd
}

func (m ModelHome) View() string {
	view := lipgloss.NewStyle()
	content := lipgloss.NewStyle()
	if !m.home.Ready {
		return "\n  Loading..."
	}
	start, end := m.updatepagination()
	m.home.Start = &start
	m.home.End = &end

	return view.Render(
		m.HeaderView()) + "\n" +
		content.Render(m.home.Viewport.View()) + "\n" +
		m.FooterView() + "\n" +
		m.paginationView()
}

func (m *ModelHome) GetSelected() string {
	return m.home.Selected
}

func (m *ModelHome) updatepagination() (int, int) {
	return 1, 1
}

func (m *ModelHome) GetDataView() string {
	data, _ := usecasePager.Execute((m.home.Viewport.Height - 1), 0)
	m.home.PageTotal = len(data)
	var (
		result []string
		maxLen = m.home.Viewport.Width
	)

	for i, v := range data {
		if m.home.Cursor == i {
			v.Desc = SelecRow.Render(v.Desc)
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
