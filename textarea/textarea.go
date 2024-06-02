package textarea

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	stopwatch stopwatch.Model
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
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
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
	var cmd tea.Cmd

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
			if !m.stopwatch.Running() {
				return m, m.stopwatch.Start()
			}
			if len(m.correct)+len(m.incorrect) == len(m.target) {
				return m, m.stopwatch.Stop()
			}
		}
	}
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	s := m.stopwatch.View() + "\n"
	s += lipgloss.StyleRunes(m.target[:m.cursor], m.correct, m.style.Correct, m.style.Incorrect) + m.target[m.cursor:]
	return s + fmt.Sprintf(" %dwpm", int(m.calcWPM()))
}
