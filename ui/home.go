package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
)

type ModelHome struct {
	home entity.Command
}

func NewHome(m entity.Command) *ModelHome {
	p := ModelHome{
		home: entity.Command{
			Content:  m.Content,
			Ready:    false,
			Selected: "",
			Viewport: viewport.Model{},
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

	return view.Render(
		m.HeaderView()) + "\n" +
		content.Render(m.home.Content.View()) + "\n" +
		view.Render(m.FooterView())
}

func (m *ModelHome) GetSelected() string {
	return m.home.Selected
}
