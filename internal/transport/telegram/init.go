package telegram

import (
	"wlbot/internal/service/authenticator"
	"wlbot/internal/service/mikrotik"

	"gopkg.in/telebot.v3"
)

type Handler struct {
	debug bool

	bot      *telebot.Bot
	mikrotik *mikrotik.Service
	auth     *authenticator.Service
}

func New(debug bool, bot *telebot.Bot, mikrotik *mikrotik.Service, auth *authenticator.Service) *Handler {
	return &Handler{
		debug:    debug,
		bot:      bot,
		mikrotik: mikrotik,
		auth:     auth,
	}
}

func (h *Handler) Start() {
	defer h.bot.Stop()

	h.setRoutes()

	h.bot.Start()
}
