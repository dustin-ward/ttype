package app

import tea "github.com/charmbracelet/bubbletea"

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
