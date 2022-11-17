package telegram

import (
	"context"
	"fmt"

	"wlbot/internal/helpers"

	tele "gopkg.in/telebot.v3"
)

func (h *Handler) approveAddIP(c tele.Context) error {
	defer c.Delete()

	comment := fmt.Sprintf("BOT %s | %s %s", c.Chat().Title, c.Sender().FirstName, c.Sender().LastName)

	err := h.mikrotik.AddIPFromChat(context.TODO(), c.Chat().ID, c.Data(), helpers.TranslitRuToEN(comment))
	if err != nil {
		_ = c.Send("Извините, что-то пошло не так, мы скоро всё исправим!")

		return err
	}

	return c.Send("IP успешно добавлен!")
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}
