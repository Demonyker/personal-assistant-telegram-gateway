package repo

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
)

type (
	// UsersMicro -.
	UsersMicro interface {
		GetByTelegramId(ctx context.Context, telegramId string) (*v1.User, error)
		CreateUser(ctx context.Context, telegramId, chatId, firstName string, lastName *string) (*v1.User, error)
	}
)
