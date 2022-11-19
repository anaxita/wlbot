package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func New(configPath string) (*App, error) {
	cfg, err := load(configPath)
	if err != nil {
		return nil, err
	}

	return &cfg, cfg.Validate()
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
