package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/helper"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/ui"
	"github.com/grrlopes/storydb/usecase/schema"
)

var (
	repositoryGorm     repositories.ISqliteRepository = sqlite.NewGormRepostory()
	usecaseMigrateGorm schema.InputBoundary           = schema.NewMigrate(repositoryGorm)
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
	usecaseMigrateGorm.Execute()

	m := model{
		home: ui.NewHome(
			&entity.CmdModel{},
		),
	}

	p := tea.NewProgram(
		&m,
		tea.WithAltScreen(),
	)

	_, err := p.Run()

	if err != nil {
		fmt.Println("could not run program:", err)
	}

	env := os.Getenv("storydb")
	if env == "" {
		log.Fatalf("%s %s", "Error", helper.ErrEnvFailed)
	}

	cmd := m.home.GetSelected()

	helper.Echobuffer(cmd, env)
}
