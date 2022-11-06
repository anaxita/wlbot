package mikrotik

import (
	"context"
	"go.uber.org/zap"
	"kms/wlbot/internal/entity"
)

type Repository interface {
	ChatWLs(ctx context.Context, chatID int64) ([]entity.ChatWL, error)
	DefaultMikroTiks(ctx context.Context) ([]entity.Mikrotik, error)
	MikroTikByID(ctx context.Context, id int64) (entity.Mikrotik, error)
}

type Device interface {
	FindIP(ctx context.Context, m entity.Mikrotik, wl string, ip string) (isDynamic bool, err error)
	AddIP(ctx context.Context, m entity.Mikrotik, ip, comment string) error
	AddIPToCustomWL(ctx context.Context, m entity.Mikrotik, wl, ip, comment string) error
	RemoveIP(ctx context.Context, m entity.Mikrotik, wl string, ip string) error
}

type Service struct {
	l *zap.SugaredLogger

	repo   Repository
	device Device
}

func New(l *zap.SugaredLogger, repo Repository, device Device) *Service {
	return &Service{l: l, repo: repo, device: device}
}
