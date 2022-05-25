package usecase

import (
	tele "gopkg.in/telebot.v3"
)

type profile struct {
	Name   string
	Level  int
	Rating int
}

func (h Handler) Profile(c tele.Context) error {
	user, err := h.wakeUser(c.Sender().Recipient())
	if err != nil {
		return err
	}

	return c.Send(
		h.layout.Text(c, "profile", &profile{
			Name:   user.Name,
			Level:  h.repository.Game.UserMaxLevel(user.ID),
			Rating: 0,
		}),
		h.layout.Markup(c, "profile"),
	)
}
