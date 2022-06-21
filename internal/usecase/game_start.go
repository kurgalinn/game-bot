package usecase

import (
	"github.com/kurgalinn/game-bot/internal/entity"
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func (h Handler) MovieGameStart(c tele.Context) error {
	user, err := h.wakeUser(c.Sender().Recipient())
	if err != nil {
		return err
	}

	// check state
	if h.state.Get(user.ID) != service.Idle {
		if _, err = h.repository.Game.ActiveByUser(user.ID); err == nil {
			return c.Send(h.layout.Text(c, "err_already_play"))
		}
	}

	game := entity.StartGame(service.GenerateID(), user.ID)

	err = h.setMovieOptions(game)
	if err != nil {
		return err
	}

	err = h.sendPoll(
		c,
		game,
		StartDuration,
	)
	if err != nil {
		return err
	}

	// persist game
	err = h.repository.Game.Create(game)
	if err != nil {
		return err
	}

	return nil
}
