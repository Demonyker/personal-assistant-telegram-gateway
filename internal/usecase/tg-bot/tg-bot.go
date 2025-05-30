package tgbot

import (
	"context"
	"fmt"
	"strconv"

	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/users/v1"
	"github.com/Demonyker/personal-assistant-telegram-gateway/internal/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// UseCase -.
type UseCase struct {
	repository repo.UsersMicro
	bot        *tgbotapi.BotAPI
}

// New -.
func New(repository repo.UsersMicro, bot *tgbotapi.BotAPI) *UseCase {
	return &UseCase{repository: repository, bot: bot}
}

func (u *UseCase) CreateUser(ctx context.Context, telegramID, chatID int64, firstName, lastName string) (*v1.User, error) {
	existedUser, err := u.repository.GetByTelegramID(ctx, strconv.Itoa(int(telegramID)))
	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		return existedUser, nil
	}

	createdUser, err := u.repository.CreateUser(ctx, strconv.Itoa(int(telegramID)), strconv.Itoa(int(chatID)), firstName, &lastName)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *UseCase) GetUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, errorsChannel chan<- error) {
	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			user, err := u.CreateUser(ctx, update.Message.From.ID, update.Message.Chat.ID, update.Message.From.FirstName, update.Message.From.LastName)

			if err != nil {
				errorsChannel <- err
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Добро пожаловать, %s %s !", user.FirstName, *user.LastName))

				_, err = u.bot.Send(msg)
				if err != nil {
					errorsChannel <- err
				}
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите команду /start")

			_, err := u.bot.Send(msg)
			if err != nil {
				errorsChannel <- err
			}
		}
	}
}
