package app

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin-ward/termtyping/internal/character"
	"github.com/dustin-ward/termtyping/internal/config"
	"github.com/dustin-ward/termtyping/internal/data"
	"github.com/dustin-ward/termtyping/internal/stats"
	"github.com/dustin-ward/termtyping/internal/statusbar"
	"github.com/dustin-ward/termtyping/internal/styles"
)

var Config config.Config

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

	status_bar statusbar.StatusBarModel

	stats *stats.Stats
}

// Number of words to use per each test
const NUM_WORDS = 20
const PUNC_CHANCE = 0.2
const CAPS_CHANCE = 0.2

func NewAppModel(init_state AppState, prev_stats *stats.Stats) tea.Model {
	// Generate new body of text
	text := data.GenText(NUM_WORDS, PUNC_CHANCE, CAPS_CHANCE)

	return NewAppModelWithText(text, init_state, prev_stats)
}

func NewAppModelWithText(text string, init_state AppState, prev_stats *stats.Stats) tea.Model {
	// Fill models array
	chars := make([]character.CharacterModel, len(text))
	for i, ch := range text {
		chars[i] = character.NewCharacter(ch)
	}

	// First character is active
	chars[0].State = character.ActiveState

	return AppModel{
		CurState:   init_state,
		chars:      chars,
		text:       text,
		pos:        0,
		status_bar: statusbar.NewStatusBar(int(init_state), prev_stats),
		stats:      stats.NewStats(),
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

	var view string
	if Config.ZenMode {
		view = styles.BorderStyle.Render(styles.Zen_TextBox.Render(view_text))
	} else {
		view = styles.BorderStyle.Render(lipgloss.JoinVertical(
			lipgloss.Left,
			styles.StatusBar.Render(m.status_bar.View()),
			styles.TextBox.Render(view_text),
		))
	}
	return view
}
