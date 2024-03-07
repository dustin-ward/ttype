package statusbar

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/ttype/internal/stats"
	"github.com/dustin-ward/ttype/internal/styles"
)

type StatusBarState int

const (
	StateDefault StatusBarState = iota
	StateTyping
)

type StatusBarModel struct {
	CurState StatusBarState
	Stats    *stats.Stats
}

func NewStatusBar(init_state int, init_stats *stats.Stats) StatusBarModel {
	return StatusBarModel{
		StatusBarState(init_state),
		init_stats,
	}
}

func (m StatusBarModel) Init() tea.Cmd {
	return nil
}

func (m StatusBarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m StatusBarModel) View() string {
	var acc, wpm float64
	if m.Stats != nil {
		acc, wpm = m.Stats.GetAccuracy(), m.Stats.GetWPM()
	}

	acc_style := getAccColour(acc)
	wpm_style := getWpmStyle(wpm)

	switch m.CurState {
	case StateDefault:
		return "Press 'Enter' to begin... (esc to exit)"
	case StateTyping:
		return styles.HiddenText.Render("|  Accuracy: ") +
			acc_style.Render(fmt.Sprintf("%0.1f%%", acc*100)) +
			styles.HiddenText.Render("  |  Speed: ") +
			wpm_style.Render(fmt.Sprintf("%0.1fwpm", wpm)) +
			styles.HiddenText.Render("  |")
	default:
		return "ERROR"
	}
}
