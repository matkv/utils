package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
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
	return fmt.Sprintf("Count: %d\n\nPress up/down to increase/decrease. Press q to quit", m.count)
}
