package usecase

import (
	"github.com/kurgalinn/game-bot/internal/entity"
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func (h Handler) movieGameOver(c tele.Context, game entity.Game) {
	h.state.Set(game.UserID, service.Idle)
	h.sets.Remove(game.UserID, "options")

	game.GameOver()
	h.repository.Game.GameOver(game)
	maxLevel := h.repository.Game.UserMaxLevel(game.UserID)

	// TODO: change locale setting
	h.layout.SetLocale(c, "ru")
	err := c.Send(
		h.layout.Text(c, "game_over", struct {
			Level    int
			MaxLevel int
		}{game.Level, maxLevel}),
		h.layout.Markup(c, "start"))
	if err != nil {
		h.OnError(err, c)
	}
}
