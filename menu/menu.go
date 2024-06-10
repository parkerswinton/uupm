package menu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ProgramStatus int

const (
	testView ProgramStatus = iota
	statsView
)

type item struct {
	title         string
	description   string
	programStatus ProgramStatus
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type Model struct {
	list   list.Model
	choice string
}

func New() Model {
	items := []list.Item{
		item{title: "test", description: "test", programStatus: testView},
		item{title: "stats", description: "stats", programStatus: statsView},
	}

	l := list.New(items, list.NewDefaultDelegate(), 20, 20)
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	return Model{list: l}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i := m.list.SelectedItem()
			m.choice = i.FilterValue()
		}
	}
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View() + m.choice
}
