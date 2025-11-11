package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const Margin = 2

type ProgramStatus int

const (
	MenuView ProgramStatus = iota
	TypingView
	StatsView
	OptionsView
)

type SwitchProgramStatusMsg struct {
	target ProgramStatus
}

type model struct {
	child     tea.Model
	primary   lipgloss.Color
	secondary lipgloss.Color

	width  int
	height int
}

func NewApp() model {
	return model{
		child:     nil,
		primary:   lipgloss.Color("#7aa2f7"),
		secondary: lipgloss.Color("#4fd6be"),
	}
}

func (m model) Init() tea.Cmd {
	return SwitchProgramStatusCmd(MenuView)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, SwitchProgramStatusCmd(MenuView)
		}
	case SwitchProgramStatusMsg:
		m.child = m.createChild(msg.target)
		cmd = m.initChild()
		return m, cmd
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	if m.child != nil {
		m.child, cmd = m.child.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	container := lipgloss.NewStyle().
		Width(m.width-Margin).
		Height(m.height-Margin).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(m.primary)

	var innerContent = ""

	if m.child != nil {
		innerContent = m.child.View()
	}

	return container.Render(innerContent)
}

func SwitchProgramStatusCmd(target ProgramStatus) tea.Cmd {
	return func() tea.Msg {
		return SwitchProgramStatusMsg{
			target: target,
		}
	}
}

func (m model) createChild(target ProgramStatus) tea.Model {
	switch target {
	case MenuView:
		return NewMenuModel()
	case TypingView:
		return NewTypingModel()
	case StatsView:
		return NewStatsModel()
	case OptionsView:
		return NewOptionsModel()
	}
	return nil
}

func (m model) initChild() tea.Cmd {
	return m.child.Init()
}
