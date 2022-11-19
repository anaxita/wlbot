package mikrotik

import (
	"context"

	"wlbot/internal/entity"

	"go.uber.org/zap"
)

var _ Provider = (*LoggedService)(nil)

type LoggedService struct {
	l *zap.SugaredLogger
	s *Service
}

func NewLogged(l *zap.SugaredLogger, s *Service) *LoggedService {
	return &LoggedService{l: l, s: s}
}

func (s *LoggedService) AddIPFromChat(ctx context.Context, chatID int64, ip string, comment string) error {
	err := s.s.AddIPFromChat(ctx, chatID, ip, comment)
	if err != nil {
		s.l.Errorw("failed to add ip from chat",
			"chat_id", chatID,
			"ip", ip,
			"comment", comment,
			"err", err)

		return err
	}

	return nil
}

func (s *LoggedService) AddIPToDefaultMikrotiks(ip, comment string) error {
	err := s.s.AddIPToDefaultMikrotiks(ip, comment)
	if err != nil {
		s.l.Errorw("failed to add ip to default mikrotiks",
			"ip", ip,
			"comment", comment,
			"err", err)

		return err
	}

	return nil
}

func (s *LoggedService) AddIPToCustomMikrotiks(
	ctx context.Context,
	wls []entity.ChatWL,
	ip, comment string,
) (err error) {
	err = s.s.AddIPToCustomMikrotiks(ctx, wls, ip, comment)
	if err != nil {
		s.l.Errorw("failed to add ip to custom mikrotiks",
			"ip", ip,
			"comment", comment,
			"err", err)

		return
	}

	return nil
}
