package usecase

import (
	"github.com/kurgalinn/game-bot/internal/entity"
	"github.com/kurgalinn/game-bot/internal/repository/sqlite"
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"time"
)

const (
	StartDuration     = 30
	BonusTimeDuration = 5
)

type Handler struct {
	repository *sqlite.Repository
	layout     *layout.Layout
	worker     service.Worker
	state      service.StatePool
	sets       service.UserSets
}

func NewHandler(r *sqlite.Repository, l *layout.Layout, s service.UserSets) *Handler {
	return &Handler{
		repository: r,
		layout:     l,
		worker:     service.NewWorker(),
		state:      service.NewStatePool(),
		sets:       s,
	}
}

func (h Handler) wakeUser(id string) (user entity.User, err error) {
	user, err = h.repository.User.Get(id)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		user = entity.NewUser(id, id)
		err = h.repository.User.Save(user)
	}

	return user, err
}

func (h Handler) sendPoll(
	c tele.Context,
	game entity.Game,
	op int,
) (err error) {
	// set state
	h.state.Set(game.UserID, service.InGame)

	// getting options from cache set
	serializer := service.Serializer[entity.Movie]()
	value, err := h.sets.PopRand(game.UserID, "options")
	if err != nil {
		return err
	}
	values, err := h.sets.GetRand(game.UserID, "options", entity.AnswersCount)
	if err != nil {
		return err
	}
	movies, err := serializer.DecodeList(values)
	if err != nil {
		return err
	}
	movies[game.CorrectAnswer], err = serializer.Decode(value)
	if err != nil {
		return err
	}

	// create game answers
	var options []tele.PollOption
	for _, movie := range movies {
		options = append(options, tele.PollOption{Text: movie.Title})
	}

	p := tele.Poll{
		Type:          tele.PollQuiz,
		Question:      h.layout.Text(c, "level", game.Level),
		OpenPeriod:    op,
		CorrectOption: game.CorrectAnswer,
		Options:       options,
	}

	// send poll
	err = c.Send(&tele.Photo{File: tele.FromURL(movies[game.CorrectAnswer].Image)})
	if err != nil {
		return err
	}

	if err = c.Send(&p); err != nil {
		return err
	} else {
		h.worker.Add(game.UserID, time.Duration(op)*time.Second, func() {
			h.movieGameOver(c, game)
		})
	}

	return nil
}

func (h Handler) setMovieOptions(game entity.Game) error {
	// getting first answers set
	movies, err := h.repository.Movie.Find(game.Complexity())
	if err != nil {
		return err
	}

	// create options cache sets
	serializer := service.Serializer[entity.Movie]()
	values, err := serializer.EncodeList(movies)
	if err != nil {
		return err
	}
	err = h.sets.Create(game.UserID, "options", values)
	if err != nil {
		return err
	}

	return nil
}
