package mikrotik

import (
	"context"

	"wlbot/internal/entity"
)

type Repository interface {
	ChatWLs(ctx context.Context, chatID int64) []entity.ChatWL
	DefaultMikroTiks() []entity.Mikrotik
	MikroTikByID(ctx context.Context, id int64) (entity.Mikrotik, error)
}

type Device interface {
	AddIP(m entity.Mikrotik, ip, comment string) error
	AddIPToCustomWL(m entity.Mikrotik, wl, ip, comment string) error
	RemoveIP(m entity.Mikrotik, wl string, ip string) error
}

// Provider provides mikrotik service methods.
type Provider interface {
	AddIPFromChat(ctx context.Context, chatID int64, ip string, comment string) (err error)
	AddIPToDefaultMikrotiks(ip, comment string) (err error)
	AddIPToCustomMikrotiks(ctx context.Context, wls []entity.ChatWL, ip, comment string) (err error)
}

type Service struct {
	repo   Repository
	device Device
}

func New(repo Repository, device Device) *Service {
	return &Service{repo: repo, device: device}
}
