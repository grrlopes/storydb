package ui

type NewListPanel struct {
	EnTitle string
	Desc  string
}

func (list NewListPanel) Title() string     { return list.EnTitle }
func (list NewListPanel) Description() string { return list.Desc }
func (list NewListPanel) FilterValue() string { return list.EnTitle }
