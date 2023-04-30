package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
)

type IHome interface {
	HeaderView() string
	FooterView() string
	WindowUpdate(msg *entity.Command)
}

type modelHome struct {
	home *entity.Command
}

func NewHome(m *entity.Command) IHome {
	return &modelHome{
		home: m,
	}
}

func (m *modelHome) WindowUpdate(msg *entity.Command) {
	m.home = msg
}

func (m modelHome) HeaderView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m modelHome) FooterView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}
