package config

// ChatWL provides the chat whitelist configuration.
type ChatWL struct {
	ChatID     int64  `yaml:"chat_id" validate:"required"`
	MikrotikID int64  `yaml:"mikrotik_id" validate:"required"`
	MikrotikWL string `yaml:"mikrotik_wl" validate:"required"`
	UseDefault bool   `yaml:"use_default"`
}
