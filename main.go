package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/parkerswinton/uupm/views"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("fatal: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(views.NewApp(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERRORS: %v", err)
		os.Exit(1)
	}
}
