package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Demonyker/personal-assistant-telegram-gateway/config"
	usersMicroRepo "github.com/Demonyker/personal-assistant-telegram-gateway/internal/repo/users-micro"
	tgbot "github.com/Demonyker/personal-assistant-telegram-gateway/internal/usecase/tg-bot"
	"github.com/Demonyker/personal-assistant-telegram-gateway/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository register
	usersMicroRepository, err := usersMicroRepo.New(cfg.Grpc.UsersMicroAddr)
	if err != nil {
		l.Fatal(err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TG.BotKey)
	if err != nil {
		l.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	tgBotUsecase := tgbot.New(usersMicroRepository, bot)
	errorsChannel := make(chan error)

	go tgBotUsecase.GetUpdates(context.Background(), updates, errorsChannel)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err := <-errorsChannel:
		l.Info("app - Run - error: %s", err.Error())
	}
}
