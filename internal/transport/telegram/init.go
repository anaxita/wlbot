package telegram

import (
	"context"
	"gopkg.in/telebot.v3"
	"kms/wlbot/internal/service/authenticator"
	"kms/wlbot/internal/service/mikrotik"
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

func (h *Handler) Start(ctx context.Context) {
	defer h.bot.Stop()

	h.setRoutes()

	go h.bot.Start()

	<-ctx.Done()
}
