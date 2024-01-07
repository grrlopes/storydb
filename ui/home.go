package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/grrlopes/storydb/entity"
	"github.com/grrlopes/storydb/helper"
	"github.com/grrlopes/storydb/repositories"
	"github.com/grrlopes/storydb/repositories/fileparse"
	"github.com/grrlopes/storydb/repositories/sqlite"
	"github.com/grrlopes/storydb/usecase/count"
	"github.com/grrlopes/storydb/usecase/favorite"
	"github.com/grrlopes/storydb/usecase/favoriteFilter"
	"github.com/grrlopes/storydb/usecase/fhistory"
	"github.com/grrlopes/storydb/usecase/finder"
	findercount "github.com/grrlopes/storydb/usecase/finderCount"
	"github.com/grrlopes/storydb/usecase/listall"
	"github.com/grrlopes/storydb/usecase/pager"
)

var (
	repositoryGorm     repositories.ISqliteRepository     = sqlite.NewGormRepostory()
	frepository        repositories.IFileParsedRepository = fileparse.NewFparsedRepository()
	usecasePager       pager.InputBoundary                = pager.NewPager(repositoryGorm)
	usecaseCount       count.InputBoundary                = count.NewCount(repositoryGorm)
	usecaseHistory     fhistory.InputBoundary             = fhistory.NewFHistory(frepository, repositoryGorm)
	usecaseFinder      finder.InputBoundary               = finder.NewFinder(repositoryGorm)
	usecaseFinderCount findercount.InputBoundary          = findercount.NewFinderCount(repositoryGorm)
	usecaseAll         listall.InputBoundary              = listall.NewListAll(repositoryGorm)
	usecaseAddFavorite favorite.InputBoundary             = favorite.NewFavorite(repositoryGorm)
	usecaseFilterFav   favoriteFilter.InputBoundary       = favoriteFilter.NewFavoriteFilter(repositoryGorm)
)

type ModelHome struct {
	home entity.CmdModel
}

func NewHome(m *entity.CmdModel) *ModelHome {
	count := usecaseCount.Execute()
	p := paginator.New()
	p.SetTotalPages(count)
	p.KeyMap.NextPage = helper.HotKeysHome.PageNext
	p.KeyMap.PrevPage = helper.HotKeysHome.PagePrev
	txt := textinput.New()
	txt.Placeholder = "type..."
	txt.CharLimit = 156
	txt.Width = 50
	txt.Prompt = "Finder: "
	txtFav := textinput.New()
	txtFav.Placeholder = "type..."
	txtFav.CharLimit = 156
	txtFav.Width = 50
	txtFav.Prompt = "Favorite: "

	h := help.New()
	spin := spinner.New()
	spin.Spinner = spinner.Monkey

	home := ModelHome{
		home: entity.CmdModel{
			Content:          m.Content,
			Ready:            false,
			Viewport:         viewport.Model{},
			PageTotal:        m.PageTotal,
			Pagination:       &p,
			Count:            &count,
			ActiveSyncScreen: false,
			StatusSyncScreen: false,
			FavoriteScreen:   false,
			Finder:           txt,
			Favorite:         txtFav,
			HomeKeys:         helper.HotKeysHome,
			FinderKeys:       helper.HotKeysFinder,
			SyncKeys:         helper.HotKeysSync,
			FavoriteKeys:     helper.HotKeysFavorite,
			Help:             h,
			Spinner:          spin,
		},
	}
	return &home
}

func (m ModelHome) HeaderView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) FooterView() string {
	line := strings.Repeat("─", Max(0, m.home.Viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func (m ModelHome) Update(msg tea.Msg) (*ModelHome, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if m.home.ActiveSyncScreen {
		m.home.Ready = true
		synced, cmd := syncUpdate(msg, m)
		return synced, cmd
	}

	switch msg := msg.(type) {
	case finderCountMsg:
		if *m.home.Count != int(msg) {
			m.home.Pagination.Page = 0
			m.home.Cursor = 0
		}
		if int(msg) == 0 {
			m.home.Pagination.SetTotalPages(1)
		} else {
			*m.home.Count = int(msg)
			m.home.Pagination.SetTotalPages(*m.home.Count)
		}
	case finderMsg:
		m.home.Store = msg
		m.home.PageTotal = len(msg)
		m.home.Viewport.Update(msg)
		m.home.Viewport.SetContent(m.GetDataView())
	case favoriteCountMsg:
		if *m.home.Count != int(msg) {
			m.home.Pagination.Page = 0
			m.home.Cursor = 0
		}
		if int(msg) == 0 {
			m.home.Pagination.SetTotalPages(1)
		} else {
			*m.home.Count = int(msg)
			m.home.Pagination.SetTotalPages(*m.home.Count)
		}
	case favoriteMsg:
		m.home.Store = msg
		m.home.PageTotal = len(msg)
		m.home.Viewport.Update(msg)
		m.home.Viewport.SetContent(m.GetDataView())
	case tea.KeyMsg:
		m.home.Warning.Active = false
		if m.home.Favorite.Focused() {
			m.home, cmd = favoriteFocused(msg, &m.home)
			*m.home.Pagination, cmd = favoritePaginatorCmd(*m.home.Pagination, msg)
			cmds = append(cmds, cmd)
			m.home.Favorite, cmd = m.home.Favorite.Update(msg)
			cmd = favoriteCount(m.home.Favorite.Value())
			cmds = append(cmds, cmd)
			m.home.Start, m.home.End = m.updatepagination()
			cmd = FavoriteCmd(m.home.Favorite, m.home.Viewport.Height-2, m.home.Start)
			cmds = append(cmds, cmd)

		} else if m.home.Finder.Focused() {
			m.home, cmd = finderFocused(msg, &m.home)
			*m.home.Pagination, cmd = finderPaginatorCmd(*m.home.Pagination, msg)
			cmds = append(cmds, cmd)
			m.home.Finder, cmd = m.home.Finder.Update(msg)
			cmd = finderCount(m.home.Finder.Value())
			cmds = append(cmds, cmd)
			m.home.Start, m.home.End = m.updatepagination()
			cmd = finderCmd(m.home.Finder, m.home.Viewport.Height-2, m.home.Start)
			cmds = append(cmds, cmd)

		} else {
			switch {
			case key.Matches(msg, helper.HotKeysHome.Quit):
				return &m, tea.Quit
			case key.Matches(msg, helper.HotKeysHome.MoveUp):
				if m.home.Cursor > 0 {
					m.home.Content = "arrow"
					m.home.Cursor--
				}
			case key.Matches(msg, helper.HotKeysHome.MoveDown):
				if m.home.Cursor < m.home.PageTotal-1 {
					m.home.Content = "arrow"
					m.home.Cursor++
				}
			case key.Matches(msg, helper.HotKeysHome.SyncScreen):
				m.home.StatusSyncScreen = false
				m.home.ActiveSyncScreen = true
				m.home.Viewport.SetContent(syncView(&m))
				return &m, cmd
			case key.Matches(msg, helper.HotKeysHome.Enter):
				m.home.RowChosen = m.home.Selected.Cmd
				return &m, tea.Quit
			case key.Matches(msg, helper.HotKeysHome.Finder):
				m.home.Finder.Focus()
				m.home.Pagination.Page = 0
				m.home.Cursor = 0
				cmd = finderCmd(m.home.Finder, m.home.Viewport.Height-2, 1)
				cmds = append(cmds, cmd)
			case key.Matches(msg, helper.HotKeysHome.Favorite):
				m.home.Favorite.Focus()
				cmd = FavoriteCmd(m.home.Favorite, m.home.Viewport.Height-2, 1)
				cmds = append(cmds, cmd)
			case key.Matches(msg, helper.HotKeysHome.PageNext):
				m.home.Cursor = 0
			case key.Matches(msg, helper.HotKeysHome.AddFav):
				favoriteInsert(m.home.Selected.ID, &m.home)
			case key.Matches(msg, helper.HotKeysHome.PagePrev):
				m.home.Cursor = 0
			}
			*m.home.Pagination, cmd = m.home.Pagination.Update(msg)
			cmds = append(cmds, cmd)
			start, end := m.updatepagination()
			m.home.Start = start
			m.home.End = end

			m.home.Store, _ = usecasePager.Execute(m.home.Viewport.Height-2, m.home.Start)
			m.home.PageTotal = len(m.home.Store)

			m.home.Viewport.Update(msg)
			m.home.Viewport.SetContent(m.GetDataView())
		}
	case tea.WindowSizeMsg:
		m.home.Content = "window"
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height - 6
		m.home.Pagination.PerPage = msg.Height - 2
		m.initStore()
		m.home.Ready = true
		m.home.Help.Width = msg.Width
		m.home.Viewport.Update(msg)
		m.home.Viewport.SetContent(m.GetDataView())

	}
	return &m, tea.Batch(cmds...)
}

func (m ModelHome) View() string {
	view := lipgloss.NewStyle()
	content := lipgloss.NewStyle()
	if !m.home.Ready {
		return "\n  Loading..."
	}
	if m.home.Finder.Focused() {
		if m.home.Warning.Active {
			return view.Render(m.HeaderView()) + "\n" +
				m.home.Finder.View() +
				content.Render(m.home.Viewport.View()) + "\n" +
				m.FooterView() + "\n" +
				m.paginationView() + " " +
				Warning(
					m.home.Viewport.Width/4,
					m.home.Warning.Color,
				).Render(m.home.Warning.Message) + "\n" +
				HelperStyle.Render(m.finderKeysView())
		}
		return view.Render(m.HeaderView()) + "\n" +
			m.home.Finder.View() +
			content.Render(m.home.Viewport.View()) + "\n" +
			m.FooterView() + "\n" +
			m.paginationView() + "\n" +
			HelperStyle.Render(m.finderKeysView())
	}
	if m.home.Favorite.Focused() {
		return view.Render(m.HeaderView()) + "\n" +
			m.home.Favorite.View() +
			content.Render(m.home.Viewport.View()) + "\n" +
			m.FooterView() + "\n" +
			m.paginationView() + "\n" +
			HelperStyle.Render(m.FavoriteKeysView())
	}
	if m.home.ActiveSyncScreen {
		return view.Render(m.HeaderView()) + "\n" +
			content.Render(m.home.Viewport.View()) + "\n" +
			m.FooterView() + "\n" +
			HelperStyle.Render(m.SyncKeysView())
	}

	return view.Render(m.HeaderView()) + "\n" +
		content.Render(m.home.Viewport.View()) + "\n" +
		m.FooterView() + "\n" +
		m.paginationView() + "\n" +
		HelperStyle.Render(m.homeKeysView())
}

func (m *ModelHome) GetSelected() string {
	return m.home.RowChosen
}

func (m *ModelHome) updatepagination() (int, int) {
	start, end := m.home.Pagination.GetSliceBounds(*m.home.Count)
	return start, end
}

func (m *ModelHome) GetDataView() string {
	var (
		maxLen = m.home.Viewport.Width
		result []string
	)

	for i, v := range m.home.Store {
		if m.home.Cursor == i {
			m.home.Selected = v
			v.Cmd = SelecRow.Render("->", v.Cmd)
		}

		if len(v.Cmd) > maxLen {
			title := ShrinkWordMiddle(v.Cmd, maxLen)
			v.Cmd = title
		}

		result = append(result, fmt.Sprintf("\n%s", v.Cmd))
	}

	rowData := strings.Trim(fmt.Sprintf("%s", result), "[]")

	return rowData
}

func (m *ModelHome) initStore() {
	m.home.Store, _ = usecasePager.Execute(m.home.Viewport.Height-2, m.home.Start)
	m.home.PageTotal = len(m.home.Store)
	m.home.Pagination.SetTotalPages(*m.home.Count)
}

func (m *ModelHome) paginationView() string {
	var b strings.Builder
	b.WriteString("  " + m.home.Pagination.View())
	return b.String()
}

func (m ModelHome) homeKeysView() string {
	return m.home.Help.View(m.home.HomeKeys)
}

func (m ModelHome) finderKeysView() string {
	return m.home.Help.View(m.home.FinderKeys)
}

func (m ModelHome) SyncKeysView() string {
	return m.home.Help.View(m.home.SyncKeys)
}

func (m ModelHome) FavoriteKeysView() string {
	return m.home.Help.View(m.home.FavoriteKeys)
}
