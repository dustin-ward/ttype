package character

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/styles"
)

type CharState int

const (
	RemainingState CharState = iota
	WrongState
	FinishedState
)

type CharacterModel struct {
	Val   string
	State CharState
}

func NewCharacter(val rune) CharacterModel {
	return CharacterModel{
		string(val),
		RemainingState,
	}
}

func (m CharacterModel) Init() tea.Cmd {
	return nil
}

func (m CharacterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m CharacterModel) View() string {
	switch m.State {
	case RemainingState:
		return styles.ActiveText.Render(m.Val)
	case WrongState:
		return styles.WrongText.Render(m.Val)
	case FinishedState:
		return styles.HiddenText.Render(m.Val)
	}

	return "?"
}
