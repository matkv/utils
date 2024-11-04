package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")). // A darker color for the border
			Padding(1, 2)

	countStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Margin(1, 0).
			Background(lipgloss.Color("63")).
			Foreground(lipgloss.Color("230")).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("63"))
)

type Model struct {
	count int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.count++
		case "down", "j":
			m.count--
		}
	}
	return m, nil
}

func (m Model) View() string {
	title := titleStyle.Render("Counter")
	count := countStyle.Render(fmt.Sprintf("Count: %d", m.count))

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		title,
		count,
		"Press up/down to increase/decrease. Press q to quit.",
	)
}
