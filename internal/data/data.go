package data

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"
	"strings"
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
func GetWord(punctuation_chance, capital_chance float64) string {
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
