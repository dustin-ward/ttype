package app

import tea "github.com/charmbracelet/bubbletea"

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
			if m.pos > 0 {
				var bs_char string
				if len(m.wrongText) > 0 {
					bs_char = string(m.wrongText[len(m.wrongText)-1])
					m.wrongText = m.wrongText[:len(m.wrongText)-1]
				} else {
					bs_char = string(m.finishedText[len(m.finishedText)-1])
					m.finishedText = m.finishedText[:len(m.finishedText)-1]
				}
				m.remainingText = bs_char + m.remainingText
				m.pos--
			}
		case "enter":
			m = NewAppModel(StateTyping).(AppModel)
		default:
			if len(m.remainingText) > 0 {
				if len(m.wrongText) == 0 && keypress == string(m.remainingText[0]) {
					m.finishedText += string(m.remainingText[0])
				} else {
					m.wrongText += string(m.remainingText[0])
				}
				m.remainingText = m.remainingText[1:]
				m.pos++
			}
		}
	}

	if len(m.remainingText)+len(m.wrongText) == 0 {
		m = NewAppModel(StateTyping).(AppModel)
	}

	cmds = append(cmds, m.Cursor.BlinkCmd())
	return m, tea.Batch(cmds...)
}
