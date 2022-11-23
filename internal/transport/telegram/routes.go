package telegram

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (h *Handler) setRoutes() {
	h.bot.Use(middleware.Recover(), middleware.AutoRespond())

	if !h.debug {
		h.bot.Use(h.mw.disallowPrivateMessages)
	}

	h.bot.Handle(&btnAddIP, h.approveAddIP)
	h.bot.Handle(&btnDeclineIP, h.declineAddIP)

	h.bot.Handle("/start", h.commandStart)
	h.bot.Handle("/chatid", h.commandChatID)
	h.bot.Handle(telebot.OnText, h.message)
}
