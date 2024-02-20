package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/parkerswinton/uupm/textarea"
)

func main() {
	p := tea.NewProgram(textarea.New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERRORS: %v", err)
		os.Exit(1)
	}
}
