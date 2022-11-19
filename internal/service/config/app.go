package config

import (
	"github.com/go-playground/validator/v10"
)

// App provides the application configuration.
type App struct {
	HTTPPort   string `yaml:"http_port" validate:"required,numeric"`
	TGBotToken string `yaml:"tg_bot_token" validate:"required"`
	Debug      bool   `yaml:"debug"`
	LogFile    string `yaml:"log_file"  validate:"required"`

	MikroTiks []Mikrotik `yaml:"mikrotiks" validate:"required,min=1"`
	ChatWLs   []ChatWL   `yaml:"chat_wls" validate:"required,min=1"`

	AdminChats []int64  `yaml:"admin_chats" validate:"required,min=1"`
	AdminUsers []string `yaml:"admin_users" validate:"required,min=1"`
}

// Validate checks the configuration for correctness.
func (a App) Validate() error {
	v := validator.New()

	err := v.Struct(a)
	if err != nil {
		return err
	}

	for _, m := range a.MikroTiks {
		err = v.Struct(m)
		if err != nil {
			return err
		}
	}

	for _, c := range a.ChatWLs {
		err = v.Struct(c)
		if err != nil {
			return err
		}
	}

	return nil
}
