package character

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/ttype/internal/styles"
)

const (
	CH_SPACE = " "
)

type CharState int

const (
	HiddenState CharState = iota
	RemainingState
	WrongState
	CorrectState
	ActiveState
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
	ch := m.Val

	// Display underscore instead of space in some states
	if ch == CH_SPACE {
		switch m.State {
		case WrongState, ActiveState:
			ch = "_"
		}
	}

	switch m.State {
	case HiddenState:
		return styles.HiddenText.Render(ch)
	case RemainingState:
		return styles.RemainingText.Render(ch)
	case WrongState:
		return styles.WrongText.Render(ch)
	case CorrectState:
		return styles.CorrectText.Render(ch)
	case ActiveState:
		return styles.ActiveText.Render(ch)
	}

	return "?"
}
