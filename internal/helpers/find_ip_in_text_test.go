package helpers

import (
	"net"
	"testing"

	"wlbot/internal/xerrors"

	"github.com/stretchr/testify/require"
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
			text:       "Hello, please add 85.249.35.209 in the white list.",
			wantIp4:    net.ParseIP("85.249.35.209"),
			withSubnet: false,
			wantErr:    nil,
		},
		{
			name:       "correct ip4 with new line",
			text:       "Hello, please add\n85.249.35.209\nin the white list.",
			wantIp4:    net.ParseIP("85.249.35.209"),
			withSubnet: false,
			wantErr:    nil,
		},
		{
			name:       "correct ip4 with subnet",
			text:       "Hello, please add 85.249.35.0/24 in the white list.",
			wantIp4:    net.ParseIP("85.249.35.0"),
			withSubnet: true,
			wantErr:    nil,
		},
		{
			name:       "zeros",
			text:       "Hello, please add 0.0.0.0 in the white list.",
			wantIp4:    nil,
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

			require.Equal(t, tt.wantIp4.String(), gotIp4.String())
			require.Equal(t, tt.withSubnet, gotWithCIDR != nil)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
