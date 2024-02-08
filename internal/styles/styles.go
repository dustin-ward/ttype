package styles

import "github.com/charmbracelet/lipgloss"

var (
	HiddenText = lipgloss.NewStyle().Foreground(Grey)
	ActiveText = lipgloss.NewStyle().Foreground(White)
	WrongText  = lipgloss.NewStyle().Foreground(Red)
)
