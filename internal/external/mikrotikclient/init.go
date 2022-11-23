package mikrotikclient

import (
	"strings"
	"sync"
	"time"

	"wlbot/internal/entity"
	"wlbot/internal/service/config"
	"wlbot/internal/xerrors"

	"go.uber.org/zap"
	"gopkg.in/routeros.v2"
)

type Client struct {
	l     *zap.SugaredLogger
	mu    sync.Mutex
	conns map[int64]*routeros.Client
}

func New(l *zap.SugaredLogger) *Client {
	return &Client{
		l:     l,
		mu:    sync.Mutex{},
		conns: make(map[int64]*routeros.Client),
	}
}

func (c *Client) HealthCheck(devices ...config.Mikrotik) error {
	errs := make([]string, 0, len(devices))

	var wg sync.WaitGroup

	wg.Add(len(devices))

	for _, v := range devices {
		go func(v config.Mikrotik) {
			defer wg.Done()

			c.l.Info("Checking connection to ", v.Address)

			c.mu.Lock()
			defer c.mu.Unlock()

			const timeout = time.Second * 3

			client, err := routeros.DialTimeout(v.Address, v.Login, v.Password, timeout)
			if err != nil {
				errs = append(errs, err.Error())

				return
			}

			c.l.Info("Connection to ", v.Address, " is OK")

			c.conns[v.ID] = client
		}(v)
	}

	wg.Wait()

	if len(errs) > 0 {
		return xerrors.Wrap(xerrors.ErrHealthCheck, strings.Join(errs, "; "))
	}

	return nil
}

func (c *Client) AddIP(m entity.Mikrotik, ip, comment string) error {
	return c.AddIPToCustomWL(m, m.DefaultWL, ip, comment)
}

func (c *Client) AddIPToCustomWL(m entity.Mikrotik, wl, ip, comment string) error {
	client, err := c.dial(m)
	if err != nil {
		return err
	}

	_, err = client.Run("/ip/firewall/address-list/add", "=list="+wl, "=address="+ip, "=comment=\""+comment+"\"")
	if err != nil {
		if strings.Contains(err.Error(), "already have such entry") {
			return xerrors.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (c *Client) RemoveIP(m entity.Mikrotik, wl string, ip string) (err error) {
	client, err := c.dial(m)
	if err != nil {
		return
	}

	findIP, err := client.Run("/ip/firewall/address-list/print", "?address="+ip, "?list="+wl)
	if err != nil {
		return err
	}

	l := len(findIP.Re)
	if l == 0 || l > 1 {
		return xerrors.ErrNotFound
	}

	ipID, ok := findIP.Re[0].Map[".id"]
	if !ok {
		return xerrors.ErrNotFound
	}

	_, err = client.Run("/ip/firewall/address-list/remove", "=.id="+ipID)
	if err != nil {
		return err
	}

	return nil
}

// dial returns a cached connection to the Mikrotik device or creates a new one.
func (c *Client) dial(m entity.Mikrotik) (*routeros.Client, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	client, ok := c.conns[m.ID]
	if ok {
		return client, nil
	}

	const timeout = time.Second * 3

	client, err := routeros.DialTimeout(m.Address, m.Login, m.Password, timeout)
	if err != nil {
		return nil, err
	}

	c.conns[m.ID] = client

	return client, nil
}
