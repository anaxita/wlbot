package telegram

import tele "gopkg.in/telebot.v3"

var (
	menu = &tele.ReplyMarkup{
		InlineKeyboard:  nil,
		ReplyKeyboard:   nil,
		ForceReply:      false,
		ResizeKeyboard:  false,
		OneTimeKeyboard: true,
		RemoveKeyboard:  false,
		Selective:       false,
		Placeholder:     "",
	}

	btnDeclineIP = menu.Data("Нет", uniqDeclineIP)
	btnAddIP     = menu.Data("Добавить", uniqAddIP)
)

const (
	uniqAddIP     = "add_ip"
	uniqDeclineIP = "decline_ip"
)
