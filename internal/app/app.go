package app

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/data"
	"github.com/dustin-ward/termtyping/internal/styles"
)

type AppState int

const (
	StateDefault AppState = iota
	StateTyping
)

type AppModel struct {
	CurState AppState
	Cursor   cursor.Model

	finishedText  string
	wrongText     string
	remainingText string
	pos           int
	quitting      bool
}

func NewAppModel(init_state AppState) tea.Model {
	const NUM_WORDS = 20
	init_text := ""
	for i := 0; i < NUM_WORDS; i++ {
		init_text += data.GetWord() + " "
	}

	return AppModel{
		CurState:      init_state,
		Cursor:        cursor.New(),
		remainingText: init_text[:len(init_text)-1],
		pos:           0,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.CurState {
	case StateDefault:
		return defaultHandler(m, msg)
	case StateTyping:
		return typingHandler(m, msg)
	default:
		log.Println("Invalid State:", m.CurState)
		return m, nil
	}
}

func (m AppModel) View() string {
	if m.quitting {
		return ""
	}

	var view_text string
	switch m.CurState {
	case StateDefault:
		view_text = fmt.Sprintf("%s%s%s",
			styles.HiddenText.Render(m.finishedText),
			styles.WrongText.Render(m.wrongText),
			styles.HiddenText.Render(m.remainingText),
		)

	case StateTyping:
		view_text = fmt.Sprintf("%s%s%s",
			styles.HiddenText.Render(m.finishedText),
			styles.WrongText.Render(m.wrongText),
			styles.ActiveText.Render(m.remainingText),
		)
	}

	return styles.TextBox.Render(view_text)
}
