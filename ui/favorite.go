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
	favoriteMsg      []entity.Commands
	favoriteCountMsg int
	favoritePagMsg   struct{}
)

func FavoriteCmd(filter textinput.Model, limit int, offset int) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = func() tea.Msg {
		data, _, _ := usecaseFinder.Execute(filter.Value(), limit, offset)
		return favoriteMsg(data)
	}

	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func favoriteCount(filter string) tea.Cmd {
	count := usecaseFinderCount.Execute(filter)
	return func() tea.Msg {
		return favoriteCountMsg(count)
	}
}

func favoriteFocused(msg tea.KeyMsg, m *entity.CmdModel) (entity.CmdModel, tea.Cmd) {
	switch {
	case key.Matches(msg, helper.HotKeysFavorite.Enter):
		m.RowChosen = m.Selected.Cmd
		return *m, tea.Quit
	case key.Matches(msg, helper.HotKeysFavorite.PageNext):
		m.Cursor = 0
	case key.Matches(msg, helper.HotKeysFavorite.ResetFinder):
		m.Favorite.Reset()
	case key.Matches(msg, helper.HotKeysFavorite.PagePrev):
		m.Cursor = 0
	case key.Matches(msg, helper.HotKeysFavorite.MoveDown):
		if m.Cursor < m.PageTotal-1 {
			m.Content = "arrow"
			m.Cursor++
		}
	case key.Matches(msg, helper.HotKeysFavorite.MoveUp):
		if m.Cursor > 0 {
			m.Content = "arrow"
			m.Cursor--
		}
	case key.Matches(msg, helper.HotKeysFavorite.Quit):
		m.Favorite.Reset()
		m.Favorite.Blur()
	}
	return *m, nil
}

func favoritePaginatorCmd(paginator paginator.Model, msg tea.Msg) (paginator.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	model, cmd := paginator.Update(msg)

	cmds = append(cmds, cmd)

	cmd = func() tea.Msg {
		return favoritePagMsg{}
	}

	cmds = append(cmds, cmd)

	return model, tea.Batch(cmds...)
}

func favoriteInsert(id uint) string {
	result := usecaseAddFavorite.Execute(id)
	return result
}
