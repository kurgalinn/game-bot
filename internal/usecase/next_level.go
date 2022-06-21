package usecase

import (
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func (h Handler) MovieNextLevel(c tele.Context) error {
	user, err := h.wakeUser(c.Sender().Recipient())
	if err != nil {
		return err
	}

	// check state
	if h.state.Get(user.ID) != service.InGame {
		return tele.Err(h.layout.Text(c, "err_game_over"))
	}

	// getting active game
	game, err := h.repository.Game.ActiveByUser(user.ID)
	if err != nil {
		return err
	}

	// check user answer
	if c.PollAnswer().Options[0] != game.CorrectAnswer {
		// execute game over job
		h.worker.Execute(user.ID)
		return nil
	}

	op := BonusTimeDuration + h.worker.TimeLeft(user.ID)
	h.worker.Remove(user.ID)

	prevComplexity := game.Complexity()
	game.LevelUp()
	if prevComplexity != game.Complexity() {
		err = h.setMovieOptions(game)
		if err != nil {
			return err
		}
	}

	err = h.sendPoll(
		c,
		game,
		op,
	)
	if err != nil {
		return err
	}

	err = h.repository.Game.Update(game)
	if err != nil {
		return err
	}

	return nil
}
