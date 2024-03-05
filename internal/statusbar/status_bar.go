package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/styles"
)

type StatusBarState int

const (
	StateDefault StatusBarState = iota
	StateTyping
)

type StatusBarModel struct {
	CurState StatusBarState
}

func NewStatusBar(init_state int) StatusBarModel {
	return StatusBarModel{
		StatusBarState(init_state),
	}
}

func (m StatusBarModel) Init() tea.Cmd {
	return nil
}

func (m StatusBarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m StatusBarModel) View() string {
	switch m.CurState {
	case StateDefault:
		return "Press 'Enter' to begin..."
	case StateTyping:
		return styles.HiddenText.Render("|  ") +
			"00.0%" +
			styles.HiddenText.Render("  |  ") +
			"00.0wpm" +
			styles.HiddenText.Render("  |")
	default:
		return "ERROR"
	}
}
