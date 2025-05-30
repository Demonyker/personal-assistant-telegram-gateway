package usecase

import (
	"context"

	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	// TgBot -.
	TgBot interface {
		CreateUser(ctx context.Context, telegramID, chatID int64, firstName string, lastName *string) (*v1.User, error)
		GetUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, errorsChannel chan<- error)
	}
)
