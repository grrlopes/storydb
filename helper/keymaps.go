package helper

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	SyncScreen  key.Binding
	Enter       key.Binding
	Finder      key.Binding
	Back        key.Binding
	Quit        key.Binding
	PageNext    key.Binding
	PagePrev    key.Binding
	ResetFinder key.Binding
	MoveUp      key.Binding
	MoveDown    key.Binding
	MoveLeft    key.Binding
	MoveRight   key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit, k.Back, k.Enter, k.SyncScreen,
		k.Finder, k.PagePrev, k.PageNext,
		k.ResetFinder, k.MoveUp, k.MoveDown,
    k.MoveLeft, k.MoveRight,
	}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.Enter, k.Quit, k.ResetFinder, k.MoveUp},
		{k.SyncScreen, k.Finder, k.PageNext, k.PagePrev, k.MoveDown},
    {k.MoveLeft, k.MoveRight},
	}
}

var HotKeysHome = Keymap{
	SyncScreen: key.NewBinding(
		key.WithKeys("alt+s"),
		key.WithHelp("alt+s", "sync"),
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
	MoveUp: key.NewBinding(
		key.WithKeys("shift+tab", "up", "k"),
		key.WithHelp("shift+tab/Up", "↑"),
	),
	MoveDown: key.NewBinding(
		key.WithKeys("tab", "down", "j"),
		key.WithHelp("tab/Down", "↓"),
	),
}

var HotKeysFinder = Keymap{
	Enter: HotKeysHome.Enter,
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	PageNext: HotKeysHome.PageNext,
	PagePrev: HotKeysHome.PagePrev,
	ResetFinder: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "Reset Finder"),
	),
	MoveUp:   HotKeysHome.MoveUp,
	MoveDown: HotKeysHome.MoveDown,
}

var HotKeysSync = Keymap{
	Enter: HotKeysHome.Enter,
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	MoveLeft: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", ""),
	),
	MoveRight: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", ""),
	),
}
