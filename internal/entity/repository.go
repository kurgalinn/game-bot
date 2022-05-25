package entity

type (
	UserRepository interface {
		Get(id string) (User, error)
		Save(user User) error
		Update(user User) error
	}

	GameRepository interface {
		ActiveByUser(userID string) (Game, error)
		UserMaxLevel(userID string) int
		Create(game Game) error
		Update(game Game) error
		GameOver(game Game)
	}

	MovieRepository interface {
		Find(popularity int) ([]Movie, error)
	}

	ScoreRepository interface {
		Top(limit int) (scores []Score, err error)
	}
)
