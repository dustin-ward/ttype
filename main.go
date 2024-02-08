package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/app"
)

func main() {
	p := tea.NewProgram(app.NewAppModel(app.StateDefault), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
