package views

import tea "github.com/charmbracelet/bubbletea"

type StatsModel struct {
}

func NewStatsModel() StatsModel {
	return StatsModel{}
}

func (m StatsModel) Init() tea.Cmd {
	return nil
}

func (m StatsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, SwitchProgramStatusCmd(MenuView)
		}
	}
	return m, nil
}

func (m StatsModel) View() string {
	return "WONDERFUL STATISTICS"
}
