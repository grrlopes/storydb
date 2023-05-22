package ui

import "github.com/charmbracelet/lipgloss"

var (
	view     = lipgloss.NewStyle()
	content  = lipgloss.NewStyle()
	winSize = lipgloss.NewStyle().
			Margin(1, 2)

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
