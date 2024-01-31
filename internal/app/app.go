package app

import tea "github.com/charmbracelet/bubbletea"

type AppModel struct {
}

func NewAppModel() tea.Model {
	return AppModel{}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		}
	}

	return m, nil
}

func (m AppModel) View() string {
	return "Hello"
}
