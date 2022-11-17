package helpers

import (
	"net"
	"regexp"
	"strings"

	"wlbot/internal/xerrors"
)

var findIPRegexp = regexp.MustCompile(`[1-9][0-9]{0,2}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)

var findIPWithCIDRRegexp = regexp.MustCompile(`[1-9][0-9]{0,2}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\/[0-9]{1,3}`)

func FindIP(text string) (ip net.IP, ipNet *net.IPNet, err error) {
	defer func() { err = xerrors.Wrap(err, "failed to find ip in text") }()

	foundIP := findIPWithCIDRRegexp.FindString(text)
	if foundIP == "" {
		foundIP = findIPRegexp.FindString(text)
	}

	foundIP = strings.TrimSpace(foundIP)

	ip, ipNet, err = net.ParseCIDR(foundIP)
	if err != nil {
		err = nil

		ip = net.ParseIP(foundIP)
		if ip == nil {
			err = xerrors.ErrNotFound
			return
		}
	}

	if ip.IsUnspecified() || ip.IsPrivate() {
		err = xerrors.ErrNotFound
		return
	}

	foundIP = ip.String()

	return
}
