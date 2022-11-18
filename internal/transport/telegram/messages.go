package telegram

import (
	"fmt"

	"wlbot/internal/helpers"

	tele "gopkg.in/telebot.v3"
)

func (h *Handler) message(c tele.Context) error {
	ip4, subnet, err := helpers.FindIP(c.Text())
	if err == nil {
		return h.askToAddIP(c, ip4, subnet) // ignore messages without ip addresses
	}

	return nil
}

func (h *Handler) askToAddIP(c tele.Context, ip4 string, subnet string) error {
	ip := ip4

	if subnet != "" {
		if !h.auth.IsAdmin(c.Chat().ID, c.Sender().Username) {
			return c.Send("Только администраторы могут добавлять подсети.")
		}

		ip = subnet
	}

	menu.Inline(
		menu.Row(
			menu.Data("Добавить", uniqAddIP, ip),
			btnDeclineIP,
		),
	)

	text := fmt.Sprintf("Вы хотите добавить IP `%s` в белый список?", ip)

	return c.Reply(text, menu)
}
