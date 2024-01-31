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
	CurText  string
}

func NewAppModel() tea.Model {
	return AppModel{
		CurState: StateDefault,
		CurText:  "Type Here:",
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
	return fmt.Sprintf("%d: %s%s", m.CurState, m.CurText, cursor)
}

func defaultHandler(m AppModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			cmd = tea.Quit
			cmds = append(cmds, cmd)

		case "enter":
			m.CurState = StateTyping

		}
	}

	return m, tea.Batch(cmds...)
}

func typingHandler(m AppModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			cmd = tea.Quit
			cmds = append(cmds, cmd)

		case "esc":
			m.CurState = StateDefault

		case "backspace":
			m.CurText = m.CurText[:max(len(m.CurText)-1, 0)]

		case "enter":
			m.CurText += "\n"
		case "tab":
			m.CurText += "\t"

		// case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		// 	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		// 	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", " ", ".", ",", "-", "_", "?", "(", ")", "[", "]", "{", "}", "/", "\t":
		default:
			m.CurText += keypress

		}
	}
	return m, tea.Batch(cmds...)
}
