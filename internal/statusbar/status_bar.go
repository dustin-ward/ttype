package statusbar

import tea "github.com/charmbracelet/bubbletea"

type StatusBarModel struct {
}

func (m StatusBarModel) Init() tea.Cmd {
	return nil
}

func (m StatusBarModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m StatusBarModel) View() string {
	return "TEST"
}
