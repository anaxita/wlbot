package telegram

import (
	"wlbot/internal/service/authenticator"
	"wlbot/internal/service/mikrotik"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	l *zap.SugaredLogger

	debug bool

	mw *Middleware

	bot      *telebot.Bot
	mikrotik mikrotik.Provider
	auth     *authenticator.Service
}

func New(
	l *zap.SugaredLogger,
	debug bool,
	mw *Middleware,
	bot *telebot.Bot,
	mikrotik mikrotik.Provider,
	auth *authenticator.Service,
) *Handler {
	return &Handler{
		l:        l,
		debug:    debug,
		mw:       mw,
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
