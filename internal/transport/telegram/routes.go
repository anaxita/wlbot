package telegram

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (h *Handler) setRoutes() {
	// middlewares
	{
		h.bot.Use(middleware.Recover(), middleware.AutoRespond())

		if !h.debug {
			h.bot.Use(h.mw.disallowPrivateMessages)
		}
	}

	// handlers
	{
		h.bot.Handle(&btnAddIP, h.approveAddIP)
		h.bot.Handle(&btnDeclineIP, h.declineAddIP)
		h.bot.Handle(tele.OnAddedToGroup, h.onAddedToGroup)

		h.bot.Handle("/start", h.commandStart)
		h.bot.Handle("/chatid", h.commandChatID)

		h.bot.Handle(tele.OnText, h.message)
	}
}
