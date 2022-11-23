package mikrotik

import (
	"context"
	"errors"
	"fmt"

	"wlbot/internal/entity"
	"wlbot/internal/xerrors"
)

func (s *Service) AddIPFromChat(ctx context.Context, chatID int64, ip string, comment string) (err error) {
	wls := s.repo.ChatWLs(ctx, chatID)

	if len(wls) > 0 {
		return s.AddIPToCustomMikrotiks(ctx, wls, ip, comment)
	}

	return s.AddIPToDefaultMikrotiks(ip, comment)
}

func (s *Service) AddIPToDefaultMikrotiks(ip, comment string) error {
	for _, m := range s.repo.DefaultMikroTiks() {
		err := s.device.RemoveIP(m, m.DefaultWL, ip)
		if err != nil && !errors.Is(err, xerrors.ErrNotFound) {
			return fmt.Errorf("%w: remove ip failed: %s; mikrotik_id = %d; wl=%s; ip=%s",
				xerrors.ErrMikrotik, err, m.ID, m.DefaultWL, ip)
		}

		err = s.device.AddIP(m, ip, comment)
		if err != nil && !errors.Is(err, xerrors.ErrAlreadyExists) {
			return fmt.Errorf("%w: add ip failed: %s; mikrotik_id = %d; wl=%s; ip=%s",
				xerrors.ErrMikrotik, err, m.ID, m.DefaultWL, ip)
		}
	}

	return nil
}

func (s *Service) AddIPToCustomMikrotiks(ctx context.Context, wls []entity.ChatWL, ip, comment string) error {
	addToDefault := false

	for _, v := range wls {
		if v.UseDefault {
			addToDefault = true
		}

		m, err := s.repo.MikroTikByID(ctx, v.MikrotikID)
		if err != nil {
			return err
		}

		err = s.device.RemoveIP(m, v.MikrotikWL, ip)
		if err != nil && !errors.Is(err, xerrors.ErrNotFound) {
			return fmt.Errorf("%w: remove ip failed: %s; mikrotik_id = %d; wl=%s; ip=%s",
				xerrors.ErrMikrotik, err, m.ID, v.MikrotikWL, ip)
		}

		err = s.device.AddIPToCustomWL(m, v.MikrotikWL, ip, comment)
		if err != nil && !errors.Is(err, xerrors.ErrAlreadyExists) {
			return fmt.Errorf("%w: add ip failed: %s; mikrotik_id = %d; wl=%s; ip=%s",
				xerrors.ErrMikrotik, err, m.ID, v.MikrotikWL, ip)
		}
	}

	if addToDefault {
		return s.AddIPToDefaultMikrotiks(ip, comment)
	}

	return nil
}
