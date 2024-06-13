package menu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ProgramStatus int

const (
	testView ProgramStatus = iota
	statsView
)

type item struct {
	title         string
	programStatus ProgramStatus
}

type model struct {
	options      []item
	activeOption int
	selected     item
}

func New() model {
	options := []item{
		{title: "test", programStatus: testView},
		{title: "stats", programStatus: statsView},
	}

	return model{options: options, activeOption: 0}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "down":
			if m.activeOption < len(m.options)-1 {
				m.activeOption++
			} else {
				m.activeOption = 0
			}
		case "k", "up":
			if m.activeOption > 0 {
				m.activeOption--
			} else {
				m.activeOption = len(m.options) - 1
			}
		case "enter":
			m.selected = m.options[m.activeOption]
		}
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	for i, option := range m.options {
		if i == m.activeOption {
			s += ">"
		}
		s += fmt.Sprintf("%s\n", option.title)
	}
	s += m.selected.title + fmt.Sprint(m.selected.programStatus)
	return s
}
