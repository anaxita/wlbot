package telegram

import (
	"context"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) approveAddIP(c tele.Context) error {
	defer c.Delete()

	return h.mikrotik.AddIPFromChat(context.TODO(), c.Chat().ID, c.Data())
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}
