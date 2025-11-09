package main

import (
	"log"

	"github.com/GoKells/tide/internal/core"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(core.NewApp(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
