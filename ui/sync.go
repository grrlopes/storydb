package ui

import (
	"fmt"
	"time"

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
			if m.home.StatusSyncScreen {
				choiceEntered = true
				cmd = syncTickCmd()
			} else {
				choiceEntered = false
			}
			m.home.StatusSyncScreen = false
			return &m, cmd
		case "q":
			if m.home.ActiveSyncScreen {
				m.home.ActiveSyncScreen = false
				m.home.Viewport.SetContent(m.GetDataView())
				choiceEntered = false
				return &m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height
		m.home.Viewport.SetContent(syncView(&m))
	case tickMsg:
		cmd = m.home.Spinner.Tick
		m.home.Viewport.SetContent(syncView(&m))
		return &m, tea.Batch(syncTickCmd(), cmd)
	}
	m.home.Spinner, cmd = m.home.Spinner.Update(msg)
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
	} else if !choiceEntered {
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
	sync := fmt.Sprintf("%s Syncing .....", m.home.Spinner.View())
	return sync
}

func syncTickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
