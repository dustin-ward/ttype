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
	chars := make([]character.CharacterModel, 0)
	text := ""

	// A bunch of gross stuff to track which words are at the
	// end of each line. We need to know this because we let
	// bubbletea do the text wrapping, but the last space on
	// each line needs to be handled differently to show the
	// '_' character.
	space_pos := -1
	start_of_cur_line := 0
	for i := 0; i < NUM_WORDS; i++ {
		// Pull random word from data
		word := data.GetWord() + " "

		text += word
		for _, ch := range word {
			chars = append(chars, character.NewCharacter(ch))
		}

		// Deal with last word of each line
		space_line := (space_pos - start_of_cur_line) / (styles.APP_WIDTH - 2)
		cur_line := (len(text) - start_of_cur_line) / (styles.APP_WIDTH - 2)
		if space_pos > 0 && (space_line != cur_line) {
			chars[space_pos].EndSpace = true
			start_of_cur_line = len(text) - (len(word) - 1)
		}

		space_pos = len(text) - 1
	}
	// Remove space from last word
	chars = chars[:len(chars)-1]
	text = text[:len(text)-1]

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
