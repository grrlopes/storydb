package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
)

type IHeader interface {
	HeaderView() string
	FooterView() string
}

type modelHeader struct {
	entity.Command
}

func Header(m entity.Command) IHeader {
	return &modelHeader{}
}

func (m *modelHeader) HeaderView() string {
	line := strings.Repeat("─", Max(0, m.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m *modelHeader) FooterView() string {
	line := strings.Repeat("─", Max(0, m.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}
