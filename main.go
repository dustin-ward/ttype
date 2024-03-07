package main

import (
	"flag"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/app"
	"github.com/dustin-ward/termtyping/internal/config"
	"github.com/muesli/termenv"
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

	// Ensure colour terminal is available
	colour_profile := termenv.ColorProfile()
	if colour_profile > termenv.ANSI {
		fmt.Println("WARNING: no colour terminal was detected. This application depends on colour support.\nPress 'enter' to continue...")
		fmt.Scanln()
	}

	p := tea.NewProgram(app.NewAppModel(app.StateDefault, nil))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
