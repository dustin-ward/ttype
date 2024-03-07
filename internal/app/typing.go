package app

import (
	"time"

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
			// Escape will reset the state of each character while keeping the same text
			return NewAppModelWithText(m.text, StateDefault, nil), func() tea.Msg { return ResetTextMsg{} }

		case "enter":
			// Enter key will generate new text
			return NewAppModel(StateTyping, nil), func() tea.Msg { return NewTextMsg{} }

		default:
			m.stats.NumKeypresses++

			if keypress == m.chars[m.pos].Val {
				m.stats.NumCorrect++
				if m.pos == 0 {
					m.stats.TimeStarted = time.Now()
				}

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
		return NewAppModel(StateTyping, m.stats.Finish()), func() tea.Msg { return NewTextMsg{} }
	}

	return m, tea.Batch(cmds...)
}
