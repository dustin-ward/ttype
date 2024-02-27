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

	TextBox     = lipgloss.NewStyle().Width(APP_WIDTH).Margin(0, 2)
	StatusBar   = lipgloss.NewStyle().Width(APP_WIDTH).Margin(0, 2).Border(lipgloss.NormalBorder()).BorderForeground(LightGrey).BorderTop(false).BorderLeft(false).BorderRight(false)
	BorderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Foreground(White)
)
