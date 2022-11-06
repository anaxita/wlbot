package mikrotikclient

import (
	"context"
	"fmt"
	"gopkg.in/routeros.v2"
	"kms/wlbot/internal/entity"
	"kms/wlbot/internal/service/config"
	"kms/wlbot/internal/xerrors"
	"strings"
	"sync"
	"time"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) FindIP(ctx context.Context, m entity.Mikrotik, wl string, ip string) (isDynamic bool, err error) {
	return true, nil // TODO replace with real implementation

	// client, err := routeros.DialTimeout(m.Address, m.Login, m.Password, time.Second*3)
	// if err != nil {
	// 	return
	// }
	// defer client.Close()
	//
	// r, err := client.Run(
	// 	"/ip/firewall/address-list/print",
	// 	"=list="+wl,
	// 	"=address="+ip,
	// )
	// if err != nil {
	// 	return
	// }
	//
	// dynamicField := r.Re[0].Map["is_dynamic"]
	//
	// isDynamic, _ = strconv.ParseBool(dynamicField)
	//
	// return
}

func (c *Client) HealthCheck(ctx context.Context, devices ...config.Mikrotik) error {
	errs := make([]string, 0, len(devices))

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(len(devices))

	for _, v := range devices {
		go func(v config.Mikrotik) {
			defer wg.Done()
			client, err := routeros.DialTimeout(v.Address, v.Login, v.Password, time.Second*3)
			if err != nil {
				mu.Lock()
				defer mu.Unlock()

				errs = append(errs, err.Error())
				return
			}

			client.Close()
		}(v)
	}

	wg.Wait()

	if len(errs) > 0 {
		return fmt.Errorf("health check failed: %s", strings.Join(errs, "; "))
	}

	return nil
}

func (c *Client) AddIP(ctx context.Context, m entity.Mikrotik, ip, comment string) error {
	return c.AddIPToCustomWL(ctx, m, m.DefaultWL, ip, comment)
}

func (c *Client) AddIPToCustomWL(ctx context.Context, m entity.Mikrotik, wl, ip, comment string) error {
	client, err := routeros.DialTimeout(m.Address, m.Login, m.Password, time.Second*3)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Run("/ip/firewall/address-list/add", "=list="+wl, "=address="+ip, "=comment=\""+comment+"\"")

	return err
}

func (c *Client) RemoveIP(ctx context.Context, m entity.Mikrotik, wl string, ip string) (err error) {
	client, err := routeros.DialTimeout(m.Address, m.Login, m.Password, time.Second*3)
	if err != nil {
		return
	}
	defer client.Close()

	r, err := client.Run(
		"/ip/firewall/address-list/print",
		"=.proplist=.id",
		"=list="+wl,
		"=address="+ip,
	)
	if err != nil {
		return
	}

	l := len(r.Re)
	if l == 0 || l > 1 {
		return xerrors.ErrNotFound
	}

	id, ok := r.Re[0].Map[".id"]
	if !ok {
		return xerrors.ErrNotFound
	}

	r, err = client.Run(
		"/ip/firewall/address-list/remove",
		"=list="+wl,
		"=.id="+id,
	)
	if err != nil {
		return
	}

	return
}
