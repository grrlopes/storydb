package helper

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	Sync     key.Binding
	Enter    key.Binding
	Finder   key.Binding
	Back     key.Binding
	Quit     key.Binding
	PageNext key.Binding
	PagePrev key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit, k.Back, k.Enter, k.Sync,
		k.Finder, k.PagePrev, k.PageNext,
	}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.Enter, k.Quit},
		{k.Sync, k.Finder, k.PageNext, k.PagePrev},
	}
}

var HotKeysHome = Keymap{
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
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	PageNext: key.NewBinding(
		key.WithKeys("ctrl+d", "Next"),
		key.WithHelp("ctrl+d", "Next Page"),
	),
	PagePrev: key.NewBinding(
		key.WithKeys("ctrl+s", "Prev"),
		key.WithHelp("ctrl+s", "Prev Page"),
	),
}

var HotKeysFinder = Keymap{
	Enter: HotKeysHome.Enter,
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c", "quit"),
	),
	PageNext: HotKeysHome.PageNext,
	PagePrev: HotKeysHome.PagePrev,
}
