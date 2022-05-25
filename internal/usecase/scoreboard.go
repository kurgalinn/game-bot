package usecase

import (
	tele "gopkg.in/telebot.v3"
)

func (h Handler) MovieGameScoreboard(c tele.Context) error {
	scores, err := h.repository.Score.Top(10)
	if err != nil {
		return err
	}

	return c.Send(h.layout.Text(c, "scoreboard", scores))
}
