package mikrotik

import (
	"context"
	"errors"
	"kms/wlbot/internal/entity"
	"kms/wlbot/internal/xerrors"
)

func (s *Service) AddIPFromChat(ctx context.Context, chatID int64, ip string) (err error) {
	defer func() { err = xerrors.Wrap(err, "failed to add ip from chat") }()

	wls, err := s.repo.ChatWLs(ctx, chatID)
	if err != nil {
		return err
	}

	comment := ""

	if len(wls) > 0 {
		return s.addIPToCustomMikrotiks(ctx, wls, ip, comment)
	}

	return s.AddIPToDefaultMikrotiks(ctx, ip, comment)
}

func (s *Service) AddIPToDefaultMikrotiks(ctx context.Context, ip, comment string) (err error) {
	defer func() { err = xerrors.Wrap(err, "failed to add ip to default mikrotiks") }()

	mikroTiks, err := s.repo.DefaultMikroTiks(ctx)
	if err != nil {
		return err
	}

	for _, m := range mikroTiks {
		s.l.Infow("add ip to default mikrotik", "mikrotik_address", m.Address, "mikrotik_id", m.ID, "wl", m.DefaultWL, "ip", ip)

		isDynamic, err := s.device.FindIP(ctx, m, m.DefaultWL, ip)
		switch {
		case err == nil:
			if isDynamic {
				s.l.Infow("found dynamic ip, try to remove", "mikrotik_id", m.ID, "wl", m.DefaultWL, "ip", ip)

				err = s.device.RemoveIP(ctx, m, m.DefaultWL, ip)
				if err != nil {
					return err
				}

				s.l.Infow("ip successfully removed", "mikrotik_id", m.ID, "wl", m.DefaultWL, "ip", ip)
			}

			err = s.device.AddIP(ctx, m, ip, comment)
			if err != nil {
				return err
			}

		case errors.Is(err, xerrors.ErrNotFound):
			s.l.Infow("ip is not found, try to add", "mikrotik_id", m.ID, "wl", m.DefaultWL, "ip", ip)

			err = s.device.AddIP(ctx, m, ip, comment)
			if err != nil {
				return err
			}

			s.l.Infow("ip successfully added", "mikrotik_id", m.ID, "wl", m.DefaultWL, "ip", ip)
		default:
			return err
		}
	}

	return nil
}

func (s *Service) addIPToCustomMikrotiks(ctx context.Context, wls []entity.ChatWL, ip, comment string) (err error) {
	defer func() { err = xerrors.Wrap(err, "failed to add ip to custom mikrotiks") }()

	addToDefault := false

	for _, v := range wls {
		s.l.Infow("add ip to custom mikrotik", "mikrotik_id", v.MikrotikID, "wl", v.MikrotikWL, "ip", ip)

		if v.UseDefault {
			addToDefault = true
		}

		m, err := s.repo.MikroTikByID(ctx, v.MikrotikID)
		if err != nil {
			return err
		}

		isDynamic, err := s.device.FindIP(ctx, m, v.MikrotikWL, ip)
		switch {
		case err == nil:
			if isDynamic {
				s.l.Infow("found dynamic ip, try to remove", "mikrotik_id", v.MikrotikID, "wl", v.MikrotikWL, "ip", ip)

				err = s.device.RemoveIP(ctx, m, v.MikrotikWL, ip)
				if err != nil {
					return err
				}

				s.l.Infow("ip successfully removed", "mikrotik_id", v.MikrotikID, "wl", v.MikrotikWL, "ip", ip)
			}

			err = s.device.AddIPToCustomWL(ctx, m, v.MikrotikWL, ip, comment)
			if err != nil {
				return err
			}

		case errors.Is(err, xerrors.ErrNotFound):
			s.l.Infow("ip is not found, try to add", "mikrotik_id", v.MikrotikID, "wl", v.MikrotikWL, "ip", ip)

			err = s.device.AddIPToCustomWL(ctx, m, v.MikrotikWL, ip, comment)
			if err != nil {
				return err
			}

			s.l.Infow("ip successfully added", "mikrotik_id", v.MikrotikID, "wl", v.MikrotikWL, "ip", ip)
		default:
			return err
		}
	}

	if addToDefault {
		s.l.Infow("try to add ip to default mikrotiks", "ip", ip)

		return s.AddIPToDefaultMikrotiks(ctx, ip, comment)
	}

	return nil
}
