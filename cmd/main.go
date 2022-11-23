package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"wlbot/internal/dal/repository"
	"wlbot/internal/external/mikrotikclient"
	"wlbot/internal/service/authenticator"
	"wlbot/internal/service/config"
	"wlbot/internal/service/mikrotik"
	"wlbot/internal/service/notificator"
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
	if err != nil {
		log.Fatal("init logger: ", err)
	}

	defer func(l *zap.SugaredLogger) {
		if err := l.Sync(); err != nil {
			log.Println("sync logger: ", err)
		}
	}(l)

	// telegram bot
	l.Info("Starting telegram bot...")

	bot, err := telebot.NewBot(telebot.Settings{
		Token:     cfg.TGBotToken,
		ParseMode: telebot.ModeMarkdown,
		OnError:   func(err error, c telebot.Context) { l.Error(zap.Error(err)) },
	})
	if err != nil {
		l.Fatal(err)
	}

	l.Info("Telegram bot started")

	// repository
	repo := repository.New(cfg.MikroTiks, cfg.ChatWLs, cfg.AdminChats)

	// external services
	mkrClient := mikrotikclient.New(l)

	// check mikrotik devices health
	err = mkrClient.HealthCheck(cfg.MikroTiks...)
	if err != nil {
		l.Fatal(err)
	}

	l.Debug("mikrotik devices health check: ok")

	// internal services
	mkr := mikrotik.New(repo, mkrClient)
	auth := authenticator.New(cfg.AdminChats, cfg.AdminUsers)
	notif := notificator.New(l, repo, bot)

	// api
	tgMw := telegram.NewMiddleware(l, auth)
	go telegram.New(l, cfg.Debug, tgMw, bot, mkr, auth).Start()

	srv := rest.NewServer(l, cfg.HTTPPort, notif, mkr)

	doneCh := make(chan struct{})
	go handleSignals(l, doneCh, bot, srv)

	l.Debug("HTTP server started at:", cfg.HTTPPort)

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		l.Fatal("HTTP server listen and serve failed:", zap.Error(err))
	}

	l.Info("HTTP server stopped")

	<-doneCh
}

func handleSignals(l *zap.SugaredLogger, doneCh chan struct{}, bot *telebot.Bot, srv *rest.Server) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	bot.Stop()
	l.Info("Telegram bot stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		l.Error("HTTP server shutdown failed:", zap.Error(err))
	}

	close(doneCh)
}
