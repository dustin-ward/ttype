package data

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"
	"strings"

	"github.com/dustin-ward/ttype/internal/styles"
)

//go:embed words_english.json
var words_english_json []byte
var words_english []string

var punctuation string = ".,?!;:"

func init() {
	if err := json.Unmarshal(words_english_json, &words_english); err != nil {
		log.Fatal("data/init:", err)
	}
}

// Return random word from dataset
func getWord(punctuation_chance, capital_chance float64) string {
	word := words_english[rand.Intn(len(words_english))]

	if rand.Float64() < capital_chance {
		word = strings.Title(word)
	}

	if rand.Float64() < punctuation_chance {
		p := punctuation[rand.Intn(len(punctuation))]
		word += string(p)
	}

	return word
}

// Generate full body of text with num_words words.
// Punctuation will be added with a punc_chance*100% chance.
// Words will be capilized with a caps_chance*100% chance.
func GenText(num_words int, punc_chance, caps_chance float64) string {
	text := ""
	line_len := 0
	lines := 0
	for i := 0; i < num_words; i++ {
		// Pull random word from data
		word := getWord(punc_chance, caps_chance) + " "

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

	return strings.TrimSpace(text)
}
