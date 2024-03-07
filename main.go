package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/app"
	"github.com/dustin-ward/termtyping/internal/config"
	"github.com/muesli/termenv"
)

var VERSION string = "v0.3.0"

var zen_flag bool
var version_flag bool

func init() {
	flag.BoolVar(&version_flag, "version", false, "print version of the application")
	flag.BoolVar(&zen_flag, "zen", false, "enable zen mode (default=false)")
	flag.Parse()
}

func main() {
	if version_flag {
		fmt.Println(VERSION)
		os.Exit(0)
	}

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
