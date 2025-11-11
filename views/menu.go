package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type MenuModel struct {
	options      []item
	activeOption int
	style        style
}

type style struct {
	width     int
	height    int
	primary   lipgloss.Color
	secondary lipgloss.Color
}

func NewMenuModel() MenuModel {
	options := []item{
		{title: "test   ", programStatus: TypingView},
		{title: "stats  ", programStatus: StatsView},
		{title: "options", programStatus: OptionsView},
	}

	return MenuModel{options: options, activeOption: 0, style: style{
		primary:   lipgloss.Color("#7aa2f7"),
		secondary: lipgloss.Color("#4fd6be"),
	}}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			cmd = SwitchProgramStatusCmd(m.options[m.activeOption].programStatus)
		}
	case tea.WindowSizeMsg:
		m.style.height = msg.Height
		m.style.width = msg.Width
	}
	return m, cmd
}

func (m MenuModel) View() string {
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

	return strings.Join(lines, "\n")
}
