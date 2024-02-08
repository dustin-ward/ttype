package app

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/data"
)

type AppState int

const (
	StateDefault AppState = iota
	StateTyping
)

type AppModel struct {
	CurState AppState
	CurText  []string
	Cursor   cursor.Model
}

func NewAppModel() tea.Model {
	const NUM_WORDS = 20
	init_text := make([]string, NUM_WORDS)
	for i := range init_text {
		init_text[i] = data.GetWord()
	}

	return AppModel{
		CurState: StateDefault,
		CurText:  init_text,
		Cursor:   cursor.New(),
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
	cursor := ""
	if m.CurState == StateTyping {
		cursor = "_"
	}

	display_text := ""
	for _, word := range m.CurText {
		display_text += " " + word
	}

	return fmt.Sprintf("%d: %s%s", m.CurState, display_text, cursor)
}
