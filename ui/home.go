package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
)

const useHighPerformanceRenderer = false

type IHome interface {
	HeaderView() string
	FooterView() string
	WindowUpdate(msg *entity.Command)
	Update(m tea.Msg) (*entity.Command, tea.Cmd)
}

type ModelHome struct {
	home entity.Command
}

func NewHome(m entity.Command) ModelHome {
	p := ModelHome{
		home: entity.Command{
			Content:  m.Content,
			Cursor:   0,
			Selected: map[int]struct{}{},
			Ready:    false,
			Viewport: viewport.Model{},
		},
	}
	return p
}

func (m ModelHome) WindowUpdate(msg entity.Command) {
	m.home = msg
}

func (m ModelHome) HeaderView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) FooterView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) Update(msg tea.Msg) (ModelHome, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.HeaderView())
		footerHeight := lipgloss.Height(m.FooterView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.home.Ready {
			m.home.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.home.Viewport.YPosition = headerHeight
			m.home.Viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.home.Viewport.SetContent(m.home.Content)
			m.home.Ready = true

			m.home.Viewport.YPosition = headerHeight + 1
		} else {
			m.home.Viewport.Width = msg.Width
			m.home.Viewport.Height = msg.Height - verticalMarginHeight
		}
		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.home.Viewport))
		}
	}

	m.WindowUpdate(m.home)

	m.home.Viewport, cmd = m.home.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	NewHome(m.home)

	return m, tea.Batch(cmds...)
}

func (m ModelHome) View() string {
	if !m.home.Ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.HeaderView(), m.home.Viewport.View(), m.FooterView())
}
