package telegram

import tele "gopkg.in/telebot.v3"

var (
	menu = &tele.ReplyMarkup{OneTimeKeyboard: true}

	btnDeclineIP = menu.Data("Нет", uniqDeclineIP)
	btnAddIP     = menu.Data("Добавить", uniqAddIP)
)

const (
	uniqAddIP     = "add_ip"
	uniqDeclineIP = "decline_ip"
)
