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
			m.CurText = append(m.CurText, '\n')
		case "tab":
			m.CurText = append(m.CurText, '\t')
		// case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		// 	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		// 	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", " ", ".", ",", "-", "_", "?", "(", ")", "[", "]", "{", "}", "/", "\t":
		default:
			m.CurText = append(m.CurText, rune(keypress[0]))

		}
	}
	return m, tea.Batch(cmds...)
}
