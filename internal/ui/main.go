package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	count int
	list  list.Model
}

func NewModel() Model {
	items := make([]list.Item, len(setUpOptions()))
	for i, option := range setUpOptions() {
		items[i] = option
	}

	// Create a new list model and set the items and title
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Options"

	return Model{
		list: l,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.list.View())
}

func setUpOptions() []item {
	return []item{
		{"Pull latest dotfiles", "Pull the latest dotfiles from the remote repository"},
		{"Sync dotfiles", "Move the config files to the correct locations"},
		{"Update book reviews", "Update the book reviews in the Hugo site"},
	}
}
