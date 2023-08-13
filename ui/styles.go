package ui

import "github.com/charmbracelet/lipgloss"

var (
	view        = lipgloss.NewStyle()
	BaseStyle   = lipgloss.NewStyle()
	SubtleStyle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

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
		b := view.Background(lipgloss.Color("#888B7E"))
		return b
	}()

	HelperStyle = func() lipgloss.Style {
		b := view.PaddingTop(1)
		return b
	}()

	DialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#b7cbbf")).
			Padding(1, 2).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	ButtonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1).
			MarginRight(2)

	ButtonDisableStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#7a0609")).
				Padding(0, 3).
				MarginTop(1).
				MarginRight(2)

	ActiveButtonStyle = ButtonStyle.Copy().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.AdaptiveColor{Light: "#b7cbbf", Dark: "#b7cbbf"}).
				MarginRight(2).
				Underline(true)
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
