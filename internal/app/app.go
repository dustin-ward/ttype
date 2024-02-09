package app

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin-ward/termtyping/internal/character"
	"github.com/dustin-ward/termtyping/internal/data"
	"github.com/dustin-ward/termtyping/internal/styles"
)

type AppState int

const (
	StateDefault AppState = iota
	StateTyping
)

type AppModel struct {
	CurState AppState
	Cursor   cursor.Model

	chars    []character.CharacterModel
	text     string
	pos      int
	quitting bool
}

func NewAppModel(init_state AppState) tea.Model {
	const NUM_WORDS = 20
	chars := make([]character.CharacterModel, 0)
	text := ""
	for i := 0; i < NUM_WORDS; i++ {
		word := data.GetWord() + " "
		text += word
		for _, ch := range word {
			chars = append(chars, character.NewCharacter(ch))
		}
	}
	chars = chars[:len(chars)-1]
	text = text[:len(text)-1]

	return AppModel{
		CurState: init_state,
		Cursor:   cursor.New(),
		chars:    chars,
		text:     text,
		pos:      0,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.CurState {
	case StateDefault:
		return defaultHandler(m, msg)
	case StateTyping:
		return typingHandler(m, msg)
	default:
		log.Println("Invalid State:", m.CurState)
		return m, nil
	}
}

func (m AppModel) View() string {
	if m.quitting {
		return ""
	}

	var view_text string
	switch m.CurState {
	case StateDefault:
		b := strings.Builder{}
		for _, ch := range m.chars {
			curState := ch.State
			ch.State = character.FinishedState
			b.WriteString(ch.View())
			ch.State = curState
		}
		view_text = b.String()

	case StateTyping:
		b := strings.Builder{}
		for _, ch := range m.chars {
			b.WriteString(ch.View())
		}
		view_text = b.String()
	}

	return styles.TextBox.Render(view_text)
}
