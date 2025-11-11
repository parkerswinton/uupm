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
	return m, nil
}

func (m StatsModel) View() string {
	return "WONDERFUL STATISTICS"
}
