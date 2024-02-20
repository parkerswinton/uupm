package textarea

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	target    string
	correct   []int
	incorrect []int
	cursor    int
	style     Style
}

type Style struct {
	Base      lipgloss.Style
	Correct   lipgloss.Style
	Incorrect lipgloss.Style
	Layout    lipgloss.Style
}

func DefaultStyles() Style {
	return Style{
		Base:      lipgloss.NewStyle().Bold(false),
		Correct:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("36")),
		Incorrect: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("52")),
		Layout: lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("36")).
			BorderStyle(lipgloss.NormalBorder()).
			Padding(1).
			Width(80),
	}
}

func New() Model {

	return Model{
		target:    "large seem give nation number think down part head one which early find possible like",
		correct:   []int{},
		incorrect: []int{},
		cursor:    0,
		style:     DefaultStyles(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if m.cursor > 0 {
				m.cursor--
				m.removePrevious()
			}
		default:
			m.addCurrent(msg.String())
			m.cursor++
			if len(m.correct)+len(m.incorrect) == len(m.target) {
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m *Model) removePrevious() {
	if len(m.correct) > 0 && m.correct[len(m.correct)-1] == m.cursor {
		m.correct = m.correct[:len(m.correct)-1]
	} else {
		m.incorrect = m.incorrect[:len(m.incorrect)-1]
	}
}

func (m *Model) addCurrent(msg string) {
	if msg == string(m.target[m.cursor]) {
		m.correct = append(m.correct, m.cursor)
	} else {
		m.incorrect = append(m.incorrect, m.cursor)
	}
}

func (m Model) View() string {
	return lipgloss.StyleRunes(m.target[:m.cursor], m.correct, m.style.Correct, m.style.Incorrect) + m.target[m.cursor:]
}
