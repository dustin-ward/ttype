package data

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"
)

//go:embed words_english.json
var words_english_json []byte
var words_english []string

func init() {
	if err := json.Unmarshal(words_english_json, &words_english); err != nil {
		log.Fatal("data/init:", err)
	}
}

// Return random word from dataset
func GetWord() string {
	return words_english[rand.Intn(len(words_english))]
}
