package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type OptionsModel struct {
}

func NewOptionsModel() OptionsModel {
	return OptionsModel{}
}

func (m OptionsModel) Init() tea.Cmd {
	return nil
}

func (m OptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m OptionsModel) View() string {
	return "insert options here"
}
