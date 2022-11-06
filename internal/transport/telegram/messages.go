package telegram

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"kms/wlbot/internal/helpers"
	"net"
)

func (h *Handler) message(c tele.Context) error {
	ip4, subnet, err := helpers.FindIP(c.Text())
	if err == nil {
		return h.askToAddIP(c, ip4, subnet) // ignore messages without ip addresses
	}

	return nil
}

func (h *Handler) askToAddIP(c tele.Context, ip4 net.IP, subnet *net.IPNet) error {
	ip := ip4.String()
	if subnet != nil {
		if !h.auth.IsAdmin(c.Chat().ID, c.Sender().Username) {
			return nil
		}

		ip = subnet.String()
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
