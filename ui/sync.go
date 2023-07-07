package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func syncUpdate(msg tea.Msg, m ModelHome) (*ModelHome, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.home.ActiveSyncScreen = true
			m.home.Viewport.SetContent(syncView(m))
			return &m, cmd
		case "right":
			m.home.ActiveSyncScreen = false
			m.home.Viewport.SetContent(syncView(m))
			return &m, cmd
		case "enter":
			m.home.StatusSyncScreen = false
			m.home.Viewport.SetContent(m.GetDataView())
			return &m, cmd
		}
	case tea.WindowSizeMsg:
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height
		m.home.Viewport.SetContent(syncView(m))
	}
	return &m, cmd
}

func syncView(m ModelHome) string {
	var okButton, cancelButton string

	if m.home.ActiveSyncScreen {
		okButton = ActiveButtonStyle.Render("Yes")
		cancelButton = ButtonStyle.Render("No, take me back")
	} else {
		okButton = ButtonStyle.Render("Yes")
		cancelButton = ActiveButtonStyle.Render("No, take me back")
	}

	question := lipgloss.NewStyle().
		Width(m.home.Viewport.Width - 50).
		Align(lipgloss.Center).
		Render("Are you sure you want to sync")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(
		(m.home.Viewport.Width - 50),
		(m.home.Viewport.Height - 50),
		lipgloss.Left, lipgloss.Center,
		DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(SubtleStyle),
	)
	return BaseStyle.PaddingLeft(20).PaddingTop((m.home.Viewport.Height / 2)).Render(dialog + "\n\n")
}
