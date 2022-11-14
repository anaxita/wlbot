package main

import (
	"log"

	"kms/wlbot/internal/dal/repository"
	"kms/wlbot/internal/external/mikrotikclient"
	"kms/wlbot/internal/service/authenticator"
	"kms/wlbot/internal/service/config"
	"kms/wlbot/internal/service/mikrotik"
	"kms/wlbot/internal/transport/rest"
	"kms/wlbot/internal/transport/telegram"
	"kms/wlbot/pkg/logging"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

const configPath = "configs/app.yml"

func main() {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Panic("loadl config: ", err)
	}

	l, err := logging.New(cfg.Debug, cfg.LogFile)
	defer l.Sync()

	// telegram bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   cfg.TGBotToken,
		OnError: func(err error, c telebot.Context) { l.Error(zap.Error(err)) },
	})
	if err != nil {
		l.Panic(err)
	}

	// repository
	repo := repository.New(cfg.MikroTiks, cfg.ChatWLs, cfg.AdminChats)

	// external services
	mkrClient := mikrotikclient.New()

	// check mikrotik devices health
	err = mkrClient.HealthCheck(cfg.MikroTiks...)
	if err != nil {
		l.Panic(err)
	}

	// internal services
	mkr := mikrotik.New(l, repo, mkrClient)
	auth := authenticator.New(cfg.AdminChats, cfg.AdminUsers)

	// api
	go telegram.New(cfg.Debug, bot, mkr, auth).Start()

	server := rest.NewServer(cfg.HTTPPort, mkr)

	l.Fatal(server.Start())
}
