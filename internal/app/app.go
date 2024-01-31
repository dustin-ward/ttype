package app

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type AppState int

const (
	StateDefault AppState = iota
	StateTyping
)

type AppModel struct {
	CurState AppState
	CurText  []rune
}

func NewAppModel() tea.Model {
	return AppModel{
		CurState: StateDefault,
		CurText:  []rune("Type Here:"),
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
	return fmt.Sprintf("%d: %s%s", m.CurState, string(m.CurText), cursor)
}
