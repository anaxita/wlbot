package helpers_test

import (
	"testing"

	"wlbot/internal/helpers"
	"wlbot/internal/xerrors"

	"github.com/stretchr/testify/assert"
)

func TestFindIP4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		text     string
		wantIP4  string
		wantCIDR string
		wantErr  error
	}{
		{
			name:     "correct ip4",
			text:     "Hello, please add 85.249.35.209 in the white list.",
			wantIP4:  "85.249.35.209",
			wantCIDR: "",
			wantErr:  nil,
		},
		{
			name:     "correct ip4 with new line",
			text:     "Hello, please add\n85.249.35.209\nin the white list.",
			wantIP4:  "85.249.35.209",
			wantCIDR: "",
			wantErr:  nil,
		},
		{
			name:     "correct ip4 with subnet",
			text:     "Hello, please add 85.249.35.0/24 in the white list.",
			wantIP4:  "85.249.35.0",
			wantCIDR: "85.249.35.0/24",
			wantErr:  nil,
		},
		{
			name:     "zeros",
			text:     "Hello, please add 0.0.0.0 in the white list.",
			wantIP4:  "",
			wantCIDR: "",
			wantErr:  xerrors.ErrNotFound,
		},
		{
			name:     "empty text",
			text:     "",
			wantIP4:  "",
			wantCIDR: "",
			wantErr:  xerrors.ErrNotFound,
		},
		{
			name:     "text without ip",
			text:     "text without ip",
			wantIP4:  "",
			wantCIDR: "",
			wantErr:  xerrors.ErrNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotIP4, gotCIDR, err := helpers.FindIP(tt.text)

			assert.Equal(t, tt.wantIP4, gotIP4, "ip4")
			assert.Equal(t, tt.wantCIDR, gotCIDR, "cidr")
			assert.ErrorIs(t, err, tt.wantErr, "err")
		})
	}
}
