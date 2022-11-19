package config

// Mikrotik provides the Mikrotik device configuration.
type Mikrotik struct {
	ID        int64  `yaml:"id" validate:"required"`
	Address   string `yaml:"address" validate:"required,hostname_port"`
	Login     string `yaml:"login" validate:"required"`
	Password  string `yaml:"password" validate:"required"`
	DefaultWL string `yaml:"default_wl" validate:"required"`
	IsDefault bool   `yaml:"is_default"`
}
