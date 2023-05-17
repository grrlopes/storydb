package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/ui"
)

type model struct {
	home *ui.ModelHome
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.home, cmd = m.home.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.home.View()
}

func main() {
	items := []list.Item{
		ui.NewListPanel{EnTitle: "hhh", Desc: "I have ’em all over my house"},
		ui.NewListPanel{EnTitle: "ls", Desc: "common list"},
		ui.NewListPanel{EnTitle: "ls -l", Desc: "list in lst hehe"},
	}

	data := list.New(items, list.NewDefaultDelegate(), 0, 0)

	m := model{
		home: ui.NewHome(
			entity.Command{
				Content: data,
			},
		),
	}

	p := tea.NewProgram(
		&m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	_, err := p.Run()

	if err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

	cmd := exec.Command(
		"xdotool",
		"type",
		m.home.GetSelected(),
	)
	_ = cmd.Run()
}
