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
			m.CurText = m.CurText[:max(len(m.CurText)-1, 0)]
		case "enter":
			m.CurText = append(m.CurText, "\n")
		case "tab":
			m.CurText = append(m.CurText, "\t")
		default:
			m.CurText = append(m.CurText, keypress)

		}
	}
	return m, tea.Batch(cmds...)
}
