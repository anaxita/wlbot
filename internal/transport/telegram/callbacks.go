package telegram

import (
	"context"
	"fmt"

	"wlbot/internal/helpers"

	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

func (h *Handler) approveAddIP(c tele.Context) error {
	defer func(c tele.Context) {
		err := c.Delete()
		if err != nil {
			h.l.Error("delete message: ", zap.Error(err))
		}
	}(c)

	comment := fmt.Sprintf("BOT %s | %s %s", c.Chat().Title, c.Sender().FirstName, c.Sender().LastName)

	err := h.mikrotik.AddIPFromChat(context.TODO(), c.Chat().ID, c.Data(), helpers.TranslitRuToEN(comment))
	if err != nil {
		h.l.Error("add ip: ", zap.Error(err))

		return c.Send("Извините, что-то пошло не так, мы скоро всё исправим!")
	}

	return c.Send("IP успешно добавлен!")
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}
