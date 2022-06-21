package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurgalinn/game-bot/internal/entity"
)

type Game struct {
	db *sqlx.DB
}

func (g Game) UserMaxLevel(userID string) (level int) {
	_ = g.db.Get(
		&level,
		`SELECT level FROM games WHERE user_id = ? ORDER BY level DESC LIMIT 1`,
		userID,
	)
	return level
}

func (g Game) GameOver(game entity.Game) {
	_, _ = g.db.NamedExec(`UPDATE games SET ended_at = :ended_at WHERE ended_at IS NULL AND user_id = :user_id`, game)
}

func (g Game) ActiveByUser(userID string) (game entity.Game, err error) {
	err = g.db.Get(&game, `SELECT * FROM games WHERE ended_at IS NULL AND user_id = ?`, userID)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (g Game) Create(game entity.Game) error {
	_, err := g.db.NamedExec(
		`
			INSERT INTO games (id, level, user_id, started_at, ended_at, correct_answer) 
			VALUES (:id, :level, :user_id, :started_at, :ended_at, :correct_answer)`,
		game,
	)
	return err
}

func (g Game) Update(game entity.Game) error {
	_, err := g.db.NamedExec(
		`
			UPDATE games 
				SET level = :level,
					correct_answer = :correct_answer,
					ended_at = :ended_at
			WHERE id = :id`,
		game,
	)
	return err
}
