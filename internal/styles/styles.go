package styles

import "github.com/charmbracelet/lipgloss"

var (
	HiddenText = lipgloss.NewStyle().Foreground(Grey)
	ActiveText = lipgloss.NewStyle().Foreground(White)
	WrongText  = lipgloss.NewStyle().Foreground(Red)

	TextBox = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Width(60).Padding(1, 2).Margin(0, 2).Foreground(White)
)
