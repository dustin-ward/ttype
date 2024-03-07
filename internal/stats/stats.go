package stats

import (
	"time"
)

type Stats struct {
	NumKeypresses int
	NumCorrect    int
	LenText       int
	TimeStarted   time.Time
	TimeEnded     time.Time
}

func NewStats() *Stats {
	return &Stats{
		NumKeypresses: 0,
		NumCorrect:    0,
		TimeStarted:   time.Now(),
	}
}

func (s *Stats) GetAccuracy() float64 {
	return float64(s.NumCorrect) / float64(s.NumKeypresses)
}

// Avg word has 5 characters
func (s *Stats) GetWPM() float64 {
	return (float64(s.NumCorrect) / 5.0) / s.TimeEnded.Sub(s.TimeStarted).Minutes()
}

func (s *Stats) Finish() *Stats {
	s.TimeEnded = time.Now()
	return s
}
