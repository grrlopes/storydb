package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/progress"
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
	"github.com/grrlopes/storydb/usecase/fhistory"
	"github.com/grrlopes/storydb/usecase/fhistorytotal"
	"github.com/grrlopes/storydb/usecase/finder"
	findercount "github.com/grrlopes/storydb/usecase/finderCount"
	"github.com/grrlopes/storydb/usecase/listall"
	"github.com/grrlopes/storydb/usecase/pager"
)

var (
	repositoryGorm      repositories.ISqliteRepository     = sqlite.NewGormRepostory()
	frepository         repositories.IFileParsedRepository = fileparse.NewFparsedRepository()
	usecasePager        pager.InputBoundary                = pager.NewPager(repositoryGorm)
	usecaseCount        count.InputBoundary                = count.NewCount(repositoryGorm)
	usecaseHistory      fhistory.InputBoundary             = fhistory.NewFHistory(frepository, repositoryGorm)
	usecaseHistoryTotal fhistorytotal.InputBoundary        = fhistorytotal.NewFHistoryTotal(frepository, repositoryGorm)
	usecaseFinder       finder.InputBoundary               = finder.NewFinder(repositoryGorm)
	usecaseFinderCount  findercount.InputBoundary          = findercount.NewFinderCount(repositoryGorm)
	usecaseAll          listall.InputBoundary              = listall.NewListAll(repositoryGorm)
)

type ModelHome struct {
	home entity.CmdModel
}

func NewHome(m *entity.CmdModel) *ModelHome {
	count := usecaseCount.Execute()
	ftotal := usecaseHistoryTotal.Execute()
	p := paginator.New()
	p.SetTotalPages(count)
	p.KeyMap.NextPage = helper.HotKeysHome.PageNext
	p.KeyMap.PrevPage = helper.HotKeysHome.PagePrev
	pro := progress.New(progress.WithDefaultGradient())
	txt := textinput.New()
	txt.Placeholder = "type..."
	txt.CharLimit = 156
	txt.Width = 50
	txt.Prompt = "Finder: "
	h := help.New()

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
			ProgressSync:     pro,
			Ftotal:           ftotal,
			Finder:           txt,
			HomeKeys:         helper.HotKeysHome,
			FinderKeys:       helper.HotKeysFinder,
			Help:             h,
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
		}
		*m.home.Count = int(msg)
	case finderMsg:
		m.home.Store = msg
		m.home.PageTotal = len(msg)
	case tea.KeyMsg:
		if m.home.Finder.Focused() {
			switch msg.String() {
			case "up", "k", "shift+tab":
				if m.home.Cursor > 0 {
					m.home.Content = "arrow"
					m.home.Cursor--
				}
			case "down", "j", "tab":
				if m.home.Cursor < m.home.PageTotal-1 {
					m.home.Content = "arrow"
					m.home.Cursor++
				}
			case "enter":
				return &m, tea.Quit
			case "ctrl+r":
				m.home.Finder.Reset()
			}
			if msg.String() == "ctrl+c" {
				m.home.Finder.Reset()
				m.home.Finder.Blur()
			}
			*m.home.Pagination, cmd = m.home.Pagination.Update(msg)
			cmds = append(cmds, cmd)
			m.home.Finder, cmd = m.home.Finder.Update(msg)
			cmd = finderCount(m.home.Finder.Value())
			cmds = append(cmds, cmd)
			m.home.Start, m.home.End = m.updatepagination()
			cmd = finderCmd(m.home.Finder, m.home.Viewport.Height-2, m.home.Start)
			cmds = append(cmds, cmd)
		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				return &m, tea.Quit
			case "up", "k", "shift+tab":
				if m.home.Cursor > 0 {
					m.home.Content = "arrow"
					m.home.Cursor--
				}
			case "down", "j", "tab":
				if m.home.Cursor < m.home.PageTotal-1 {
					m.home.Content = "arrow"
					m.home.Cursor++
				}
			case "ctrl+g":
				m.home.Cursor = m.home.PageTotal - 1
			case "s":
				m.home.StatusSyncScreen = false
				m.home.ActiveSyncScreen = true
				m.home.Viewport.SetContent(syncView(&m))
				return &m, cmd
			case "ctrl+u":
				m.home.Cursor = 0
			case "enter":
				return &m, tea.Quit
			case "f":
				m.home.Finder.Focus()
				m.home.Pagination.Page = 0
			}
		}
	case tea.WindowSizeMsg:
		m.home.Content = "window"
		m.home.Viewport.Width = msg.Width
		m.home.Viewport.Height = msg.Height - 6
		m.home.Pagination.PerPage = msg.Height - 2
		m.home.Viewport.SetContent(m.GetDataView())
		m.home.Ready = true
		m.home.Help.Width = msg.Width
	}

	if !m.home.Finder.Focused() {
		*m.home.Pagination, cmd = m.home.Pagination.Update(msg)
		cmds = append(cmds, cmd)
		start, end := m.updatepagination()
		m.home.Start = start
		m.home.End = end
	}

	m.home.Viewport.Update(msg)
	m.home.Viewport.SetContent(m.GetDataView())
	return &m, tea.Batch(cmds...)
}

func (m ModelHome) View() string {
	view := lipgloss.NewStyle()
	content := lipgloss.NewStyle()
	if !m.home.Ready {
		return "\n  Loading..."
	}
	if m.home.Finder.Focused() {
		return view.Render(m.HeaderView()) + "\n" +
			m.home.Finder.View() +
			content.Render(m.home.Viewport.View()) + "\n" +
			m.FooterView() + "\n" +
			m.paginationView() + "\n" +
			HelperStyle.Render(m.finderKeysView())
	}
	return view.Render(m.HeaderView()) + "\n" +
		content.Render(m.home.Viewport.View()) + "\n" +
		m.FooterView() + "\n" +
		m.paginationView() + "\n" +
		HelperStyle.Render(m.homeKeysView())
}

func (m *ModelHome) GetSelected() string {
	return m.home.Selected
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

	if !m.home.Finder.Focused() {
		m.home.Store, _ = usecasePager.Execute(m.home.Viewport.Height-2, m.home.Start)
		m.home.PageTotal = len(m.home.Store)
	}
	m.home.Pagination.SetTotalPages(*m.home.Count)

	for i, v := range m.home.Store {
		if m.home.Cursor == i {
			m.home.Selected = v.Cmd
			v.Cmd = SelecRow.Render(v.Cmd)
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
