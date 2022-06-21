package usecase

import (
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
	"log"
)

func (h Handler) OnError(err error, c tele.Context) {
	if c != nil {
		switch h.state.Get(c.Sender().Recipient()) {
		case service.InGame:
			user, err := h.wakeUser(c.Sender().Recipient())
			if err != nil {
				return
			}
			h.state.Set(user.ID, service.Idle)

			game, err := h.repository.Game.ActiveByUser(user.ID)
			if err == nil {
				game.GameOver()
				h.repository.Game.GameOver(game)
			}

			// TODO: change locale setting
			h.layout.SetLocale(c, "ru")
			_ = c.Send(
				h.layout.Text(c, "err_game_fatal"),
				h.layout.Markup(c, "start"),
			)
		}
		log.Println(c.Sender().Recipient(), err)
	} else {
		log.Println(err)
	}
}
