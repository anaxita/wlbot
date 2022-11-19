package config_test

import (
	"testing"

	"wlbot/internal/service/config"

	"github.com/stretchr/testify/assert"
)

var (
	correctCfg = config.App{
		HTTPPort:   "8080",
		TGBotToken: "mytoken",
		Debug:      true,
		LogFile:    "app.log",
		MikroTiks: []config.Mikrotik{
			correctMkr,
			correctMkr,
			correctMkr,
		},
		ChatWLs: []config.ChatWL{
			correctChatWL,
			correctChatWL,
			correctChatWL,
		},
		AdminChats: make([]int64, 1),
		AdminUsers: make([]string, 1),
	}

	correctMkr = config.Mikrotik{
		ID:        1,
		Address:   "192.168.1.1:80",
		Login:     "dev",
		Password:  "dev",
		DefaultWL: "dev",
		IsDefault: true,
	}

	correctChatWL = config.ChatWL{
		ChatID:     1,
		MikrotikID: 1,
		MikrotikWL: "wl",
		UseDefault: false,
	}
)

func TestApp_Validate_Positive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		config config.App
	}{
		{
			name:   "correct",
			config: correctCfg,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.Validate()
			assert.NoError(t, err)
		})
	}
}

func TestApp_Validate_Negative(t *testing.T) {
	t.Parallel()

	// empty mikrotiks
	emptyMikrotiks := correctCfg
	emptyMikrotiks.MikroTiks = nil

	// incorrect mikrotik address
	incorrectMikrotiksAddr := correctCfg
	mikr := correctMkr
	mikr.Address = "incorrect"
	incorrectMikrotiksAddr.MikroTiks = []config.Mikrotik{mikr}

	// empty chat wls
	emptyChatWLs := correctCfg
	emptyChatWLs.ChatWLs = nil

	tests := []struct {
		name   string
		config config.App
	}{
		{
			name:   "all fields are empty",
			config: config.App{},
		},
		{
			name:   "empty mikroticks",
			config: emptyMikrotiks,
		},
		{
			name:   "empty chat wls",
			config: emptyChatWLs,
		},
		{
			name:   "incorrect mikrotiks address",
			config: incorrectMikrotiksAddr,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.Validate()
			assert.Error(t, err)
		})
	}
}
