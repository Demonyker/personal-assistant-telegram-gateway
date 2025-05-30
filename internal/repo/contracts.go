package repo

import (
	"context"

	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
)

type (
	// UsersMicro -.
	UsersMicro interface {
		GetByTelegramID(ctx context.Context, telegramID string) (*v1.User, error)
		CreateUser(ctx context.Context, telegramID, chatID, firstName string, lastName *string) (*v1.User, error)
	}
)
