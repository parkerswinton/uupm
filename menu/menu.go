package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProgramStatus int

const (
	testView ProgramStatus = iota
	statsView
	optionsView
)

const banner = ` _   _ _   _ ____  ____
| | | | | | |  _ \|    \ 
| |_| | |_| | |_| | | | |
|____/|____/|  __/|_|_|_|
            |_|          `

const Margin = 4

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
	width     int
	height    int
	primary   lipgloss.Color
	secondary lipgloss.Color
}

func New() model {
	options := []item{
		{title: "test", programStatus: testView},
		{title: "stats", programStatus: statsView},
		{title: "options", programStatus: optionsView},
	}

	return model{options: options, activeOption: 0, style: style{
		primary:   lipgloss.Color("#7aa2f7"),
		secondary: lipgloss.Color("#4fd6be"),
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
	case tea.WindowSizeMsg:
		m.style.height = msg.Height
		m.style.width = msg.Width
	}
	return m, cmd
}

func (m model) View() string {

	container := lipgloss.NewStyle().
		Width(m.style.width-Margin).
		Height(m.style.height-Margin).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(m.style.primary)

	content := lipgloss.NewStyle().Foreground(m.style.primary)
	hovered := lipgloss.NewStyle().Foreground(m.style.secondary).Bold(true)

	lines := []string{content.Render(banner)}
	for i, option := range m.options {
		if i == m.activeOption {
			lines = append(lines, hovered.Render("> "+option.title))
		} else {
			lines = append(lines, content.Render(option.title))
		}
	}

	return container.Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
}
