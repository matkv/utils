package ui

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// CommandItem represents a Cobra command in the TUI list.
type CommandItem struct {
	Name string
	Cmd  *cobra.Command
}

func (i CommandItem) Title() string       { return i.Name }
func (i CommandItem) Description() string { return i.Cmd.Short }
func (i CommandItem) FilterValue() string { return i.Title() }

type model struct {
	list      list.Model
	executing bool
	output    string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			if m.executing {
				m.executing = false
				m.output = ""
				return m, nil // Reset UI after viewing output
			}

			selectedItem := m.list.SelectedItem().(CommandItem)
			cmd := selectedItem.Cmd

			execCmd := exec.Command(os.Args[0], cmd.Use)
			output, err := execCmd.CombinedOutput()
			if err != nil {
				m.output = "Error executing command: " + err.Error()
			} else {
				m.output = string(output)
			}

			m.executing = true
			return m, nil
		}
	}

	if !m.executing {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.executing {
		return m.output
	}

	return m.list.View()
}

func RunTUI(rootCmd *cobra.Command) {
	var items []list.Item

	for _, cmd := range rootCmd.Commands() {
		if !cmd.Hidden {
			items = append(items, CommandItem{Name: cmd.Use, Cmd: cmd})
		}
	}

	cmdList := list.New(items, list.NewDefaultDelegate(), 50, 20)
	cmdList.Title = "Available Commands"
	cmdList.SetShowPagination(false)       // Disable pagination display
	cmdList.Paginator.PerPage = len(items) // Show all items on one page

	m := model{list: cmdList}
	p := tea.NewProgram(m)
	p.Run()
}
