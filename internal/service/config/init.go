package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// App provides the application configuration.
type App struct {
	HTTPPort   string `yaml:"http_port"`
	TGBotToken string `yaml:"tg_bot_token"`
	Debug      bool   `yaml:"debug"`
	LogFile    string `yaml:"log_file" `

	MikroTiks []Mikrotik `yaml:"mikrotiks"`
	ChatWLs   []ChatWL   `yaml:"chat_wls"`

	AdminChats []int64  `yaml:"admin_chats"`
	AdminUsers []string `yaml:"admin_users"`
}

// Mikrotik provides the Mikrotik device configuration.
type Mikrotik struct {
	ID        int64  `yaml:"id"`
	Address   string `yaml:"address"`
	Login     string `yaml:"login"`
	Password  string `yaml:"password"`
	DefaultWL string `yaml:"default_wl"`
	IsDefault bool   `yaml:"is_default"`
}

// ChatWL provides the chat whitelist configuration.
type ChatWL struct {
	ChatID     int64  `yaml:"chat_id"`
	MikrotikID int64  `yaml:"mikrotik_id"`
	MikrotikWL string `yaml:"mikrotik_wl"`
	UseDefault bool   `yaml:"use_default"`
}

func New(configPath string) (*App, error) {
	cfg, err := load(configPath)
	if err != nil {
		return nil, err
	}

	return &cfg, cfg.validate()
}

func load(configPath string) (cfg App, err error) {
	f, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&cfg)

	return
}

// validate checks the configuration for correctness.
func (a App) validate() error {
	mapErrs := make(map[string]string)

	if a.HTTPPort == "" {
		mapErrs["http_port"] = "empty"
	}

	if a.TGBotToken == "" {
		mapErrs["tg_bot_token"] = "empty"
	}

	if a.LogFile == "" {
		mapErrs["logfile"] = "empty"
	}

	// if len(a.MikroTiks) == 0 {
	// 	mapErrs["mikrotiks"] = "empty"
	// }
	//
	// if len(a.ChatWLs) == 0 {
	// 	mapErrs["chat_wls"] = "empty"
	// }
	//
	// if len(a.AdminChats) == 0 {
	// 	mapErrs["admin_chats"] = "empty"
	// }
	//
	// if len(a.AdminUsers) == 0 {
	// 	mapErrs["admin_users"] = "empty"
	// }

	if len(mapErrs) > 0 {
		var b strings.Builder
		b.Grow(len(mapErrs))

		for k, v := range mapErrs {
			b.WriteString("\n")
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(v)
			b.WriteString("; ")
		}

		return fmt.Errorf("validation failed: " + b.String())
	}

	return nil
}
