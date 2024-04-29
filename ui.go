package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StartUI() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// stores the application's state

type utilsModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() utilsModel { // doesn't have to be an actual function
	return utilsModel{
		choices:  []string{"pull", "sync"},
		selected: map[int]struct{}{},
	}
}

func (m utilsModel) Init() tea.Cmd {
	return nil
}

// is called when "things happen" -> updates the model and returns it

func (m utilsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j", "down":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "k", "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

// renders the model to the screen
func (m utilsModel) View() string {
	s := "Please provide a valid command: pull or sync\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += cursor + " [" + checked + "] " + choice + "\n"
	}

	s += "\nPress 'q' to quit.\n"

	return s
}
