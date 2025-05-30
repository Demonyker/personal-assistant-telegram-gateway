package usersMicroRepo

import (
	"context"
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
	"google.golang.org/grpc"
)

type UsersMicro struct {
	Client v1.UsersClient
}

func New(connectUrl string) (*UsersMicro, error) {
	conn, err := grpc.NewClient(connectUrl)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	client := v1.NewUsersClient(conn)
	return &UsersMicro{client}, nil
}

func (um *UsersMicro) GetByTelegramId(ctx context.Context, telegramId string) (*v1.User, error) {
	response, err := um.Client.GetUserByTgId(ctx, &v1.GetUserByTgIdRequest{TelegramId: telegramId})

	if err != nil {
		return nil, err
	}

	return response.User, nil
}

func (um *UsersMicro) CreateUser(ctx context.Context, telegramId, chatId, firstName string, lastName *string) (*v1.User, error) {
	response, err := um.Client.CreateUser(ctx, &v1.CreateUserRequest{TelegramId: telegramId, ChatId: chatId, FirstName: firstName, LastName: lastName})

	if err != nil {
		return nil, err
	}

	return response.User, nil
}
