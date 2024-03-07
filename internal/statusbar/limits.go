package statusbar

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin-ward/termtyping/internal/styles"
)

// Colour gradient stops for acc metric
func getAccColour(acc float64) lipgloss.Style {
	if acc < 0.01 {
		return styles.HiddenText
	}

	c := styles.Accuracy30
	if acc > 0.3 {
		c = styles.Accuracy50
	}
	if acc > 0.5 {
		c = styles.Accuracy70
	}
	if acc > 0.7 {
		c = styles.Accuracy80
	}
	if acc > 0.8 {
		c = styles.Accuracy90
	}
	if acc > 0.9 {
		c = styles.Accuracy95
	}
	if acc > 0.95 {
		c = styles.Accuracy98
	}
	if acc > 0.98 {
		c = styles.Accuracy100
	}

	return lipgloss.NewStyle().Foreground(c)
}

// Colour gradient stops for wpm metric
func getWpmStyle(wpm float64) lipgloss.Style {
	if wpm < 0.1 {
		return styles.HiddenText
	}

	c := styles.Accuracy30
	if wpm > 30.0 {
		c = styles.Accuracy50
	}
	if wpm > 40.0 {
		c = styles.Accuracy70
	}
	if wpm > 50.0 {
		c = styles.Accuracy80
	}
	if wpm > 60.0 {
		c = styles.Accuracy90
	}
	if wpm > 80.0 {
		c = styles.Accuracy95
	}
	if wpm > 90.0 {
		c = styles.Accuracy98
	}
	if wpm > 100.0 {
		c = styles.Accuracy100
	}

	return lipgloss.NewStyle().Foreground(c)
}
