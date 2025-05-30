package usersmicrorepo

import (
	"context"

	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
	"google.golang.org/grpc"
)

type UsersMicro struct {
	Client v1.UsersClient
}

func New(connectURL string) (*UsersMicro, error) {
	conn, err := grpc.NewClient(connectURL)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := v1.NewUsersClient(conn)

	return &UsersMicro{client}, nil
}

func (um *UsersMicro) GetByTelegramID(ctx context.Context, telegramID string) (*v1.User, error) {
	response, err := um.Client.GetUserByTgId(ctx, &v1.GetUserByTgIdRequest{TelegramId: telegramID})
	if err != nil {
		return nil, err
	}

	return response.User, nil
}

func (um *UsersMicro) CreateUser(ctx context.Context, telegramID, chatID, firstName string, lastName *string) (*v1.User, error) {
	response, err := um.Client.CreateUser(ctx, &v1.CreateUserRequest{TelegramId: telegramID, ChatId: chatID, FirstName: firstName, LastName: lastName})
	if err != nil {
		return nil, err
	}

	return response.User, nil
}
