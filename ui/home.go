package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/usecase/pager"
)

var (
	repository   repositories.ISqliteRepository = sqlite.NewSQLiteRepository()
	usecasePager pager.InputBoundary            = pager.NewPager(repository)
)

type ModelHome struct {
	home entity.Command
}

func NewHome(m entity.Command) *ModelHome {
	p := ModelHome{
		home: entity.Command{
			Content:   m.Content,
			Ready:     false,
			Selected:  "",
			Viewport:  viewport.Model{},
			PageTotal: m.PageTotal,
		},
	}
	return &p
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
		if msg.String() == "ctrl+c" {
			return &m, tea.Quit
		}
		if msg.String() == "enter" {
			command := m.home.Content.SelectedItem().FilterValue()
			m.home.Selected = command
			return &m, cmd
		}
		if msg.String() == "o" {
			data, _ := usecasePager.Execute(9, 18)
			items := []list.Item{}

			for _, value := range data {
				items = append(
					items,
					NewListPanel{
						SqliteCommand: entity.SqliteCommand(value),
					},
				)
			}
			m.home.Content.SetItems(items)
			m.home.Content.ResetFilter()
			return &m, cmd
		}
	case tea.WindowSizeMsg:
		h, v := winSize.GetFrameSize()
		m.home.Content.SetSize(msg.Width-h, msg.Height-v)
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height
		m.home.Ready = true
	}

	m.home.Content, cmd = m.home.Content.Update(msg)
	return &m, cmd
}

func (m ModelHome) View() string {
	view := lipgloss.NewStyle()
	content := lipgloss.NewStyle()
	if !m.home.Ready {
		return "\n  Loading..."
	}
	start, end := m.updatepagination()
	m.home.Start = start
	m.home.End = end

	return view.Render(
		m.HeaderView()) + "\n" +
		content.Render(m.home.Content.View()) + "\n" +
		view.Render(m.FooterView())
}

func (m *ModelHome) GetSelected() string {
	return m.home.Selected
}

func (m *ModelHome) updatepagination() (int, int) {
	start, end := m.home.Content.Paginator.GetSliceBounds(m.home.PageTotal)
	m.home.Content.Paginator.SetTotalPages(m.home.PageTotal)
	return start, end

}
