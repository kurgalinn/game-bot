package usecase

import (
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func (h Handler) UserInput(c tele.Context) error {
	switch h.state.Get(c.Sender().Recipient()) {
	case service.EditsName:
		return h.editName(c)
	}

	h.state.Set(c.Sender().Recipient(), service.Idle)
	return nil
}
