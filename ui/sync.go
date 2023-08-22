package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var choiceEntered bool = false

type tickMsg time.Time

func syncUpdate(msg tea.Msg, m ModelHome) (*ModelHome, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			m.home.StatusSyncScreen = true
			m.home.Viewport.SetContent(syncView(&m))
			return &m, nil
		case "right":
			m.home.StatusSyncScreen = false
			m.home.Viewport.SetContent(syncView(&m))
			return &m, nil
		case "enter":
			choiceEntered = true
			m.home.StatusSyncScreen = false
			return &m, syncTickCmd()
		case "q":
			if m.home.ActiveSyncScreen {
				m.home.ActiveSyncScreen = false
				m.home.ProgressSync = progress.NewModel(progress.WithDefaultGradient())
				m.home.Viewport.SetContent(m.GetDataView())
				choiceEntered = false
				return &m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height
		m.home.ProgressSync.Width = msg.Width - padding*2 - 4
		if m.home.ProgressSync.Width > maxWidth {
			m.home.ProgressSync.Width = maxWidth
		}
		m.home.Viewport.SetContent(syncView(&m))
	case tickMsg:
		// fposition := usecaseHistory.Execute()
		m.home.Viewport.SetContent(syncView(&m))
		return &m, cmd
	}
	return &m, cmd
}

func syncView(m *ModelHome) string {
	var okButton, cancelButton string

	if m.home.StatusSyncScreen && !choiceEntered {
		okButton = ActiveButtonStyle.Render("Yes")
		cancelButton = ButtonStyle.Render("No, take me back")
	} else if choiceEntered {
		okButton = ButtonDisableStyle.Render("Yes")
		cancelButton = ButtonDisableStyle.Render("No, take me back")
	} else {
		okButton = ButtonStyle.Render("Yes")
		cancelButton = ActiveButtonStyle.Render("No, take me back")
	}

	question := lipgloss.NewStyle().
		Width(m.home.Viewport.Width / 2).
		Align(lipgloss.Center).
		Render("Are you sure to sync")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(
		(m.home.Viewport.Width / 2),
		(m.home.Viewport.Height - 50),
		lipgloss.Left, lipgloss.Center,
		DialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(SubtleStyle),
	)

	return BaseStyle.MarginLeft(m.home.Viewport.Width/4).
		PaddingTop((m.home.Viewport.Height / 3)).
		Render(dialog+"\n\n", syncProgressView(m))
}

func syncProgressView(m *ModelHome) string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + "--------" + "\n"
}

func syncTickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
