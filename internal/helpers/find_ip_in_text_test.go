package helpers

import (
	"github.com/stretchr/testify/require"
	"kms/wlbot/internal/xerrors"
	"net"
	"testing"
)

func TestFindIP4(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		wantIp4    net.IP
		withSubnet bool
		wantErr    error
	}{
		{
			name:       "correct ip4",
			text:       "Hello, please add 192.168.1.1 in the white list.",
			wantIp4:    net.ParseIP("192.168.1.1"),
			withSubnet: false,
			wantErr:    nil,
		},
		{
			name:       "correct ip4 with new line",
			text:       "Hello, please add\n192.168.1.1\nin the white list.",
			wantIp4:    net.ParseIP("192.168.1.1"),
			withSubnet: false,
			wantErr:    nil,
		},
		{
			name:       "correct ip4 with subnet",
			text:       "Hello, please add 192.168.1.0/24 in the white list.",
			wantIp4:    net.ParseIP("192.168.1.0"),
			withSubnet: true,
			wantErr:    nil,
		},
		{
			name:       "zeros",
			text:       "Hello, please add 0.0.0.0 in the white list.",
			wantIp4:    net.IPv4zero,
			withSubnet: false,
			wantErr:    xerrors.ErrNotFound,
		},
		{
			name:       "empty text",
			text:       "",
			wantIp4:    nil,
			withSubnet: false,
			wantErr:    xerrors.ErrNotFound,
		},
		{
			name:       "text without ip",
			text:       "text without ip",
			wantIp4:    nil,
			withSubnet: false,
			wantErr:    xerrors.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIp4, gotWithCIDR, err := FindIP(tt.text)

			require.Equal(t, tt.wantIp4, gotIp4)
			require.Equal(t, tt.withSubnet, gotWithCIDR != nil)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
