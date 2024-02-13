package app

import (
	"log"
	"strings"

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

	chars    []character.CharacterModel
	text     string
	pos      int
	quitting bool
}

// Number of words to use per each test
const NUM_WORDS = 20

func NewAppModel(init_state AppState) tea.Model {
	text := ""
	line_len := 0
	lines := 0
	for i := 0; i < NUM_WORDS; i++ {
		// Pull random word from data
		word := data.GetWord() + " "

		// Manually insert newlines
		if line_len+len(word) >= styles.APP_WIDTH-4 {
			text += "\n"
			line_len = len(word)
			lines++

			if lines == styles.MAX_LINES {
				break
			}
		} else {
			line_len += len(word)
		}
		text += word
	}
	text = strings.TrimSpace(text)

	// Fill models array
	chars := make([]character.CharacterModel, len(text))
	for i, ch := range text {
		chars[i] = character.NewCharacter(ch)
	}

	// First character is active
	chars[0].State = character.ActiveState

	return AppModel{
		CurState: init_state,
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
		// All words are greyed out in default state
		b := strings.Builder{}
		for _, ch := range m.chars {
			curState := ch.State
			ch.State = character.HiddenState
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
