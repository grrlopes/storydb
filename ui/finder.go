package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/helper"
)

type (
	finderMsg      []entity.Commands
	finderCountMsg int
	finderPagMsg   struct{}
)

func finderCmd(filter textinput.Model, limit int, offset int) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = func() tea.Msg {
		data, _, _ := usecaseFinder.Execute(filter.Value(), limit, offset)
		return finderMsg(data)
	}

	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func finderCount(filter string) tea.Cmd {
	count := usecaseFinderCount.Execute(filter)
	return func() tea.Msg {
		return finderCountMsg(count)
	}
}

func finderFocused(msg tea.KeyMsg, m *entity.CmdModel) (entity.CmdModel, tea.Cmd) {
	switch {
	case key.Matches(msg, helper.HotKeysFinder.Enter):
		m.RowChosen = m.Selected.Cmd
		return *m, tea.Quit
	case key.Matches(msg, helper.HotKeysFinder.PageNext):
		m.Cursor = 0
	case key.Matches(msg, helper.HotKeysFinder.ResetFinder):
		m.Finder.Reset()
	case key.Matches(msg, helper.HotKeysFinder.PagePrev):
		m.Cursor = 0
	case key.Matches(msg, helper.HotKeysFinder.AddFav):
		favoriteInsert(m.Selected.ID)
	case key.Matches(msg, helper.HotKeysFinder.MoveDown):
		if m.Cursor < m.PageTotal-1 {
			m.Content = "arrow"
			m.Cursor++
		}
	case key.Matches(msg, helper.HotKeysFinder.MoveUp):
		if m.Cursor > 0 {
			m.Content = "arrow"
			m.Cursor--
		}
	case key.Matches(msg, helper.HotKeysFinder.Quit):
		m.Finder.Reset()
		m.Finder.Blur()
	}

	return *m, nil
}

func finderPaginatorCmd(paginator paginator.Model, msg tea.Msg) (paginator.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	model, cmd := paginator.Update(msg)
	cmds = append(cmds, cmd)
	cmd = func() tea.Msg {
		return finderPagMsg{}
	}

	cmds = append(cmds, cmd)
	return model, tea.Batch(cmds...)
}
