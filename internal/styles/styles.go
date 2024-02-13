package styles

import "github.com/charmbracelet/lipgloss"

const (
	APP_WIDTH = 60
	MAX_LINES = 3
)

var (
	HiddenText    = lipgloss.NewStyle().Foreground(Grey)
	CorrectText   = lipgloss.NewStyle().Foreground(Green)
	WrongText     = lipgloss.NewStyle().Foreground(Red)
	RemainingText = lipgloss.NewStyle().Foreground(LightGrey)
	ActiveText    = lipgloss.NewStyle().Foreground(White)

	TextBox = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Width(APP_WIDTH).Padding(1, 2).Margin(0, 2).Foreground(White)
)
