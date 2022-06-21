package entity

import (
	"database/sql"
	"math"
	"math/rand"
	"time"
)

const AnswersCount = 4

type Game struct {
	ID            string       `db:"id"`
	Level         int          `db:"level"`
	UserID        string       `db:"user_id"`
	StartedAt     time.Time    `db:"started_at"`
	EndedAt       sql.NullTime `db:"ended_at"`
	CorrectAnswer int          `db:"correct_answer"`
}

func StartGame(ID string, userID string) Game {
	return Game{
		ID:            ID,
		Level:         1,
		UserID:        userID,
		StartedAt:     time.Now(),
		CorrectAnswer: rand.Intn(AnswersCount),
	}
}

// Complexity When 100 is easy and 0 is very hard with step 5
func (g Game) Complexity() int {
	return int(100 - math.Floor(float64(g.Level/30))*5)
}

func (g *Game) LevelUp() {
	g.Level++
	g.CorrectAnswer = rand.Intn(AnswersCount)
}

func (g *Game) GameOver() {
	g.EndedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}
