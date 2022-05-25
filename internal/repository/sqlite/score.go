package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurgalinn/game-bot/internal/entity"
)

type Score struct {
	db *sqlx.DB
}

func (s Score) Top(limit int) (scores []entity.Score, err error) {
	err = s.db.Select(
		&scores,
		`
		SELECT u.name, max(g.level) score
		FROM games g
			INNER JOIN users u on u.id = g.user_id
		GROUP BY g.user_id
		ORDER BY score DESC 
		LIMIT ?`,
		limit,
	)
	if err != nil {
		return nil, err
	}

	return scores, nil
}
