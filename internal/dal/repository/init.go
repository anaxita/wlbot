package repository

import (
	"context"

	"wlbot/internal/entity"
	"wlbot/internal/service/config"
	"wlbot/internal/xerrors"
)

type R struct {
	devices        []entity.Mikrotik
	defaultDevices []entity.Mikrotik
	chatWLs        []entity.ChatWL
	adminChatIDs   []int64
}

func New(devices []config.Mikrotik, chatWLs []config.ChatWL, adminChatIDs []int64) *R {
	var r R

	r.adminChatIDs = adminChatIDs

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

func (r *R) ChatWLs(ctx context.Context, chatID int64) []entity.ChatWL {
	var chats []entity.ChatWL

	for _, c := range r.chatWLs {
		if c.ChatID == chatID {
			chats = append(chats, c)
		}
	}

	return chats
}

func (r *R) DefaultMikroTiks() []entity.Mikrotik {
	return r.defaultDevices
}

func (r *R) MikroTikByID(ctx context.Context, id int64) (entity.Mikrotik, error) {
	for _, d := range r.devices {
		if d.ID == id {
			return d, nil
		}
	}

	return entity.Mikrotik{}, xerrors.ErrNotFound
}

func (r *R) AdminChatIDs() []int64 {
	return r.adminChatIDs
}
