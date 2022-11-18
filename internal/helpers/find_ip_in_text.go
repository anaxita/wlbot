package helpers

import (
	"net"
	"regexp"
	"strings"

	"wlbot/internal/xerrors"
)

var findIPRegexp = regexp.MustCompile(`[1-9][0-9]{0,2}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)

var findIPWithCIDRRegexp = regexp.MustCompile(`[1-9][0-9]{0,2}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\/[0-9]{1,3}`)

func FindIP(text string) (ip string, ipNet string, err error) {
	defer func() { err = xerrors.Wrap(err, "failed to find ip4 in text") }()

	foundIP := findIPWithCIDRRegexp.FindString(text)
	if foundIP == "" {
		foundIP = findIPRegexp.FindString(text)
	}

	foundIP = strings.TrimSpace(foundIP)

	ip4, cidr, err := net.ParseCIDR(foundIP)
	if err != nil {
		err = nil

		ip4 = net.ParseIP(foundIP)
		if ip4 == nil {
			return "", "", xerrors.ErrNotFound
		}
	}

	if ip4.IsUnspecified() || ip4.IsPrivate() {
		return "", "", xerrors.ErrNotFound
	}

	cidrStr := ""
	if cidr != nil {
		cidrStr = cidr.String()
	}

	return ip4.String(), cidrStr, nil
}
