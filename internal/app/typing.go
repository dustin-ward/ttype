package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/character"
)

type NewTextMsg struct{}
type ResetTextMsg struct{}

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
			return NewAppModelWithText(m.text, StateDefault), func() tea.Msg { return ResetTextMsg{} }

		case "enter":
			// Enter key will generate new text
			return NewAppModel(StateTyping), func() tea.Msg { return NewTextMsg{} }
		default:
			if keypress == m.chars[m.pos].Val {
				m.chars[m.pos].State = character.CorrectState
				m.pos++
				if m.pos < len(m.text) {
					if m.text[m.pos] == '\n' {
						m.pos++
					}
					m.chars[m.pos].State = character.ActiveState
				}
			} else {
				m.chars[m.pos].State = character.WrongState
			}
		}
	}

	if m.pos == len(m.text) {
		// All of the current text is completed. Reset
		return NewAppModel(StateTyping), func() tea.Msg { return NewTextMsg{} }
	}

	return m, tea.Batch(cmds...)
}
