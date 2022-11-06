package scheduler

import (
	"context"
	"go.uber.org/zap"
	"kms/wlbot/internal/service/config"
	"os"
	"time"
)

type Service struct {
	l      *zap.SugaredLogger
	config *config.Service
}

func New(l *zap.SugaredLogger, cfg *config.Service) *Service {
	return &Service{
		l:      l,
		config: cfg,
	}
}

func (s *Service) Start(ctx context.Context) {
	go s.updateConfig(ctx)
}

func (s *Service) updateConfig(ctx context.Context) {
	t := time.NewTicker(time.Second * 30)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}

		stat, err := os.Stat(s.config.ConfigPath())
		if err != nil {
			s.l.Error(zap.Error(err))
			continue
		}

		modTime := stat.ModTime()
		lastUpdateTime := s.config.LastUpdateAt()

		if modTime.After(lastUpdateTime) {
			s.l.Debug("updating config...")

			err = s.config.Reload()
			if err != nil {
				s.l.Error("reload: ", zap.Error(err))
				continue
			}

			s.l.Debugln("last updated at:", s.config.LastUpdateAt())
			continue
		}

		s.l.Debugln("config has no changes")
	}
}
