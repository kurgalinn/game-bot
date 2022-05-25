package usecase

import (
	tele "gopkg.in/telebot.v3"
)

func (h Handler) Help(c tele.Context) error {
	return c.Send(
		h.layout.Text(c, "help"),
	)
}
