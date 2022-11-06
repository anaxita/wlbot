package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"time"
)

type Service struct {
	configPath string

	mu        sync.Mutex
	config    App
	updatedAt time.Time
}

func (s *Service) ConfigPath() string {
	return s.configPath
}

func (s *Service) Config() App {
	s.mu.Lock()
	s.mu.Unlock()

	return s.config
}

func (s *Service) LastUpdateAt() time.Time {
	return s.updatedAt
}

type App struct {
	TGBotToken string `yaml:"tg_bot_token"`
	Debug      bool   `yaml:"debug"`
	LogFile    string `yaml:"log_file" `

	MikroTiks []Mikrotik `yaml:"mikrotiks"`
	ChatWLs   []ChatWL   `yaml:"chat_wls"`

	AdminChats []int64  `yaml:"admin_chats"`
	AdminUsers []string `yaml:"admin_users"`

	Database Database `yaml:"database"`
}

type Mikrotik struct {
	ID        int64  `yaml:"id"`
	Address   string `yaml:"address"`
	Login     string `yaml:"login"`
	Password  string `yaml:"password"`
	DefaultWL string `yaml:"default_wl"`
	IsDefault bool   `yaml:"is_default"`
}

type ChatWL struct {
	ChatID     int64  `yaml:"chat_id"`
	MikrotikID int64  `yaml:"mikrotik_id"`
	MikrotikWL string `yaml:"mikrotik_wl"`
	UseDefault bool   `yaml:"use_default"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func New(configPath string) (*Service, error) {
	cfg, err := load(configPath)
	if err != nil {
		return nil, err
	}

	return &Service{
		configPath: configPath,
		config:     cfg,
		updatedAt:  time.Now(),
	}, nil
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

func (s *Service) Reload() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cfg, err := load(s.configPath)
	if err != nil {
		return err
	}

	s.updatedAt = time.Now()
	s.config = cfg

	return nil
}
