package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/app"
)

func main() {
	p := tea.NewProgram(app.NewAppModel(app.StateDefault))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
