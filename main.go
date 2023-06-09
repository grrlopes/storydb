package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/ui"
	"github.com/grrlopes/storydb/usecase/listall"
	"github.com/grrlopes/storydb/usecase/schema"
)

var (
	repository     repositories.ISqliteRepository = sqlite.NewSQLiteRepository()
	usecaseMigrate schema.InputBoundary           = schema.NewMigrate(repository)
	usecaseAll     listall.InputBoundary          = listall.NewListAll(repository)
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
	m := model{
		home: ui.NewHome(
			&entity.Command{},
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
	fmt.Printf("%+v\n\n", " --")
}
