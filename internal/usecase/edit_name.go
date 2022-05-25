package usecase

import (
	"github.com/kurgalinn/game-bot/internal/service"
	tele "gopkg.in/telebot.v3"
)

func (h Handler) EditName(c tele.Context) error {
	h.state.Set(c.Sender().Recipient(), service.EditsName)
	return c.Send(h.layout.Text(c, "edit_nickname"))
}

func (h Handler) editName(c tele.Context) error {
	name := c.Text()
	err := c.Delete()
	if err != nil {
		return err
	}

	if name == "" || len(name) > 10 {
		return c.Send(h.layout.Text(c, "err_bad_name"))
	}

	user, err := h.wakeUser(c.Sender().Recipient())
	if err != nil {
		return err
	}

	user.SetName(name)
	err = h.repository.User.Update(user)
	if err != nil {
		return err
	}

	return h.Profile(c)
}
