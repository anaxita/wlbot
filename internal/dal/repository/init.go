package repository

import (
	"context"
	"kms/wlbot/internal/entity"
	"kms/wlbot/internal/service/config"
	"kms/wlbot/internal/xerrors"
)

type R struct {
	devices        []entity.Mikrotik
	defaultDevices []entity.Mikrotik
	chatWLs        []entity.ChatWL
}

func New(devices []config.Mikrotik, chatWLs []config.ChatWL) *R {
	var r R

	for _, d := range devices {
		m := entity.Mikrotik(d)

		if d.IsDefault {
			r.defaultDevices = append(r.defaultDevices, m)
		}

		r.devices = append(r.devices, m)
	}

	for _, c := range chatWLs {
		r.chatWLs = append(r.chatWLs, entity.ChatWL(c))
	}

	return &r
}

func (r *R) ChatWLs(ctx context.Context, chatID int64) ([]entity.ChatWL, error) {
	return r.chatWLs, nil
}

func (r *R) DefaultMikroTiks(ctx context.Context) ([]entity.Mikrotik, error) {
	return r.defaultDevices, nil
}

func (r *R) MikroTikByID(ctx context.Context, id int64) (entity.Mikrotik, error) {
	for _, d := range r.devices {
		if d.ID == id {
			return d, nil
		}
	}

	return entity.Mikrotik{}, xerrors.ErrNotFound
}
