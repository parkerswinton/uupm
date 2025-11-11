package views

import tea "github.com/charmbracelet/bubbletea"

type TypingModel struct {
}

func NewTypingModel() TypingModel {
	return TypingModel{}
}

func (m TypingModel) Init() tea.Cmd {
	return nil
}

func (m TypingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m TypingModel) View() string {
	return "Typing Typing"
}
