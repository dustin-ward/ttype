package character

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/styles"
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
	Val      string
	State    CharState
	EndSpace bool
}

func NewCharacter(val rune) CharacterModel {
	return CharacterModel{
		string(val),
		RemainingState,
		false,
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
	// For the last word of each line, we need to make sure
	// the space character is double width. This is because
	// when its in its active state it will display the '_'
	// character instead.
	if ch == CH_SPACE {
		switch m.State {
		case HiddenState, RemainingState, CorrectState:
			if m.EndSpace {
				ch = "  "
			} else {
				ch = " "
			}
		case ActiveState, WrongState:
			if m.EndSpace {
				ch = "_ "
			} else {
				ch = "_"
			}
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
