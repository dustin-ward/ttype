package main

import (
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/app"
	"github.com/dustin-ward/termtyping/internal/config"
)

var zen_flag bool

func init() {
	flag.BoolVar(&zen_flag, "zen", false, "enable zen mode (default=false)")
	flag.Parse()
}

func main() {
	config := config.Config{
		ZenMode: zen_flag,
	}
	app.Config = config

	p := tea.NewProgram(app.NewAppModel(app.StateDefault, nil))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
