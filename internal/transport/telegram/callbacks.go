package telegram

import (
	"context"
	"fmt"

	"wlbot/internal/helpers"

	tele "gopkg.in/telebot.v3"
)

func (h *Handler) approveAddIP(c tele.Context) error {
	comment := fmt.Sprintf("BOT %s | %s %s", c.Chat().Title, c.Sender().FirstName, c.Sender().LastName)

	err := h.mikrotik.AddIPFromChat(context.TODO(), c.Chat().ID, c.Data(), helpers.TranslitRuToEN(comment))
	if err != nil {
		return err
	}

	return c.Reply("IP успешно добавлен")
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}
