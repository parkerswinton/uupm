package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProgramStatus int

const (
	testView ProgramStatus = iota
	statsView
)

const banner = ` _   _ _   _ ____  ____
| | | | | | |  _ \|    \ 
| |_| | |_| | |_| | | | |
|____/|____/|  __/|_|_|_|
            |_|          `

type item struct {
	title         string
	programStatus ProgramStatus
}

type model struct {
	options      []item
	activeOption int
	selected     item
	style        style
}

type style struct {
	base   lipgloss.Style
	active lipgloss.Style
}

func New() model {
	options := []item{
		{title: "test", programStatus: testView},
		{title: "stats", programStatus: statsView},
	}

	return model{options: options, activeOption: 0, style: style{
		base:   lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color("#7aa2f7")),
		active: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#4fd6be")),
	}}
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
	lines := []string{m.style.base.Render(banner)}
	for i, option := range m.options {
		if i == m.activeOption {
			lines = append(lines, m.style.active.Render("> "+option.title))
		} else {
			lines = append(lines, m.style.base.Render(option.title))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
