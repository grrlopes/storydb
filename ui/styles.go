package ui

import "github.com/charmbracelet/lipgloss"

var (
	view = lipgloss.NewStyle()

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

	SelecRow = func() lipgloss.Style {
		b := view.Background(lipgloss.Color("#00B377"))
		return b
	}()
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ShrinkWordMiddle(text string, maxLen int) string {
	var style = "..."

	halfLen := (maxLen - len(style)) / 2

	start := text[:halfLen]
	end := text[len(text)-halfLen:]

	return start + style + end
}
