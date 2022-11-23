package telegram

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (h *Handler) commandStart(c telebot.Context) error {
	return nil
}

func (h *Handler) commandChatID(c telebot.Context) error {
	return c.Reply(fmt.Sprintf("Chat ID: `%d`", c.Chat().ID))
}
