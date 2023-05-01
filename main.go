package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/ui"
)

const useHighPerformanceRenderer = false

type model struct {
	home ui.ModelHome
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.home, cmd = m.home.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.home.View()
}

func main() {
	content, err := os.ReadFile("/tmp/text.md")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	m := model{home: ui.NewHome(entity.Command{Content: string(content)})}
	p := tea.NewProgram(
		&m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
