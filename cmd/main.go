package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"wlbot/internal/dal/repository"
	"wlbot/internal/external/mikrotikclient"
	"wlbot/internal/service/authenticator"
	"wlbot/internal/service/config"
	"wlbot/internal/service/mikrotik"
	"wlbot/internal/transport/rest"
	"wlbot/internal/transport/telegram"
	"wlbot/pkg/logging"

	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
)

const configPath = "configs/app.yml"

func main() {
	cfg, err := config.New(configPath)
	if err != nil {
		log.Panic("load config: ", err)
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
	l.Debug("mikrotik devices health check: ok")

	// internal services
	mkr := mikrotik.New(l, repo, mkrClient)
	auth := authenticator.New(cfg.AdminChats, cfg.AdminUsers)

	// api
	go telegram.New(cfg.Debug, bot, mkr, auth).Start()

	srv := rest.NewServer(cfg.HTTPPort, mkr)

	doneCh := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		bot.Stop()
		l.Info("Telegram bot stopped")

		if err := srv.Shutdown(context.Background()); err != nil {
			l.Error("HTTP server shutdown failed:", zap.Error(err))
		}

		close(doneCh)
	}()

	l.Debug("HTTP server started at:", cfg.HTTPPort)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		l.Fatal("HTTP server listen and serve failed:", zap.Error(err))
	}

	<-doneCh
}
