package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/character"
)

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

		case "enter":
			m = NewAppModel(StateTyping).(AppModel)
		default:
			if keypress == m.chars[m.pos].Val {
				m.chars[m.pos].State = character.FinishedState
				m.pos++
			} else {
				m.chars[m.pos].State = character.WrongState
			}
		}
	}

	if m.pos == len(m.text) {
		m = NewAppModel(StateTyping).(AppModel)
	}

	cmds = append(cmds, m.Cursor.BlinkCmd())
	return m, tea.Batch(cmds...)
}
