package telegram

import (
	"wlbot/internal/service/authenticator"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type Middleware struct {
	l    *zap.SugaredLogger
	auth *authenticator.Service
}

func NewMiddleware(l *zap.SugaredLogger, auth *authenticator.Service) *Middleware {
	return &Middleware{l: l, auth: auth}
}

func (m *Middleware) disallowPrivateMessages(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Chat().Type == telebot.ChatPrivate {
			return c.Send("Я не отвечаю в личные сообщения, обратитесь в чат тех. поддержки.")
		}

		return next(c)
	}
}
