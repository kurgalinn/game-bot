package sqlite

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/kurgalinn/game-bot/internal/entity"
)

type User struct {
	db *sqlx.DB
}

func (p User) Get(id string) (user entity.User, err error) {
	err = p.db.Get(
		&user,
		`SELECT * FROM users WHERE id = ?`,
		id,
	)
	if err == sql.ErrNoRows {
		return user, nil
	}
	return user, err
}

func (p User) Save(user entity.User) error {
	_, err := p.db.Exec(
		`INSERT INTO users (id, name) VALUES (?, ?)`,
		user.ID,
		user.Name,
	)
	return err
}

func (p User) Update(user entity.User) error {
	_, err := p.db.Exec(
		`UPDATE users SET name = ? WHERE id = ?`,
		user.Name,
		user.ID,
	)
	return err
}
