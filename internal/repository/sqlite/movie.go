package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurgalinn/game-bot/internal/entity"
)

type Movie struct {
	db *sqlx.DB
}

func (m Movie) Find(popularity int) (movies []entity.Movie, err error) {
	err = m.db.Select(
		&movies,
		`SELECT m.name_ru as title,
					  (SELECT i.url FROM images i WHERE i.id = m.id ORDER BY random() LIMIT 1) as image
			   FROM movies m
			   WHERE image NOT NULL AND m.popularity = ?`,
		popularity,
	)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
