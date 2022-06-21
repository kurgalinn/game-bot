package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/kurgalinn/game-bot/internal/entity"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	Game  entity.GameRepository
	User  entity.UserRepository
	Movie entity.MovieRepository
	Score entity.ScoreRepository
}

func NewRepository(driver string, source string) (*Repository, error) {
	db, err := sqlx.Open(driver, source)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return &Repository{
		Game:  &Game{db: db},
		User:  &User{db: db},
		Movie: &Movie{db: db},
		Score: &Score{db: db},
	}, nil
}
