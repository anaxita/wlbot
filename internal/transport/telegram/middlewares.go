package telegram

import "gopkg.in/telebot.v3"

func disallowPrivateMessages(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Chat().Type == telebot.ChatPrivate {
			return c.Send("Я не отвечаю в личные сообщения, обратитесь в чат тех. поддержки.")
		}

		return next(c)
	}
}
