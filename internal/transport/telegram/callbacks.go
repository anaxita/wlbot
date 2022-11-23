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

	err := h.mikrotik.AddIPFromChat(context.Background(), c.Chat().ID, c.Data(), helpers.TranslitRuToEN(comment))
	if err != nil {
		h.l.Errorw("failed add IP to a chat: ",
			"error", err,
			"chat_id", c.Chat().ID,
			"chat_title", c.Chat().Title)

		return c.Send("Извините, что-то пошло не так, мы скоро всё исправим!")
	}

	return c.Send("IP успешно добавлен!")
}

func (h *Handler) declineAddIP(c tele.Context) error {
	return c.Delete()
}

func (h *Handler) onAddedToGroup(c tele.Context) error {
	if !h.auth.IsAdmin(c.Chat().ID, c.Sender().Username) {
		h.l.Warnw("someone tried to add the bot to group",
			"user_id", c.Sender().ID,
			"chat_id", c.Chat().ID,
			"chat_title", c.Chat().Title)

		return c.Bot().Leave(c.Chat())
	}

	h.l.Infow("bot was added to group",
		"user_id", c.Sender().ID,
		"user_name", c.Sender().Username,
		"chat_id", c.Chat().ID,
		"chat_title", c.Chat().Title)

	return nil
}
