package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/ui"
)

const useHighPerformanceRenderer = false

var (
	command          = entity.NewCmd()
	home    ui.IHome = ui.NewHome(command)
)

type model struct {
	home *entity.Command
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(home.HeaderView())
		footerHeight := lipgloss.Height(home.FooterView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.home.Ready {
			m.home.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.home.Viewport.YPosition = headerHeight
			m.home.Viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.home.Viewport.SetContent(m.home.Content)
			m.home.Ready = true

			m.home.Viewport.YPosition = headerHeight + 1
		} else {
			m.home.Viewport.Width = msg.Width
			m.home.Viewport.Height = msg.Height - verticalMarginHeight
		}
		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.home.Viewport))
		}
	}

	home.WindowUpdate(m.home)

	m.home.Viewport, cmd = m.home.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	ui.NewHome(m.home)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.home.Ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", home.HeaderView(), m.home.Viewport.View(), home.FooterView())
}

func main() {
	content, err := os.ReadFile("/tmp/text.md")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}
	m := model{home: entity.NewCmd()}
	m.home.Content = string(content)
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
