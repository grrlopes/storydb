package helper

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	Sync   key.Binding
	Enter  key.Binding
	Finder key.Binding
	Back   key.Binding
	Quit   key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Back, k.Enter, k.Sync, k.Finder}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.Enter, k.Quit},
		{k.Sync, k.Finder},
	}
}

var HotKeys = Keymap{
	Sync: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "sync"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Finder: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "finder"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}
