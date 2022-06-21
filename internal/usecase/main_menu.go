package usecase

import (
	tele "gopkg.in/telebot.v3"
)

func (h Handler) MainMenu(c tele.Context) error {
	return c.Send(
		h.layout.Text(c, "start"),
		h.layout.Markup(c, "start"),
	)
}
