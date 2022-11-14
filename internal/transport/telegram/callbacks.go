package telegram

import (
	"context"

	tele "gopkg.in/telebot.v3"
)

func (h *Handler) approveAddIP(c tele.Context) error {
	err := h.mikrotik.AddIPFromChat(context.TODO(), c.Chat().ID, c.Data())
	if err != nil {
		return err
	}

	return c.Reply("IP успешно добавлен")
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}
