package notificator

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

type Repository interface {
	AdminChatIDs() []int64
}

type Service struct {
	l *zap.SugaredLogger

	repo Repository
	bot  *telebot.Bot
}

func New(l *zap.SugaredLogger, repo Repository, bot *telebot.Bot) *Service {
	return &Service{l: l, repo: repo, bot: bot}
}

func (s *Service) SendToAdminChats(text string) error {
	errs := make([]string, 0)

	for _, chatID := range s.repo.AdminChatIDs() {
		_, err := s.bot.Send(&telebot.Chat{ID: chatID}, text)
		if err != nil {
			s.l.Errorw("failed to send message to admin chat", "chatID", chatID, "error", err)
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to send message to admin chats: %s", strings.Join(errs, "; "))
	}

	return nil
}
